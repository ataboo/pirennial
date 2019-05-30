package reader

import "github.com/ataboo/pirennial/hardware/sensors/sensor"

type SensorReader interface {
	Update([]sensor.Sensor) error
	Cleanup()
	Sleep() error
}
