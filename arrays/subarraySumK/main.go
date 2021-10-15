package main

func subarraySum(nums []int, k int) int {
	l := 0
	s := 0
	// sum to count
	m := map[int]int{s: 1}

	for i := 0; i < len(nums); i++ {
		s = s + nums[i]
		l += m[s-k]

		m[s]++

	}
	return l

}
