package main

func toLowerCase(str string) string {
	var result string

	diff := 'a' - 'A'
	for _, i := range str {
		if i >= 'A' && i <= 'Z' {
			i = i + diff
		}

		result += string(i)
	}

	return result
}

func main() {
	lowerCase := toLowerCase("Hello")
	println(lowerCase)
}
