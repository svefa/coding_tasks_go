package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//   1
//  2  3
// (1,2,3)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == q && root == p {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	//I: lca(2,3,2) -> lca(nil,2,3) -> nil
	if l != nil {
		return l
	}
	r := lowestCommonAncestor(root.Right, p, q)
	//II: lca(3,2,3) -> lca(nil, 2,3 ) -> nil
	if r != nil {
		return r
	}
	if p == q {
		return nil
	}
	//III & IV: lca(1, 2,2) && lca(1,3,3)
	if lowestCommonAncestor(root, p, p) != nil && lowestCommonAncestor(root, q, q) != nil {
		return root
	}
	return nil
}
