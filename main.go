package main

import (
	"github.com/ataboo/pirennial/hardware/pumps"
	"github.com/ataboo/pirennial/services/config"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirenneal")
}

func main() {
	ps := pumps.CreatePumpControl(*config.Cfg())

}
