package main

import (
	"time"

	"github.com/ataboo/pirennial/hardware/sensors"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirenneal")
}

func main() {
	// pumps.CreatePumpControl(*config.Cfg())

	tick := time.Tick(time.Millisecond * 500)
	for {
		select {
		case <-tick:
			ret, err := sensors.ReadSerial()
			if err != nil {
				logger.Error(err)
			} else {
				logger.Infof("Return: %+v", ret)
			}
		}
	}
}
