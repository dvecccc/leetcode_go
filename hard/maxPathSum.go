package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
}


func maxPathSum(root *TreeNode) int {
	var res = math.MinInt64

	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var f func(node *TreeNode) int
	f = func (root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := Max(0, f(root.Left))
		right := Max(0, f(root.Right))
		res = Max(res, left + right + root.Val)
		return Max(left, right) + root.Val
	}
	f(root)
	return res
}


func traverse(root *TreeNode)  {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	traverse(root.Left)
	traverse(root.Right)
}

func main()  {
	secondL1 := &TreeNode{Val: -10}
	secondR1 := &TreeNode{Val: 5}
	secondL2 := &TreeNode{Val: 6}
	secondR2 := &TreeNode{Val: 9}
	firstL := &TreeNode{Val: 13, Left: secondL1, Right: secondR1}
	firstR := &TreeNode{Val: 11,Left: secondL2, Right: secondR2}
	root := &TreeNode{Val: 12, Left: firstL, Right: firstR}
	fmt.Println(maxPathSum(root))
}
