package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/core"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/model"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type mqttClientModule struct {
	client MQTT.Client
	topics []string
}

func NewMqttClientRunner(client MQTT.Client) core.ModuleRunner {
	module := &mqttClientModule{client: client, topics: []string{}}
	return module
}

func (mod *mqttClientModule) Run() {
	if token := mod.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, apiary := range core.GetCfgModel().Apiaries {
		for _, hive := range apiary.Hives {
			topic := fmt.Sprintf("apiary/%s/hive/%s", apiary.Id, hive.Id)
			mod.topics = append(mod.topics, topic)

			go func(topic string) {
				for {
					m := model.Metrics{
						ClientId:    core.GetCfgModel().ClientId,
						ApiaryId:    apiary.Id,
						HiveId:      hive.Id,
						Temperature: fmt.Sprintf("%.2f", rand.Float64()*100),
						Humidity:    fmt.Sprintf("%.2f", rand.Float64()*100),
						Weight:      fmt.Sprintf("%d", rand.Intn(10000)),
					}
					serializedMetrics, _ := json.Marshal(m)
					token := mod.client.Publish(topic, 0, false, serializedMetrics)
					token.Wait()
					time.Sleep(time.Duration(core.GetCfgModel().UploadInterval) * time.Second)
				}
			}(topic)

			if core.GetCfgModel().Debug {
				subscribe(mod.client, topic)
			}

			time.Sleep(time.Duration(core.GetCfgModel().InitializationInterval) * time.Second)
		}
	}
}

func (mod *mqttClientModule) CleanUp() {
	if core.GetCfgModel().Debug {
		for _, topic := range mod.topics {
			unsubscribe(mod.client, topic)
		}
	}

	mod.client.Disconnect(250)
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
