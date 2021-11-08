package main

func checkSubarraySum(nums []int, k int) bool {
	tmp := make(map[int]int)
	tmp[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if prevIndex, ok := tmp[sum]; ok {
			if i-prevIndex > 1 {
				return true
			}
		} else {
			tmp[sum] = i
		}
	}
	return false
}
