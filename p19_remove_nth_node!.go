package main

//倒数第一个时有特殊情况
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	lastNode := head
//	for i := 1; i < n; i++ {
//		lastNode = lastNode.Next
//	}
//
//	lastNthNode := head
//	prevNode := head
//
//	for ; lastNode.Next != nil; {
//		if prevNode != lastNthNode{
//			prevNode = prevNode.Next
//		}
//		lastNthNode = lastNthNode.Next
//		lastNode = lastNode.Next
//	}
//
//	//当倒数第n个节点就是第一个节点时，需要特殊处理
//	if lastNthNode == head {
//		return lastNthNode.Next
//	}
//
//	prevNode.Next = lastNthNode.Next
//	lastNthNode.Next = nil
//
//	return head
//}

//找倒数第n个并删除之，实际上就是找倒数第n个节点的前一个节点
//可以加上一个头结点简化运算
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{0, head}
	left := &dummy
	right := &dummy

	for i := 0; i < n; i++ {
		right = right.Next
	}

	for ; right.Next != nil; {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}