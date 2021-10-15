package main

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		var data []byte

		b := make([]byte, 4)
		for len(data) < n {
			c := read4(b)
			data = append(data, b[:c]...)
		}

		return copy(buf, data[:n])
	}
}
