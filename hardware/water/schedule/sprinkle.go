package schedule

// TODO: this should end up in a higher business logic layer with an interface for the pump model

import (
	"sort"
	"time"

	"github.com/ataboo/pirennial/environment/clock"
)

// Sprinkle representation of a watering event performed by a pump
type Sprinkle struct {
	startTime time.Time
	volume    float64
}

// SprinkleLog slice of Sprinkles
type SprinkleLog []Sprinkle

// SumWithinDurationAgo get the total sprinkle amounts within duration before now
func (l SprinkleLog) SumWithinDurationAgo(duration time.Duration) float64 {
	sum := float64(0)
	cutoff := clock.Now().Add(-duration)

	for _, s := range l {
		if s.startTime.Before(cutoff) {
			break
		}

		sum += s.volume
	}

	return sum
}

// GCBeforeDurationAgo remove sprinkles more than `duration` before now
func (l *SprinkleLog) GCBeforeDurationAgo(duration time.Duration) {
	cutoff := clock.Now().Add(-duration)

	for i := len(*l); i >= 0; i-- {
		if (*l)[i].startTime.Before(cutoff) {
			l.popBack()
		} else {
			break
		}
	}
}

// AddLog add a sprinkle amount to the log
func (l *SprinkleLog) LogPumpVolume(volume float64) {
	s := Sprinkle{
		startTime: clock.Now(),
		volume:    volume,
	}

	l.push(s)
}

func (l *SprinkleLog) Sort() {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].startTime.Before((*l)[j].startTime)
	})
}

// Push a Sprinkle to the front of the log
func (l *SprinkleLog) push(s Sprinkle) {
	*l = append(*l, s)
}

// popBack pop a Sprinkle off the back of the log
func (l *SprinkleLog) popBack() *Sprinkle {
	back := (*l)[len(*l)-1]
	*l = (*l)[0 : len(*l)-1]

	return &back
}
