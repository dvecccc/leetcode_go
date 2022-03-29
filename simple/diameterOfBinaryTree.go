package main

import (
	"fmt"
	"math"
)

/*func diameterOfBinaryTree(root *TreeNode) int {
	var maxCount = math.MinInt64
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		leftMax := maxDepthOfTree(root.Left)
		rightMax := maxDepthOfTree(root.Right)
		temp := leftMax + rightMax
		maxCount = max(temp, maxCount)
		f(root.Left)
		f(root.Right)
	}
	f(root)
	return maxCount
}

func maxDepthOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthOfTree(root.Left)
	right := maxDepthOfTree(root.Right)
	return max(left, right) + 1
}*/

func diameterOfBinaryTree(root *TreeNode) int {
	var MaxCount = math.MinInt64
	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := f(root.Left)
		rightMax := f(root.Right)
		temp := leftMax + rightMax
		MaxCount = max(temp, MaxCount)
		return max(leftMax, rightMax) + 1
	}
	f(root)
 	return MaxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	/*t3l := &TreeNode{9, nil, nil}
	t2r := &TreeNode{5, t3l, nil}
	t2l := &TreeNode{4, nil, nil}
	t1r := &TreeNode{3, nil, nil}
	t1l := &TreeNode{2, t2l, t2r}*/
	t1l := &TreeNode{2, nil, nil}
	t1 := &TreeNode{1, t1l, nil}
	//res1 := maxDepthOfTree(t1)
	//res2 := maxDepthOfTree(t1l)
	//fmt.Println("res1:", res1, "res2:", res2)
	fmt.Println(diameterOfBinaryTree(t1))
}