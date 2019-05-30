package water

import (
	"github.com/ataboo/pirennial/environment/config"
	"github.com/ataboo/pirennial/hardware/water/pump"
)

type PumpService struct {
	pumps []pump.Pump
}

func CreatePumpService(cfg config.HardwareConfig) (s *PumpService) {
	s.pumps = make([]pump.Pump, len(cfg.GPIO.Pumps))

	for i, pumpCfg := range cfg.GPIO.Pumps {
		if config.GPIOActive {
			s.pumps[i] = pump.CreatePumpGPIO(pumpCfg)
		} else {
			s.pumps[i] = pump.CreatePumpMock(pumpCfg)
		}
	}

	return s
}

func (s *PumpService) Cleanup() {
	for _, pump := range s.pumps {
		pump.Cleanup()
	}
}
