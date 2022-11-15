package main

import (
		"testing"
		"bytes"
		"reflect"
)

func TestCountDown(t *testing.T) {

	t.Run("prints 3,2,1,Go!", func (t *testing.T) {
		buffer := &bytes.Buffer{}
		
		CountDown(buffer, &SpyCountdownOperations{})
		
		got := buffer.String()
		want := `3
2
1
Go!`
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	
	t.Run("sleep before every print", func (t *testing.T) {
		spyWriteSleep := &SpyCountdownOperations{}
		
		CountDown(spyWriteSleep, spyWriteSleep)
		
		want := []string{
				write,
				sleep,
				write,
				sleep,
				write,
				sleep,
				write,
		}
		
		if !reflect.DeepEqual(want, spyWriteSleep.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spyWriteSleep.Calls)
		}
	})
}

