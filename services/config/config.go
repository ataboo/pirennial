package config

import (
	"io/ioutil"
	"log"

	"github.com/op/go-logging"

	"github.com/pelletier/go-toml"
)

var loadedCfg Config
var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirennial")
	c, err := loadConfig()
	if err != nil {
		log.Fatal("failed to load config: " + err.Error())
	}

	loadedCfg = c
}

// Cfg get the app's loaded config
func Cfg() *Config {
	return &loadedCfg
}

// Config for the app
type Config struct {
	Pumps  []Pump
	Serial Serial
}

// Pump config for a pump
type Pump struct {
	Pin             uint
	FlowLPM         float64
	PrimeTimeMillis int
}

type Serial struct {
	PortName          string
	BaudRate          uint
	RetryDelaySeconds uint
	BufferSize        uint
}

func loadConfig() (Config, error) {
	cfg := Config{}
	cfgPath, err := AssetPath("config.toml")
	if err != nil {
		return cfg, err
	}

	buff, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		logger.Errorf("failed to load config: ", err.Error())
		return cfg, err
	}

	err = toml.Unmarshal(buff, &cfg)
	if err != nil {
		logger.Errorf("failed to unmarshal config: ", err.Error())
		return cfg, err
	}

	return cfg, nil
}
