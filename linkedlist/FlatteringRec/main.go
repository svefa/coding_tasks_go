package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flattenTail(root *Node) *Node {
	tail := root
	for root != nil {
		tail = root
		child, next := root.Child, root.Next
		root.Child = nil
		if child != nil {
			tail = flattenTail(child)
			root.Next, child.Prev = child, root
			if next != nil {
				tail.Next, next.Prev = next, tail
			}
		}
		root = next
	}
	return tail
}

func flatten(root *Node) *Node {
	flattenTail(root)
	return root
}
