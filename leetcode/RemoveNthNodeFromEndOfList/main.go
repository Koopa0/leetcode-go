package RemoveNthNodeFromEndofList

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 創建虛擬頭節點，統一處理所有情況
	dummy := &ListNode{0, head}

	// 初始化快慢指針
	fast, slow := dummy, dummy

	// 快指針先走 n 步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快慢指針一起走，直到快指針達到鏈表末尾
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 刪除倒數第 n 個節點
	slow.Next = slow.Next.Next

	// 返回新的頭節點
	return dummy.Next
}
