package main

import (
	config "github.com/beesbuddy/beesbuddy-config"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/cmd"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/models"
)

func main() {
	appConfig := config.NewConfig[models.AppConfig](0).Cfg
	cmd.RunProducers(appConfig)
}
