package movie

import (
	"math"
)

// Movie returns the first n where ceil(systemB) < ceil(systemA)
func Movie(card, ticket int, perc float64) int {
	var systemA float64
	systemB := float64(card)
	n := 0

	for systemA <= math.Ceil(systemB) {
		n++
		systemA += float64(ticket)
		systemB += math.Pow(perc, float64(n)) * float64(ticket)
	}

	return n
}
