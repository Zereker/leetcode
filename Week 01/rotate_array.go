package main

func rotate(nums []int, k int) {
	length := len(nums)

	for i := k; i < length; i++ {
		curr := i
		swap := k + i
		if swap > length {
			swap = curr + k - length
		}

		nums[i], nums[swap] = nums[swap], nums[i]

	}

}

func main() {

}
