package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	step := 1
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			tmp := queue[0]
			if tmp.Left == nil && tmp.Right == nil{
				step++
				return step
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			queue = queue[1:]
		}
		step++
	}
	return step
}

func main()  {

	l2 := &TreeNode{4, nil, nil}
	r2 := &TreeNode{5, nil, nil}
	l1 := &TreeNode{2, nil, r2}
	r1 := &TreeNode{3, l2, nil}
	root := &TreeNode{1, l1, r1}
	res := minDepth(root)
	fmt.Println(res)
}
