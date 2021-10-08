package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	// [[1,0],[2,0],[3,1],[3,2]]
	n := numCourses
	g := make([][]int, n)
	in := make([]int, n)
	q := []int{}
	for _, c := range prerequisites {
		//g[0]=[1,2] g[1]=[3] g[2]=[3]
		// 0->1->3  0->2->3
		g[c[1]] = append(g[c[1]], c[0])
		in[c[0]]++
		// in = [0112]
	}

	for i := 0; i < n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	r := []int{}
	// q=[1]
	/// q=[1,2]
	for len(q) > 0 {
		c := q[0]
		// c=0
		/// c=1
		//// c=2
		///// c=3
		r = append(r, c)
		//res=[0]
		///res=[0,1]
		//// res=[0,1,2]
		///// res=[0,1,2,3]
		q = q[1:]
		// q=[]
		/// q=[2]
		//// q=[]
		for _, i := range g[c] {
			// i=1; i=2
			/// i=3
			//// i= 3
			in[i]--
			// in=[0002]
			/// in =0001
			//// in = 0000
			if in[i] == 0 {
				q = append(q, i)
				// q=[1,2]
				/// q=[2]
				//// q=[3]
			}
		}
	}
	if len(r) == n {
		return r
	}
	return []int{}
}

/**************************/

func main() {
	//Input: numCourses = 2, prerequisites = [[1,0]]
	//Output: [0,1]
	//Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
	//Output: [0,2,1,3]
	// Input: numCourses = 1, prerequisites = []
	// Output: [0]

}
