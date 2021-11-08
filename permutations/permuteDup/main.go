package main

func permuteUnique(nums []int) (res [][]int) {
	if len(nums) == 1 {
		res = [][]int{{nums[0]}}
	}
	m := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[i], nums[0] = nums[0], nums[i]
		P := permuteUnique(nums[1:])
		for _, p := range P {
			res = append(res, append(p, nums[0]))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return
}
