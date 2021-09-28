package main

// TIMELiMIT  ?

// 1611. Minimum One Bit Operations to Make Integers Zero
// op: Change the rightmost (0th) bit in the binary representation of n.
// op: Change the ith bit in the binary representation of n if the (i-1)th bit is set to 1 and the (i-2)th through 0th bits are set to 0.
//  *****a -> *****b
//  ****a10000 -> *****b100000
// examples:  0 ; 1 ~> 0; 10 ~> 11 -> 01 ~> 0;
// 11 -> 01 ~> 0;
// 100 ~> 101 -> 111 ~> 110 -> 010 ~> 011 ->001 ~> 000
// 100 -> 1100 ~> 1101 -> 1111 ~> 1110 -> 1010 ~> 1011 -> 1001 ~> 1000 ...
//
// 0-> 1 -> 11 -> 10
// 110 -> 111 -> 101 -> 100
// 1100 -> 1101 -> 1111 -> 1110->1010 -> 1011 -> 1001 -> 1000
//  11000 -> 11001 -> 11011 - 11010 - 11110 -> 11111 -> 11101 - 11100 -> 10100 -> 10101 -> 10111 -> 10110->10010 -> 10011 -> 10001 -> 10000

func minimumOneBitOperations(n int) int {
	a := 0
	o := 0
	// n == 3
	for a != n {
		// a = 0
		if a%2 == 1 {
			a--
		} else {
			a++
		}
		// a = 1
		if a == n {
			return 2*o + 1
		}
		b := a
		c := 2
		// b = 1
		for b%2 == 0 {
			b = b / 2
			c = c * 2
		}
		// a = 1 c=2
		if b%4 == 1 {
			a = a + c
		} else if b%4 == 3 {
			a = a - c
		}
		// a = 3
		o += 1
	}
	return 2 * o

}
