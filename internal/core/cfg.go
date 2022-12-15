package core

import (
	config "github.com/beesbuddy/beesbuddy-config"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/model"
)

var appConfig model.Config

func GetConfig() model.Config {
	return appConfig
}

func InitializeConfig() {
	appConfig = config.NewConfig[model.Config](0).Cfg
}
