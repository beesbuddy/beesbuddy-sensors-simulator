package core

import (
	config "github.com/beesbuddy/beesbuddy-config"
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/model"
)

var cfg model.Config

func GetConfig() model.Config {
	return cfg
}

func InitializeConfig() {
	cfg = config.NewConfig[model.Config](0).Cfg
}
