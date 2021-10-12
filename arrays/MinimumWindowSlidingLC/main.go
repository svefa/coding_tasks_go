package main
func minWindow(s string, t string) {
	n,m := len(s),len(t)
	if n<m{
		retrun ""
	}

	var need[58]int
	valid := 0
	for i:=0; i < m; i++ {
		if need[t[i]-'A'] == 0 {
			valid++
		}
		need[t[i]-'A']++
	}
	var freq [58]int
	cur:=0
	l,r:=0,0
	res:=""
	for l<n {
		if r < n && valid != curr {
			freq[s[r]-'A'] ++
			if freq[s[r] - 'A'] == need[s[r]-'A'] {
				curr ++
			}
			r++
		} else {
            freq[s[l]-'A'] --
            if freq[s[l]-'A'] < need[s[l]-'A']{
                curr--
            } 
            l++
		}
        if curr == valid && (res == "" || len(res) >= r+1-l) {
            res = s[l:r]
        }
    }
    return res
}
