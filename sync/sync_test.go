package main

import (
	"testing"
	"sync"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter three times and get value" , func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		
		assertCounter(t, counter, 3)
	})

	t.Run("concurrent incrementing", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i< wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		
		assertCounter(t, counter, wantedCount)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
	
