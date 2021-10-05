package algorithms

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLeakyBucket_Allow(t *testing.T) {
	go test(t)
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Time Out!")
		return
	}

}
func test(t *testing.T) {
	r := leakyBucket(2)
	for range time.Tick(time.Millisecond * 600) {
		al := r.Allow()
		assert.True(t, al)
	}
}
func TestTokenBucket_Allow(t *testing.T) {
	go testToken(t)
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Time Out!")
		return
	}

}
func testToken(t *testing.T) {
	r := tokenBucket(time.Second*5, 2)
	for range time.Tick(time.Second) {
		al := r.Allow()
		assert.True(t, al)
	}
}

func TestFixedWindow_Allow(t *testing.T) {
	go testFixedWindow(t)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Time Out!")
		return
	}
}
func testFixedWindow(t *testing.T) {
	r := fixedWindow(2)
	for range time.Tick(time.Millisecond * 300) {
		al := r.Allow()
		assert.True(t, al)
	}
}
func TestSlidingWindowLog_Allow(t *testing.T) {
	go testSlidingWindowLog(t)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Time Out!")
		return
	}
}
func testSlidingWindowLog(t *testing.T) {
	r := slidingWindowLog(2)
	for range time.Tick(time.Millisecond * 600) {
		al := r.Allow()
		assert.True(t, al)
	}
}
