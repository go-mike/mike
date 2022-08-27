package clocks

import "time"

// FixedClock is a Clock that always returns the same time.
type FixedClock struct{
	time time.Time
}

func NewFixedClock(time time.Time) *FixedClock {
	return &FixedClock{
		time: time,
	}
}

var _ Clock = &FixedClock{}

// Now implements Clock.Now for FixedClock, returning the time from constructor.
func (clock *FixedClock) Now() time.Time {
	return clock.time
}
