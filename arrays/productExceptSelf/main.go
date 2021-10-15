package main

func productExceptSelf(nums []int) []int {
	// nums =[1,2,3]
	n := len(nums)
	// n= 3
	res := make([]int, n)
	// res = [0,0,0]
	res[0] = 1
	// res =[1,0,0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
		// res[1]=res[0]*nums[0] = 1
		// res[2]= 1*2=2
	}
	//res = [1,1,2]
	k := 1
	for i := n - 1; i >= 0; i-- {

		res[i] *= k
		k *= nums[i]
		// res[2] = 2 k = 3
		// res[1] = 3 k = 6
		// res[0] = 6 k = 6
	}
	return res
}
