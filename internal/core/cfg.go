package core

import (
	config "github.com/beesbuddy/beesbuddy-config"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/model"
)

var appConfig *model.Config

func GetConfig() *model.Config {
	return appConfig
}

func InitializeConfig() {
	cfg, err := config.Init[model.Config](config.WithName("dev"))

	if err != nil {
		panic("Unable to load config")
	}

	appConfig = cfg.GetCfg()
}
