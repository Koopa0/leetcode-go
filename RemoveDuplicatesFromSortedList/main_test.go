package RemoveDuplicatesFromSortedList

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{
			name:     "空串列",
			input:    []int{},
			expected: []int{},
			desc:     "測試空串列",
		},
		{
			name:     "單一節點",
			input:    []int{1},
			expected: []int{1},
			desc:     "測試只有一個節點的串列",
		},
		{
			name:     "無重複",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
			desc:     "測試無重複元素的串列",
		},
		{
			name:     "有重複",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
			desc:     "測試有重複元素的串列",
		},
		{
			name:     "全部重複",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
			desc:     "測試所有元素都相同的串列",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 建立鏈結串列
			head := createLinkedList(tt.input)

			// 執行函數
			result := deleteDuplicatesTwoPointers(head)

			// 將結果轉換回陣列
			got := linkedListToSlice(result)

			// 比較結果
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 輔助函數：建立鏈結串列
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return dummy.Next
}

// 輔助函數：將鏈結串列轉換為陣列
func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func BenchmarkSolutions(b *testing.B) {
	// 測試資料：長串列，重複元素較多
	testData := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 5, 6, 7, 7, 7, 8, 9, 9}

	b.Run("迭代解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := createLinkedList(testData)
			deleteDuplicatesIterative(head)
		}
	})

	b.Run("遞迴解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := createLinkedList(testData)
			deleteDuplicatesRecursive(head)
		}
	})

	b.Run("雙指針解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := createLinkedList(testData)
			deleteDuplicatesTwoPointers(head)
		}
	})
}
