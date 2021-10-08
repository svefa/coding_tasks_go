package main

import (
	"fmt"
	"math"
)

/* 787 bellman-ford dynamic programming */

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([]int, n)
	pred := make([]int, n)
	return bellmanFord(flights, src, dst, n, k, dist, pred)
}

func bellmanFord(flights [][]int, src, dst, n, k int, dist, pred []int) int {
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		pred[i] = -1
	}

	dist[src] = 0

	for i := 0; i <= k; i++ {
		temp := append([]int{}, dist...)
		for _, flight := range flights {
			u := flight[0]
			v := flight[1]
			w := flight[2]
			if dist[u] == math.MaxInt32 {
				continue
			} else if dist[u]+w < temp[v] {
				temp[v] = dist[u] + w
				pred[v] = u
			}
		}
		dist = temp

	}

	// for _, flight := range flights {
	//     u := flight[0]
	//     v := flight[1]
	//     w := flight[2]
	//     if dist[u] + w < dist[v] {
	//         return -1
	//     }
	// }
	fmt.Println(dist, pred)

	if dist[dst] == math.MaxInt32 {
		return -1
	}
	return dist[dst]
}

/**********************/
/*
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Output: 200
*/
/* Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Output: 500
*/
