package main

import (
	"io"
	"log"
	"net/http"

	"github.com/em-le/rate_limiter/internal/middleware"
)

func main() {
	http.Handle("/health_check", middleware.RateLimiter(HeathCheck))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HeathCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK!")
}
