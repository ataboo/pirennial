package pumps

import (
	"github.com/ataboo/pirennial/hardware/pump"
	"github.com/ataboo/pirennial/services/config"
)

type PumpControl struct {
	cfg   config.HardwareConfig
	pumps []pump.Pump
}

func CreatePumpControl(cfg config.HardwareConfig) *PumpControl {
	pumps := make([]pump.Pump, len(cfg.Pumps))

	// for i := 0; i < len(cfg.Pumps); i++ {
	// 	pumps[i] = NewPump(cfg.Pumps[i])
	// }

	pc := PumpControl{
		cfg:   cfg,
		pumps: pumps,
	}

	return &pc
}
