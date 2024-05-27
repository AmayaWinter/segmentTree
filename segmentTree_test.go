package segmentTree

import (
	"testing"
)

func TestBuildTreeEmptyList(t *testing.T) {
	if _, err := New([]int{}); err == nil {
		t.Errorf("failed to return error on empty list")
		return
	}
}

func TestBuildTreeSingleItem(t *testing.T) {
	segTree, err := New([]int{1})
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	expected := []int{0, 1}

	if len(segTree.tree) != len(expected) {
		t.Errorf("built tree is incorrect: expected %v but got %v", expected, segTree.tree)
		return
	}

	for i, actualNum := range segTree.tree {
		if expected[i] != actualNum {
			t.Errorf("built tree is incorrect: expected %v but got %v", expected, segTree.tree)
		}
	}

}

func TestBuildTree(t *testing.T) {
	segTree, err := New([]int{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	expected := []int{0, 21, 18, 3, 7, 11, 1, 2, 3, 4, 5, 6}

	if len(segTree.tree) != len(expected) {
		t.Errorf("built tree is incorrect: expected %v but got %v", expected, segTree.tree)
		return
	}

	for i, actualNum := range segTree.tree {
		if expected[i] != actualNum {
			t.Errorf("built tree is incorrect")
		}
	}
}

func TestRangeSum(t *testing.T) {
	segTree, _ := New([]int{1, 2, 3, 4, 5, 6})

	testSets := [][]int{
		[]int{0, 5, 21},
		[]int{1, 5, 20},
		[]int{2, 3, 7},
	}

	for _, testParams := range testSets {
		left, right, expectedSum := testParams[0], testParams[1], testParams[2]
		if ans, err := segTree.RangeSum(left, right); err == nil {
			if ans != expectedSum {
				t.Errorf("expected sum %d != actual sum %d", expectedSum, ans)
				return
			}
		} else {
			t.Errorf("unexpected error %v", err)
		}
	}
}

func TestUpdate(t *testing.T) {
	segTree, _ := New([]int{1, 2, 3, 4, 5, 6})
	testSets := [][]int{
		[]int{0, 0, 20},
		[]int{5, 10, 24},
		[]int{3, 1, 21},
	}

	for _, testParams := range testSets {
		index, newVal, expectedSum := testParams[0], testParams[1], testParams[2]
		if err := segTree.Update(index, newVal); err != nil {
			t.Errorf("unexpected error when updating tree: %v", err)
			return
		}
		if ans, err := segTree.RangeSum(0, 5); err == nil {
			if ans != expectedSum {
				t.Errorf("expected sum %d != actual sum %d", expectedSum, ans)
			}
		} else {
			t.Errorf("unexpected error %v", err)
		}
	}
}