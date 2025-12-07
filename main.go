package consumerproject
package main

import (














}	}		fmt.Println("Request denied")	} else {		fmt.Println("Request allowed")	if rl.Allow() {	fmt.Println("Rate limiter created")	rl := ratelimiter.NewRateLimiter(10, time.Minute)func main() {)	ratelimiter "github.com/DakshithaS/test-policy-hub/rate-limiter/v1.0.0"	"time"	"fmt"