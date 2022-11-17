package main

import (
		"testing"
		"net/http/httptest"
		"net/http"
		"time"
)

func TestRacer(t *testing.T) {
	t.Run("return url that respond faster", func(t *testing.T) {
		slowServer := makeDelayedServer(10 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		
		defer slowServer.Close()
		defer fastServer.Close()
		
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		
		want := fastURL
		got, _ := Racer(slowURL, fastURL)
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if server doesn't respond within 10 sec", func(t *testing.T) {
	     serverA := makeDelayedServer(11 * time.Second)
	     serverB := makeDelayedServer(12 * time.Second)

	     defer serverA.Close()
	     defer serverB.Close()

	     _, err := Racer(serverA.URL, serverB.URL)

	     if err == nil {
	     	t.Error("exected an error but didn't get one")
	     }
     	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
