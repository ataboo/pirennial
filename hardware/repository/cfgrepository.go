package repository

import (
	"github.com/ataboo/pirennial/hardware/pump"
	"github.com/ataboo/pirennial/hardware/soilsensor"
	"github.com/ataboo/pirennial/services/config"
)

type CfgRepository struct {
	cfg config.HardwareConfig
}

func CreateCfgRepository() Repository {
	cfg, err := config.LoadHardwareConfig()

	repo := CfgRepository{
		cfg: cfg,
	}

	return &repo
}

func (r *CfgRepository) GetPumps() []pump.Pump {
	cfg := config.Cfg().Pumps
}

func (r *CfgRepository) GetSoilSensors() []soilsensor.SoilSensor {

}

// // NewPump create a new pump
// func (r *CfgRepository) newPump(pumpCfg config.Pump) *Pump {
// 	pump := Pump{
// 		relay:           relay,
// 		flowLPM:         pumpCfg.FlowLPM,
// 		primeTimeMillis: pumpCfg.PrimeTimeMillis,
// 	}

// 	return &pump
// }
