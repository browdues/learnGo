package brute_test

import (
	"path"
	"testing"

	. "wordle-helper/pkg/brute"

	"github.com/stretchr/testify/assert"
)

var fpath = path.Clean("../sgb-words.txt")

func TestWordleHelper(t *testing.T) {
	t.Run("FirstExample", func(t *testing.T) {
		grays := "adieupsychtr"
		yellows := "nk"
		greens := "o3"

		want := map[int]string{72: "known", 3350: "knoll"}
		got := WordleHelper(fpath, grays, yellows, greens)

		assert.Equal(t, want, got)
	})

	t.Run("SecondExample", func(t *testing.T) {
		grays := "uicngt"
		yellows := "d"
		greens := "r2e5"

		want := map[int]string{312: "drove", 2496: "erode", 2785: "drape", 3309: "drake"}
		got := WordleHelper(fpath, grays, yellows, greens)

		assert.Equal(t, want, got)
	})
}
