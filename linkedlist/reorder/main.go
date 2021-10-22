package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	v1, v2 := head, head
	// find mid
	for v2 != nil && v2.Next != nil {
		println(v2, v2.Next)
		println(v1.Val, v2.Val)
		v1, v2 = v1.Next, v2.Next.Next
	}

	// 2, 3
	// reverse right
	// 1 -> 2 <- 3 -< 4
	//      v1   v2
	/// 1->2 <-3
	v2 = v1.Next
	v1.Next = nil
	//println(v1.Val)
	for v2 != nil {
		t := v2.Next
		v2.Next = v1
		v1 = v2
		v2 = t
	}
	// v1 points end
	v2 = v1
	v1 = head

	for v1 != nil && v2 != nil {
		t2, t1 := v2.Next, v1.Next
		println(t1, t2)

		v1.Next = v2
		if v2 != t1 {
			v2.Next = t1
		}
		v1, v2 = t1, t2
		// 1->4->2  // v1 = 2 v2=3
		// 1->4->2->3->nil
		/// 1->3 ->2
	}
	h := head
	for h != nil {
		println(h.Val)
		h = h.Next
	}

}

func main() {
	///Input: head = [1,2,3,4]
	//Output: [1,4,2,3]
	h := &ListNode{Val: 1}
	c := h
	for i := 2; i < 4; i++ {
		c.Next = &ListNode{Val: i}
		c = c.Next
	}
	reorderList(h)

}
