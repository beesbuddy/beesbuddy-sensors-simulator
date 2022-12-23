package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/beesbuddy/beesbuddy-sensors-simulator/cmd"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/core"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	core.InitializeConfig()
	opts := MQTT.NewClientOptions().AddBroker(core.GetCfgModel().BrokerTCPUrl)
	opts.SetClientID(core.GetCfgModel().ClientId)
	opts.SetDefaultPublishHandler(internal.DefaultMessageHandler)
	client := MQTT.NewClient(opts)
	mqttClientRunner := cmd.NewMqttClientRunner(client)
	mqttClientRunner.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	mqttClientRunner.CleanUp()
}
