package stopwatch

import (
	"testing"
	"time"
)

func TestStopWatchTimeout(t *testing.T) {
	done := make(chan struct{})
	timeout := make(chan struct{})

	go StopWatch(10*time.Millisecond, done, timeout)

	select {
	case <-timeout:
		// O cronômetro sinalizou o timeout corretamente.

	case <-time.After(1 * time.Second):
		t.Fatal("o cronômetro não sinalizou o timeout")
	}
}

func TestStopWatchCancel(t *testing.T) {
	done := make(chan struct{})
	timeout := make(chan struct{})
	finished := make(chan struct{})

	go func() {
		defer close(finished)
		StopWatch(1*time.Second, done, timeout)
	}()

	close(done)

	select {
	case <-finished:
		// A goroutine terminou após o cancelamento.

	case <-time.After(1 * time.Second):
		t.Fatal("o cronômetro não encerrou após o cancelamento")
	}

	select {
	case <-timeout:
		t.Fatal("o cronômetro sinalizou timeout após o cancelamento")

	default:
		// Nenhum timeout foi sinalizado.
	}
}
