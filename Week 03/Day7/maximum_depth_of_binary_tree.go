package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)

	if depthLeft > depthRight {
		return depthLeft + 1
	} else {
		return depthRight + 1
	}
}
