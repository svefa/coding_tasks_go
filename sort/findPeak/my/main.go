package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	// [3,4,3,2,1]
	l, r := 0, len(nums)-1
	// l = 0 r =4
	for l <= r {
		m := l + (r-l)/2
		// m = 0 + 2 = 2
		// [4, 3, 2]
		if nums[m-1] < nums[m] && nums[m] > nums[m+1] {
			return m
		} else if nums[m-1] > nums[m] && nums[m] > nums[m+1] {
			r = m
			// l = 0 r = 2
		} else if nums[m-1] < nums[m] && nums[m] < nums[m+1] {
			l = m
		} else {
			l = m
		}
	}
	return -1

}
