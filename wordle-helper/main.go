package main

import (
	"fmt"
	"os"
	"wordle-helper/pkg/optimized"
	"wordle-helper/pkg/remove"
)

func main() {
	corpus := remove.ParseCorpusMap(
		optimized.ReduceCorpusWithGrays(protectArg(1),
			optimized.ParseCorpusAlphabetBins("pkg/sgb-words.txt")),
	)
	possibilities := remove.WordleHelper(
		protectArg(1),
		protectArg(2),
		protectArg(3),
		corpus,
	)

	for _, v := range possibilities {
		fmt.Println(v)
	}
}

var colorMap = map[int]string{
	1: "gray",
	2: "yellow",
	3: "green",
}

func protectArg(argIndex int) string {
	var arg string
	defer func() {
		if recover() != nil {
			fmt.Printf("WARN: No %s input. Using empty string.\n", colorMap[argIndex])
		}
	}()

	arg = os.Args[argIndex]
	return arg
}
