package environment

import "time"

var currentClock Clock

func init() {
	systemClock := SystemClock{}
	SetClock(&systemClock)
}

// Now return the current time
func Now() time.Time {
	return currentClock.Now()
}

// After return a channel that will activate after the duration has passed
func After(duration time.Duration) <-chan time.Time {
	return currentClock.After(duration)
}

// SetClock set the current time provider
func SetClock(c Clock) {
	currentClock = c
}

// Clock provides mockable time interface
type Clock interface {
	Now() time.Time
	After(delay time.Duration) <-chan time.Time
}

// SystemClock provides default time behaviour
type SystemClock struct{}

// Now return the current time
func (t *SystemClock) Now() time.Time {
	return time.Now()
}

// After return a channel that will activate after the duration has passed
func (t *SystemClock) After(delay time.Duration) <-chan time.Time {
	return time.After(delay)
}

// CreateFakeClock create a FakeClock
func CreateFakeClock() *FakeClock {
	clock := FakeClock{
		AfterChan: make(chan time.Time),
	}

	return &clock
}

// FakeClock helpful when mocking the Clock interface
type FakeClock struct {
	AfterChan    chan time.Time
	AfterClosure func(time.Duration) <-chan time.Time
	NowClosure   func() time.Time
}

// Now use NowClosure if set, else use default time
func (t *FakeClock) Now() time.Time {
	if t.NowClosure == nil {
		return time.Now()
	}

	return t.NowClosure()
}

// After use AfterClosure if set, else return the AfterChan
func (t *FakeClock) After(delay time.Duration) <-chan time.Time {
	if t.AfterClosure == nil {
		return t.AfterChan
	}

	return t.AfterClosure(delay)
}
