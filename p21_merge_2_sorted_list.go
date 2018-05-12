package main


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := ListNode{0,nil}
	p := &head

	for l1 != nil && l2 != nil {
		curNode := l1
		if curNode.Val > l2.Val {
			curNode = l2
			l2 = l2.Next
		}else{
			l1 = l1.Next
		}
		p.Next = curNode
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	} else if l2 != nil{
		p.Next = l2
	}

	return head.Next
}