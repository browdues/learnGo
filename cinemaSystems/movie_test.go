package movie

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestMovieExample1(t *testing.T) {
	n := Movie(500, 15, 0.9)
	Equal(t, 43, n)
}

func TestMovieExample2(t *testing.T) {
	n := Movie(100, 10, 0.95)
	Equal(t, 24, n)
}

func TestMovieFreeCard(t *testing.T) {
	n := Movie(0, 10, 0.95)
	Equal(t, 2, n)
}
