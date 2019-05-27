package sensors

import (
	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/remote/connection"
	"github.com/ataboo/pirennial/hardware/remote/reader"
	"github.com/ataboo/pirennial/hardware/remote/sensor"
)

type SensorService struct {
	reader reader.SensorReader
}

func CreateSensorService(cfg config.Serial) (service *SensorService, err error) {
	sensors := make([]sensor.Sensor, len(cfg.SoilSensors))
	for i := 0; i < len(cfg.SoilSensors); i++ {
		sensors[i] = sensor.CreateSoilSensorSerial(cfg.SoilSensors[i])
	}

	if config.GPIOActive {
		conn, err := connection.CreateConnectionArduino(cfg)
		if err != nil {
			return nil, err
		}
		service.reader = reader.CreateSensorReaderSerial(cfg, conn)
	} else {
		service.reader = reader.CreateSensorReaderMock(cfg, 1)
	}

	return service, err
}
