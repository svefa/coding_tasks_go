package main

type Tree struct {
	Left  *Tree
	Right *Tree
	Val   int
}

func TreeRange(root, min, max *Tree) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if !TreeRange(root.Left, min, root) {
		return false
	}
	if !TreeRange(root.Right, root, max) {
		return false
	}
	return true

}

func isValidBST(root *Tree) bool {
	isValid := TreeRange(root, nil, nil)
	return isValid
}
