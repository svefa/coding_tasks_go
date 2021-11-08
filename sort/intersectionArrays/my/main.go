package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	res := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	i1, i2 := 0, 0
	curr := nums1[0] - 1
	if nums2[0] < curr {
		curr = nums2[0] - 1
	}
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] == curr {
			i1++
		} else if nums1[i1] == curr {
			i2++
		} else if nums1[i1] > nums2[i2] {
			i2++
		} else if nums2[i2] > nums1[i1] {
			i1++
		} else { //if nums1[i1] == nums[i2]
			res = append(res, nums1[i1])
			curr = nums1[i1]
			i1++
			i2++
		}
	}
	return res

}
