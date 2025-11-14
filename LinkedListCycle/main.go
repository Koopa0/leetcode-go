package main

// ListNode 定義鏈結串列節點
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 使用 Floyd's Cycle Detection Algorithm 檢測鏈結串列是否有環
// 使用快慢指標：slow 每次走一步，fast 每次走兩步
// 如果存在環，fast 最終會追上 slow
// 時間複雜度: O(n)，其中 n 是鏈結串列長度
// 空間複雜度: O(1)，只使用兩個指標
func hasCycle(head *ListNode) bool {
	// 邊界條件：空串列或只有一個節點且無環
	if head == nil || head.Next == nil {
		return false
	}

	// 初始化快慢指標
	// slow 從 head 開始，fast 從 head.Next 開始
	slow := head
	fast := head.Next

	// 當 fast 和 slow 不相等時繼續移動
	for slow != fast {
		// 如果 fast 到達末尾，表示無環
		if fast == nil || fast.Next == nil {
			return false
		}
		// slow 走一步，fast 走兩步
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow 和 fast 相遇，表示有環
	return true
}

// hasCycleAlternative 另一種實作方式：兩個指標都從 head 開始
// 時間複雜度: O(n)
// 空間複雜度: O(1)
func hasCycleAlternative(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	// 先移動再檢查
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 移動後檢查是否相遇
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	// 範例 1：有環的鏈結串列 [3,2,0,-4]，環從位置 1 開始
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // 形成環

	println("範例 1: [3,2,0,-4], pos = 1")
	println("是否有環:", hasCycle(node1)) // true

	// 範例 2：有環的鏈結串列 [1,2]，環從位置 0 開始
	node5 := &ListNode{Val: 1}
	node6 := &ListNode{Val: 2}
	node5.Next = node6
	node6.Next = node5 // 形成環

	println("\n範例 2: [1,2], pos = 0")
	println("是否有環:", hasCycle(node5)) // true

	// 範例 3：無環的鏈結串列 [1]
	node7 := &ListNode{Val: 1}

	println("\n範例 3: [1], pos = -1")
	println("是否有環:", hasCycle(node7)) // false
}
