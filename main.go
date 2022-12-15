package main

import (
	config "github.com/beesbuddy/beesbuddy-config"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/cmd"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/model"
)

func main() {
	appConfig := config.NewConfig[model.AppConfig](0).Cfg
	cmd.RunProducers(appConfig)
}
