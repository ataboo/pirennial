package main

import (
	"log"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/water"
	"github.com/ataboo/pirennial/hardware/sensors"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirenneal")
}

func main() {
	hardwareCfg, err := config.LoadHardwareConfig()
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	sensors, err := sensors.CreateSensorService(hardwareCfg)
	if err != nil {
		log.Fatal("failed to connect to sensor service", err)
	}

	pumps := water.CreatePumpService(hardwareCfg)

	defer func() {
		sensors.Cleanup()
		pumps.Cleanup()
	}()

}
