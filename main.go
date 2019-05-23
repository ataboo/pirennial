package main

import (
	"github.com/ataboo/pirennial/config"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirenneal")
}

func main() {
	cfg := config.Cfg()

}
