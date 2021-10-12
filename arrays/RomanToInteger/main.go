package main

func romanToInt(s string) int {
	m := map[byte]int{'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	m2 := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	i, res := 0, 0
	for i < len(s)-1 {
		cc := s[i : i+2]
		if m2[cc] != 0 {
			res += m2[cc]
			i += 2
		} else {
			res += m[s[i]]
			i++
		}
	}
	if i == len(s)-1 {
		res += m[s[i]]
	}
	return res
}
