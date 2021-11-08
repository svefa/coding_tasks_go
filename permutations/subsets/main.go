package main

func subsets(nums []int) (res [][]int) {
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		k := len(res)
		for j := 0; j < k; j++ {
			res = append(res, append([]int{nums[i]}, res[j]...))
		}
	}
	return
}
