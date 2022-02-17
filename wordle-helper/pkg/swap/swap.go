package swap

import (
	"strconv"
	"strings"
)

func WordleHelper(grays, yellows, greens string, corpus []string) []string {
	startIndex := 0
	corpusSize := len(corpus) - 1

	check := func(startIndex *int, corpusSize *int) bool {
		for _, word := range corpus[*startIndex : *corpusSize+1] {
			if *corpusSize == *startIndex {
				return true
			}

			for _, grayLet := range grays {
				if strings.ContainsRune(word, grayLet) {
					corpus[*startIndex], corpus[*corpusSize] = corpus[*corpusSize], corpus[*startIndex]
					*corpusSize--
					return false
				}
			}

			for _, yellowLet := range yellows {
				if !strings.ContainsRune(word, yellowLet) {
					corpus[*startIndex], corpus[*corpusSize] = corpus[*corpusSize], corpus[*startIndex]
					*corpusSize--
					return false
				}
			}

			for i := 0; i < len(greens); i += 2 {
				let := greens[i : i+1]
				pos, _ := strconv.Atoi(greens[i+1 : i+2])
				if corpus[*startIndex][pos-1:pos] != let {
					corpus[*startIndex], corpus[*corpusSize] = corpus[*corpusSize], corpus[*startIndex]
					*corpusSize--
					return false
				}
			}

			*startIndex++
		}
		return true
	}

	var finished bool
	for !finished {
		finished = check(&startIndex, &corpusSize)
	}

	return corpus[:corpusSize]
}
