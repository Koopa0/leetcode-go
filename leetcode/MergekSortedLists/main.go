package MergekSortedLists

import (
	"container/heap"
)

// ListNode 定義鏈表節點結構
type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap 定義一個最小堆，用於存儲鏈表節點
type NodeHeap []*ListNode

// Len 實現 heap.Interface 所需的方法
func (h NodeHeap) Len() int { return len(h) }

// Less 比較兩個節點的值，用於確定堆的排序
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

// Swap 交換堆中的兩個元素
func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 向堆中添加新元素
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// Pop 從堆中彈出最小元素
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeKLists 合併 k 個排序鏈表
func MergeKLists(lists []*ListNode) *ListNode {
	// 處理空輸入
	if len(lists) == 0 {
		return nil
	}

	// 創建一個虛擬頭節點，用於構建結果鏈表
	dummy := &ListNode{}
	curr := dummy

	// 初始化最小堆
	h := &NodeHeap{}
	heap.Init(h)

	// 將所有非空鏈表的頭節點加入堆中
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	// 不斷從堆中取出最小節點，加入結果鏈表
	for h.Len() > 0 {
		// 彈出堆頂元素（當前最小節點）
		minNode := heap.Pop(h).(*ListNode)
		curr.Next = minNode
		curr = curr.Next

		// 如果彈出的節點有下一個節點，則將其加入堆中
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	return dummy.Next
}

// 以下是兩個輔助合併方法，用於比較

// MergeKListsLinear 使用線性合併方法
func MergeKListsLinear(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}

	return result
}

// MergeKListsDivideConquer 使用分治法合併
func MergeKListsDivideConquer(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	mid := n / 2
	left := MergeKListsDivideConquer(lists[:mid])
	right := MergeKListsDivideConquer(lists[mid:])

	return mergeTwoLists(left, right)
}

// mergeTwoLists 合併兩個排序鏈表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// 將剩餘節點接上
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}
