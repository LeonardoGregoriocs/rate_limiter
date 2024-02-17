package limiter

import "net/http"

type RateLimiter struct{}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your rate limiting logic here
		next.ServeHTTP(w, r)
	})
}
