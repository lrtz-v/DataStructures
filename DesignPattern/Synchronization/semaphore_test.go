package synchronizationdesignpattern

import (
	"testing"
	"time"
)

func TestSemaphoreWithTimeout(t *testing.T) {
	tickets, timeout := 1, 3*time.Second
	s := NewSemaphore(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	if err := s.Release(); err != nil {
		panic(err)
	}
}

func TestSemaphoreWithoutTimeout(t *testing.T) {
	tickets, timeout := 1, 0*time.Second
	s := NewSemaphore(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	if err := s.Release(); err != nil {
		panic(err)
	}
}
