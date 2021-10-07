package main

// 1584. Min Cost to Connect All Points
//  You are given an array points representing integer coordinates
// of some points on a 2D-plane, where points[i] = [xi, yi].
//The cost of connecting two points [xi, yi] and [xj, yj] is the
// manhattan distance between them: |xi - xj| + |yi - yj|,
// where |val| denotes the absolute value of val.
//Return the minimum cost to make all points connected.
//All points are connected if there is exactly one simple path
// between any two points.

//	[0,0}, {2,2}, {3,10}, {5,2}, {7,0]]
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	s := 0
	// best distance to visited
	d := make([]int, n)
	// visited
	v := make([]bool, n)
	for i := 0; i < n; i++ {
		v[i] = false
		d[i] = 1 << 32
	}
	d[0] = 0
	k := 0
	for i := 0; i < n; i++ {
		// update dist for unvisited
		v[k] = true
		s += d[k]
		for j := 0; j < n; j++ {
			if v[j] == false {
				l := abs(points[k][0]-points[j][0]) + abs(points[k][1]-points[j][1])

				if d[j] > l {
					d[j] = l
				}
			}
		}

		// choose nearest
		k = 0
		for j := 0; j < n; j++ {
			if v[j] == false {
				if k == 0 || d[j] < d[k] {
					k = j
				}
			}
		}

		// fmt.Println(i, k, d[k], d)

	}

	return s
}

/********************************************/

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

/*****************************************************/
func main() {
	//points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	//s := minCostConnectPoints(points)
	//println(s)  // 20

	points2 := [][]int{{-8, 14}, {16, -18}, {-19, -13}, {-18, 19}, {20, 20}, {13, -20}, {-15, 9}, {-4, -8}}
	s2 := minCostConnectPoints(points2)
	println(s2) // 139
}
