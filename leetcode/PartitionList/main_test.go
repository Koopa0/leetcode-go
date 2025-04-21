package PartitionList

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		x        int
		expected *ListNode
		desc     string // 測試目的描述
	}{
		{
			name:     "標準案例1",
			head:     createLinkedList([]int{1, 4, 3, 2, 5, 2}),
			x:        3,
			expected: createLinkedList([]int{1, 2, 2, 4, 3, 5}),
			desc:     "驗證標準分割，確保保持相對順序",
		},
		{
			name:     "標準案例2",
			head:     createLinkedList([]int{2, 1}),
			x:        2,
			expected: createLinkedList([]int{1, 2}),
			desc:     "測試邊界值情況",
		},
		{
			name:     "空鏈表",
			head:     nil,
			x:        3,
			expected: nil,
			desc:     "測試空鏈表的處理",
		},
		{
			name:     "所有節點都小於x",
			head:     createLinkedList([]int{1, 2, 2}),
			x:        3,
			expected: createLinkedList([]int{1, 2, 2}),
			desc:     "測試所有節點都小於分割值的情況",
		},
		{
			name:     "所有節點都大於等於x",
			head:     createLinkedList([]int{3, 4, 5}),
			x:        3,
			expected: createLinkedList([]int{3, 4, 5}),
			desc:     "測試所有節點都大於等於分割值的情況",
		},
		{
			name:     "單節點",
			head:     createLinkedList([]int{1}),
			x:        0,
			expected: createLinkedList([]int{1}),
			desc:     "測試單節點鏈表的處理",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.head, tt.x)
			if !isLinkedListEqual(got, tt.expected) {
				t.Errorf("partition() = %v, want %v", linkedListToSlice(got), linkedListToSlice(tt.expected))
			}
		})
	}
}

// 輔助函數：創建鏈表
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// 輔助函數：比較兩個鏈表是否相等
func isLinkedListEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	// 檢查兩個鏈表是否都到達結尾
	return l1 == nil && l2 == nil
}

// 輔助函數：將鏈表轉換為切片（用於錯誤消息）
func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func BenchmarkPartition(b *testing.B) {
	// 準備測試資料，包含不同規模的鏈表
	testCases := []struct {
		name string
		size int
		x    int
	}{
		{"小規模", 10, 5},
		{"中規模", 100, 50},
		{"大規模", 1000, 500},
	}

	for _, tc := range testCases {
		// 創建具有隨機值的鏈表
		values := make([]int, tc.size)
		for i := 0; i < tc.size; i++ {
			values[i] = rand.Intn(1000) - 500 // 生成 -500 到 499 的隨機數
		}
		original := createLinkedList(values)

		b.Run(fmt.Sprintf("%s_雙鏈表解決方案", tc.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 重新複製一份鏈表，避免多次測試時鏈表結構被修改
				testHead := cloneLinkedList(original)
				partition(testHead, tc.x)
			}
		})
	}
}

// 輔助函數：複製鏈表
func cloneLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Val: head.Val}
	newCurrent := newHead
	current := head.Next

	for current != nil {
		newCurrent.Next = &ListNode{Val: current.Val}
		newCurrent = newCurrent.Next
		current = current.Next
	}

	return newHead
}
