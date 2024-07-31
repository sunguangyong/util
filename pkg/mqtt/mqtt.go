package mqtt

import (
	"crypto/tls"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/threading"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	Qos0 = iota // 最多一次 消息可能会丢失
	Qos1        // 确保消息至少到达一次 可能会重复
	Qos2        // 确保消息只到达一次 可能会有延迟
)

func NewMqtt(mqttConfig *MqttConfig, clientId string) (svr *MqttServer, err error) {
	svr = new(MqttServer)
	mqttName := "MQTT"
	svr.Host = mqttConfig.Host
	svr.UserName = mqttConfig.UserName
	svr.Password = mqttConfig.Password
	svr.Name = mqttName
	svr.Topic = mqttConfig.Topic

	if clientId == "" {
		clientId = uuid.New().String()
	}

	// ClientId必须是唯一的，并且不同的MQTT客户端必须使用不同的ClientId
	svr.ClientId = clientId

	client, err := svr.NewClient()

	if err != nil {
		log.Println("mqtt client err ==== ", err)
		return nil, err
	}

	log.Println("mqtt success")
	svr.Client = client
	return svr, nil
}

func (svr *MqttServer) NewClient() (client MQTT.Client, err error) {
	connOpts := MQTT.NewClientOptions().AddBroker(svr.Host).SetClientID(svr.ClientId).SetCleanSession(true)
	if svr.UserName != "" {
		connOpts.SetUsername(svr.UserName)
		if svr.Password != "" {
			connOpts.SetPassword(svr.Password)
		}
	}
	connOpts.SetAutoReconnect(true)                    // 启用自动重连
	connOpts.SetConnectRetry(true)                     // 允许在建立连接之前发送数据
	connOpts.SetMaxReconnectInterval(10 * time.Second) // 设置最大重连间隔

	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)
	client = MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		log.Fatalln(err)
	}
	return client, err
}

func (svr *MqttServer) Write(topic, payload string) {
	threading.GoSafe(func() {
		// 发布MQTT消息
		token := svr.Client.Publish(topic, byte(svr.Qos), true, payload)
		ok := token.Wait()
		if ok {
			log.Printf("write success ===%s \n", payload)
		} else {
			log.Printf("write fail ===%s \n", payload)
		}
	})
}

func (svr *MqttServer) Read(topic string, callback MQTT.MessageHandler) {
	if token := svr.Client.Subscribe(topic, byte(Qos0), callback); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
}
