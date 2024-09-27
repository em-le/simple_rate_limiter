package middleware

import (
	"io"
	"net/http"

	"github.com/em-le/rate_limiter/internal/limiter"
)

type NextFunc func(http.ResponseWriter, *http.Request)

func RateLimiter(next NextFunc) http.Handler {
	rateLimiter := limiter.NewLimiter()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok, _ := rateLimiter.Allow(); !ok {
			w.WriteHeader(http.StatusTooManyRequests)
			io.WriteString(w, "Too Many Requests!")
			return
		}
		next(w, r)
	})
}
