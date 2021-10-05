package algorithms

import "time"

type linkedList struct {
	values []int64
}

func newList() *linkedList {
	return &linkedList{
		values: make([]int64, 0),
	}
}

func (l linkedList) poll() *int64 {
	if len(l.values) == 0 {
		return nil
	}
	h := l.values[0]
	if len(l.values) > 1 {
		l.values = l.values[1:]
	} else {
		l.values = make([]int64, 0)
	}
	return &h
}

func (l linkedList) element() *int64 {
	if len(l.values) == 0 {
		return nil
	}
	return &l.values[0]
}

func (l linkedList) isEmpty() bool {
	return len(l.values) == 0
}

type SlidingWindowLog struct {
	maxRequestPerSec int
	list             *linkedList
}

// (t - 1, t]

func (b *SlidingWindowLog) Allow() bool {
	curTime := time.Now().UnixMilli()
	t := curTime - 1000
	for !b.list.isEmpty() && *b.list.element() <= t {
		b.list.poll()
	}
	b.list.values = append(b.list.values, curTime)
	return len(b.list.values) <= b.maxRequestPerSec
}

func slidingWindowLog(maxRequestPerSec int) *SlidingWindowLog {
	return &SlidingWindowLog{
		maxRequestPerSec: maxRequestPerSec,
		list:             newList(),
	}
}
