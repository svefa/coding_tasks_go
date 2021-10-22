package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// [[7,null],[13,0],[11,4],[10,2],[1,0]]
	if head == nil {
		return nil
	}

	c := head

	// c = 7
	for c != nil {
		c.Next = &Node{Val: c.Val, Random: c.Random, Next: c.Next}
		c = c.Next.Next
	}

	c = head
	c2 := head.Next
	for c != nil {
		if c.Next.Random != nil {
			c.Next.Random = c.Next.Random.Next
		}
		c = c.Next.Next
	}
	c = head
	for c != nil {
		if c.Next.Next != nil {
			c.Next, c.Next.Next = c.Next.Next, c.Next.Next.Next
		} else {
			c.Next = nil
		}
		c = c.Next
	}
	return c2

}

/*
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//	head2 := &Node{Val: head.Val, Next: head.Next, Random: head.Random}
	//	head.Next = head2
	c := head

	// c = 7
	for c != nil {
		c.Next = &Node{Val: c.Val, Random: c.Random, Next: c.Next}
		c = c.Next.Next
	}
	printLL(head)

	c = head
	x := head.Next
	for c != nil {
		p := c.Next
		if c.Next.Random != nil {
			c.Next.Random = c.Next.Random.Next
		}
		c.Next = p.Next
		if p.Next != nil {
			p.Next = p.Next.Next
		}
		c = c.Next
	}
	printLL(head)

	return x
}
*/
func printLL(n *Node) {
	fmt.Println()
	for n != nil {
		a, b := -1, -1
		if n.Next != nil {
			a = n.Next.Val
		}
		if n.Random != nil {
			b = n.Random.Val
		}
		fmt.Println(n.Val, a, b, *n)
		n = n.Next
	}
}

func main() {
	// Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
	// Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	a := [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	fmt.Println(a)
	n := make([]Node, len(a))
	for i, x := range a {
		fmt.Println(i, x)
		n[i].Val = x[0]
		if i < len(a)-1 {
			n[i].Next = &n[i+1]
		}
		if x[1] != -1 {
			n[i].Random = &n[x[1]]
		}
	}
	printLL(&n[0])

	r := copyRandomList(&n[0])
	printLL(r)
}
