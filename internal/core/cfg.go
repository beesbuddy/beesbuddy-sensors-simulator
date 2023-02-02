package core

import (
	"github.com/beesbuddy/beesbuddy-sensors-simulator/internal/models"
	c "github.com/leonidasdeim/goconfig"
	fh "github.com/leonidasdeim/goconfig/pkg/filehandler"
)

var cfgObject *c.Config[models.Config]

func init() {
	h, _ := fh.New(handler.WithName("dev"))
	cfg, err := c.Init[models.Config](h)

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
