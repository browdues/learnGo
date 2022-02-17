package remove

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WordleHelper(grays, yellows, greens string, corpus map[int]string) map[int]string {
	for wordIndex := range corpus {
		if checkYellows(yellows, corpus, wordIndex) {
			continue
		}

		if checkGrays(grays, corpus, wordIndex) {
			continue
		}

		if checkGreens(greens, corpus, wordIndex) {
			continue
		}
	}
	return corpus
}

func ParseCorpusSlice(fpath string) []string {
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var corpus []string

	for scanner.Scan() {
		corpus = append(corpus, strings.Split(scanner.Text(), " ")[0])
	}
	return corpus
}

func ParseCorpusMap(corpusSlice []string) map[int]string {
	corpusMap := make(map[int]string, len(corpusSlice))
	for i, word := range corpusSlice {
		corpusMap[i] = word
	}
	return corpusMap
}

func checkGrays(grays string, corpus map[int]string, wordIndex int) (deleted bool) {
	for _, grayLet := range grays {
		if strings.ContainsRune(corpus[wordIndex], grayLet) {
			delete(corpus, wordIndex)
			return true
		}
	}
	return false
}

func checkYellows(yellows string, corpus map[int]string, wordIndex int) (deleted bool) {
	for _, yellowLet := range yellows {
		if !strings.ContainsRune(corpus[wordIndex], yellowLet) {
			delete(corpus, wordIndex)
			return true
		}
	}
	return false
}

func checkGreens(greens string, corpus map[int]string, wordIndex int) (deleted bool) {
	for i := 0; i < len(greens); i += 2 {
		let := greens[i : i+1]
		pos, _ := strconv.Atoi(greens[i+1 : i+2])
		if corpus[wordIndex][pos-1:pos] != let {
			delete(corpus, wordIndex)
			return true
		}
	}
	return false
}
