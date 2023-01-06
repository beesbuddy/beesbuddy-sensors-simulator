package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/core"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/models"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type producerCmd struct {
	client MQTT.Client
	topics []string
}

func ProducerRunner(client MQTT.Client) core.CmdRunner {
	module := &producerCmd{client: client, topics: []string{}}
	return module
}

func (cmd *producerCmd) Run() {
	if token := cmd.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, apiary := range core.GetCfg().Apiaries {
		for _, hive := range apiary.Hives {
			topic := fmt.Sprintf("apiary/%s/hive/%s", apiary.Id, hive.Id)
			cmd.topics = append(cmd.topics, topic)

			go func(topic string) {
				for {
					m := models.Metrics{
						ClientId:    core.GetCfg().ClientId,
						ApiaryId:    apiary.Id,
						HiveId:      hive.Id,
						Temperature: fmt.Sprintf("%.2f", rand.Float64()*100),
						Humidity:    fmt.Sprintf("%.2f", rand.Float64()*100),
						Weight:      fmt.Sprintf("%d", rand.Intn(10000)),
					}
					serializedMetrics, _ := json.Marshal(m)

					if !core.GetCfg().Debug {
						log.Println("Publishing to topic:", topic, ", metrics: ", m)
					}

					token := cmd.client.Publish(topic, 0, false, serializedMetrics)
					token.Wait()
					time.Sleep(time.Duration(core.GetCfg().UploadInterval) * time.Second)
				}
			}(topic)

			if core.GetCfg().Debug {
				subscribe(cmd.client, topic)
			}

			time.Sleep(time.Duration(core.GetCfg().InitializationInterval) * time.Second)
		}
	}
}

func (cmd *producerCmd) CleanUp() {
	if core.GetCfg().Debug {
		for _, topic := range cmd.topics {
			unsubscribe(cmd.client, topic)
		}
	}

	cmd.client.Disconnect(250)
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
