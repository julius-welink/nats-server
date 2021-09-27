package server

import (
	"sync"
	"time"
)

type rateCounter struct {
	limit    int64
	count    int64
	start    time.Time
	interval time.Duration
	mu       sync.Mutex
}

func newRateCounter(limit int64) *rateCounter {
	return &rateCounter{
		limit:    limit,
		interval: time.Second,
	}
}

func (r *rateCounter) allow() bool {
	now := time.Now()

	r.mu.Lock()

	if now.After(r.start.Add(r.interval)) {
		r.count = 0
		r.start = now
	} else {
		r.count++
	}

	r.mu.Unlock()

	return r.count < r.limit
}
