package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultHead := result
	var sum,left,addition int
	for ; l1 != nil && l2 != nil; {
		sum = l1.Val + l2.Val + addition
		left = sum%10
		result.Val = left
		addition = sum/10
		if l1.Next != nil && l2.Next != nil {
			result.Next = &ListNode{}
			result = result.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	//统一使用l1处理剩余的数
	if l2 != nil {
		l1 = l2
	}
	for ; l1 != nil; l1 = l1.Next {
		result.Next = &ListNode{}
		result = result.Next
		sum = l1.Val + addition
		result.Val = sum % 10
		addition = sum / 10
	}
	//位数相同但有进位的情况
	if l1 == nil && addition != 0{
		result.Next = &ListNode{Val:addition}
	}
	return resultHead
}


