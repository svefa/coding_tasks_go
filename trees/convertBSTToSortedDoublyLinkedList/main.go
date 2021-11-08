package main

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func rec(root *Node) (*Node, *Node) {

	l1, l2, r1, r2 := root, root, root, root
	if root.Left != nil {
		l1, l2 = rec(root.Left)
		root.Left = l2
		l2.Right = root

	}
	if root.Right != nil {
		r1, r2 = rec(root.Right)
		root.Right = r1
		r1.Left = root
	}
	return l1, r2

}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	n1, n2 := rec(root)
	n2.Right = n1
	n1.Left = n2
	return n1
}
