package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	q := []*Node{node}
	n2c := map[*Node]*Node{}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if _, ok := n2c[n]; ok {
			continue
		}
		c := &Node{Val: n.Val, Neighbors: []*Node{}}
		n2c[n] = c
		for _, nei := range n.Neighbors {
			if cn, ok := n2c[nei]; ok {
				c.Neighbors = append(c.Neighbors, cn)
				cn.Neighbors = append(cn.Neighbors, c)
			} else {
				q = append(q, nei)
			}
		}
	}

	return n2c[node]
}
