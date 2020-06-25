package main

import "fmt"

func moveZeroes(nums []int) {
	length := len(nums)
	var j int

	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("%v\n", nums)
}
