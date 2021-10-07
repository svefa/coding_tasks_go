package main

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	l := []*Node{root}
	//res = append(res, []int{root.Val})

	for len(l) > 0 {
		r := make([]int, 0, len(l))
		for _, node := range l {
			r = append(r, node.Val)
		}
		res = append(res, r)
		nl := make([]*Node, 0, len(l))
		for _, node := range l {
			if len(node.Children) != 0 {
				nl = append(nl, node.Children...)
			}
		}
		l = nl
	}
	return res
}

func main() {
	r := Node{Val: 1}
	a := levelOrder(&r)
	fmt.Println(a)
}
