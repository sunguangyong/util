package mqtt

import MQTT "github.com/eclipse/paho.mqtt.golang"

type MqttConfig struct {
	Host     string
	UserName string
	Password string
	Topic    string
}

type MqttServer struct {
	Host     string
	Name     string
	Topic    string
	Qos      int
	ClientId string
	UserName string
	Password string
	Client   MQTT.Client
}
