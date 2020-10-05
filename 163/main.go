package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	num  int
}

type People []Person

func main() {
	p := People{
		{name: "Joe", num: 9},
		{name: "Cami", num: 3},
		{name: "Cassidy", num: 4},
	}

	r := gimmePizza(p, 8)
	fmt.Println(r)
}

func (p People) TotalSlices() int {
	total := 0
	for _, req := range p {
		total += req.num
	}

	return total
}

func gimmePizza(arr People, num int) int {
	return int(math.Ceil(float64(arr.TotalSlices()) / float64(num)))
}
