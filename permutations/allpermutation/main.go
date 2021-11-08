package main

func permute(nums []int) (res [][]int) {
	if len(nums) == 1 {
		res = [][]int{{nums[0]}}
		return
	}
	for i := 0; i < len(nums); i++ {
		nums[i], nums[0] = nums[0], nums[i]
		R := permute(nums[1:])
		for _, r := range R {
			res = append(res, append(r, nums[0]))
		}
		nums[i], nums[0] = nums[0], nums[i]
	}
	return res
}
