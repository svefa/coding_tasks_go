package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				l++
				for l < r-1 && nums[l] == nums[l-1] {
					l++
				}
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else if n1+n2+n3 > 0 {
				r--
			} else {
				l++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	Output := threeSum(nums)
	fmt.Println(Output)
	// [[-1,-1,2],[-1,0,1]]
}
