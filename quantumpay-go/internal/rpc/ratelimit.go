package rpc

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientState
	limit   int
	window  time.Duration
}

type clientState struct {
	count     int
	resetTime time.Time
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		clients: make(map[string]*clientState),
		limit:   limit,
		window:  window,
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	state, exists := rl.clients[ip]
	if !exists || now.After(state.resetTime) {
		rl.clients[ip] = &clientState{
			count:     1,
			resetTime: now.Add(rl.window),
		}
		return true
	}

	if state.count >= rl.limit {
		return false
	}

	state.count++
	return true
}

func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func RateLimitMiddleware(limit int, window time.Duration) func(http.Handler) http.Handler {
	limiter := newRateLimiter(limit, window)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getIP(r)

			if !limiter.allow(ip) {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte(`{"error":"rate limit exceeded"}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
