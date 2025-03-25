package AddTwoNumbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 創建虛擬頭節點，以簡化操作
	dummy := &ListNode{Val: 0}
	current := dummy
	carry := 0 // 初始進位為0

	// 只要還有節點沒遍歷完或者還有進位，就繼續循環
	for l1 != nil || l2 != nil || carry > 0 {
		// 獲取當前兩個鏈表對應位置的值（如果節點為空則值為0）
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// 計算當前位的和，包括上一位的進位
		sum := val1 + val2 + carry
		carry = sum / 10 // 計算新的進位

		// 創建新節點並鏈接到結果鏈表
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	// 返回結果鏈表（跳過虛擬頭節點）
	return dummy.Next
}
