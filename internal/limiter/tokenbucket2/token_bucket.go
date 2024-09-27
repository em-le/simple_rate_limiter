package tokenbucket2

import (
	"sync"
	"time"
)

func NewLimiter(capacity int64, period time.Duration) *tokenBucketLimiter {
	fillInterval := period / time.Duration(capacity)
	bucket := &tokenBucketLimiter{
		buckets:      make(chan bool, capacity),
		fillInterval: fillInterval,
	}
	go bucket.fillToken()
	return bucket
}

type tokenBucketLimiter struct {
	buckets      chan bool
	fillInterval time.Duration
	mu           sync.Mutex
}

func (b *tokenBucketLimiter) Allow() (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	wait := make(chan bool)

	go func() {
		time.Sleep(b.fillInterval / 2)
		wait <- false
	}()

	for {
		select {
		case <-b.buckets:
			return true, nil
		case <-wait:
			return false, nil
		}
	}
}

func (b *tokenBucketLimiter) fillToken() {
	ticker := time.NewTicker(b.fillInterval)
	for {
		select {
		case <-ticker.C:
			select {
			case b.buckets <- true:
			default:
			}
		}
	}
}
