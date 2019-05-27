package sensors

import "github.com/ataboo/pirennial/hardware/sensors/soil"

type SensorService struct {
	sensorreader.SensorReader
}

func CreateSensorService() {
	sensors := make([]*soil.SoilSensor, len(cfg.SoilSensors))
	for i := 0; i < len(cfg.SoilSensors); i++ {
		sensors[i] = soil.CreateSoilSensorSerial(cfg.SoilSensors[i])
	}
}
