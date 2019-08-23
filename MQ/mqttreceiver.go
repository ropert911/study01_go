package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"strings"
	"study01_go/MQ/mqtt"
)

func onIncomingDataReceived(client MQTT.Client, message MQTT.Message) {
	var response map[string]interface{}
	json.Unmarshal(message.Payload(), &response)

	fmt.Printf("Topic [%s] receiv==%s\n", message.Topic(), string(message.Payload()))
}

func main() {
	var scheme = "tcp"
	var brokerUrl = "192.168.20.46"
	var brokerPort = 1883
	var MqttClientId = "clientIDReceiv"
	var username = "admin"
	var password = "admin"
	var keepAlive = 3600
	var topic = "dev/+"
	var qos = byte(1)

	uri := &url.URL{
		Scheme: strings.ToLower(scheme),
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	client, err := mqtt.CreateMQTTClient(MqttClientId, uri, keepAlive)
	if err != nil {
		fmt.Printf("createMQTTClient error:%v\n", err)
		return
	}
	//defer client.Disconnect(5000)

	token := client.Subscribe(topic, qos, onIncomingDataReceived)
	if token.Wait() && token.Error() != nil {
		fmt.Printf("[Response listener] Stop command response listening. Cause:%v", token.Error())
		return
	}

	fmt.Printf("[Response listener] Start command response listening. \n")
	select {}
}
