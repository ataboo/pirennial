package reader

import (
	"fmt"
	"io"
	"strings"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/remote/connection"
	"github.com/ataboo/pirennial/hardware/remote/sensor"
	"github.com/jacobsa/go-serial/serial"
	"github.com/op/go-logging"
)

const (
	GET = 0x2
	SLEEP = 0x3
)

var logger logging.Logger

type SensorReaderSerial struct {
	cfg        config.Serial
	connection io.ReadWriteCloser
	buffer []byte
}

func CreateSensorReaderArduino(cfg config.Serial) SensorReader {
	r := SensorReaderSerial{
		cfg:        cfg,
		connection: connection.CreateSensorReaderArduino(cfg),
		buffer: make([]byte, cfg.BufferSize)
	}

	return &r
}

func CreateSensorReaderMock(cfg config.Serial) SensorReader {
	r := SensorReaderSerial {
		cfg config.Serial,
		connection: connection.CreateSensorReaderMock(cfg),
	}

	return &r
}

func (r *SensorReaderSerial) Update(sensors []sensor.Sensor) error {
	values, err := r.getSensorData()
	if err != nil {
		return fmt.Errorf("failed to read serial sensors:\n%s", err)
	}

	hadErr := false

	for _, sensor := range sensors {
		val, ok := values[sensor.Data().InputPin]
		if !ok {
			hadErr = true
			logger.Errorf("Failed to get pin %d from serial return", sensor.Data().InputPin)
		}
		sensor.Data().Value = val
	}

	if hadErr {
		return fmt.Errorf("some pins could not be read")
	}

	return nil
}

func (r *SensorReaderSerial) Cleanup() {
	if r.connection != nil {
		r.connection.Close()
		r.connection = nil
	}
}

func (r *SensorReaderSerial) Sleep() (err error) {
	n, err := r.connection.Write([]byte{SLEEP})

	return err
}

func (r *SensorReaderSerial) getSensorData() (out map[uint]int, err error) {	
	_, err := r.connection.Write([]byte{GET})
	
	n, err := r.connection.Read(r.buffer)
	if err != nil {
		return out, err
	}

	trimmed := buf[0:n]
	err = json.Unmarshal(trimmed, &out)
	return out, err
}
