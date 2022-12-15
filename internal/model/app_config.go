package model

type Hive struct {
	Id string `required:"true"`
}

type Apiary struct {
	Id    string `required:"true"`
	Hives []Hive `required:"true"`
}

type AppConfig struct {
	AppName                string   `required:"true"`
	ClientId               string   `required:"true"`
	BrokerTCPUrl           string   `requred:"true"`
	UploadInterval         int      `default:"15"`
	InitializationInterval int      `default:"3"`
	Apiaries               []Apiary `required:"true"`
	Debug                  bool     `default:"false"`
}
