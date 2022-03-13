package main

import "fmt"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var f func([]int, []int) *TreeNode
	f = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		var root = &TreeNode{Val: preorder[0]}
		if len(preorder) == 1 {
			return root
		}
		i1, i2 := inorderDivide(inorder, preorder[0])
		//fmt.Println("i1: ", i1)
		//fmt.Println("i2: ", i2)
		p1, p2 := preorderDivide(preorder,len(i1))
		//fmt.Println("p1: ", p1)
		//fmt.Println("p2: ", p2)
		root.Left = f(p1, i1)
		root.Right = f(p2, i2)
		return root
	}
	root := f(preorder, inorder)
	return root
}

func inorderDivide(inorder []int, flag int) ([]int, []int) {
	if len(inorder) <= 1 {
		return nil, nil
	}
	var d1, d2 = make([]int, 0), make([]int, 0)
	for index, value := range inorder {
		if value == flag {
			d1 = inorder[:index]
			d2 = inorder[index + 1:]
			return d1, d2
		}
	}
	return d1, d2
}

func preorderDivide(preorder []int, length int) ([]int, []int) {
	if len(preorder) <= 1 {
		return nil, nil
	}
	var d1, d2 = make([]int, 0), make([]int, 0)
	d1 = preorder[1:length + 1]
	d2 = preorder[length + 1:]
	return d1, d2
}

func traverse(root *TreeNode)  {
	if root == nil {
		return
	}

	traverse(root.Left)
	fmt.Printf("%d  ", root.Val)
	traverse(root.Right)
}

func main()  {
	preorder := []int{3,9,11,12,13,20,15,6,7}
	inorder := []int{11,9,13,12,3,15,6,20,7}
	/*r3, _ :=inorderDivide(inorder, preorder[0])
	r1, r2 := preorderDivide(preorder, len(r3))
	fmt.Println(r1)
	fmt.Println(r2)*/
	root := buildTree(preorder, inorder)
	traverse(root)
}
