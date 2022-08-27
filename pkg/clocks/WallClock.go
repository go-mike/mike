package clocks

import "time"

// WallClock is a Clock that always returns the system time.
type WallClock struct{}

func NewWallClock() *WallClock {
	return &WallClock{}
}

var _ Clock = &WallClock{}

// Now implements Clock.Now for WallClock, returning the system time.
func (*WallClock) Now() time.Time {
	return time.Now()
}
