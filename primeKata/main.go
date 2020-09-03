package main

import (
	"fmt"
)

//PrimeFactors when given any number, returns that number's prime factors s.t
//the product of those prime factors is that number.
//Else it returns an empty array
func PrimeFactors(n int) []int {
	var result []int
	var nCopy = n

	for factor := 2; factor <= nCopy; {
		if nCopy%factor == 0 && n != factor {
			result = append(result, factor)
			nCopy = nCopy / factor
			factor = 2
		} else {
			factor++
		}
	}

	return result
}

func main() {
	fmt.Println(PrimeFactors(27))
}
