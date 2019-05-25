package sensors

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

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
	cfg := config.Cfg().Serial

	if serialRead == nil {
		serialRead, err = connect(cfg)
		if err != nil {
			return out, fmt.Errorf("failed to connect to serial: %s", err.Error())
		}
	}

	buf := make([]byte, cfg.BufferSize)
	serialRead.Write([]byte("1"))
	_, err = serialRead.Read(buf)
	if err != nil {
		serialRead.Close()
		serialRead = nil
		return out, fmt.Errorf("failed to read from serial: %s", err)
	}

	buf = []byte(strings.Trim(string(buf), "\r\n\x00"))

	json.Unmarshal(buf, &out.Sensors)
	out.LastRead = time.Now()

	if len(out.Sensors) == 0 {
		err = fmt.Errorf("no sensor data in received")
	} else {
		err = nil
	}

	return out, err
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
