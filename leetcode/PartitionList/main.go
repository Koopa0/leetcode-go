package PartitionList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 雙鏈表解決方案
func partition(head *ListNode, x int) *ListNode {
	// 1. 初始化兩個虛擬頭節點
	smallerHead := &ListNode{Val: 0}
	greaterOrEqualHead := &ListNode{Val: 0}

	// 2. 初始化兩個指針指向兩個鏈表的尾部
	smallerPtr := smallerHead
	greaterOrEqualPtr := greaterOrEqualHead

	// 3. 遍歷原始鏈表
	current := head
	for current != nil {
		if current.Val < x {
			// 3a. 將小於 x 的節點添加到 smaller 鏈表
			smallerPtr.Next = current
			smallerPtr = smallerPtr.Next
		} else {
			// 3b. 將大於等於 x 的節點添加到 greaterOrEqual 鏈表
			greaterOrEqualPtr.Next = current
			greaterOrEqualPtr = greaterOrEqualPtr.Next
		}
		// 移動到下一個節點
		current = current.Next
	}

	// 4. 將 smaller 鏈表的尾部連接到 greaterOrEqual 鏈表的頭部
	smallerPtr.Next = greaterOrEqualHead.Next

	// 5. 確保 greaterOrEqual 鏈表的末尾是 nil
	greaterOrEqualPtr.Next = nil

	// 返回新鏈表的頭部
	return smallerHead.Next
}
