package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flattenTail(root *Node) *Node {

	l, r := root.Child, root.Next
	if l != nil {
		root.Child = nil
		root.Next, l.Prev = l, root
		l = flattenTail(l)
		if r != nil {
			l.Next, r.Prev = r, l
		}
	}
	if r != nil {
		r = flattenTail(r)
	}

	if r != nil {
		return r
	} else if l != nil {
		return l
	} else {
		return root
	}

}

func flatten(root *Node) *Node {
	if root != nil {
		flattenTail(root)
	}
	return root
}
