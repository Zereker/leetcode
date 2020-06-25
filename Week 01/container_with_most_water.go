package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := area(height, left, right)

	for left != right {
		if height[left] > height[right] {
			right--
		} else {
			left++
		}

		result = max(area(height, left, right), result)
	}

	return result
}

func area(height []int, left, right int) int {
	return min(height[left], height[right]) * (right - left)
}

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
