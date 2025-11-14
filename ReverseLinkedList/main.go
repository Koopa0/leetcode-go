package ReverseLinkedList

// ListNode 定義單向鏈結串列的節點
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 使用迭代法反轉鏈結串列
// 時間複雜度: O(n)，其中 n 是鏈結串列的節點數
// 空間複雜度: O(1)，只使用固定的額外空間
func reverseList(head *ListNode) *ListNode {
	// 處理空鏈結串列或單節點情況
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode // 前一個節點，初始為 nil
	curr := head       // 當前節點，從頭節點開始

	// 遍歷整個鏈結串列
	for curr != nil {
		next := curr.Next // 暫存下一個節點，避免斷鏈
		curr.Next = prev  // 將當前節點的 next 指向前一個節點（反轉）
		prev = curr       // 移動 prev 指標到當前節點
		curr = next       // 移動 curr 指標到下一個節點
	}

	// 最後 prev 指向新的頭節點
	return prev
}

// reverseListRecursive 使用遞迴法反轉鏈結串列
// 時間複雜度: O(n)，需要訪問每個節點一次
// 空間複雜度: O(n)，遞迴呼叫堆疊的深度
func reverseListRecursive(head *ListNode) *ListNode {
	// 基礎情況：空鏈結串列或只有一個節點
	if head == nil || head.Next == nil {
		return head
	}

	// 遞迴反轉從第二個節點開始的子鏈結串列
	// newHead 是反轉後子鏈結串列的頭節點
	newHead := reverseListRecursive(head.Next)

	// 將下一個節點的 next 指向當前節點（反轉連結）
	// 例如：1 -> 2 -> 3 反轉後段得到 1 -> 2 <- 3
	head.Next.Next = head

	// 將當前節點的 next 設為 nil（斷開原有連結）
	// 完成反轉：1 <- 2 <- 3
	head.Next = nil

	// 返回新的頭節點
	return newHead
}
