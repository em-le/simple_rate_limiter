package leakybucket2

import (
	"errors"
	"sync"
	"time"
)

func NewLimiter(capacity int64, period time.Duration) *leakyBucketLimiter {
	leakRate := period / time.Duration(capacity)
	limiter := &leakyBucketLimiter{
		capacity: capacity,
		leakRate: leakRate,
	}
	defer limiter.ticker()

	return limiter
}

type leakyBucketLimiter struct {
	last time.Time
	mu   sync.Mutex

	// maximum number of tokens in the bucket
	capacity int64

	// number of tokens added
	length int64

	// one process per leakRate duration => free a slot in length
	leakRate time.Duration
}

func (l *leakyBucketLimiter) Allow() (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.length > l.capacity {
		return false, errors.New("Limited Bucket")
	}
	l.length++
	return true, nil
}

func (l *leakyBucketLimiter) ticker() <-chan bool {
	ticker := time.NewTicker(l.leakRate)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				l.freeSlot()
			}
		}
	}()

	return done
}

func (l *leakyBucketLimiter) freeSlot() {
	if l.length == 0 {
		return
	}
	l.length -= 1
}
