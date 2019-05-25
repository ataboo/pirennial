package config

import (
	"io/ioutil"

	"github.com/op/go-logging"

	"github.com/pelletier/go-toml"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirennial")
	// c, err := loadConfig()
	// if err != nil {
	// 	log.Fatal("failed to load config: " + err.Error())
	// }

	// loadedCfg = c
}

// HardwareConfig configuration for the hardware
type HardwareConfig struct {
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

func LoadHardwareConfig() (HardwareConfig, error) {
	cfg := HardwareConfig{}

	err := LoadTOMLFile("hardware_config.toml", &cfg)

	return cfg, err
}

func LoadTOMLFile(assetPath string, output interface{}) error {
	cfgPath, err := AssetPath(assetPath)
	if err != nil {
		return err
	}

	buff, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		logger.Errorf("failed to load config: ", err.Error())
		return err
	}

	err = toml.Unmarshal(buff, output)
	if err != nil {
		logger.Errorf("failed to unmarshal config: ", err.Error())
		return err
	}

	return nil
}
