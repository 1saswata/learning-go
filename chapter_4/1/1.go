package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	s := make([]int, 0, 100)
	for range 100 {
		s = append(s, rand.IntN(100))
	}
	fmt.Println(s)
}
