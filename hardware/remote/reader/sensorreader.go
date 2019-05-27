package reader

import "github.com/ataboo/pirennial/hardware/remote/sensor"

type SensorReader interface {
	Update([]sensor.Sensor) error
	Cleanup()
}
