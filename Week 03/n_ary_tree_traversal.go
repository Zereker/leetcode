package main

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}

	result = append(result, root.Val)

	return result
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, root.Val)
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}

	return result
}
