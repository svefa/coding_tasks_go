package main

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		b := make([]byte, 26)
		for _, c := range str {
			b[c-'a']++
		}
		m[string(b)] = append(m[string(b)], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
