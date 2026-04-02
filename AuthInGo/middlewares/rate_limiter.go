package middlewares

import (
	"net/http"
	"time"
	"golang.org/x/time/rate"
)



// RateLimiter is a middleware that limits the number of requests a client can make in a given time period.

// It uses the golang.org/x/time/rate package to implement a token bucket algorithm for rate limiting.

// The limiter variable is a global variable that is initialized with a rate of 1 request per minute and a burst size of 5.

//burst size is the maximum number of requests that can be made in a short period of time, while the rate is the average number of requests that can be made per second.

// The middleware checks if the client has exceeded the rate limit by calling the Allow() method on the limiter variable. If the client has exceeded the rate limit, it returns a 429 Too Many Requests response. Otherwise, it allows the request to proceed to the next handler.

//limiter.Allow() returns true if a token is available and false if the rate limit has been exceeded. The limiter will automatically refill tokens at the specified rate, allowing clients to make requests again after a certain amount of time has passed.

var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 5)

func RateLimiteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}