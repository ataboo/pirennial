package repository

import (
	"github.com/ataboo/pirennial/hardware/pump"
	"github.com/ataboo/pirennial/hardware/soilsensor"
)

type Repository interface {
	GetPumps() []pump.Pump
	GetSoilSensors() []soilsensor.SoilSensor
}
