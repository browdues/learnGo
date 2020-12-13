package daythree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	lines := []string{
		"test",
		"",
	}
	filteredLines := filter(lines, func(l string) bool {
		return l != ""
	})

	assert.Equal(t, 1, len(filteredLines))
}

func TestReadMap(t *testing.T) {
	lines := readMap("./day-three-map.txt")

	assert.Equal(t, 323, len(lines))
}

func TestFindTrees(t *testing.T) {
	lines := readMap("./day-three-map.txt")
	sledloc := loc{0, 0}
	trees := sledloc.travel(lines, 3, 1)

	assert.Equal(t, 278, trees)
}

func TestFindTreesPart2(t *testing.T) {
	lines := readMap("./day-three-map.txt")
	result := 1

	slopeTests := []struct {
		inc  loc
		want int
	}{
		{loc{1, 1}, 90},
		{loc{3, 1}, 278},
		{loc{5, 1}, 88},
		{loc{7, 1}, 98},
		{loc{1, 2}, 45},
	}

	for _, st := range slopeTests {
		sledloc := loc{0, 0}
		nTrees := sledloc.travel(lines, st.inc.row, st.inc.col)
		assert.Equal(t, st.want, nTrees)
		result *= nTrees
	}
	fmt.Println(result)
}
