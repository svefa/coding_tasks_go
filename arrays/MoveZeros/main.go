package main

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

/*func moveZeroes(nums []int)  {
    // [0,1,0,3,12]
    n := len(nums)
    // n=5
    i := 0
    for i < n && nums[i] != 0 { i++}
    j := i
    for j < n {
        for j < n && nums[j] == 0 { j++}
        // j= 1 // j = 3
        k := j
        for k < n && nums[k] != 0 { k++}
        // k= 2 // k=5
        copy(nums[i:],nums[j:k])
        // nums[0:2] <-nums[1:2]  1,1,0,3,12
        // nums[1:3] <- nums[3:5] 1,3,12,3,12
        i = k - j + i
        j = k
        // i = 2-1+0= 1
        // j = 2
    }
    copy(nums[i:n], make([]int,n-i))

}*/
