package main

import "fmt"

func main() {
	s := "oh heavens"
	c := 'h'

	count := numChars(s, c)
	fmt.Printf("numChars(%s, %c): %d\n", s, c, count)
}

func numChars(s string, c rune) int {
	count := 0

	for _, char := range s {
		if char == c {
			count++
		}
	}

	return count
}
