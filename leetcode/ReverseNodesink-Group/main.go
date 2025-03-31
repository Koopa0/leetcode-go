package ReverseNodesink_Group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		// 檢查是否有 k 個節點
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next // 不足 k 個節點，返回結果
			}
		}

		// 保存下一組的頭節點
		next := tail.Next

		// 反轉當前組
		head, tail = reverseSubList(head, tail)

		// 連接前一組和當前組
		prev.Next = head

		// 連接當前組和下一組
		tail.Next = next

		// 更新 prev 和 head
		prev = tail
		head = next
	}

	return dummy.Next
}

// 反轉從 head 到 tail 的子鏈表，返回新的頭和尾
func reverseSubList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head

	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return tail, head
}
