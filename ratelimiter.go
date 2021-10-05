package main

type RateLimiterIFace interface {
	RateLimiter(int)
	Allow() bool
}
