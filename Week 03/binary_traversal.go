package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0, 10000)
	result = append(result, inOrderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inOrderTraversal(root.Right)...)
	return result
}

func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, root.Val)
	result = append(result, preOrderTraversal(root.Left)...)
	result = append(result, preOrderTraversal(root.Right)...)
	return result
}

func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, postOrderTraversal(root.Left)...)
	result = append(result, postOrderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}
