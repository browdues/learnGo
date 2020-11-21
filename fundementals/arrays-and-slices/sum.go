package sum

// Sum returns the sum of the array
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// All takes n slices and returns n sums
func All(numbersToSum ...[]int) []int {
	var sums []int

	for _, numberSlice := range numbersToSum {
		sums = append(sums, Sum(numberSlice))
	}

	return sums
}

// AllTails takes n slices and returns the n tail sums
// a tail sum is the sum of all numbers in a slice except the first (its head)
func AllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, sumSlice := range numbersToSum {
		if len(sumSlice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := sumSlice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
