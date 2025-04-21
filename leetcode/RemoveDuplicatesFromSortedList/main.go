package RemoveDuplicatesFromSortedList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代解法
func deleteDuplicatesIterative(head *ListNode) *ListNode {
	// 處理空串列或只有一個節點的情況
	if head == nil || head.Next == nil {
		return head
	}

	// 從頭節點開始遍歷
	current := head

	// 當 current 和下一個節點都存在時
	for current != nil && current.Next != nil {
		// 如果當前節點值等於下一個節點值
		if current.Val == current.Next.Val {
			// 跳過下一個節點
			current.Next = current.Next.Next
		} else {
			// 移動到下一個節點
			current = current.Next
		}
	}

	return head
}

// 遞迴解法
func deleteDuplicatesRecursive(head *ListNode) *ListNode {
	// 基礎情況：空串列或只有一個節點
	if head == nil || head.Next == nil {
		return head
	}

	// 遞迴處理剩餘部分
	head.Next = deleteDuplicatesRecursive(head.Next)

	// 如果當前節點與下一個節點值相同，返回下一個節點
	if head.Val == head.Next.Val {
		return head.Next
	}

	// 否則返回當前節點
	return head
}

// 雙指針解法
func deleteDuplicatesTwoPointers(head *ListNode) *ListNode {
	// 處理空串列情況
	if head == nil {
		return nil
	}

	prev := head          // 指向最後一個唯一值節點
	explorer := head.Next // 用於探索的指針

	for explorer != nil {
		// 如果找到新的唯一值
		if prev.Val != explorer.Val {
			// 將 prev 的下一個節點設為 explorer
			prev.Next = explorer
			// 移動 prev 到新的唯一值
			prev = explorer
		}
		// 繼續探索
		explorer = explorer.Next
	}

	// 確保串列正確終止，移除可能的尾部重複
	prev.Next = nil

	return head
}
