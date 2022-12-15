package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/models"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var defaultMessageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func RunProducers(appConfig models.AppConfig) {
	opts := MQTT.NewClientOptions().AddBroker(appConfig.BrokerTCPUrl)
	opts.SetClientID(appConfig.ClientId)
	opts.SetDefaultPublishHandler(defaultMessageHandler)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topics := []string{}

	for _, apiary := range appConfig.Apiaries {
		for _, hive := range apiary.Hives {
			topic := fmt.Sprintf("apiary/%s/hive/%s", apiary.Id, hive.Id)
			topics = append(topics, topic)

			go func(topic string) {
				for {
					s := models.Sensor{
						ClientId:    appConfig.ClientId,
						ApiaryId:    apiary.Id,
						HiveId:      hive.Id,
						Temperature: fmt.Sprintf("%.2f", rand.Float64()*100),
						Humidity:    fmt.Sprintf("%.2f", rand.Float64()*100),
						Weight:      fmt.Sprintf("%d", rand.Intn(10000)),
					}
					j, _ := json.Marshal(s)
					token := c.Publish(topic, 0, false, j)
					token.Wait()
					time.Sleep(time.Duration(appConfig.UploadInterval) * time.Second)
				}
			}(topic)

			if appConfig.Debug {
				subscribe(c, topic)
			}

			time.Sleep(time.Duration(appConfig.InitializationInterval) * time.Second)
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	if appConfig.Debug {
		for _, topic := range topics {
			unsubscribe(c, topic)
		}
	}

	c.Disconnect(250)
}

func unsubscribe(c MQTT.Client, topic string) {
	if token := c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func subscribe(c MQTT.Client, topic string) {
	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
