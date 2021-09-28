package main

import (
	"fmt"
)

func main() {
	n := 0
	m := minimumOneBitOperations(n)
	fmt.Println(n, m)

	n = 3
	m = minimumOneBitOperations(n)
	fmt.Println(n, m)
	n = 6
	m = minimumOneBitOperations(n)
	fmt.Println(n, m)
	n = 9
	m = minimumOneBitOperations(n)
	fmt.Println(n, m)
	n = 333
	m = minimumOneBitOperations(n)
	fmt.Println(n, m)
}
