package clocks

import (
	"testing"
	"time"
)

// TestNewWallClock tests the NewWallClock function.
func TestNewWallClock(t *testing.T) {
	// Given
	clock := NewWallClock()
	// Given
	expectedTime := time.Now()
	// When
	actualTime := clock.Now()
	// Then
	if actualTime.Before(expectedTime) || actualTime.After(expectedTime.Add(time.Second)) {
		t.Errorf("actualTime = %v, want %v", actualTime, expectedTime)
	}
}
