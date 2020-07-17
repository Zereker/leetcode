package main

func largestRectangleArea(heights []int) int {
	var maxArea int

	for i := 0; i < len(heights)-2; i++ {
		for j := i + 1; j < len(heights)-1; j++ {
			var area int
			if heights[i] > heights[j] {
				break
			}

			area = heights[i] * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
