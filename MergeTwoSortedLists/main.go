package MergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 基本情況：其中一個鏈表為空
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 遞迴情況：比較兩個鏈表頭節點的值
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
