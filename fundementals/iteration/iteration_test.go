package iteration

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestRepeat5x(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	Equal(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("n", 3)
	fmt.Println(repeated)
	// Output: nnn
}
