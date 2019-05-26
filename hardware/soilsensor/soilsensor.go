package soilsensor

import (
	"github.com/ataboo/pirennial/environment/config"
)

type SoilSensor interface {
	InputPin() uint
	PowerPin() uint
	RawValue() uint
	SetActive(on bool)
}

func LoadSoilSensors() ([]SoilSensor, error) {
	var sensors []SoilSensor

	cfg, err := config.LoadHardwareConfig()
	if err != nil {
		return sensors, err
	}

	for i, sensor := range cfg.Serial.SoilSensors {
		if config.GPIOActive {
			sensors[i] = CreateSoilSensorSerial(sensor)
		} else {
			sensors[i] = CreateSoilSensorMock(sensor)
		}
	}

	return sensors, nil
}
