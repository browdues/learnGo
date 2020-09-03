package main

import (
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	factors := PrimeFactors(27)
	expected := []int{3, 3, 3}
	for i, v := range expected {
		if v != factors[i] {
			t.Errorf("A factor was incorrect, got: %d, want %d.", factors[i], v)
		}
	}
}

func TestPrimeFactorsNada(t *testing.T) {
	factors := PrimeFactors(0)
	expected := []int{}
	for i, v := range expected {
		if v != factors[i] {
			t.Errorf("A factor was incorrect, got: %d, want %d.", factors[i], v)
		}
	}
}

func TestPrimeFactorsPrime(t *testing.T) {
	factors := PrimeFactors(7)
	expected := []int{}
	for i, v := range expected {
		if v != factors[i] {
			t.Errorf("A factor was incorrect, got: %d, want %d.", factors[i], v)
		}
	}
}
