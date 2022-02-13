package main

import (
	"fmt"
	"os"
	"wordle-helper/pkg/brute"
)

func main() {
	possibilitiesMap := brute.WordleHelper("pkg/sgb-words.txt", os.Args[1], os.Args[2], os.Args[3])

	for _, v := range possibilitiesMap {
		fmt.Println(v)
	}
}
