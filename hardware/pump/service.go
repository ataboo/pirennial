package pump

import (
	"github.com/ataboo/pirennial/environment/config"
)

type PumpService struct {
	pumps []Pump
}

func CreatePumpService(cfg config.HardwareConfig) (s *PumpService) {
	s.pumps = make([]Pump, len(cfg.GPIO.Pumps))

	for i, pumpCfg := range cfg.GPIO.Pumps {
		if config.GPIOActive {
			s.pumps[i] = CreatePumpGPIO(pumpCfg)
		} else {
			s.pumps[i] = CreatePumpMock(pumpCfg)
		}
	}

	return s
}

func (s *PumpService) Cleanup() {
	for _, pump := range s.pumps {
		pump.Cleanup()
	}
}
