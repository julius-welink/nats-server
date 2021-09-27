package server

import (
	"testing"
	"time"
)

func TestRateCounter(t *testing.T) {
	counter := newRateCounter(10)

	var i int
	for i = 0; i <= 10; i++ {
		if !counter.allow() {
			break
		}
	}

	if i != 10 {
		t.Errorf("Expected i = 10, got %d", i)
	}

	time.Sleep(1100 * time.Millisecond)

	if !counter.allow() {
		t.Errorf("Expected true after current time window expired")
	}
}
