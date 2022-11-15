package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying 'Thank you' to someone", func(t *testing.T) {
		got := Hello("Dummy", "")
		want := "Hello, Dummy"
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
	t.Run("What if input is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, bro"
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
	t.Run("In Russian", func(t *testing.T) {
		got := Hello("Dummy", "russian")
		want := "Привет, Dummy"
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
	t.Run("in french", func(t *testing.T) {
	got := Hello("Dummy", "french")
	want := "Bonjour, Dummy"
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	})
}

