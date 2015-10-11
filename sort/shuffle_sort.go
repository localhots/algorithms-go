package sort

import (
	"math/rand"
)

// ShuffleSort is an implementation of shuffle sort algorithm.
func ShuffleSort(a []int) []int {
	isSorted := func(a []int) bool {
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				return false
			}
		}

		return true
	}

	for {
		if isSorted(a) {
			return a
		}

		for i := range a {
			j := rand.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}
	}
}
