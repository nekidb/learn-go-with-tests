package main

import (
		"fmt"
		"io"
		"os"
		"time"
)

type Sleeper interface {
	Sleep()
}


type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}


type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

const finalWord = "Go!"
const countdownStart = 3

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	
	fmt.Fprint(out, finalWord)
}

func main() {
	CountDown(os.Stdout, &DefaultSleeper{})
}