package remove_test

import (
	"path"
	"testing"

	. "wordle-helper/pkg/remove"

	"github.com/stretchr/testify/assert"
)

var fpath = path.Clean("../sgb-words.txt")

func TestWordleHelper(t *testing.T) {
	t.Run("FirstExample", func(t *testing.T) {
		grays := "adieupsychtr"
		yellows := "nk"
		greens := "o3"

		want := map[int]string{72: "known", 3350: "knoll"}
		got := WordleHelper(grays, yellows, greens, ParseCorpusMap(ParseCorpusSlice(fpath)))

		assert.EqualValues(t, want, got)
	})

	t.Run("SecondExample", func(t *testing.T) {
		grays := "uicngt"
		yellows := "d"
		greens := "r2e5"

		want := map[int]string{312: "drove", 2496: "erode", 2785: "drape", 3309: "drake"}
		got := WordleHelper(grays, yellows, greens, ParseCorpusMap(ParseCorpusSlice(fpath)))

		assert.EqualValues(t, want, got)
	})

	t.Run("UlcerExample", func(t *testing.T) {
		grays := "tidbaksp"
		yellows := "u"
		greens := "e4r5"

		want := map[int]string{411: "ruler", 775: "queer", 2137: "ulcer", 3161: "huger", 5201: "lurer", 5296: "fumer", 5456: "urger", 5642: "gluer", 5731: "curer"}
		got := WordleHelper(grays, yellows, greens, ParseCorpusMap(ParseCorpusSlice(fpath)))

		assert.EqualValues(t, want, got)
	})

	t.Run("UltraExample", func(t *testing.T) {
		grays := "iednpsck"
		yellows := ""
		greens := "t1r2a3"

		want := map[int]string{2028: "trawl"}
		got := WordleHelper(grays, yellows, greens, ParseCorpusMap(ParseCorpusSlice(fpath)))

		assert.EqualValues(t, want, got)
	})

	t.Run("RobinExample", func(t *testing.T) {
		grays := "tedpsmay"
		yellows := "in"
		greens := "r1"

		want := map[int]string{3341: "rhino", 5473: "ruing", 1091: "robin", 3551: "runic"}
		got := WordleHelper(grays, yellows, greens, ParseCorpusMap(ParseCorpusSlice(fpath)))

		assert.EqualValues(t, want, got)
	})
}
