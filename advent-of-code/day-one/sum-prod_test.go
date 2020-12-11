package sumprod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	t.Run("can read lines from a text file", func(t *testing.T) {
		fpath := "./input.txt"
		lines, err := LinesToArray(fpath)

		assert.NoError(t, err)
		assert.Equal(t, 200, len(lines))
	})
}

var (
	fpath    = "./input.txt"
	lines, _ = LinesToArray(fpath)
)

func TestFind2020Sum(t *testing.T) {
	t.Run("find two numbers that sum to 2020", func(t *testing.T) {
		s, err := Find2020Sum(lines)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(s))
		assert.Equal(t, int64(2020), s[0]+s[1])
	})

	t.Run("find three numbers that sum to 2020", func(t *testing.T) {
		s, err := Find2020TripleSum(lines)

		assert.NoError(t, err)
		assert.Equal(t, 3, len(s))
		assert.Equal(t, int64(2020), s[0]+s[1]+s[2])
	})
}
