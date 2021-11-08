package main

func checkSubarraySum(nums []int, k int) bool {
	// 1 0 1 0 1
	// k = 4
	//[2,2,4,6,6]
	//7
	if len(nums) > 2*k+1 {
		return true
	}
	s := 0
	m := map[int]int{}
	m[0] = 0
	for i := 0; i < len(nums); i++ {
		// i   = 0  1  2  3  4
		// s%k = 2  4  1  0  6   k = 7
		// s%k = 1  1  2  2  3
		// m   = 0  1
		s += nums[i]
		s = s % k
		// i = 3

		if m[s] > 0 {
			if i != m[s] {
				return true
			}
		} else if s == 0 {
			if i > 0 {
				return true
			}
		} else {
			m[s] = i + 1
		}

	}
	return false

}
