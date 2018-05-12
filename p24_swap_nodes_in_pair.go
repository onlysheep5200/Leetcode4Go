package main


func swapPairs(head *ListNode) *ListNode {
	if head  == nil{
		return head
	}
	var prev,pPrev,next *ListNode

	dummy := &ListNode{0, head}
	prev = head
	pPrev = dummy
	next = head.Next

	for prev != nil && next != nil {
		pPrev.Next = next
		prev.Next = next.Next
		next.Next = prev
		prev,next = next,prev
		pPrev = next
		prev = next.Next
		//判空需注意
		if prev == nil {
			continue
		}
		next = prev.Next
	}

	return dummy.Next
}