package main

func detectCycle(head *ListNode) *ListNode {
	var fast, low = head, head

	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	low = head
	for low != fast {
		low = low.Next
		fast = fast.Next
	}
	return low
}

func main()  {

}
