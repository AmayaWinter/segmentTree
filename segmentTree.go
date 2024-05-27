package segmentTree

import (
	"errors"
)

type SegmentTree struct {
	tree    []int
	lenNums int
}

func New(nums []int) (SegmentTree, error) {
	if len(nums) == 0 {
		return SegmentTree{}, errors.New("input cannot be empty")
	}
	// create an 1 based segment tree using an array
	// Size is based on the fact that given a complete tree, the number of nodes will be
	// more than the total number of nodes on every level above

	//   	 1        1 node
	//  	/ \
	// 	   2   3      2 nodes
	//	  / \ / \
	//   4  5 6  7    4 nodes
	tree := make([]int, len(nums))
	tree = append(tree, nums...)

	// start right before the last level, get the children of each node
	// and build up the sum towards the root
	for i := len(nums) - 1; i > 0; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}

	return SegmentTree{
		tree:    tree,
		lenNums: len(nums),
	}, nil
}

func (self *SegmentTree) RangeSum(left int, right int) (int, error) {
	if left < 0 || right >= self.lenNums {
		return 0, errors.New("range is out of bounds")
	}

	// move the indices into the part of the array where the original numbers are stored
	left += self.lenNums
	right += self.lenNums
	rangeSum := 0

	for left <= right {

		if left%2 == 1 {
			rangeSum += self.tree[left]
			left += 1
		}
		if right%2 == 0 {
			rangeSum += self.tree[left]
			right -= 1
		}
		left /= 2
		right /= 2
	}

	return rangeSum, nil
}

func (self *SegmentTree) Update(index int, newVal int) error {
	if index < 0 || index >= self.lenNums {
		return errors.New("index is out of range")
	}

	index += self.lenNums
	self.tree[index] = newVal

	for index > 0 {
		left, right := index, index

		// if this is a left child, increment right by 1 to get the sibling right child
		if index%2 == 0 {
			right += 1
		} else {
			// decrement to get the sibling left child
			left -= 1
		}

		index /= 2
		if index == 0 {
			break
		}
		// set parent to sum of children
		self.tree[index] = self.tree[left] + self.tree[right]
	}

	return nil
}