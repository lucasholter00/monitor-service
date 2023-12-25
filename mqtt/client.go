package mqtt

import (
	"Monitoring-service/controllers"
	"fmt"
	"log"
	"net/url"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqtt_client mqtt.Client

func GetInstance() mqtt.Client {

	if mqtt_client == nil {
		mqtt_client = mqtt.NewClient(getOptions())
		if token := mqtt_client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	return mqtt_client

}

func getOptions() *mqtt.ClientOptions {
	broker := os.Getenv("BROKER_URL")
	url, err := url.Parse(broker)
	if err != nil {
		log.Fatal(err)
	}
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", url))
	opts.SetClientID("go_mqh2nbjk,n,u1h2ieh12iuett_client")
	opts.SetUsername("")
	opts.SetPassword("")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("MQTT client is connected")
	controllers.InitialiseLogger(client)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
