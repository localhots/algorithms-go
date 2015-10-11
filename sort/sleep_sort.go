package sort

import (
	"time"
)

const (
	// SleepSortTimeout is the default sleep sort timeout in millisecond.
	SleepSortTimeout = 5
)

// SleepSort is an implementation of sleep sort algorithm.
func SleepSort(a []int) []int {
	return SleepSortWithTimeout(a, SleepSortTimeout)
}

// SleepSortWithTimeout is an implementation of sleep sort algorithm with user
// defined sleep timeout.
func SleepSortWithTimeout(a []int, timeout time.Duration) []int {
	start := make(chan struct{})
	results := make(chan int, len(a))

	for _, val := range a {
		go func(val int) {
			<-start
			<-time.After(time.Duration(val) * time.Millisecond)
			results <- val
		}(val)
	}
	close(start)

	for i := range a {
		a[i] = <-results
	}

	return a
}
