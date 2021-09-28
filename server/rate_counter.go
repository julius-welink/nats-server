// Copyright 2012-2021 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"sync"
	"time"
)

type rateCounter struct {
	limit    int64
	count    int64
	end      time.Time
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
	defer r.mu.Unlock()

	if now.After(r.end) {
		r.count = 0
		r.end = now.Add(r.interval)
	} else {
		r.count++
	}

	return r.count < r.limit
}
