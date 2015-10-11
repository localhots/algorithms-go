package sort

// BinaryInsertionSort is an implementation of binary insertion sort algorithm.
// Wikipedia: https://en.wikipedia.org/wiki/Insertion_sort#Variants
func BinaryInsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		v := a[i]
		first := 0
		last := i

		for first < last {
			mid := first + (last-first)/2
			if v < a[mid] {
				last = mid
			} else {
				first = mid + 1
			}
		}

		for j := i; j > first; j-- {
			a[j] = a[j-1]
		}

		a[first] = v
	}

	return a
}
