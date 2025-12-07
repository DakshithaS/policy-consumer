package main

import (
	"fmt"
	"time"

	ratelimiter "github.com/DakshithaS/test-policy-hub/rate-limiter/v1.0.0"
)

func main() {
	rl := ratelimiter.NewRateLimiter(10, time.Minute)
	fmt.Println("Rate limiter created")
	if rl.Allow() {
		fmt.Println("Request allowed")
	} else {
		fmt.Println("Request denied")
	}
}