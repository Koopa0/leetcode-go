package ReverseLinkedListII

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {

	// 從陣列建立鏈結串列的輔助函數
	createList := func(vals []int) *ListNode {
		dummy := &ListNode{}
		curr := dummy
		for _, val := range vals {
			curr.Next = &ListNode{Val: val}
			curr = curr.Next
		}
		return dummy.Next
	}

	// 將鏈結串列轉換為陣列的輔助函數
	listToArray := func(head *ListNode) []int {
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
		left     int
		right    int
		expected []int
		desc     string
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
			desc:     "反轉鏈結串列中的第 2 到第 4 個節點",
		},
		{
			name:     "Example 2",
			input:    []int{5},
			left:     1,
			right:    1,
			expected: []int{5},
			desc:     "只有一個節點，無需反轉",
		},
		{
			name:     "反轉整個鏈結串列",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    5,
			expected: []int{5, 4, 3, 2, 1},
			desc:     "從頭到尾完全反轉鏈結串列",
		},
		{
			name:     "反轉從頭開始的部分",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    3,
			expected: []int{3, 2, 1, 4, 5},
			desc:     "反轉從頭節點開始的部分",
		},
		{
			name:     "反轉到尾部的部分",
			input:    []int{1, 2, 3, 4, 5},
			left:     3,
			right:    5,
			expected: []int{1, 2, 5, 4, 3},
			desc:     "反轉到尾節點的部分",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			result := reverseBetween(head, tt.left, tt.right)
			got := listToArray(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 準備長度為 100 的測試鏈結串列
	createLongList := func(length int) *ListNode {
		dummy := &ListNode{}
		curr := dummy
		for i := 1; i <= length; i++ {
			curr.Next = &ListNode{Val: i}
			curr = curr.Next
		}
		return dummy.Next
	}

	// 測試於中間位置反轉的情況
	//head := createLongList(100)
	left, right := 30, 70

	b.Run("兩次遍歷解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製一份鏈結串列避免修改原始測試資料
			listCopy := createLongList(100)
			reverseBetweenTwoPass(listCopy, left, right)
		}
	})

	b.Run("一次遍歷解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := createLongList(100)
			reverseBetweenOnePass(listCopy, left, right)
		}
	})

	b.Run("最優解法 (頭插法)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := createLongList(100)
			reverseBetween(listCopy, left, right)
		}
	})
}
