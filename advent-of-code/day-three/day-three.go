package daythree

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func filter(arr []string, f func(s string) bool) []string {
	output := make([]string, 0)

	for _, item := range arr {
		if f(item) {
			output = append(output, item)
		}
	}

	return output
}

func readMap(fpath string) []string {
	file, err := ioutil.ReadFile(fpath)

	if err != nil {
		fmt.Printf("Some err occured reading map:\n%q", err)
		os.Exit(1)
	}

	lines := strings.Split(string(file), "\n")
	filteredLines := filter(lines, func(s string) bool {
		return s != ""
	})

	return filteredLines
}

type loc struct {
	row int
	col int
}

func (t *loc) travel(lines []string, rInc int, cInc int) int {

	width := len(lines[0])
	nTrees := 0

	for i := cInc; i < len(lines); i += cInc {
		t.row += rInc
		t.row = t.row % width
		if lines[i][t.row] == byte('#') {
			nTrees++
		}
	}
	return nTrees
}
