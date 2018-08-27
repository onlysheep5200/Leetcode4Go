package main

import "encoding/asn1"

/**
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	preHead := ListNode{0, head}
	p := head
	nodeCount := 0
	var tail *ListNode

	//计算链表长度，找到尾节点
	for p != nil{
		p = p.Next
		nodeCount++
		if p!=nil && p.Next == nil {
			tail = p
		}
	}

	k %= nodeCount
	if k == 0 {
		return head
	}
	before := k

	p = head
	beforeP := head
	preBeforeP := &preHead

	for p != nil {
		p = p.Next
		if before <= 0 {
			 beforeP = beforeP.Next
			 preBeforeP = preBeforeP.Next
		}
		before--
	}

	preBeforeP.Next = nil
	preHead.Next = beforeP
	tail.Next = head

	return preHead.Next
}
