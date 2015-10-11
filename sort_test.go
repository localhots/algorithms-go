package algorithms

import (
	"testing"
)

func testSort(t *testing.T, sortFun func([]int) []int, a, exp []int) {
	s := sortFun(a)
	if len(s) != len(exp) {
		t.Fatal("Array sizes don't match")
	}
	for i := 0; i < len(exp); i++ {
		if s[i] != exp[i] {
			t.Fatalf("Expected sorted array to equal %v, got %v", exp, s)
		}
	}
}

func unsortedArray() []int {
	return []int{
		57, 64, 83, 25, 26, 10, 55, 22, 76, 61,
		28, 77, 56, 32, 63, 17, 91, 20, 58, 16,
		1, 51, 88, 82, 24, 70, 81, 35, 49, 39,
		89, 30, 46, 6, 41, 19, 43, 67, 53, 97,
		65, 37, 13, 23, 29, 69, 0, 73, 9, 59,
		96, 34, 66, 79, 27, 14, 40, 80, 98, 2,
		5, 45, 50, 4, 85, 18, 86, 7, 87, 31,
		95, 47, 68, 36, 15, 48, 8, 92, 11, 74,
		78, 52, 44, 42, 54, 84, 12, 21, 38, 99,
		72, 33, 71, 93, 60, 62, 90, 94, 3, 75,
	}
}
