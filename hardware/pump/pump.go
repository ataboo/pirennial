package pump

import "github.com/ataboo/pirennial/environment/config"

type Pump interface {
	Sprinkle(liters float64) error
	Stop() error
	Cleanup()
	IsOn() bool
}

func LoadPumps() ([]Pump, error) {
	var pumps []Pump

	cfg, err := config.LoadHardwareConfig()
	if err != nil {
		return pumps, err
	}

	for i, pumpCfg := range cfg.GPIO.Pumps {
		if config.GPIOActive {
			pumps[i] = CreatePumpGPIO(pumpCfg)
		} else {
			pumps[i] = CreatePumpMock(pumpCfg)
		}
	}

	return pumps, nil
}
