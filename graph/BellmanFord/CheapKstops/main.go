package main

/* dynamic programming 787 */

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// n=3 [[0,1,100],[1,2,100],[0,2,500]] src=0 dst=2 k =1
	dp := make([]int, n)
	tmp := make([]int, n)

	g := make(map[int]map[int]int, n)
	INF := 1 << 32

	for i := 0; i < n; i++ {
		g[i] = make(map[int]int)
		dp[i] = INF
		tmp[i] = INF
	}
	// g[0][1] = 100, g[1][2] =100 g[0][2]=500
	for _, f := range flights {
		g[f[0]][f[1]] = f[2]
	}
	dp[src] = 0
	tmp[src] = 0
	// dp=[0 inf inf]
	// k= 1
	for k >= 0 {
		// i=0
		// i=1
		for i := 0; i < n; i++ {
			// (i=0,j=1,v=100) ; (i=0,j=2, v=500)
			// (i=1,j=2; v=100)
			for j, v := range g[i] {
				// dp=[0 inf inf]
				// dp=[0 100 500]
				if dp[i]+v < tmp[j] {
					tmp[j] = dp[i] + v

				}
			}
			// tmp = [0 100 500]
			// tmp = [0 100 200]
		}
		k--
		// k=0
		for i := 0; i < n; i++ {
			dp[i] = tmp[i]
		}
	}

	if tmp[dst] == INF {
		return -1
	}
	return tmp[dst]
}

/**********************/
/*
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Output: 200
*/
/* Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Output: 500
*/
