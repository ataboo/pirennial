package reader

import (
	"fmt"
	"io"
	"strings"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/remote/sensor"
	"github.com/jacobsa/go-serial/serial"
	"github.com/op/go-logging"
)

var logger logging.Logger

type SensorReaderSerial struct {
	cfg        config.Serial
	connection io.ReadWriteCloser
}

func CreateSensorReaderSerial(cfg config.Serial) SensorReader {
	r := SensorReaderSerial{
		cfg: cfg,
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

func (r *SensorReaderSerial) getSensorData() (map[uint]int, error) {
	var vals map[uint]int

	return vals, nil
}

func (r *SensorReaderSerial) connect() (err error) {
	options := serial.OpenOptions{
		PortName:        r.cfg.PortName,
		BaudRate:        r.cfg.BaudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	r.connection, err = serial.Open(options)

	return err
}

func (r *SensorReaderSerial) getSerialRaw() (buf []byte, err error) {
	if r.connection == nil {
		err = r.connect()
		if err != nil {
			return nil, fmt.Errorf("failed to connect to serial: %s", err.Error())
		}
	}

	buf = make([]byte, r.cfg.BufferSize)
	r.connection.Write([]byte("1"))
	_, err = r.connection.Read(buf)
	if err != nil {
		r.connection.Close()
		r.connection = nil
		return nil, fmt.Errorf("failed to read from serial: %s", err)
	}

	buf = []byte(strings.Trim(string(buf), "\r\n\x00"))

	return buf, nil
}
