package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	child := flatten(root.Child)
	next := flatten(root.Next)
	if child != nil {
		root.Next = child
		child.Prev = root
		for child.Next != nil {
			child = child.Next
		}
		if next != nil {
			child.Next = next
			next.Prev = child
		}
	}
	root.Child = nil
	return root
}
