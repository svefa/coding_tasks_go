package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestPath(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	l := longestPath(root.Left, diameter)
	r := longestPath(root.Right, diameter)
	if l+r > *diameter {
		*diameter = l + r
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	longestPath(root, &d)
	return d
}
