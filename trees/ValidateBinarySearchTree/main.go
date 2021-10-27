package main

type Tree struct {
	Left  *Tree
	Right *Tree
	Val   int
}

func TreeRange(root *Tree) (bool, *Tree, *Tree) {
	l, r := root, root
	if root.Left != nil {
		res, low, high := TreeRange(root.Left)
		if !res || root.Val <= high.Val {
			return false, nil, nil
		}
		l = low
	}
	if root.Right != nil {
		res, low, high := TreeRange(root.Right)
		if !res || root.Val >= low.Val {
			return false, nil, nil
		}
		r = high
	}
	return true, l, r
}

func isValidBST(root *Tree) bool {
	isValid, _, _ := TreeRange(root)
	return isValid
}
