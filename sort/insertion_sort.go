package sort

// InsertionSort is an implementation of insertion sort algorithm.
// Wikipedia: https://en.wikipedia.org/wiki/Insertion_sort
//
// Insertion sort iterates, consuming one input element each repetition, and
// growing a sorted output list. Each iteration, insertion sort removes one
// element from the input data, finds the location it belongs within the sorted
// list, and inserts it there. It repeats until no input elements remain.
func InsertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}

	return a
}

// InsertionSortOptimized is an optimized implementation of insertion sort
// algorithm.
func InsertionSortOptimized(a []int) []int {
	// Moving the "sentinel" value to the first position
	for i := len(a) - 1; i > 0; i-- {
		if a[i] < a[i-1] {
			a[i-1], a[i] = a[i], a[i-1]
		}
	}
	if len(a) < 3 {
		return a
	}

	for i := 2; i < len(a); i++ {
		val := a[i]
		j := i
		for ; val < a[j-1]; j-- {
			a[j] = a[j-1]
		}
		a[j] = val
	}

	return a
}
