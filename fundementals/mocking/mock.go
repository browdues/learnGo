package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

// Sleeper sleeps a lot
type Sleeper interface {
	Sleep()
}

// SpySleeper tracks sleep
type SpySleeper struct {
	Calls int
}

// Sleep tracks sleep
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper is the def sleeper
type DefaultSleeper struct{}

// Sleep sleeps for one second
func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper struct
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep sleeps
func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime spies sometimes
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep sleeps
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept += duration
}

const (
	write = "write"
	sleep = "sleep"
)

// CountDownOperationsSpy spys on counting down
type CountDownOperationsSpy struct {
	Calls []string
}

// Sleep spys on sleep calls
func (s *CountDownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write spies on write calls
func (s *CountDownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// Countdown counts down
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
