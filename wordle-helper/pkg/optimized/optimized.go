package optimized

import (
	"sort"
	"wordle-helper/pkg/remove"
)

// ReduceCorpusWithGrays seeks to reduce the corpus upfront by assuming
// a preloaded map of letter -> []sortedWord exists
func ReduceCorpusWithGrays(grays string, alphaBins map[byte][]string) []string {
	grays = sortString(grays)

	finalSlice := []string{}
	cnt := 0

	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		cnt++
		for i, grayLet := range grays {
			if grayLet != letter {
				finalSlice = append(finalSlice, alphaBins[byte(letter)]...)
				break
			}
			if grayLet == letter {
				if len(grays) > i {
					grays = string(grays[i+1:])
				} else {
					grays = ""
				}
				break
			}
		}
	}
	return finalSlice
}

// ParseCorpusAlphabetBins maps words by starting letter
func ParseCorpusAlphabetBins(fpath string) map[byte][]string {
	alphaBins := make(map[byte][]string, 26)
	corpus := remove.ParseCorpusSlice(fpath)
	sort.Slice(corpus, func(curr, next int) bool { return corpus[curr] < corpus[next] })

	currLetter := byte('a')
	for _, word := range corpus {
		if word[0] != currLetter {
			currLetter = word[0]
		}
		alphaBins[currLetter] = append(alphaBins[currLetter], word)
	}

	return alphaBins
}

// sortString sorts a string alphabetically
func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}
