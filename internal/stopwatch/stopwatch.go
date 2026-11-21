package stopwatch

import (
	"time"
)

func StopWatch(minutes int, done <-chan struct{}, timeout chan<- struct{}) {
	duration := time.Duration(minutes * int(time.Minute))

	timer := time.NewTimer(duration)

	defer timer.Stop()

	select {
	case <-timer.C:
		close(timeout)

	case <-done:
		return
	}
}
