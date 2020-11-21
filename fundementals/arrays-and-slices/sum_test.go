package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5} // fixed capacity

	want := 15
	got := Sum(numbers)

	if got != want {
		t.Errorf("got %v want %d given, %v", got, want, numbers)
	}
}
