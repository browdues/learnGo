package swap

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WordleHelper(fpath, grays, yellows, greens string) []string {
	corpus, err := parseCorpusSlice(fpath)
	if err != nil {
		fmt.Printf("err reading corpus %v", err)
		os.Exit(0)
	}

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

func parseCorpusSlice(fpath string) ([]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var corpus []string

	for scanner.Scan() {
		corpus = append(corpus, strings.Split(scanner.Text(), " ")[0])
	}
	return corpus, nil
}
