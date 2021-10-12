package main

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, si, d := 0, 1, 0

	for i < len(s) && (s[i] == ' ') {
		i++
	}
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			si = -1
		}
		i++
	}
	for i < len(s) && '0' <= s[i] && s[i] <= '9' {

		d = d*10 + int(s[i]-'0')
		i++
		if si*d < -(1 << 31) {
			return -(1 << 31)
		}
		if si*d > (1<<31)-1 {
			return (1 << 31) - 1
		}

	}
	return si * d
}

func main() {
	s := " "
	n := myAtoi(s)
	println(n)
}
