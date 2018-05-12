package main

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k<=1{
		 return head
	}

	dummy := &ListNode{0, head}
	prev := dummy
	points := make([]*ListNode, k, k)
	start := head
	var p *ListNode

	for start != nil{
		for i := 0; i < k; i++ {
			if start == nil {
				goto end
			}
			points[i] = start
			start = start.Next
		}

		//exchange
		for i:= 0; i<k/2; i++ {
			other := k - i - 1

			//交换next
			p = points[i].Next
			points[i].Next = points[other].Next
			if points[other] != p {
				points[other].Next = p
			}else{
				points[other].Next = points[i]
			}

			//前序节点
			if i != 0 {
				prev = points[i-1]
			}
			prev.Next = points[other]
			if points[k-i-2] != points[i] {
				points[k-i-2].Next = points[i]
			}

			points[i],points[k-i-1] = points[k-i-1],points[i]
		}
		prev = points[k-1]
	}

	end : return dummy.Next
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	reverseKGroup(head,2)
}