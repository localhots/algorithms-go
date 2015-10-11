package algorithms

// BinaryInsertionSort is an implementation of binary insertion sort algorithm.
// Wikipedia: https://en.wikipedia.org/wiki/Insertion_sort#Variants
func BinaryInsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		v := a[i]
		low := 0
		high := i

		for low < high {
			mid := low + (high-low)/2
			if v < a[mid] {
				high = mid
			} else {
				low = mid + 1
			}
		}

		for j := i; j > low; j-- {
			a[j] = a[j-1]
		}

		a[low] = v
	}

	return a
}
