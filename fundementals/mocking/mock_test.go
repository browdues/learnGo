package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {

	t.Run("counts down", func(t *testing.T) {
		b := bytes.Buffer{}
		spySleeper := SpySleeper{}

		Countdown(&b, &spySleeper)

		got := b.String()
		want := `3
2
1
Go!`

		assert.Equal(t, want, got)
		assert.Equal(t, 3, spySleeper.Calls)
	})

	t.Run("counts down correctly", func(t *testing.T) {
		// b := by/tes.Buffer{}
		spy := CountDownOperationsSpy{}

		Countdown(&spy, &spy)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		got := spy.Calls

		assert.Equal(t, want, got)
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	assert.Equal(t, sleepTime, spyTime.durationSlept)
}
