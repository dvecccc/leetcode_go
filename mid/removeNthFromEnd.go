package main

import "fmt"

type ListNode struct {
	Val   int
	Next  *ListNode
}

//simple
/*func removeNthFromEnd(head *listNode, n int) *listNode {
	var pre, root2 = head, head
	var length int
	for head != nil {
		length++
		head = head.next
	}
	if n == length {
		return pre.next
	}
	for length - n-1 > 0 {
		pre = pre.next
		length--
	}
	pre.next = pre.next.next
	return root2
}*/

//two ptr
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, low = head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		low = low.Next
	}
	low.Next = low.Next.Next
	return head
}

func main()  {
	//var l5 = &ListNode{6, nil}
	//var l4 = &ListNode{5, l5}
	//var l3 = &ListNode{4, l4}
	//var l2 = &ListNode{3, l3}
	var l1 = &ListNode{2, nil}
	var head = &ListNode{1, l1}
	res := removeNthFromEnd(head,2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
