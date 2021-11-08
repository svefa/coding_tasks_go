package main

func intersection(nums1 []int, nums2 []int) []int {
	n1, n2 := short(nums1, nums2)

	m := make(map[int]struct{})
	// map of the first slice
	for _, v := range n1 {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	var res []int
	for _, k := range n2 {
		if _, ok := m[k]; ok {
			delete(m, k)
			res = append(res, k)
		}
	}

	return res
}

func short(nums1, nums2 []int) ([]int, []int) {
	if len(nums1) < len(nums2) {
		return nums1, nums2
	}

	return nums2, nums1
}
