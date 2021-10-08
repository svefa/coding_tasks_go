func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 4
	// [[1,0],[1,2],[1,3]]
	in := make([]int, n)
	g := make([][]int, n)
	q := []int{}
	nq := []int{}

	for _, e := range edges {
		i, j := e[0], e[1]
		g[i] = append(g[i], j)
		g[j] = append(g[j], i)
		in[i]++
		in[j]++
	}

	for i, _ := range in {
		if in[i] == 1 {
			q = append(q, i)
		}
	}
	// g[0] = [1] g[1] =[0,2,3] g[2]=[1] g[3]=[1]
	// in = [1311]

	for len(q) > 0 {
		nq = []int{}
		// q=[0,2,3] nq=[] in=[]
		for _, i := range q {
			for _, j := range g[i] {
				if in[j] > 0 {
					in[j]--
					in[i]--
					if in[j] == 1 {
						nq = append(nq, j)
					}
					//in: 0211 nq:[]
					//in: 0101 nq:[1]
					//in: 0000 nq:[1]
				}
			}
		}
		q, nq = nq, q
		// q=[1] nq=[023]
	}
	return nq

}