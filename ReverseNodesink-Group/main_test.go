package ReverseNodesink_Group

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1: k=2",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "Example 2: k=3",
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "Example 3: k=1 (no reversal)",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty list",
			input:    []int{},
			k:        2,
			expected: []int{},
		},
		{
			name:     "Single node",
			input:    []int{1},
			k:        2,
			expected: []int{1},
		},
		{
			name:     "Exactly k nodes",
			input:    []int{1, 2, 3},
			k:        3,
			expected: []int{3, 2, 1},
		},
		{
			name:     "Multiple complete groups",
			input:    []int{1, 2, 3, 4, 5, 6},
			k:        2,
			expected: []int{2, 1, 4, 3, 6, 5},
		},
		{
			name:     "Multiple groups with remainder",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{3, 2, 1, 6, 5, 4, 7},
		},
		{
			name:     "Large k equal to list length",
			input:    []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 構建輸入鏈表
			head := buildList(tt.input)

			// 執行函數
			result := reverseKGroup(head, tt.k)

			// 將結果轉換回數組
			resultArr := listToArray(result)

			if len(resultArr) == 0 && len(tt.expected) == 0 {
				return
			}

			// 驗證結果
			if !reflect.DeepEqual(resultArr, tt.expected) {
				t.Errorf("reverseKGroup() = %v, want %v", resultArr, tt.expected)
			}
		})
	}
}

// 輔助函數：將整數數組轉換為鏈表
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}

	return dummy.Next
}

// 輔助函數：將鏈表轉換為整數數組
func listToArray(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
