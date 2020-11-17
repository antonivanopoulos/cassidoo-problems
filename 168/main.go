package main

import "fmt"

func main() {
	hit := CharacterJump(3, []int{0, 1, 0, 0, 0, 1, 0})

	fmt.Printf("CharacterJump(3, []int{0, 1, 0, 0, 0, 1, 0}) = %t\n", hit)
}

// CharacterJump takes in a step and an array that represents the layout of obstacles, and returns whether
// or not the character will hit any of the obstacles when stepping all the way through.
func CharacterJump(step int, hurdles []int) bool {
	for i := 0; i < len(hurdles)-1; i += step {
		hurdle := hurdles[i]

		if hurdle == 1 {
			return false
		}
	}

	return true
}
