package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sums any size collection", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		want := 6
		got := Sum(numbers)

		if got != want {
			t.Errorf("got %v want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums any size slice collection", func(t *testing.T) {
		numbersA := []int{1, 2, 3}
		numbersB := []int{1, 1, 1}

		want := []int{6, 3}
		got := All(numbersA, numbersB)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make sums of some slices", func(t *testing.T) {
		numbersA := []int{1, 2, 3}
		numbersB := []int{1, 1, 1}

		want := []int{5, 2}
		got := AllTails(numbersA, numbersB)

		checkSums(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		numbersA := []int{}
		numbersB := []int{1, 1, 1}

		want := []int{0, 2}
		got := AllTails(numbersA, numbersB)

		checkSums(t, got, want)
	})
}
