package sensor

type SoilSensor struct {
	data SensorData
}

func CreateSoilSensorSerial(pin uint) Sensor {
	s := SoilSensor{
		data: SensorData{
			InputPin: pin,
			Value:    0,
		},
	}

	return &s
}

func (s *SoilSensor) Data() *SensorData {
	return &s.data
}
