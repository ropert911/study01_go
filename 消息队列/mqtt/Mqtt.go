package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"time"
)

func CreateMQTTClient(clientID string, uri *url.URL, keepAlive int) (MQTT.Client, error) {
	fmt.Printf("Create MQTT client and connection: uri=%v clientID=%v \n", uri.String(), clientID)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s", uri.Scheme, uri.Host))
	opts.SetClientID(clientID)
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetKeepAlive(time.Second * time.Duration(keepAlive))
	opts.SetConnectionLostHandler(func(client MQTT.Client, e error) {
		fmt.Printf("Connection lost : %v\n", e)
		token := client.Connect()
		if token.Wait() && token.Error() != nil {
			fmt.Printf("Reconnection failed : %v\n", token.Error())
		} else {
			fmt.Printf("Reconnection sucessful\n")
		}
	})

	client := MQTT.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return client, token.Error()
	}

	return client, nil
}
