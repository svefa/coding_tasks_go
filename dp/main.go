package main

import (
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	s := trap(height) // 6
	fmt.Println(s)

	height = []int{4, 2, 0, 3, 2, 5}
	s = trapX(height) // 9
	fmt.Println(s)
}
