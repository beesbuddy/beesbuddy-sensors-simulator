package models

type Hive struct {
	Id string `validate:"required"`
}

type Apiary struct {
	Id    string `validate:"required"`
	Hives []Hive `validate:"required"`
}

type Config struct {
	AppName                string   `validate:"required"`
	ClientId               string   `validate:"required"`
	BrokerTCPUrl           string   `validate:"required"`
	UploadInterval         int      `default:"15"`
	InitializationInterval int      `default:"3"`
	Apiaries               []Apiary `validate:"required"`
	Debug                  bool     `default:"false"`
}
