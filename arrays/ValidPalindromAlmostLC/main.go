package main

/*
   "aba"

   "abca"

    ab a
    a ca
     ij
*/
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return palindrome(s[i:j]) || palindrome(s[i+1:j+1])
		}
		i++
		j--
	}
	return true
}

func palindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
