package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSumRec(root *TreeNode, res *int) int {
	s := 0
	if root != nil {
		p := 0
		l := maxPathSumRec(root.Left, res)
		if s < l {
			s = l
		}
		if l > 0 {
			p += l
		}
		r := maxPathSumRec(root.Right, res)
		if s < r {
			s = r
		}
		if r > 0 {
			p += r
		}

		s += root.Val
		p += root.Val
		if *res < p {
			*res = p
		}
	}
	return s
}

func maxPathSum(root *TreeNode) int {
	var res int
	if root != nil {
		res = root.Val
	}
	maxPathSumRec(root, &res)
	return res
}
