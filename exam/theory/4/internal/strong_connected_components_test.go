package internal

import "testing"

func Test_Base(t *testing.T) {
	g := [][]int{
		{1, 2},
		{3},
		{4},
		{4},
		{1},
	}

	gT := [][]int{
		{},
		{0, 4},
		{0},
		{1},
		{2, 3},
	}

	result := FindStrongConnectedComponents(g, gT)

	expected := [][]int{
		{0},
		{2},
		{1, 4, 3},
	}

	if len(result) != len(expected) {
		t.Errorf("FindStrongConnectedComponents() result length is %v, expected %v", len(result), expected)
	}

	for i := range result {
		if len(result[i]) != len(expected[i]) {
			t.Errorf("FindStrongConnectedComponents() result[%d] length is %v, expected %v", i, len(result), expected)
		}
		for j := range result[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("FindStrongConnectedComponents() result elements at index %d differs, got: %v, expected %v", j, result[i], expected[i])
			}
		}
	}
}
