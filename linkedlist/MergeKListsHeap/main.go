package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func heappush(heap []*ListNode, node *ListNode) []*ListNode {
	heap = append(heap, node)
	// 0, 1, 2
	i := len(heap) - 1
	for i >= 1 && heap[(i-1)/2].Val > heap[i].Val {
		heap[(i-1)/2], heap[i] = heap[i], heap[(i-1)/2]
		i = (i - 1) / 2
	}
	return heap
}

func heappop(heap []*ListNode) (*ListNode, []*ListNode) {
	node := heap[0]
	l := len(heap)
	heap[0] = heap[l-1]
	heap = heap[:l-1]
	i := 0
	for i < l-1 {
		j := i
		if 2*i+2 < l-1 && heap[j].Val > heap[2*i+2].Val {
			j = 2*i + 2
		}
		if 2*i+1 < l-1 && heap[j].Val > heap[2*i+1].Val {
			j = 2*i + 1
		}
		if i == j {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
	return node, heap
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := []*ListNode{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap = heappush(heap, lists[i])
		}
	}

	cur := &ListNode{}
	head, node := cur, cur

	for len(heap) > 0 {
		node, heap = heappop(heap)
		if node.Next != nil {
			heap = heappush(heap, node.Next)
		}
		cur.Next, cur = node, node
	}

	return head.Next
}
