package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var v, res int
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if v == k {
			return
		}
		if root == nil {
			return
		}
		f(root.Left)
		v++
		if v == k {
			res = root.Val
		}
		f(root.Right)
		return
	}
	f(root)
	return res
}

func main()  {

}
