package main

// ?
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		var data []byte
		b := make([]byte, 4)
		totalChars := 0
		readChars := 4
		for readChars == 4 {
			readChars = read4(b)
			data = append(data, b[:readChars]...)
			totalChars += readChars
		}
		return copy(buf, data[:min(totalChars, n)])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
