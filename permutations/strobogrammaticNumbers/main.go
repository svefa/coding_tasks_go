package main

import "fmt"

func findStrobogrammatic(n int) []string {
	pd := map[string]string{"9": "6", "6": "9", "1": "1", "8": "8"}
	t := []string{}
	if n%2 == 0 {
		t = append(t, "")
	} else {
		t = append(t, "0", "1", "8")
	}
	for i := 0; i < n/2; i++ {
		nt := []string{}
		for _, s := range t {
			for p, d := range pd {
				nt = append(nt, p+s+d)
			}
			if i < n/2-1 {
				nt = append(nt, "0"+s+"0")
			}
		}
		fmt.Println(t)
		t = nt
	}
	return t
}

func main() {
	x := findStrobogrammatic(4)
	fmt.Println(x)
}
