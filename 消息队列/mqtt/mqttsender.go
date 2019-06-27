package mqtt

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type MQTTInfo struct {
	Protocol     string
	Host         string
	Port         int
	Username     string
	Password     string
	Qos          int
	KeepAlive    int
	MqttClientId string
	Topic        string
}

func (t *MQTTInfo) initMqtt(clientId string, topic string) {
	t.Protocol = "tcp"
	t.Host = "192.168.20.46"
	t.Port = 1883
	t.Username = "admin"
	t.Password = "admin"
	t.Qos = 0
	t.KeepAlive = 3600
	t.MqttClientId = clientId
	t.Topic = topic
	fmt.Println("MqttClient==", clientId)
	fmt.Println("MqttTopic==", topic)
}

func main() {
	var scheme = "tcp"
	var brokerUrl = "192.168.20.46"
	var brokerPort = 1883
	var MqttClientId = "clientID"
	var username = "admin"
	var password = "admin"
	var keepAlive = 3600
	var topic = "dev/"
	var qos = byte(1)

	uri := &url.URL{
		Scheme: strings.ToLower(scheme),
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	client, err := CreateMQTTClient(MqttClientId, uri, keepAlive)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := client.Publish(topic+"mac", qos, false, "test data")
	token.Wait()

	time.Sleep(6 * time.Second)
	client.Disconnect(250)
}
