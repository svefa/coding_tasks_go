package main

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	L, R, M := l, r, -1
	for l <= r {
		m := l - (l-r)/2
		if nums[m] < target {
			L = m
			l = m + 1
		} else if nums[m] > target {
			R = m
			r = m - 1
		} else { // nums[m] == target
			M = m
			if m+1 == len(nums) || nums[m+1] > target {
				R = m + 1
				break
			}
			l = m + 1
		}
	}
	l, r = L, M
	for l <= r {
		m := l - (l-r)/2
		if nums[m] < target {
			L = m
			l = m + 1
		} else if nums[m] > target {
			R = m
			r = m - 1
		} else { // nums[m] == target
			M = m
			if m == 0 || nums[m-1] < target {
				L = m - 1
				break
			}
			r = m - 1
		}
	}
	if M == -1 {
		return []int{-1, -1}
	}
	return []int{L + 1, R - 1}

}
