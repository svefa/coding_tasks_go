package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	// s= "eceba" k= 2
	//"ab" k=1
	if k == 0 {
		return 0
	}
	n := 0

	j := 0
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 {
			k--
		}
		m[s[i]]++

		if k < 0 {
			//  i =3 m {e:2 c=1 b:1}
			for ; j <= i; j++ {
				m[s[j]]--
				if m[s[j]] == 0 {
					k++
					j++
					break
				}
			}
		}
		if n < i-j+1 {
			n = i - j + 1
		}
	}
	return n
}
