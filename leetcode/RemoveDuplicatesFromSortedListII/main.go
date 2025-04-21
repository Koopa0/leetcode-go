package RemoveDuplicatesFromSortedListII

/**
 * 鏈表節點定義
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicateBasic(head *ListNode) *ListNode {
	// 建立前導節點，簡化對頭節點的處理
	dummy := &ListNode{0, head}

	// prev 指針用來追蹤上一個非重複節點
	prev := dummy

	// 當還有節點可以檢查時
	for prev.Next != nil && prev.Next.Next != nil {
		// 如果發現兩個連續節點值相同（可能是重複值的開始）
		if prev.Next.Val == prev.Next.Next.Val {
			// 記錄這個重複值
			duplicate := prev.Next.Val
			// 向前移動直到找到不同的值
			for prev.Next != nil && prev.Next.Val == duplicate {
				prev.Next = prev.Next.Next
			}
		} else {
			// 沒有重複，前進到下一個節點
			prev = prev.Next
		}
	}

	return dummy.Next
}

func deleteDuplicatesBacktracking(head *ListNode) *ListNode {
	// 基本情況：空鏈表或只有一個節點
	if head == nil || head.Next == nil {
		return head
	}

	// 如果頭節點是重複值的開始
	if head.Val == head.Next.Val {
		// 向前移動直到找到不同的值
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		// 遞迴處理剩餘部分
		return deleteDuplicates(head.Next)
	} else {
		// 頭節點不是重複的，保留它並遞迴處理剩餘部分
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 處理空鏈表
	if head == nil {
		return nil
	}

	// 建立前導節點
	dummy := &ListNode{0, head}
	prev := dummy
	curr := head

	for curr != nil {
		// 檢測重複序列的開始
		isDuplicate := false

		// 當下一個節點存在且與當前節點值相同
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
			isDuplicate = true
		}

		if isDuplicate {
			// 找到了重複序列，跳過整個序列
			prev.Next = curr.Next
		} else {
			// 當前節點不是重複的，前進 prev
			prev = prev.Next
		}

		// 前進到下一個節點
		curr = curr.Next
	}

	return dummy.Next
}
