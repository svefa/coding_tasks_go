package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	stack := []*TreeNode{}
	node := root
	for node != nil {
		if node.Left != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			node.Right = node.Left
			node.Left = nil
		} else if node.Right == nil && len(stack) > 0 {
			node.Right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		node = node.Right
	}

}
