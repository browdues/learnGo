package valid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	lines, err := ReadLines("./input.txt")

	assert.NoError(t, err)
	assert.Equal(t, 1000, len(lines))
}

var lines, _ = ReadLines("./input.txt")

func TestParsePasswords(t *testing.T) {
	passwords, err := ParsePassword(lines)

	assert.NoError(t, err)
	assert.Equal(t, len(lines), len(passwords))
}

func TestValidatePassword(t *testing.T) {
	passwords, _ := ParsePassword(lines)
	validPasswords := ValidatePassword(passwords)

	assert.Equal(t, 655, validPasswords)
}

func TestValidatePasswordPart2(t *testing.T) {
	passwords, _ := ParsePassword(lines)
	validPasswords := ValidatePasswordPart2(passwords)

	assert.Equal(t, 673, validPasswords)
}
