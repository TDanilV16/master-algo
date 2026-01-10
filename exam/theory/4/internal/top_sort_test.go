package internal

import "testing"

func TestBase(t *testing.T) {
	graph := [][]int{
		{},
		{0, 2},
		{0},
	}

	answer := TopSort(graph)
	expected := []int{1, 2, 0}

	if len(answer) != len(expected) {
		t.Errorf("TopSort(graph) output length = %d, want %d", len(answer), len(expected))
	}

	for i := 0; i < len(answer); i++ {
		if answer[i] != expected[i] {
			t.Errorf("TopSort(graph) output differs = %d, want %d at index %d", answer[i], expected[i], i)
		}
	}

}
