package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	tail := head
	head = nil
	for tail != nil {
		cut := tail
		tail = cut.Next
		cut.Next = head
		head = cut
	}
	return head

}
