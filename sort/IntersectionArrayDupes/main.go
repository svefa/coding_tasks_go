package main

func intersect(nums1 []int, nums2 []int) (res []int) {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	cnt := map[int]int{}
	for _, i := range nums1 {
		cnt[i]++
	}
	for _, i := range nums2 {
		if cnt[i] > 0 {
			res = append(res, i)
			cnt[i]--
		}
	}
	return
}
