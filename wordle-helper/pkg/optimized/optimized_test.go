package optimized_test

import (
	"path"
	"testing"

	. "wordle-helper/pkg/optimized"
	"wordle-helper/pkg/remove"
	"wordle-helper/pkg/swap"

	"github.com/stretchr/testify/assert"
)

var fpath = path.Clean("../sgb-words.txt")

func TestImprove(t *testing.T) {
	alphabetBins := ParseCorpusAlphabetBins(fpath)

	t.Run("FirstExample", func(t *testing.T) {
		grays := "adieupsychtr"
		yellows := "nk"
		greens := "o3"

		want := []string{"known", "knoll"}
		got := swap.WordleHelper(
			grays,
			yellows,
			greens,
			ReduceCorpusWithGrays(grays, alphabetBins),
		)

		assert.ElementsMatch(t, want, got)
	})
}

func BenchmarkWordleHelper(b *testing.B) {
	grays := "tidbaksp"
	yellows := "u"
	greens := "e4r5"

	b.Run("RemoveWordleHelper", func(b *testing.B) {
		corpus := remove.ParseCorpusMap(remove.ParseCorpusSlice(fpath))

		for i := 0; i < b.N; i++ {
			remove.WordleHelper(
				grays,
				yellows,
				greens,
				corpus,
			)
		}
	})

	b.Run("OptimizedRemoveWordleHelper", func(b *testing.B) {
		optimizedCorpus := remove.ParseCorpusMap(
			ReduceCorpusWithGrays(grays, ParseCorpusAlphabetBins(fpath)),
		)

		for i := 0; i < b.N; i++ {
			remove.WordleHelper(
				grays,
				yellows,
				greens,
				optimizedCorpus,
			)
		}
	})

	b.Run("SwapWordleHelper", func(b *testing.B) {
		corpus := remove.ParseCorpusSlice(fpath)

		for i := 0; i < b.N; i++ {
			swap.WordleHelper(
				grays,
				yellows,
				greens,
				corpus,
			)
		}
	})

	b.Run("OptimizedSwapWordleHelper", func(b *testing.B) {
		optimizedCorpus := ReduceCorpusWithGrays(grays, ParseCorpusAlphabetBins(fpath))

		for i := 0; i < b.N; i++ {
			swap.WordleHelper(
				grays,
				yellows,
				greens,
				optimizedCorpus,
			)
		}
	})
}
