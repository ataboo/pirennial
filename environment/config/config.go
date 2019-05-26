package config

import (
	"io/ioutil"
	"runtime"

	"github.com/ataboo/pirennial/environment/filestorage"
	"github.com/op/go-logging"

	"github.com/pelletier/go-toml"
)

var logger *logging.Logger

// GPIOActive if the arch is arm (pi)
var GPIOActive bool

func init() {
	logger = logging.MustGetLogger("pirennial")
	GPIOActive = runtime.GOARCH == "arm"
}

func LoadTOMLFile(assetPath string, output interface{}) error {
	cfgPath, err := filestorage.AssetPath(assetPath)
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
