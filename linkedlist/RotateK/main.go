package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	l := 1
	for cur.Next != nil {
		l++
		cur = cur.Next
	}

	k = l - k%l
	if k == 0 {
		return head
	}
	// cur = 2->1 head =1->2 k=1
	cur.Next = head
	cur = head
	for k > 1 {
		cur = cur.Next
		k--

	}
	head = cur.Next
	cur.Next = nil
	return head
}
