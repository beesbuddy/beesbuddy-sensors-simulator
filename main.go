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
	opts := MQTT.NewClientOptions().AddBroker(core.GetCfg().BrokerTCPUrl)
	opts.SetClientID(core.GetCfg().ClientId)
	opts.SetDefaultPublishHandler(internal.DefaultMessageHandler)
	client := MQTT.NewClient(opts)
	mqttClientRunner := cmd.ProducerRunner(client)
	mqttClientRunner.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	mqttClientRunner.CleanUp()
}
