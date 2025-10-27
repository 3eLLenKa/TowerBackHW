package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Insert(value int) *TreeNode {
	if tn == nil {
		return &TreeNode{
			Value: value,
		}
	}

	if value < tn.Value {
		tn.Left = tn.Left.Insert(value)
	} else if value > tn.Value {
		tn.Right = tn.Right.Insert(value)
	}

	return tn
}

func (tn *TreeNode) Remove(value int) *TreeNode {
	if tn == nil {
		return nil
	}

	if value < tn.Value {
		tn.Left = tn.Left.Remove(value)
	} else if value > tn.Value {
		tn.Right = tn.Right.Remove(value)
	} else {
		if tn.Left == nil && tn.Right == nil {
			return nil
		} else if tn.Left == nil {
			return tn.Right
		} else if tn.Right == nil {
			return tn.Left
		} else {
			minRight := tn.Right.min()
			tn.Value = minRight.Value
			tn.Right = tn.Right.Remove(minRight.Value)
		}
	}

	return tn
}

func (tn *TreeNode) min() *TreeNode {
	current := tn
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (tn *TreeNode) Find(value int) *TreeNode {
	if tn == nil || tn.Value == value {
		return tn
	}

	if value < tn.Value {
		return tn.Left.Find(value)
	}

	return tn.Right.Find(value)
}

func (tn *TreeNode) Depth() int {
	if tn == nil {
		return 0
	}

	leftDepth := tn.Left.Depth()
	rightDepth := tn.Right.Depth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func main() {
	var tree *TreeNode

	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, v := range values {
		tree = tree.Insert(v)
	}

	fmt.Println("Depth: ", tree.Depth())
	fmt.Println("Find (80): ", tree.Find(80) != nil)
	fmt.Println("Find (52): ", tree.Find(52) != nil)

	tree = tree.Remove(40)
	fmt.Println("find 40 after remove:", tree.Find(40) != nil)
}
