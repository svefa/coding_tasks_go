package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sendPair(queue chan struct {
	*TreeNode
	int
}, n *TreeNode, i int) {
	queue <- struct {
		*TreeNode
		int
	}{n, i}
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	order := map[int][]int{}
	queue := make(chan struct {
		*TreeNode
		int
	})
	//go
	go sendPair(queue, root, 0)
	//  que := [] *TreeNode, int { root,0 }
	mi, ma := 0, 0
	//for len(que) > 0{
	for {
		// n, o = que[0]
		// que = que[1:]
		pair, ok := <-queue
		fmt.Println(pair, ok)
		if !ok {
			break
		}
		n, o := pair.TreeNode, pair.int
		if o < mi {
			mi = o
		}
		if o > ma {
			ma = o
		}

		order[o] = append(order[o], n.Val)
		if n.Left != nil {
			//  que = append(que, n.Left, o-1)
			//    queue <- struct {*TreeNode; int}{n.Left, o-1}
			//go
			sendPair(queue, n.Left, o-1)
		}
		if n.Right != nil {
			//  que = append(que, n.Right, o+1)
			//   queue <- struct {*TreeNode; int}{n.Right, o+1}
			//go
			sendPair(queue, n.Right, o+1)
		}

	}

	res := [][]int{}

	for i := mi; i <= ma; i++ {
		res = append(res, order[i])
	}

	return res

}

func main() {
	//[3,9,8,4,0,1,7,null,null,null,2,5]
	//[[4],[9,5],[3,0,1],[8,2],[7]]
	root := &TreeNode{Val: 3}
	res := verticalOrder(root)
	fmt.Println(res)
}
