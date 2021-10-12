// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	// s= "abcabcbb"
	/// "tmmzuxt"
	cnt := make(map[byte]int, 26)
	i, j, l := 0, 0, 0
	for i < len(s) {
		cnt[s[i]]++
		// cnt: { a:2, b:1, c:1}
		/// { t:2 m:1 z:1 u:1 x:1 }
		if cnt[s[i]] > 1 {
			// i = 3 , j=0
			/// i = 2 j = 0
			/// i=6 j =2
			for s[j] != s[i] {
				cnt[s[j]]--
				j++
				/// j= 1
			}
			cnt[s[j]]--
			// {a:1 b:1 c:1}
			j++
			// j=1
			/// j=2
		}
		// i=4
		i++
		/// i=3
		if l < i-j {
			l = i - j
			// l=1 l=2 l=3
			/// l =1
		}
	}
	return l
}