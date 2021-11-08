package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	y := 1.0
	factor := x
	for n > 0 {
		if n%2 == 1 {
			y = y * factor
		}
		n = n / 2
		factor = factor * factor
	}
	return y
}
