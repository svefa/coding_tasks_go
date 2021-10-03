package mainC

import "fmt"

/**
 * Definition for a Node.
**/
type Node struct {
	Val       int
	Neighbors []*Node
}

var n2c map[*Node]*Node

func cloneGraph(node *Node) *Node {
	if n2c == nil {
		n2c = make(map[*Node]*Node)
	}
	if c, ok := n2c[node]; ok {
		return c
	}
	cloneNode := Node{Val: node.Val}
	n2c[node] = &cloneNode
	if len(node.Neighbors) > 0 {
		cloneNode.Neighbors = make([]*Node, len(node.Neighbors))
		for i, nei := range node.Neighbors {
			cloneNode.Neighbors[i] = cloneGraph(nei)
		}
	}
	return &cloneNode
}

func main() {
	al := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	n := make([]Node, len(al))
	for i, neis := range al {
		n[i].Val = i + 1
		n[i].Neighbors = make([]*Node, len(neis))
		for j, nei := range neis {
			n[i].Neighbors[j] = &n[nei-1]
		}
	}

	clone := cloneGraph(&n[0])
	fmt.Println(clone, *clone.Neighbors[0], *clone.Neighbors[1])
}
