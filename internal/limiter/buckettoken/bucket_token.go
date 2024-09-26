package buckettoken

import (
	"time"
)

func NewLimiter(capacity int64, period time.Duration) *bucketTokenLimiter {
	return &bucketTokenLimiter{}
}

type bucketTokenLimiter struct {
}

func (b *bucketTokenLimiter) Allow() (bool, error) {
	return true, nil
}
