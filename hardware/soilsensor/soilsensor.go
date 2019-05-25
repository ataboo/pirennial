package soilsensor

type SoilSensor interface {
	PinNumber() uint
	RawValue() uint
}
