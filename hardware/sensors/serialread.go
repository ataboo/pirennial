package sensors

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ataboo/pirennial/services/config"
	"github.com/jacobsa/go-serial/serial"
	"github.com/op/go-logging"
)

var logger *logging.Logger
var serialRead io.ReadWriteCloser

func init() {
	logger = logging.MustGetLogger("pirennial")
}

func ReadSerial() (SerialReturn, error) {
	out := SerialReturn{}
	var err error

	if serialRead == nil {
		serialRead, err = connect(config.Cfg().Serial)
		if err != nil {
			return out, fmt.Errorf("failed to connect to serial: %s", err.Error())
		}
	}

	buf := make([]byte, 100000)
	_, err = serialRead.Read(buf)
	if err != nil {
		serialRead.Close()
		serialRead = nil
		return out, fmt.Errorf("failed to read from serial: %s", err)
	}

	json.Unmarshal(buf, &out)

	return out, nil
}

func Cleanup() {
	if serialRead != nil {
		serialRead.Close()
		serialRead = nil
	}
}

func connect(cfg config.Serial) (io.ReadWriteCloser, error) {
	options := serial.OpenOptions{
		PortName:        cfg.PortName,
		BaudRate:        cfg.BaudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	return serial.Open(options)
}
