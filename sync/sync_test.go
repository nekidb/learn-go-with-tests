package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment counter three times and get value" , func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		
		assertCounter(t, counter, 3)
	})
}

func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
	
