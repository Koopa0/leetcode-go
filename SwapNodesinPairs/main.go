package SwapNodesinPairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 兩兩交換鏈表中的節點
// 時間複雜度：O(n)，其中 n 是鏈表的長度
// 空間複雜度：O(1)，只使用常數額外空間
func swapPairs(head *ListNode) *ListNode {
	// 創建啞節點，指向頭節點
	// 啞節點可以幫助我們統一處理頭節點和其他節點
	dummy := &ListNode{Next: head}

	// prev 指針用於跟蹤當前處理的一對節點之前的節點
	prev := dummy

	// 當還有至少兩個節點可以交換時，繼續循環
	for head != nil && head.Next != nil {
		// 獲取要交換的兩個節點
		first := head       // 第一個節點
		second := head.Next // 第二個節點

		// 執行交換操作
		// 步驟 1: 前一個節點指向第二個節點
		prev.Next = second

		// 步驟 2: 第一個節點指向第二個節點之後的節點
		first.Next = second.Next

		// 步驟 3: 第二個節點指向第一個節點
		second.Next = first

		// 更新指針，準備處理下一對節點
		prev = first      // prev 移動到當前這一對中的第一個節點（交換後的第二個位置）
		head = first.Next // head 移動到下一對的第一個節點
	}

	// 返回交換後的新頭節點（即啞節點的下一個節點）
	return dummy.Next
}
