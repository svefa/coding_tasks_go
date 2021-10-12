package main

import "fmt"

func nextPermutation(nums []int) {
	// 123 132 213 231 312 321
	// 1234 1243 1324 1342 1423 1432
	// 2134 2143 2314 2341 2413 2431
	// 3 ....
	// 4 ....
	// we find the first element from right decreasing:         a < b > c > d > e
	// place  in this small element into increasing sequence:   b > c > d > a > e
	// NEXT ELEMENT will be the first in thee sequesnce d:
	//  other elements in increasing order:                     d ; e < a < c < b
	// algo:  reverse                                           a ;  e < d  < c < b  // e d c b a
	// algo swap a and d                                        d; e < a < c < b

	// 132  1<3>2
	/// [5,4,7,5,3,2]
	i := len(nums) - 1
	// i=2
	for i > 0 && nums[i-1] >= nums[i] {
		// nums[1]  > nums[2]
		i--
	}
	fmt.Println(i, nums)
	// i = 1
	// j = 1 j < 1 j++
	/// i= 2  nums[1]=5  nums[2]=7  >= nums[3]=5
	// j = 2 j <= 2 j++
	for j := 0; j <= (len(nums)-i-1)/2; j++ {
		nums[i+j], nums[len(nums)-j-1] = nums[len(nums)-j-1], nums[j+i]
		fmt.Println(i+j, nums)
	}
	fmt.Println(nums)
	// nums 123
	if i == 0 {
		return
	}
	//  nums[i-1] < nums[i]
	for j := i; j < len(nums); j++ {
		if nums[i-1] < nums[j] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			fmt.Println(nums)
			// 213 break
			break
		}
	}
}

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
