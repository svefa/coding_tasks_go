package main

func divide(dividend int, divisor int) int {
	// no idea so far
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	if dividend < 0 {
		return -divide(-dividend, divisor)
	}
	if divisor < 0 {
		return -divide(dividend, -divisor)
	}
	y, x := dividend, divisor
	//y=10 x=3
	res := 0
	for y >= x {
		println(y, x, res)
		c := 1
		// 6 < 10  X 12 < 10
		// X 6  < 4
		for (x << 1) <= y {
			x = x << 1
			c = c << 1
			// x =6 c=2
		}
		// y = 10 -6 = 4
		// y = 4 - 3 = 1
		y -= x
		// x = 3

		x = divisor
		// res = 2
		// res = 2 +1 = 3
		res += c
		println(y, x, res, "c", c)
	}
	return res
}

func main() {
	//println(divide(10, 3))
	//println(divide(1, 1))
	println(divide(-2147483648, -1))
	//println(divide(10, 3))
	//println(divide(10, 3))
}
