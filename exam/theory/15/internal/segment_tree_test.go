package internal

import "testing"

func Test_Sum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := NewSegmentTreeSum(a)

	result := tree.Sum(1, 8)
	if result != 44 {
		t.Errorf("Sum(%d, %d) = %d, want %d", 1, 8, result, 44)
	}
}

func Test_Update(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := NewSegmentTreeSum(a)

	tree.Update(1, 8)
	result := tree.Sum(1, 8)
	if result != 50 {
		t.Errorf("Sum(%d, %d) = %d, want %d after Update(%d, %d)", 1, 8, result, 50, 1, 8)
	}
}
