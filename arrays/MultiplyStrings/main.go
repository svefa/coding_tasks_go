package main

import "fmt"

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	// "2"*"3"  l1,l2 = 1,1
	p := make([]byte, l1+l2)
	for i := range p {
		p[i] = '0'
	}
	// [xx]
	for i1, d1 := range num1 {
		for i2, d2 := range num2 {
			x := byte((d1 - '0') * (d2 - '0'))
			// "2"*"3" x = 6
			// xAxx By zzzcCzzz
			// 1 [4] 0[2]   4 [8]  3-1  2-1 2*1   8-3 = 5

			j := l1 + l2 - (l1 - i1 - 1) - (l2 - i2 - 1) - 1 //23 * 36  0,0 ->  1  1,1 ->3 0000
			// j = 1+1 - (1-0-1) - (1-0-1) -1 = 1

			for x > 0 {
				a := p[j] + x - '0'
				p[j] = (a)%10 + '0'
				x = a / 10
				j--
			}
		}
	}

	if p[0] == '0' {
		return string(p[1:])
	}
	return string(p)
}

func main() {
	//n :=
	n := multiply("22", "33")
	fmt.Println(n)
}
