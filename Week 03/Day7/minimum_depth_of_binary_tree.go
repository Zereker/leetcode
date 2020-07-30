package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}

	if root.Right != nil && root.Left == nil {
		return minDepth(root.Right) + 1
	}

	dl := minDepth(root.Left)
	dr := minDepth(root.Right)

	if dl > dr {
		return dr + 1
	} else {
		return dl + 1
	}
}
