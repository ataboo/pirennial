package sensors

import (
	"time"

	"github.com/op/go-logging"

	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/sensors/connection"
	"github.com/ataboo/pirennial/hardware/sensors/reader"
	"github.com/ataboo/pirennial/hardware/sensors/sensor"
)

var logger logging.Logger

type SensorService struct {
	sensors  []sensor.Sensor
	reader   reader.SensorReader
	active   bool
	stopChan chan int
	cfg      config.Serial
}

func init() {
	logger = *logging.MustGetLogger("pirennial")
}

func CreateSensorService(cfg config.HardwareConfig) (service *SensorService, err error) {
	service.cfg = cfg.Serial
	service.sensors = make([]sensor.Sensor, len(cfg.GPIO.Pumps))
	for i := 0; i < len(cfg.GPIO.Pumps); i++ {
		service.sensors[i] = sensor.CreateSoilSensorSerial(cfg.GPIO.Pumps[i].SensorPin)
	}

	if config.GPIOActive {
		conn, err := connection.CreateConnectionArduino(cfg.Serial)
		if err != nil {
			return nil, err
		}
		service.reader = reader.CreateSensorReaderSerial(cfg.Serial, conn)
	} else {
		service.reader = reader.CreateSensorReaderMock(cfg.Serial, 1)
	}

	go service.readRoutine()

	return service, err
}

func (s *SensorService) Sensors() []sensor.Sensor {
	return s.sensors
}

func (s *SensorService) Sleep() error {
	s.active = false

	return s.reader.Sleep()
}

func (s *SensorService) Wake() error {
	s.active = true

	return s.reader.Update(s.sensors)
}

func (s *SensorService) readRoutine() {
	tick := time.Tick(time.Millisecond * time.Duration(s.cfg.SensorUpdateMillis))

	for {
		select {
		case <-tick:
			if s.active {
				err := s.reader.Update(s.sensors)
				if err != nil {
					logger.Error("failed to update sensors", err)
				}
			}
		case <-s.stopChan:
			return
		}

	}
}

func (s *SensorService) Cleanup() {
	s.reader.Cleanup()
	s.stopChan <- 0
}
