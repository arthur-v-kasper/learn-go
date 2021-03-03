package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const contDownStart = 3
const write = "write"
const sleep = "sleep"

func (s *CountDownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := contDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleep struct{}

func (d *DefaultSleep) Sleep() {
	time.Sleep(1 * time.Second)
}

type CountDownOperationsSpy struct {
	Calls []string
}

func (s *CountDownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
