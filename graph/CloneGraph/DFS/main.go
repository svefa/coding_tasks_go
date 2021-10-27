package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

var n2c map[*Node]*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
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
