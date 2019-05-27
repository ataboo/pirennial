package sensor

type SensorData struct {
	InputPin uint `json:"pin"`
	Value    int  `json:"value"`
}

type Sensor interface {
	Data() *SensorData
}
