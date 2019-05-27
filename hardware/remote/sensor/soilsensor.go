package sensor

import "github.com/ataboo/pirennial/environment/config"

type SoilSensor struct {
	data SensorData
}

func CreateSoilSensorSerial(cfg config.SoilSensor) Sensor {
	s := SoilSensor{
		data: SensorData{
			InputPin: cfg.InputPin,
			Value:    0,
		},
	}

	return &s
}

func (s *SoilSensor) Data() *SensorData {
	return &s.data
}
