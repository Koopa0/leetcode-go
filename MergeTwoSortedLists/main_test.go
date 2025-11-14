package MergeTwoSortedLists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// 輔助函數：從切片創建鏈表
	createList := func(values []int) *ListNode {
		dummy := &ListNode{}
		current := dummy
		for _, val := range values {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		return dummy.Next
	}

	// 輔助函數：將鏈表轉換為切片，便於比較
	listToSlice := func(head *ListNode) []int {
		result := []int{}
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	// 測試用例
	testCases := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "Example 1: Normal case",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Example 2: Both empty",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "Example 3: One empty, one has value",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
		{
			name:     "List1 longer than list2",
			list1:    []int{1, 3, 5, 7, 9},
			list2:    []int{2, 4},
			expected: []int{1, 2, 3, 4, 5, 7, 9},
		},
		{
			name:     "List2 longer than list1",
			list1:    []int{1, 3},
			list2:    []int{2, 4, 6, 8, 10},
			expected: []int{1, 2, 3, 4, 6, 8, 10},
		},
		{
			name:     "Same values",
			list1:    []int{1, 1, 1},
			list2:    []int{1, 1, 1},
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:     "Negative values",
			list1:    []int{-3, -1, 1},
			list2:    []int{-2, 0, 2},
			expected: []int{-3, -2, -1, 0, 1, 2},
		},
		{
			name:     "Edge case: Large difference in values",
			list1:    []int{-100, 0, 100},
			list2:    []int{-50, 50},
			expected: []int{-100, -50, 0, 50, 100},
		},
	}

	// 執行測試
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list1 := createList(tc.list1)
			list2 := createList(tc.list2)

			result := mergeTwoLists(list1, list2)

			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, resultSlice)
			}
		})
	}
}
