package passport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPassports(t *testing.T) {
	passports := getPassports("./day-four-input.txt")

	assert.Equal(t, 286, len(passports))
}

func TestValid(t *testing.T) {
	validTests := []struct {
		passport string
		isValid  bool
	}{
		{"ecl:hzl hgt:66in byr:2000 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"iyr:2012 pid:749770535 byr:1969 cid:148 hcl:#733820 hgt:180cm eyr:2021", false},
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
	}

	for _, test := range validTests {
		assert.Equal(t, test.isValid, validFields(test.passport))
	}
}

func TestGetValidPassports(t *testing.T) {
	passports := getPassports("./day-four-input.txt")
	validPassports := getValidPassports(passports)

	assert.Equal(t, 242, len(validPassports))
}
