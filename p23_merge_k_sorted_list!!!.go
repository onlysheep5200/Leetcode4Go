package main

import (
	"fmt"
	"container/heap"
)

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int{
	return len(h)
}

func (h ListNodeHeap) Less(i,j int) bool  {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	top := (*h)[len(*h) - 1]
	*h = (*h)[0:len(*h) - 1]
	return top
}


//使用堆来优化比较
//归并的思路，logk复杂度
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	head := &ListNode{0, nil}
	current := head

	nodeHeap := &ListNodeHeap{}
	heap.Init(nodeHeap)

	for i := 0; i < len(lists); i++ {
		//此处可能为空，需进行判空处理
		if lists[i] == nil {
			continue
		}
		heap.Push(nodeHeap, lists[i])
	}

	for nodeHeap.Len() != 0{
		pop := heap.Pop(nodeHeap)
		current.Next = pop.(*ListNode)
		current = current.Next
		if current.Next != nil {
			heap.Push(nodeHeap, current.Next)
		}
	}
	current.Next = nil

	return head.Next
}

func main() {
	head := ListNode{0,nil}
	fmt.Println(mergeKLists([]*ListNode{&head}))
}