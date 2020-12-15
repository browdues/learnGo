package passport

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

func validFields(p string) bool {
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

func getValidPassports(passports []string) []string {
	validPassports := make([]string, 0)
	for _, passport := range passports {
		if validFields(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}
