package RemoveNthNodeFromEndofList

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// 輔助函數：將切片轉換為鏈表
	createList := func(values []int) *ListNode {
		dummy := &ListNode{}
		current := dummy
		for _, val := range values {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		return dummy.Next
	}

	// 輔助函數：將鏈表轉換為切片，方便比較
	listToSlice := func(head *ListNode) []int {
		var result []int
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
		desc     string
	}{
		{
			name:     "例子1：移除中間節點",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
			desc:     "移除倒數第2個節點（值為4）",
		},
		{
			name:     "例子2：只有一個節點的鏈表",
			input:    []int{1},
			n:        1,
			expected: []int{},
			desc:     "鏈表只有一個節點，移除後為空",
		},
		{
			name:     "例子3：移除尾節點",
			input:    []int{1, 2},
			n:        1,
			expected: []int{1},
			desc:     "移除倒數第1個節點（尾節點）",
		},
		{
			name:     "例子4：移除頭節點",
			input:    []int{1, 2, 3},
			n:        3,
			expected: []int{2, 3},
			desc:     "移除倒數第3個節點（頭節點）",
		},
		{
			name:     "例子5：較長的鏈表",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:        5,
			expected: []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
			desc:     "移除中間節點（值為6）",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			result := removeNthFromEnd(head, tt.n)
			resultSlice := listToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v; want %v (%s)",
					tt.input, tt.n, resultSlice, tt.expected, tt.desc)
			}
		})
	}
}
