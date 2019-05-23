package pumps

import (
	"time"
)

// Sprinkle representation of a watering event performed by a pump
type Sprinkle struct {
	startTime time.Time
	amount    float32
}

// SprinkleLog slice of Sprinkles
type SprinkleLog []Sprinkle

// SumWithinDurationAgo get the total sprinkle amounts within duration before now
func (l SprinkleLog) SumWithinDurationAgo(duration time.Duration) float32 {
	sum := float32(0)
	cutoff := time.Now().Add(-duration)

	for _, s := range l {
		if s.startTime.Before(cutoff) {
			break
		}

		sum += s.amount
	}

	return sum
}

// GCBeforeDurationAgo remove sprinkles more than `duration` before now
func (l *SprinkleLog) GCBeforeDurationAgo(duration time.Duration) {
	cutoff := time.Now().Add(-duration)

	for i := len(*l); i >= 0; i-- {
		if (*l)[i].startTime.Before(cutoff) {
			l.PopBack()
		} else {
			break
		}
	}
}

// AddLog add a sprinkle amount to the log
func (l *SprinkleLog) AddLog(amount float32) {
	s := Sprinkle{
		startTime: time.Now(),
		amount:    amount,
	}

	l.Push(s)
}

// Push a Sprinkle to the front of the log
func (l *SprinkleLog) Push(s Sprinkle) {
	*l = append(*l, s)
}

// PopBack pop a Sprinkle off the back of the log
func (l *SprinkleLog) PopBack() *Sprinkle {
	back := (*l)[len(*l)-1]
	*l = (*l)[0 : len(*l)-1]

	return &back
}
