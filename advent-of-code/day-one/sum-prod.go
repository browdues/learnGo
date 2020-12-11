package sumprod

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Find2020Sum takes a list of ints and checks if any two sum to 2020
func Find2020Sum(numbers []int64) ([]int64, error) {
	attempted := make([]int64, 0)
	for _, num := range numbers {
		if num > 2020 {
			continue
		}

		for _, attempt := range attempted {
			if num+attempt == 2020 {
				return []int64{num, attempt}, nil
			}
		}
		attempted = append(attempted, num)
	}
	return nil, errors.New("Couldn't find two numbers that sum to 2020")
}

// Find2020TripleSum takes a list of ints and checks if any three sum to 2020
func Find2020TripleSum(numbers []int64) ([]int64, error) {
	firstAttempt := make([]int64, 0)
	secondAttempt := make([]int64, 0)

	for _, num1 := range numbers {
		if num1 > 2020 {
			continue
		}

		for _, num2 := range firstAttempt {
			for _, num3 := range secondAttempt {
				if num1+num2+num3 == 2020 {
					return []int64{num1, num2, num3}, nil
				}
			}
			secondAttempt = append(secondAttempt, num2)
		}

		firstAttempt = append(firstAttempt, num1)
	}
	return nil, errors.New("Couldn't find two numbers that sum to 2020")
}

// LinesToArray accepts a path to a text file and returns an array where each element is a line of the txt file.
func LinesToArray(fpath string) ([]int64, error) {

	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []int64

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")[0]
		intLine, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			return nil, err
		}
		lines = append(lines, intLine)
	}

	return lines, nil
}
