package main

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//Notice that the solution set must not contain duplicate triplets.

func threeSum(nums []int) [][]int {
	m := map[int]int{}
	r := [][]int{}
	for _, a := range nums {
		m[a]++
	}
	for k1, _ := range m {
		for k2, _ := range m {
			k3 := -(k1 + k2)
			if k1 < k2 && k2 < k3 && m[k3] > 0 {
				r = append(r, []int{k1, k2, k3})
			}
			if k1 == k2 && k2 < k3 && m[k3] > 0 && m[k2] > 1 {
				r = append(r, []int{k1, k2, k3})
			}
			if k1 < k2 && k2 == k3 && m[k3] > 1 {
				r = append(r, []int{k1, k2, k3})
			}
			if k1 == k2 && k2 == k3 && m[k3] > 2 {
				r = append(r, []int{k1, k2, k3})
			}
		}
	}
	return r
}
