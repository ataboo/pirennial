package config

import (
	"io/ioutil"

	"github.com/op/go-logging"

	"github.com/pelletier/go-toml"
)

var cfg Config
var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirennial")
	cfg = loadConfig()
}

func Cfg() *Config {
	return &cfg
}

func loadConfig() Config {
	cfg := Config{}
	buff, err := ioutil.ReadFile(AssetPath("config.toml"))
	if err != nil {
		logger.Errorf("failed to load config: ", err.Error())
	}

	err = toml.Unmarshal(buff, &cfg)
	if err != nil {
		logger.Errorf("failed to unmarshal config: ", err.Error())
	}

	return cfg
}

type Config struct {
	Pumps []Pump
}

type Pump struct {
	Pin       uint
	Flow      float32
	PrimeTime float32
}
