package algorithms

import (
	"testing"
)

func TestBinarySearchEmptyArray(t *testing.T) {
	if _, ok := BinarySearch([]int{}, 1); ok {
		t.Error("Position should not be found")
	}
}

func TestBinarySearchLowerThanFirstValue(t *testing.T) {
	if _, ok := BinarySearch([]int{1, 2, 3, 4, 5}, 0); ok {
		t.Error("Position should not be found")
	}
}

func TestBinarySearchHigherThanLastValue(t *testing.T) {
	if _, ok := BinarySearch([]int{1, 2, 3, 4, 5}, 6); ok {
		t.Error("Position should not be found")
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	if _, ok := BinarySearch([]int{1, 2, 4, 5}, 3); ok {
		t.Error("Position should not be found")
	}
}

func TestBinarySearchOddSize(t *testing.T) {
	pos, ok := BinarySearch([]int{1, 2, 3, 4, 5}, 1)
	if !ok {
		t.Fatalf("Position should be found")
	}
	if pos != 0 {
		t.Errorf("Expected position %d, got %d", 0, pos)
	}
}

func TestBinarySearchEvenSize(t *testing.T) {
	pos, ok := BinarySearch([]int{1, 2, 3, 4, 5, 6}, 6)
	if !ok {
		t.Fatalf("Position should be found")
	}
	if pos != 5 {
		t.Errorf("Expected position %d, got %d", 5, pos)
	}
}
