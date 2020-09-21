package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	println(longestCommonPrefix([]string{"a"}))
	println(longestCommonPrefix([]string{"c", "c"}))
}
