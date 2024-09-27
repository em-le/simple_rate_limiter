package limiter

import (
	"os"
	"time"

	"github.com/em-le/rate_limiter/internal/limiter/leakybucket"
	"github.com/em-le/rate_limiter/internal/limiter/tokenbucket"
	"github.com/em-le/rate_limiter/internal/limiter/tokenbucket2"
)

const (
	capacity = 5
	period   = 5 * time.Second
)

func NewLimiter() Rate {
	switch os.Getenv("LIMITER") {
	case "TOKEN_BUCKET_2":
		return tokenbucket2.NewLimiter(capacity, period)
	case "TOKEN_BUCKET":
		return tokenbucket.NewLimiter(capacity, period)
	default:
		return leakybucket.NewLimiter(capacity, period)
	}
}
