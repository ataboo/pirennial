package connection

import (
	"io"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/jacobsa/go-serial/serial"
)

type ConnectionArduino struct {
	serial io.ReadWriteCloser
}

func CreateConnectionArduino(cfg config.Serial) (connection Connection, err error) {
	options := serial.OpenOptions{
		PortName:        cfg.PortName,
		BaudRate:        cfg.BaudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	serial, err := serial.Open(options)
	if err != nil {
		return serial, err
	}

	c := ConnectionArduino{
		serial: serial,
	}

	return &c, nil
}

func (c *ConnectionArduino) Write(val []byte) (n int, err error) {
	return c.serial.Write(val)
}

func (c *ConnectionArduino) Read(buf []byte) (n int, err error) {
	return c.serial.Read(buf)
}

func (c *ConnectionArduino) Close() error {
	return c.Close()
}
