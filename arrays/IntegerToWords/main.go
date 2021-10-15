package main

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	big := []string{"Billion", "Million", "Thousand", ""}
	b := []int{num / (1000 * 1000 * 1000), num / (1000 * 1000) % 1000, num / 1000 % 1000, num % 1000}
	res := []string{}
	dig := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	ten := []string{"", "x", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	for i := 0; i < len(big); i++ {
		if b[i] == 0 {
			continue
		}
		h := b[i] / 100
		t := (b[i] / 10) % 10
		d := b[i] % 10
		if h > 0 {
			res = append(res, dig[h], "Hundred")
		}
		if t > 1 {
			res = append(res, ten[t])
		}
		if t == 1 {
			res = append(res, dig[d+10])
		} else if d > 0 {
			res = append(res, dig[d])
		}
		if len(big[i]) > 0 {
			res = append(res, big[i])
		}
	}
	return strings.Join(res, " ")

}
