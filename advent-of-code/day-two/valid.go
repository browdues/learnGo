package valid

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// ReadLines returns the lines from a text file as an array of strings
func ReadLines(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}

// Password data struct
type Password struct {
	min      int
	max      int
	letter   string
	password string
}

// ParsePassword parses password data
func ParsePassword(lines []string) ([]Password, error) {
	passwordList := make([]Password, len(lines))

	re := regexp.MustCompile(`(\d+)-(\d+)\s([a-zA-Z]{1}):\s(\w+)$`)

	for i, line := range lines {
		if line == "" {
			continue
		}
		m := re.FindStringSubmatch(line)
		if len(m) != 5 {
			str := fmt.Sprintf("Expected rules and a password, but something went wrong %q", line)
			return nil, errors.New(str)
		}
		min, _ := strconv.Atoi(m[1])
		max, _ := strconv.Atoi(m[2])

		passwordList[i] = Password{min, max, m[3], m[4]}
	}
	return passwordList, nil
}

// ValidatePassword returns the number of valid passwords that are passed to the function
func ValidatePassword(passwords []Password) int {
	validPasswords := 0
	for _, pass := range passwords {
		nLetters := 0
		for _, l := range pass.password {
			l := string(l)
			if l == pass.letter {
				nLetters++
			}
		}
		if nLetters >= pass.min && nLetters <= pass.max {
			validPasswords++
		}
	}
	return validPasswords
}

// ValidatePasswordPart2 returns the number of valid passwords that are
// passed to the function given updated rules
func ValidatePasswordPart2(passwords []Password) int {
	nValidPasswords := 0
	for _, pass := range passwords {
		letter1 := string(pass.password[pass.min-1])
		letter2 := string(pass.password[pass.max-1])

		eitherMatch := letter1 == pass.letter || letter2 == pass.letter
		bothAreDifferent := letter1 != letter2

		if eitherMatch && bothAreDifferent {
			nValidPasswords++
		}
	}
	return nValidPasswords
}
