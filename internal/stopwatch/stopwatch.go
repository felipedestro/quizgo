package stopwatch

import (
	"fmt"
	"os"
	"time"
)

func StopWatch(minutes int, done <-chan struct{}) {
	duration := time.Duration(minutes * int(time.Minute))

	timer := time.NewTimer(duration)

	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("Tempo esgotado! Tente novamente.")
		os.Exit(0)

	case <-done:
		return
	}
}
