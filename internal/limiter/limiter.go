package limiter

import (
	"fmt"
	"os"
	"time"

	"github.com/em-le/rate_limiter/internal/limiter/buckettoken"
	"github.com/em-le/rate_limiter/internal/limiter/leakybucket"
	"github.com/em-le/rate_limiter/internal/limiter/leakybucket2"
)

const (
	capacity = 5
	period   = 5 * time.Second
)

func NewLimiter() Rate {
	fmt.Println(os.Getenv("LIMITER"))
	switch os.Getenv("LIMITER") {
	case "BUCKET_TOKEN":
		return buckettoken.NewLimiter(capacity, period)
	case "LEAKY_BUCKET_2":
		return leakybucket2.NewLimiter(capacity, period)
	default:
		return leakybucket.NewLimiter(capacity, period)
	}
}
