package pumps

import (
	"fmt"
	"sync"
	"time"

	"github.com/ataboo/pirennial/services/clock"
	"github.com/ataboo/pirennial/services/config"
)

// Pump provides water to plant
type Pump struct {
	relay           *Relay
	runStart        time.Time
	flowLPM         float64
	primeTimeMillis int
	sprinkles       SprinkleLog
	lock            sync.Mutex
	stopChan        chan int
}

// NewPump create a new pump
func NewPump(pumpCfg config.Pump) *Pump {
	pump := Pump{
		relay:           NewRelay(pumpCfg.Pin),
		flowLPM:         pumpCfg.FlowLPM,
		primeTimeMillis: pumpCfg.PrimeTimeMillis,
	}

	return &pump
}

// IsOn determine if the pump is currently running
func (p *Pump) IsOn() bool {
	return p.relay.IsOn()
}

// Sprinkle pump a number of liters
func (p *Pump) Sprinkle(liters float64) error {
	duration := p.timeToPumpVolume(liters)
	if err := p.pumpForTime(duration); err != nil {
		return err
	}

	p.sprinkles.AddLog(liters)
	return nil
}

// Stop pumping
func (p *Pump) Stop() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.stopChan == nil {
		return fmt.Errorf("stop chan not active")
	}

	p.stopChan <- 0
	p.stopChan = nil

	return nil
}

// Cleanup the relay pin
func (p *Pump) Cleanup() {
	p.relay.Cleanup()
}

func (p *Pump) pumpForTime(duration time.Duration) error {
	if p.IsOn() {
		return fmt.Errorf("pump on pin %d is already running", p.relay.PinNumber())
	}

	if err := p.initStop(); err != nil {
		return err
	}

	done := clock.After(duration)
	p.relay.On()
	go func() {
		defer p.relay.Off()
		for {
			select {
			case <-p.stopChan:
				logger.Debugf("pump (%d) stopped by stopChan", p.relay.PinNumber())
				return
			case <-done:
				p.cancelStop()
				return
			}
		}
	}()

	return nil
}

func (p *Pump) timeToPumpVolume(liters float64) time.Duration {
	seconds := float32(liters/p.flowLPM) * 60.0

	return time.Duration(seconds*float32(time.Second)) + time.Duration(p.primeTimeMillis)*time.Millisecond
}

func (p *Pump) initStop() error {
	defer p.lock.Unlock()
	p.lock.Lock()

	if p.stopChan != nil {
		return fmt.Errorf("stop chan already active")
	}

	p.stopChan = make(chan int)

	return nil
}

func (p *Pump) cancelStop() error {
	defer p.lock.Unlock()
	p.lock.Lock()

	if p.stopChan == nil {
		return fmt.Errorf("stop chan not active")
	}

	p.stopChan = nil

	return nil
}
