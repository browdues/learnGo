package sumprod

import (
	"bufio"
	"os"
)

// LinesToArray accepts a path to a text file and returns an array where each element is a line of the txt file.
func LinesToArray(fpath string) ([]string, error) {

	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
