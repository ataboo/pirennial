package soilhumidity

import "github.com/ataboo/pirennial/environment/config"

type SoilSensor struct {
	InputPin uint `json:"pin"`
	PowerPin uint
	Value    uint `json:"value"`
	Active   bool
}

func CreateSoilSensorSerial(cfg config.SoilSensor) &SoilSensor {
	s := SoilSensor{
		inputPin: cfg.InputPin,
		powerPin: cfg.PowerPin,
	}

	return &s
}
