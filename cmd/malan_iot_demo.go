package main

import (
	"context"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-queue/kq"
	"net/http"
	"util/pkg/kafkamq"
	"util/pkg/mqtt"

	"github.com/gin-gonic/gin"
)

type IotParams struct {
	Mid       int64  `json:"mid"`
	Type      string `json:"type"`
	Expire    int    `json:"expire"`
	App       string `json:"app"`
	Timestamp string `json:"timestamp"`
	DeviceId  string `json:"deviceId"`
	Param     struct {
		Method string `json:"method"`
		Paras  struct {
			Payload map[string]interface {
			} `json:"payload"`
			ExtendParam struct {
				Token       string `json:"token"`
				Url         string `json:"url"`
				ContentType string `json:"contentType"`
				FileValue   struct {
					FileName string `json:"fileName"`
					FileSize string `json:"fileSize"`
					FilePath string `json:"filePath"`
					FileMd5  string `json:"fileMd5"`
				} `json:"fileValue"`
			} `json:"extendParam"`
		} `json:"paras"`
	} `json:"param"`
}

type appRequest struct {
	Mid       int64  `json:"mid"`
	DeviceId  string `json:"deviceId"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Code      int    `json:"code"`
	App       string `json:"app"`
	Msg       string `json:"msg"`
	Param     struct {
		Method string `json:"method"`
		Paras  interface {
		} `json:"paras"`
	} `json:"param"`
}

var (
	kqClient kafkamq.KafkaClient
	mqClient *mqtt.MqttServer
)

func init() {
	kqClient = NewKqClient()
	NewMqtt()
}

func main() {

	mqttConfig := &mqtt.MqttConfig{
		Host:     "82.156.56.237:1883",
		Password: "EMQemq@1172",
		UserName: "emqx_u",
		Topic:    "iot/test",
	}
	mqttServer, err := mqtt.NewMqtt(mqttConfig, "")
	if err != nil {

	}

	mqttServer.Read("iot/app/pub", MqProxy)

	router := gin.Default()
	router.POST("/iot/v1/service/response", addResult)

	router.Run(":8080")
}

func addResult(c *gin.Context) {
	var params appRequest
	c.ShouldBindJSON(&params)
	body, _ := json.Marshal(&params)
	// 将数据推送到mqtt
	mqClient.Write("iot/app/sub", string(body))

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Gin!",
	})
}

func MqProxy(client MQTT.Client, message MQTT.Message) {
	ctx := context.Background()
	fmt.Printf("Received message: %s\n", message.Payload())
	iotParams := new(IotParams)
	json.Unmarshal(message.Payload(), iotParams)
	fmt.Println(iotParams)
	// 推送至kafka
	err := kqClient.KqPusherClient.KPush(ctx, "", string(message.Payload()))
	if err != nil {
		fmt.Println("err kafka", err)
	}
}

func NewKqClient() kafkamq.KafkaClient {
	topic := "service_request"
	config := kafkamq.Config{
		KqPusherConf: kq.KqConf{
			Topic:   topic,
			Brokers: []string{"127.0.0.1:9092"},
		},
	}
	kqClient := kafkamq.NewKqClient(config)
	return kqClient
}

func NewMqtt() {
	mqttConfig := &mqtt.MqttConfig{
		Host:     "82.156.56.237:1883",
		Password: "EMQemq@1172",
		UserName: "emqx_u",
		Topic:    "iot/test",
	}
	mqClient, _ = mqtt.NewMqtt(mqttConfig, "")
}
