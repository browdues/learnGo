package passport

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getPassports(fpath string) []string {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reEmpty := regexp.MustCompile(`^\s*$`)
	passportStr := ""
	passports := make([]string, 0)

	fileLines := strings.Split(string(file), "\n")
	for _, line := range fileLines {
		if reEmpty.MatchString(line) {
			passports = append(passports, strings.TrimSpace(passportStr))
			passportStr = ""
			continue
		}
		passportStr += line + " "
	}

	return passports
}

func validFieldsPart1(p string) bool {
	isValid := false
	reOptionalField := regexp.MustCompile(`(cid:\S+)\s`)

	fields := strings.Split(p, " ")
	validNumber := len(fields) > 6
	containsOptional := reOptionalField.MatchString(p)

	if len(fields) == 8 {
		isValid = true
	} else if validNumber && !containsOptional {
		isValid = true
	}
	return isValid
}

func getValidPassportsPart1(passports []string) []string {
	validPassports := make([]string, 0)
	for _, passport := range passports {
		if validFieldsPart1(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

func validYear(p string) bool {
	reBirth := regexp.MustCompile(`(byr:)(\S+)\s`)
	reIssue := regexp.MustCompile(`(iyr:)(\S+)\s`)
	reExpiration := regexp.MustCompile(`(eyr:)(\S+)\s`)
	isValid := false

	yearDetectors := []struct {
		re     *regexp.Regexp
		bounds []int
	}{
		{reBirth, []int{1920, 2002}},
		{reIssue, []int{2010, 2020}},
		{reExpiration, []int{2020, 2030}},
	}

	for _, yd := range yearDetectors {
		if yd.re.MatchString(string(p)) {
			keyVal := yd.re.FindStringSubmatch(p)
			year, _ := strconv.Atoi(keyVal[2])
			yearInBounds := year >= yd.bounds[0] && year <= yd.bounds[1]
			if yearInBounds {
				isValid = true
			} else {
				isValid = false
				break
			}
		}
	}

	return isValid
}

// validYear
// validHeight
// validHairColor
// validEyeColor
// validPassportId
