package core

import (
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/models"
	c "github.com/leonidasdeim/goconfig"
)

var cfgObject *c.Config[models.Config]

func init() {
	cfg, err := c.Init[models.Config](c.WithName("dev"))

	if err != nil {
		panic("Unable to load config")
	}

	cfgObject = cfg
}

func GetCfg() models.Config {
	return cfgObject.GetCfg()
}

func GetCfgObject() *c.Config[models.Config] {
	return cfgObject
}
