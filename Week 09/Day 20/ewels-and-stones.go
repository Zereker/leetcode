package main

func numJewelsInStones(J string, S string) int {
	var jewels []rune

	for _, i := range J {
		jewels = append(jewels, i)
	}

	var result int
	for _, i := range S {
		for _, jewel := range jewels {
			if i == jewel {
				result++
			}
		}
	}

	return result
}

func main() {
	println(numJewelsInStones("aA", "aAAbbbb"))
	println(numJewelsInStones("z", "ZZ"))
}
