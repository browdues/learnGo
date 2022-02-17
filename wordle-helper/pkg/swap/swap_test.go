package swap_test

import (
	"path"
	"testing"

	"wordle-helper/pkg/remove"
	. "wordle-helper/pkg/swap"

	"github.com/stretchr/testify/assert"
)

var fpath = path.Clean("../sgb-words.txt")

func TestWordleHelper(t *testing.T) {
	t.Run("FirstExample", func(t *testing.T) {
		grays := "adieupsychtr"
		yellows := "nk"
		greens := "o3"

		want := []string{"known", "knoll"}
		got := WordleHelper(grays, yellows, greens, remove.ParseCorpusSlice(fpath))

		assert.ElementsMatch(t, want, got)
	})

	t.Run("SecondExample", func(t *testing.T) {
		grays := "uicngt"
		yellows := "d"
		greens := "r2e5"

		want := []string{"drove", "erode", "drape", "drake"}
		got := WordleHelper(grays, yellows, greens, remove.ParseCorpusSlice(fpath))

		assert.ElementsMatch(t, want, got)
	})

	t.Run("UlcerExample", func(t *testing.T) {
		grays := "tidbaksp"
		yellows := "u"
		greens := "e4r5"

		want := []string{"ruler", "queer", "ulcer", "huger", "lurer", "fumer", "urger", "gluer", "curer"}
		got := WordleHelper(grays, yellows, greens, remove.ParseCorpusSlice(fpath))

		assert.ElementsMatch(t, want, got)
	})

	t.Run("UltraExample", func(t *testing.T) {
		grays := "iednpsck"
		yellows := ""
		greens := "t1r2a3"

		want := []string{"trawl"}
		got := WordleHelper(grays, yellows, greens, remove.ParseCorpusSlice(fpath))

		assert.ElementsMatch(t, want, got)
	})

	t.Run("RobinExample", func(t *testing.T) {
		grays := "tedpsmay"
		yellows := "in"
		greens := "r1"

		want := []string{"rhino", "ruing", "robin", "runic"}
		got := WordleHelper(grays, yellows, greens, remove.ParseCorpusSlice(fpath))

		assert.ElementsMatch(t, want, got)
	})
}
