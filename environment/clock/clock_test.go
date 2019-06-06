package clock

import (
	"math"
	"testing"
	"time"
)

func TestFakeClockNow(t *testing.T) {
	fakeTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	clock := FakeClock{}
	currentClock = &clock

	diff := math.Abs(float64(Now().Sub(time.Now())))
	if diff > float64(time.Second) {
		t.Errorf("nows should be equal")
	}

	clock.NowClosure = func() time.Time {
		return fakeTime
	}

	if !fakeTime.Equal(clock.Now()) {
		t.Errorf("times should be equal")
	}
}

func TestFakeClockAfterChan(t *testing.T) {
	clock := CreateFakeClock()
	currentClock = clock

	go func() {
		timeOut := time.After(time.Second)
		select {
		case <-timeOut:
			t.Error("timed out waiting in go routine")
			return
		case <-After(time.Second):
			return
		}
	}()

	timeOut := time.After(time.Second)
	select {
	case clock.AfterChan <- time.Now():
		break
	case <-timeOut:
		t.Errorf("timed out waiting to send to channel")
		break
	}
}

func TestFakeClockAfterClosure(t *testing.T) {
	clock := CreateFakeClock()
	otherChan := make(chan time.Time)

	clock.AfterClosure = func(duration time.Duration) <-chan time.Time {
		return otherChan
	}

	otherChanReturned := clock.After(time.Second)
	if otherChanReturned != otherChan {
		t.Error("closure should have returned otherChan")
	}

	if otherChanReturned == clock.AfterChan {
		t.Error("channel returned by closure should be unique")
	}
}

func TestSystemClock(t *testing.T) {
	clock := SystemClock{}

	diff := math.Abs(float64(clock.Now().Sub(time.Now())))
	if diff > float64(time.Second) {
		t.Errorf("system clock should return time.Now")
	}

	afterChan := clock.After(time.Second)

	if afterChan == nil {
		t.Errorf("chan should not be nil")
	}
}
