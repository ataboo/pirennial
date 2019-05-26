package soilsensor

import "github.com/ataboo/pirennial/environment/config"

type SoilSensorMock struct {
	inputPin uint
	powerPin uint
	value    uint
	active   bool
}

func CreateSoilSensorMock(cfg config.SoilSensor) SoilSensor {
	s := SoilSensorMock{
		inputPin: cfg.InputPin,
		powerPin: cfg.PowerPin,
	}

	return &s
}

func (s *SoilSensorMock) InputPin() uint {
	return s.inputPin
}

func (s *SoilSensorMock) PowerPin() uint {
	return s.powerPin
}

func (s *SoilSensorMock) RawValue() uint {
	return s.value
}

func (s *SoilSensorMock) SetActive(on bool) {
	s.active = on
}
