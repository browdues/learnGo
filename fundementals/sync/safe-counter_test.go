package safecount

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing 3x results in value of 3", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		got := counter.Value()

		assert.Equal(t, want, got)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assert.Equal(t, wantedCount, counter.Value())
	})
}
