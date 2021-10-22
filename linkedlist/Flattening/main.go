package main

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// 1---2---3---4---5---6--NULL
//         |
//         7---8---9---10--NULL
//             |
//             11--12--NULL
// 1
// 2
// 3
// 4 7
// 4 8
// 4 9 11
// 4 9 12
// 4 9
//

// seq 1 2 3 (4) 7 8 (9) 11 12
// (9) 10
// (4) 5 6
// what we nedd change   12 -> 9 , 8 -> 11

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	var p *Node
	for len(stack) > 0 {

		l := len(stack)
		n := stack[l-1]
		if p != nil && (p.Child != nil || p.Next == nil) {
			n.Prev = p
			p.Next = n
			p.Child = nil
		}
		stack = stack[:l-1]
		if n.Next != nil {
			stack = append(stack, n.Next)
		}
		if n.Child != nil {
			stack = append(stack, n.Child)
		}
		p = n
	}

	return root
}
