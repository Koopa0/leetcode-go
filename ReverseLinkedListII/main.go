package ReverseLinkedListII

type ListNode struct {
	Val  int
	Next *ListNode
}

// 兩次遍歷解法
func reverseBetweenTwoPass(head *ListNode, left int, right int) *ListNode {
	// 創建虛擬頭節點，簡化處理 left=1 的情況
	dummy := &ListNode{Val: 0, Next: head}

	// 1. 找到 prevL 節點 (第 left-1 個節點)
	prevL := dummy
	for i := 1; i < left; i++ {
		prevL = prevL.Next
	}

	// 2. 找到 leftNode (第 left 個節點)
	leftNode := prevL.Next

	// 3. 反轉 left 到 right 的部分
	current := leftNode
	var prev *ListNode = nil
	for i := left; i <= right; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// 4. 重新連接
	prevL.Next = prev       // prevL 連接到反轉後的新頭（原來的 rightNode）
	leftNode.Next = current // 原來的 leftNode (現在是反轉後的尾) 連接到 nextR

	return dummy.Next
}

// 一次遍歷解法
func reverseBetweenOnePass(head *ListNode, left int, right int) *ListNode {
	// 創建虛擬頭節點
	dummy := &ListNode{Val: 0, Next: head}

	// 找到 prevL 節點
	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// 當前節點 (開始反轉的第一個節點)
	curr := prev.Next

	// 使用頭插法反轉
	for i := 0; i < right-left; i++ {
		// 取得下一個要被移動的節點
		nextNode := curr.Next

		// 將 nextNode 從鏈結串列中移除
		curr.Next = nextNode.Next

		// 將 nextNode 插入到 prevL 之後
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}

// 最優解法 (頭插法優化)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 邊界情況：空鏈結串列或無需反轉
	if head == nil || left == right {
		return head
	}

	// 創建虛擬頭節點
	dummy := &ListNode{Val: 0, Next: head}

	// 找到反轉區段前的節點
	prevL := dummy
	for i := 1; i < left; i++ {
		prevL = prevL.Next
	}

	// 反轉起始節點 (第 left 個節點)
	leftNode := prevL.Next

	// 使用頭插法反轉
	for i := 0; i < right-left; i++ {
		// 要被移動的節點
		nodeToMove := leftNode.Next

		// 從原位置移除節點
		leftNode.Next = nodeToMove.Next

		// 將節點插入到反轉區段的開頭
		nodeToMove.Next = prevL.Next
		prevL.Next = nodeToMove
	}

	return dummy.Next
}
