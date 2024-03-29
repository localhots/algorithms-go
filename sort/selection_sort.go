package sort

// SelectionSort is an impelementaion of selection sort algorithm.
// Wikipedia: https://en.wikipedia.org/wiki/Selection_sort
//
// The algorithm divides the input list into two parts: the sublist of items
// already sorted, which is built up from left to right at the front (left) of
// the list, and the sublist of items remaining to be sorted that occupy the
// rest of the list. Initially, the sorted sublist is empty and the unsorted
// sublist is the entire input list. The algorithm proceeds by finding the
// smallest (or largest, depending on sorting order) element in the unsorted
// sublist, exchanging (swapping) it with the leftmost unsorted element (putting
// it in sorted order), and moving the sublist boundaries one element to the
// right.
func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

	return a
}
