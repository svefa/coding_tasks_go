package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	for root != nil {
		if root.Left != nil {
			rightMost := root.Left
			for rightMost.Right != nil {
				rightMost = rightMost.Right
			}
			rightMost.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}

		root = root.Right
	}

}
