package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// head = 1->2->2.->1.->nil
	//~ 1->2->1->nil
	if head == nil || head.Next == nil {
		return true
	}
	res := true
	mid, node := head, head
	// mid = 1->2... node =2->2. ....
	//~ mid = 1->2  node = 2->1
	for node != nil && node.Next != nil {
		mid, node = mid.Next, node.Next.Next
		// mid=2 node=1. ->nil
		//~ mid=2->1 node= nil
	}
	// mid=2. node =1->2->
	node, mid = mid, nil
	for node != nil {
		node.Next, node, mid = mid, node.Next, node
	}
	node = head
	node2 := mid

	for node2 != nil {
		if node2.Val != node.Val {
			res = false
			break
		}
		node2, node = node2.Next, node.Next
	}

	node, mid = mid, nil
	// node = 1->2 mid=2->nil
	for node != nil {
		node.Next, node, mid = mid, node.Next, node
	}
	return res
}

func main() {
	a := []int{1, 2}
	var l *ListNode

	for i := len(a) - 1; i >= 0; i-- {
		l = &ListNode{Val: a[i], Next: l}
	}

	b := isPalindrome(l)
	//printLL(l)
	println(b)

}
