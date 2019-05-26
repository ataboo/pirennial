package sensorreader

type SensorReader interface {
	Update() error
	SoilSensors() []SoilSensor
}
