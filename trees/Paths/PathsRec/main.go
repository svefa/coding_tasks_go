package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprint(root.Val)}
	}
	x := append(binaryTreePaths(root.Left), binaryTreePaths(root.Right)...)
	res := []string{}
	for _, s := range x {
		res = append(res, fmt.Sprint(root.Val)+"->"+s)
	}
	return res

}
func main() {
	//   1
	// 2   3
	//  5
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	fmt.Println(root, root.Left, root.Right)

	x := binaryTreePaths(root)
	fmt.Println(fmt.Sprint(root.Val), x)
	// ["1->2->5","1->3"]
}
