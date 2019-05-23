package pumps

import (
	"time"

	"github.com/ataboo/pirennial/config"
)

type Pump struct {
	Relay     *Relay
	runStart  time.Time
	flow      float32
	primeTime float32
}

func NewPump(pumpCfg config.Pump) *Pump {
	pump := Pump{
		Relay:     NewRelay(pumpCfg.Pin),
		flow:      pumpCfg.Flow,
		primeTime: pumpCfg.PrimeTime,
	}

	return &pump
}
