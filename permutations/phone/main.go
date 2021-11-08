package main

func letterCombinations(digits string) (res []string) {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		switch d := digits[0]; d {
		case '2':
			return []string{"a", "b", "c"}
		case '3':
			return []string{"d", "e", "f"}
		case '4':
			return []string{"g", "h", "i"}
		case '5':
			return []string{"j", "k", "l"}
		case '6':
			return []string{"m", "n", "o"}
		case '7':
			return []string{"p", "q", "r", "s"}
		case '8':
			return []string{"t", "u", "v"}
		case '9':
			return []string{"w", "x", "y", "z"}

		}
	}

	a := letterCombinations(digits[:n/2])
	b := letterCombinations(digits[n/2:])
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			res = append(res, a[i]+b[j])
		}
	}
	return res
}
