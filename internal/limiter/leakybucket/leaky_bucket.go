package leakybucket

import (
	"errors"
	"sync"
	"time"
)

func NewLimiter(capacity int64, period time.Duration) *leakyBucketLimiter {
	leakRate := period / time.Duration(capacity)
	return &leakyBucketLimiter{
		capacity: capacity,
		leakRate: leakRate,
	}
}

type leakyBucketLimiter struct {
	last time.Time
	mu   sync.Mutex
	// maximum number of tokens in the bucket
	capacity int64
	// one process per leakRate duration
	leakRate time.Duration
}

func (l *leakyBucketLimiter) Allow() (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	last := l.last
	now := time.Now()
	if now.Before(last) {
		last = last.Add(l.leakRate)
	} else {
		var offset time.Duration
		duration := now.Sub(last)
		if duration < l.leakRate {
			offset = l.leakRate - duration
		}
		last = now.Add(offset)
	}
	wait := last.Sub(now).Milliseconds()
	if wait/l.leakRate.Milliseconds() >= l.capacity {
		return false, errors.New("Limited Bucket")
	}
	l.last = last
	return true, nil
}
