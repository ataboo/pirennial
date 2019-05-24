package pumps

import "github.com/ataboo/pirennial/services/config"

type PumpControl struct {
	cfg   config.Config
	pumps []*Pump
}

func CreatePumpControl(cfg config.Config) *PumpControl {
	pumps := make([]*Pump, len(cfg.Pumps))

	for i := 0; i < len(cfg.Pumps); i++ {
		pumps[i] = NewPump(cfg.Pumps[i])
	}

	pc := PumpControl{
		cfg:   cfg,
		pumps: pumps,
	}

	return &pc
}
