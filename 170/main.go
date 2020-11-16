package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1, 1, 3}
	n := specialPairs(arr)

	fmt.Printf("specialPairs(%v): %d\n", arr, n)
}

func specialPairs(arr []int) int {
	matching := make(map[int][]int)

	// Build map of matching valus
	for k, v := range arr {
		matching[v] = append(matching[v], k)
	}

	// Calculate the number of matching pairs
	n := 0
	for _, v := range matching {
		n += len(v) * (len(v) - 1) / 2
	}

	return n
}
