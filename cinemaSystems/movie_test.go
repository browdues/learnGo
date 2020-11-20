package movie

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestAddWithTestify(t *testing.T) {
	n := Movie(100, 10, .1)
	Equal(t, 0, n)
}
