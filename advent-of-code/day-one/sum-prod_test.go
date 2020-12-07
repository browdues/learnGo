package sumprod

import (
	"fmt"
	"testing"
)

func TestFixture(t *testing.T) {
	t.Run("can read lines from a text file", func(t *testing.T) {
		fpath := "./input.txt"
		lines, err := LinesToArray(fpath)

		if err != nil {
			fmt.Println("unexpected err")
		}
	})
}
