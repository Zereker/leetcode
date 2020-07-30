package main

import "math"

func isValidBST(root *TreeNode) bool {
	return recurse(root, math.MinInt64, math.MaxInt64)
}

func recurse(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}

	if node.Val <= lower || node.Val >= upper {
		return false
	}

	return recurse(node.Left, lower, node.Val) && recurse(node.Right, node.Val, upper)
}
