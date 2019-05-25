package pump

// import (
// 	"math"
// 	"testing"
// 	"time"

// 	"github.com/ataboo/pirennial/services/clock"
// 	"github.com/ataboo/pirennial/services/config"
// )

// func TestPumpForTimeAndStop(t *testing.T) {
// 	p := NewPump(config.Pump{})

// 	p.pumpForTime(time.Second)

// 	if !p.IsOn() {
// 		t.Error("pump expected to be on")
// 	}

// 	if p.stopChan == nil {
// 		t.Error("stop chan should be active")
// 	}

// 	err := p.pumpForTime(time.Second)
// 	if err == nil {
// 		t.Error("should get error when trying to start pump")
// 	}

// 	err = p.Stop()
// 	if err != nil {
// 		t.Error("should not receive error on stop")
// 	}

// 	if p.stopChan != nil {
// 		t.Error("stop chan should not be active")
// 	}

// 	if p.IsOn() {
// 		t.Error("pump should not longer be on")
// 	}

// 	err = p.Stop()
// 	if err == nil {
// 		t.Error("second stop should receive error")
// 	}

// 	p.Cleanup()
// }

// func TestPumpForTimeWithoutStop(t *testing.T) {
// 	fakeClock := clock.CreateFakeClock()
// 	clock.SetClock(fakeClock)
// 	afterChan := fakeClock.AfterChan
// 	p := NewPump(config.Pump{})

// 	err := p.pumpForTime(time.Second)
// 	if err != nil {
// 		t.Error("unnexpected error:", err)
// 	}

// 	if !p.IsOn() {
// 		t.Error("pump should be running")
// 	}

// 	afterChan <- time.Now()

// 	if p.IsOn() {
// 		t.Error("pump should be stopped")
// 	}

// 	if p.stopChan != nil {
// 		t.Error("stop chan should be innactive")
// 	}

// 	err = p.Stop()
// 	if err == nil {
// 		t.Error("stop call should generate error")
// 	}

// 	err = p.pumpForTime(time.Second)
// 	if err != nil {
// 		t.Error("should be able to start pump again")
// 	}

// 	if !p.IsOn() {
// 		t.Error("pump should be running")
// 	}

// 	afterChan <- time.Now()

// 	p.Cleanup()
// }

// func TestTimeToPumpVolume(t *testing.T) {
// 	rows := []struct {
// 		primeTimeMillis int
// 		flowLPM         float64
// 		liters          float64
// 		expected        time.Duration
// 	}{
// 		{1000, 1, 1, time.Second * 61},
// 		{1000, 10, 1, time.Second * 7},
// 	}

// 	cfg := config.Pump{
// 		PrimeTimeMillis: 500,
// 		FlowLPM:         2,
// 	}

// 	var p *Pump

// 	for _, row := range rows {
// 		cfg.PrimeTimeMillis = row.primeTimeMillis
// 		cfg.FlowLPM = row.flowLPM
// 		p = NewPump(cfg)

// 		duration := p.timeToPumpVolume(row.liters)

// 		t_assertTimesWithinNanos(duration, row.expected, 1e6, t)

// 		p.Cleanup()
// 	}

// }

// func TestSprinkle(t *testing.T) {
// 	fakeClock := clock.CreateFakeClock()
// 	clock.SetClock(fakeClock)
// 	afterChan := make(chan time.Time)
// 	expectedDuration := time.Second * 31

// 	fakeClock.AfterClosure = func(d time.Duration) <-chan time.Time {
// 		t_assertTimesWithinNanos(d, expectedDuration, 1e6, t)

// 		return afterChan
// 	}

// 	cfg := config.Pump{
// 		PrimeTimeMillis: 1000,
// 		FlowLPM:         2,
// 	}

// 	p := NewPump(cfg)

// 	p.Sprinkle(1)

// 	if !p.IsOn() {
// 		t.Error("pump should be running")
// 	}

// 	afterChan <- time.Now()

// 	if p.IsOn() {
// 		t.Error("pump should not be running")
// 	}

// 	p.Cleanup()
// }

// func t_assertTimesWithinNanos(a time.Duration, b time.Duration, maxDiff int64, t *testing.T) {
// 	nanoDiff := a.Nanoseconds() - b.Nanoseconds()
// 	nanoDiff = int64(math.Abs(float64(nanoDiff)))

// 	if nanoDiff > maxDiff {
// 		t.Errorf("unnexpected time %s, %s", a, b)
// 	}
// }
