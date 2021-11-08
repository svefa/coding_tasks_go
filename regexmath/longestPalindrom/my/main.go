package main

func longestPalindrome(s string) string {
	s = s
	memo := map[int]int{}
	queue := []int{}
	for i := 1; i < len(s); i++ {
		memo[2*i] = 1
		queue = append(queue, 2*i)
	}
	for i := 0; i < len(s); i++ {
		memo[2*i+1] = 2
		queue = append(queue, 2*i+1)
	}
	///  0   1   2    3
	///  a   b   b   c
	///0 1 2 3 4 5 6 7 8
	///
	/// queu 2:0 4:0 6:0
	/// queu 1:1  3:1  5:1  7:1
	//
	//  2:0  a:b  i = 2  [(2-2)/2 : 2/2]
	//  4:0  b:b  i= 4   1:2
	//  3:1  0 2          (3-1)/2   (3+1)/2

	curr := 1
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		l := memo[i]
		if i-l >= 0 && i+l < 2*len(s) {
			if s[(i-l)/2] == s[(i+l)/2] {
				memo[i] = l + 2
				queue = append(queue, i)
				curr = i
			}
		}
	}
	return s[(curr-memo[curr]+2)/2 : ((curr+memo[curr]-2)/2 + 1)]
}
