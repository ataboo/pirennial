package soilsensor

import "github.com/ataboo/pirennial/environment/config"

type SoilSensorSerial struct {
	inputPin uint `json:"pin"`
	powerPin uint
	value    uint `json:"value"`
	active   bool
}

func CreateSoilSensorSerial(cfg config.SoilSensor) SoilSensor {
	s := SoilSensorSerial{
		inputPin: cfg.InputPin,
		powerPin: cfg.PowerPin,
	}

	return &s
}

func (s *SoilSensorSerial) InputPin() uint {
	return s.inputPin
}

func (s *SoilSensorSerial) PowerPin() uint {
	return s.powerPin
}

func (s *SoilSensorSerial) RawValue() uint {
	return s.value
}

func (s *SoilSensorSerial) SetActive(on bool) {
	s.active = on
}
