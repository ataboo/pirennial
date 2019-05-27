package reader

import (
	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/remote/sensor"
)

type SensorReaderMock struct {
	cfg    config.Serial
	Values map[uint]int
}

func CreateSensorReaderMock(cfg config.Serial, startVal int) SensorReader {
	var vals map[uint]int

	for i := 0; i < 30; i++ {
		vals[uint(i)] = startVal
	}

	r := SensorReaderMock{
		cfg:    cfg,
		Values: vals,
	}

	return &r
}

func (r *SensorReaderMock) Update(sensors []sensor.Sensor) error {
	for _, sensor := range sensors {
		sensor.Data().Value = r.Values[uint(sensor.Data().InputPin)]
	}

	return nil
}

func (r *SensorReaderMock) Cleanup() {
	//
}
