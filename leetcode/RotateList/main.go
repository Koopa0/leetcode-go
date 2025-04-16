package RotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力解法
func bruteForceSolution(head *ListNode, k int) *ListNode {
	// 處理邊緣情況：空串列或只有一個節點
	if head == nil || head.Next == nil {
		return head
	}

	// 計算鏈結串列長度
	length := 1
	current := head
	for current.Next != nil {
		length++
		current = current.Next
	}

	// 優化：k 可能很大，取模減少實際需要的旋轉次數
	k = k % length

	// 如果 k 等於 0，不需要旋轉
	if k == 0 {
		return head
	}

	// 執行 k 次旋轉
	for i := 0; i < k; i++ {
		// 找到最後一個和倒數第二個節點
		last := head
		secondLast := &ListNode{}

		for last.Next != nil {
			secondLast = last
			last = last.Next
		}

		// 重新連接
		last.Next = head
		secondLast.Next = nil
		head = last
	}

	return head
}

func rotateRightOptimized(head *ListNode, k int) *ListNode {
	// 處理邊緣情況
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 計算鏈結串列長度並找到尾節點
	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	// 計算實際需要旋轉的次數
	k = k % length
	if k == 0 {
		return head
	}

	// 找到新的尾節點位置（從頭開始的第 length-k-1 個節點）
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}

	// 重新連接
	newHead := newTail.Next // 新的頭節點
	newTail.Next = nil      // 切斷新的尾節點
	tail.Next = head        // 原尾節點指向原頭節點

	return newHead
}

func rotateRight(head *ListNode, k int) *ListNode {
	// 處理邊緣情況
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 計算鏈結串列長度並找到尾節點
	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	// 計算實際需要旋轉的次數
	k = k % length
	if k == 0 {
		return head
	}

	// 將串列首尾相連形成環
	tail.Next = head

	// 找到新的尾節點（需要向前走 length-k 步）
	current := head
	for i := 0; i < length-k-1; i++ {
		current = current.Next
	}

	// 切斷環並重新定義頭節點
	newHead := current.Next
	current.Next = nil

	return newHead
}
