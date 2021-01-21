package passport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPassports(t *testing.T) {
	passports := getPassports("./day-four-input.txt")

	assert.Equal(t, 286, len(passports))
}

func TestValidPart1(t *testing.T) {
	validTests := []struct {
		passport string
		isValid  bool
	}{
		{"ecl:hzl hgt:66in byr:2000 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"iyr:2012 pid:749770535 byr:1969 cid:148 hcl:#733820 hgt:180cm eyr:2021", false},
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
	}

	for _, test := range validTests {
		assert.Equal(t, test.isValid, validFieldsPart1(test.passport))
	}
}

func TestGetValidPassports(t *testing.T) {
	passports := getPassports("./day-four-input.txt")
	validPassports := getValidPassportsPart1(passports)

	assert.Equal(t, 242, len(validPassports))
}

func TestValidYear(t *testing.T) {
	validYearTests := []struct {
		name     string
		passport string
		isValid  bool
	}{
		{"birth year oldest", "ecl:hzl hgt:66in byr:1920 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"birth year too old", "ecl:hzl hgt:66in byr:1919 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", false},
		{"birth year youngest", "ecl:hzl hgt:66in byr:2002 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"birth year too young", "ecl:hzl hgt:66in byr:2003 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", false},
		{"issue year oldest", "ecl:hzl hgt:66in byr:1920 iyr:2010 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"issue year too old", "ecl:hzl hgt:66in byr:1920 iyr:2009 eyr:2020 pid:162973694 hcl:#a97842", false},
		{"issue year youngest", "ecl:hzl hgt:66in byr:1920 iyr:2020 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"issue year too young", "ecl:hzl hgt:66in byr:1920 iyr:2021 eyr:2020 pid:162973694 hcl:#a97842", false},
		{"expire year oldest", "ecl:hzl hgt:66in byr:1920 iyr:2017 eyr:2020 pid:162973694 hcl:#a97842", true},
		{"expire year too old", "ecl:hzl hgt:66in byr:1920 iyr:2017 eyr:2019 pid:162973694 hcl:#a97842", false},
		{"expire year youngest", "ecl:hzl hgt:66in byr:1920 iyr:2017 eyr:2030 pid:162973694 hcl:#a97842", true},
		{"expire year too young", "ecl:hzl hgt:66in byr:1920 iyr:2017 eyr:2031 pid:162973694 hcl:#a97842", false},
	}

	for _, test := range validYearTests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.isValid, validYear(test.passport))
		})
	}
}

func TestValidEyeColor(t *testing.T) {
	passport := "ecl:hzl hgt:66in byr:1920 iyr:2017 eyr:2031 pid:162973694 hcl:#a97842"

	assert.True(t, validEyeColor(passport))
}
