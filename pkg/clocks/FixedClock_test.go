package clocks

import (
	"testing"
	"time"
)

func TestNewFixedClock(t *testing.T) {
	// Given
	fixedTime := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	// Given
	clock := NewFixedClock(fixedTime)
	// When
	actualTime := clock.Now()
	// Then
	if actualTime != fixedTime {
		t.Errorf("actualTime = %v, want %v", actualTime, fixedTime)
	}
}
