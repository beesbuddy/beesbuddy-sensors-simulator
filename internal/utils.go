package internal

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var DefaultMessageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
