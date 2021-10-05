package algorithms

import (
	"fmt"
	"time"
)

type LeakyBucket struct {
	nextAllowedTime       int64
	requestIntervalMillis int64
}

func leakyBucket(maxRequestPerSec int64) *LeakyBucket {
	return &LeakyBucket{
		nextAllowedTime:       time.Now().UnixMilli(),
		requestIntervalMillis: 1000 / maxRequestPerSec,
	}
}

func (b *LeakyBucket) Allow() bool {
	curTime := time.Now().UnixMilli()
	fmt.Println(curTime - b.nextAllowedTime)
	if curTime >= b.nextAllowedTime {
		b.nextAllowedTime = curTime + b.requestIntervalMillis
		return true
	}
	return false
}
