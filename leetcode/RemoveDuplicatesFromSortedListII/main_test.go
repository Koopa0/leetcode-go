package RemoveDuplicatesFromSortedListII

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	// 建立測試用的鏈表結構
	createList := func(vals []int) *ListNode {
		dummy := &ListNode{}
		curr := dummy
		for _, val := range vals {
			curr.Next = &ListNode{Val: val}
			curr = curr.Next
		}
		return dummy.Next
	}

	// 將鏈表轉換為數組，方便比較
	listToSlice := func(head *ListNode) []int {
		result := []int{}
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	tests := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{
			name:     "範例1",
			input:    []int{1, 2, 3, 3, 4, 4, 5},
			expected: []int{1, 2, 5},
			desc:     "移除中間的重複元素",
		},
		{
			name:     "範例2",
			input:    []int{1, 1, 1, 2, 3},
			expected: []int{2, 3},
			desc:     "移除開頭的重複元素",
		},
		{
			name:     "空鏈表",
			input:    []int{},
			expected: []int{},
			desc:     "處理空鏈表",
		},
		{
			name:     "全部重複",
			input:    []int{1, 1, 2, 2, 3, 3},
			expected: []int{},
			desc:     "所有元素都是重複的",
		},
		{
			name:     "尾部重複",
			input:    []int{1, 2, 3, 3},
			expected: []int{1, 2},
			desc:     "移除尾部的重複元素",
		},
		{
			name:     "無重複",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
			desc:     "處理沒有重複的鏈表",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := createList(tt.input)
			result := deleteDuplicates(input)
			got := listToSlice(result)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	// 建立各種大小的測試用例
	createTestCase := func(size int, duplicateRatio float64) *ListNode {
		dummy := &ListNode{}
		curr := dummy

		val := 1
		for i := 0; i < size; i++ {
			curr.Next = &ListNode{Val: val}
			curr = curr.Next

			// 根據重複比率決定是否生成重複值
			if rand.Float64() > duplicateRatio {
				val++
			}
		}

		return dummy.Next
	}

	// 小型鏈表，約 20 個節點
	smallList := createTestCase(20, 0.3)
	// 中型鏈表，約 100 個節點
	mediumList := createTestCase(100, 0.3)
	// 大型鏈表，約 300 個節點
	largeList := createTestCase(300, 0.3)

	b.Run("小型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製鏈表以免修改原始測試數據
			listCopy := deepCopyList(smallList)
			deleteDuplicateBasic(listCopy)
		}
	})

	b.Run("中型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := deepCopyList(mediumList)
			deleteDuplicateBasic(listCopy)
		}
	})

	b.Run("大型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := deepCopyList(largeList)
			deleteDuplicateBasic(listCopy)
		}
	})

	b.Run("小型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製鏈表以免修改原始測試數據
			listCopy := deepCopyList(smallList)
			deleteDuplicatesBacktracking(listCopy)
		}
	})

	b.Run("中型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := deepCopyList(mediumList)
			deleteDuplicatesBacktracking(listCopy)
		}
	})

	b.Run("大型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := deepCopyList(largeList)
			deleteDuplicatesBacktracking(listCopy)
		}
	})

	b.Run("小型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製鏈表以免修改原始測試數據
			listCopy := deepCopyList(smallList)
			deleteDuplicates(listCopy)
		}
	})

	b.Run("中型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := deepCopyList(mediumList)
			deleteDuplicates(listCopy)
		}
	})

	b.Run("大型鏈表", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			listCopy := deepCopyList(largeList)
			deleteDuplicates(listCopy)
		}
	})
}

// 深度複製鏈表
func deepCopyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Val: head.Val}
	curr := newHead
	node := head.Next

	for node != nil {
		curr.Next = &ListNode{Val: node.Val}
		curr = curr.Next
		node = node.Next
	}

	return newHead
}
