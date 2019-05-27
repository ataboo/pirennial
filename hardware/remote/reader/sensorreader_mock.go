package reader

import (
	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/remote/sensor"
)

type SensorReaderMock struct {
	cfg     config.Serial
	FakeVal int
}

func CreateSensorReaderMock(cfg config.Serial, fakeVal int) SensorReader {
	r := SensorReaderMock{
		cfg:     cfg,
		FakeVal: fakeVal,
	}

	return &r
}

func (r *SensorReaderMock) Update(sensors []sensor.Sensor) error {
	for _, sensor := range sensors {
		sensor.Data().Value = r.FakeVal
	}

	return nil
}

func (r *SensorReaderMock) Cleanup() {
	//
}
