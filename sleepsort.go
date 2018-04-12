package sleepsort

import "time"

// SleepSort sorts integers in ascending order
func SleepSort(is []int) []int {
	c := make(chan int, len(is))
	for i := range is {
		// goroutine aware immutable storage
		val := is[i]
		time.AfterFunc(time.Duration(val)*time.Second, func() {
			c <- val
		})
	}
	var os []int
	for range is {
		os = append(os, <-c)
	}
	return os
}
