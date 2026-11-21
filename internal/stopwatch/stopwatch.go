package stopwatch

import (
	"time"
)

func StopWatch(duration time.Duration, done <-chan struct{}, timeout chan<- struct{}) {
	timer := time.NewTimer(duration)

	defer timer.Stop()

	select {
	case <-timer.C:
		close(timeout)

	case <-done:
		return
	}
}
