package algorithms

// BinarySearch is an implementation of binary search algorithm.
// Wikipedia: https://en.wikipedia.org/wiki/Binary_search_algorithm
//
// The binary search algorithm begins by comparing the target value to the value
// of the middle element of the sorted array. If the target value is equal to
// the middle element's value, then the position is returned and the search is
// finished. If the target value is less than the middle element's value, then
// the search continues on the lower half of the array; or if the target value
// is greater than the middle element's value, then the search continues on the
// upper half of the array. This process continues, eliminating half of the
// elements, and comparing the target value to the value of the middle element
// of the remaining elements - until the target value is either found (and its
// associated element position is returned), or until the entire array has been
// searched (and "not found" is returned).
func BinarySearch(haystack []int, needle int) (pos int, ok bool) {
	if len(haystack) == 0 {
		return 0, false
	}
	if haystack[0] > needle {
		return 0, false
	}
	if haystack[len(haystack)-1] < needle {
		return 0, false
	}

	first := 0
	last := len(haystack)
	for first < last {
		mid := first + (last-first)/2

		if haystack[mid] == needle {
			return mid, true
		} else if haystack[mid] > needle {
			last = mid
		} else {
			first = mid + 1
		}
	}

	return 0, false
}
