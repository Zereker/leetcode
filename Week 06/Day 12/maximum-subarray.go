package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1], 0) + nums[i]

		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

func main() {

}
