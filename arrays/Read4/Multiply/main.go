package main

// moving vars into solution helps
// before: not working when submit, working when run : the same test case

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	var buf4 []byte
	var j int
	var k int
	// implement read below.
	return func(buf []byte, n int) int {
		if buf4 == nil {
			buf4 = make([]byte, 4)
			j = 0
			k = 0
		}
		i := 0
		for j < k && i < n {
			buf[i] = buf4[j]
			i++
			j++
		}
		for i < n {
			j = 0
			k = read4(buf4)

			for j < k && i < n {
				buf[i] = buf4[j]
				i++
				j++
			}
			if k < 4 {
				break
			}
		}
		return i

	}
}
