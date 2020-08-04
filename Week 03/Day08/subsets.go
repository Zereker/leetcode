package main

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var result [][]int

	for _, num := range nums {
		var newsets []int

		for _, subset := range result {
			newsets = append(newsets, num)
			
		}

		result = append(result, newsets)
	}
}
