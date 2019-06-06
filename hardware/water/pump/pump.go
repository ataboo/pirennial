package pump

type Pump interface {
	Sprinkle(liters float64) error
	Stop() error
	Cleanup()
	IsOn() bool
}

type PumpLogger interface {
	LogPumpVolume(liters float64)
}
