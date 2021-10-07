package main

import "fmt"

// 116. Populating Next Right Pointers in Each Node
/** You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
 * The binary tree has the following definition:  Node
 * Populate each next pointer to point to its next right node.
 * If there is no next right node, the next pointer should be set to
	1
  2    3
4  5  6  7

1 2 3 4 5 6 7
- 3 - 5 6 7 -
 **/

/* Definition for a Node. */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root != nil {
		h := root
		for h.Left != nil {
			head := h
			for head.Next != nil {
				fmt.Println(h.Val, head.Val)
				head.Left.Next = head.Right
				head.Right.Next = head.Next.Left
				head = head.Next
			}
			head.Left.Next = head.Right
			h = h.Left
		}
	}
	return root
}

/************************************* */

func main() {
	// Input:
	root := []int{1, 2, 3, 4, 5, 6, 7}
	// Output: [1,#,2,3,#,4,5,6,7,#]
	nodes := make([]Node, len(root))
	for i, _ := range nodes {
		nodes[i].Val = root[i]
		if 2*i+2 < len(root) {
			nodes[i].Left = &nodes[2*i+1]
			nodes[i].Right = &nodes[2*i+2]
		}
		fmt.Println(i, nodes[i])
	}

	connect(&nodes[0])
	fmt.Println(nodes)
}
