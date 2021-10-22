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
//   1 2 3 7 8 11 12 null
//       4   9
//   9 10 null
//   4 5 6 null

// seq 1 2 3 (4) 7 8 (9) 11 12
// (9) 10
// (4) 5 6
// what we nedd change   12 -> 9 , 8 -> 11

func flatten(root *Node) *Node {
	stack := []*Node{}
	n := root

	for n != nil {
		if n.Child != nil {
			if n.Next != nil {
				stack = append(stack, n.Next)
			}
			n.Child.Prev = n
			n.Next = n.Child
			n.Child = nil
		}
		if n.Next == nil && len(stack) > 0 {
			n.Next = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n.Next.Prev = n
		}
		n = n.Next

	}

	return root
}
