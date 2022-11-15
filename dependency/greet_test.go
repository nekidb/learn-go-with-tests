package main

import "testing"
import "bytes"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	
	Greet(&buffer, "Lolkek")
	
	got := buffer.String()
	want := "Hello, Lolkek"
	
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}