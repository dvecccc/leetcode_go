package main

func middleNode(head *ListNode) *ListNode {
	var fast, low = head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		low = low.Next
	}
	return low
}
