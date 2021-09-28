package main

import "fmt"

func main() {
	s := "babad"
	out := longestPalindrome_n_n(s)
	fmt.Println(s, out)

	s = "cbaxccxabd"
	out = longestPalindrome_n_n(s)
	fmt.Println(s, out)

	s = "a"
	out = longestPalindrome_n_n(s)
	fmt.Println(s, out)

}
