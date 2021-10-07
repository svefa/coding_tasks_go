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

type edge struct {
	i int
	j int
	d int
}

//	[0,0}, {2,2}, {3,10}, {5,2}, {7,0]]
func minCostConnectPoints(points [][]int) int {
	n := len(points)

	v := make(map[int]bool, n)
	v[0] = true

	hq := make([]edge, 0, n*n)
	for j := 1; j < n; j++ {
		v[j] = false
		heappush(&hq, edge{i: 0, j: j, d: abs(points[0][0]-points[j][0]) + abs(points[0][1]-points[j][1])})
		// 0, 1: 4
		// 0, 4: 7
		// 0, 2: 13
		// 0, 3: 7
	}
	s := 0

	for cnt := 0; cnt < n-1; cnt++ {
		e := heappop(&hq)
		// {0, 1, }{        //  1,3: 3    // 3,4 4   //  1,2: 9
		// v {1, 0,0,0,0}
		for v[e.j] {
			e = heappop(&hq)
		}

		s += e.d
		// s = 4  // s = 7  // s = 11 // s =20
		j := e.j
		// j = 1
		v[j] = true
		//	 =11000             // v=11010
		for i := 0; i < n; i++ {
			if !v[i] {
				// i =  2,3,4   //i=2,4
				heappush(&hq, edge{i: j, j: i, d: abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
				// 3,2 : 10
				// 3,4 :7
			}
		}
		// - {0, 14}
		// - 1,3: 3
		// 3, 4: 4
		// 0, 3:
		// 0,4 :
		// 1,2: 9
		// 1,4: 7
		// 3,2 : 10
		// 0, 2: 13
		// 4, 3, 14
	}
	return s
}

/*****************************************/

func heappush(pq *[]edge, e edge) {
	q := *pq
	q = append(q, e)

	i, j := len(q)-1, len(q)/2-1 // l=3 0,1 ; l =2 1,0

	for j >= 0 && q[j].d > q[i].d {
		q[j], q[i] = q[i], q[j]
		i, j = j, (j-1)/2 // ? -1/2 == 0
	}

	*pq = q
}

func heappop(pq *[]edge) edge {
	q := *pq
	n := len(q)
	e := q[0]
	q[0] = q[n-1]
	q = q[:n-1]
	i := 0

	for i < n-1 {
		if 2*i+2 < n-1 && (q[i].d > q[2*i+2].d || q[i].d > q[2*i+1].d) {
			if q[2*i+2].d < q[2*i+1].d {
				q[2*i+2], q[i] = q[i], q[2*i+2]
				i = 2*i + 2
			} else {
				q[2*i+1], q[i] = q[i], q[2*i+1]
				i = 2*i + 1
			}

		} else if 2*i+1 < n-1 && q[i].d > q[2*i+1].d {

			q[2*i+1], q[i] = q[i], q[2*i+1]
			i = 2*i + 1
		} else {
			break
		}
	}
	*pq = q
	return e
}

/********************************************/

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
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
