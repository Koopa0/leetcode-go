package RemoveDuplicatesFromSortedArrayII

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		result   []int
		desc     string
	}{
		{
			name:     "Example 1",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: 5,
			result:   []int{1, 1, 2, 2, 3},
			desc:     "基本功能測試，包含三個連續重複的元素",
		},
		{
			name:     "Example 2",
			input:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: 7,
			result:   []int{0, 0, 1, 1, 2, 3, 3},
			desc:     "基本功能測試，包含四個連續重複的元素",
		},
		{
			name:     "Empty Array",
			input:    []int{},
			expected: 0,
			result:   []int{},
			desc:     "邊界情況：空陣列",
		},
		{
			name:     "Single Element",
			input:    []int{1},
			expected: 1,
			result:   []int{1},
			desc:     "邊界情況：只有一個元素",
		},
		{
			name:     "Two Elements - Same",
			input:    []int{1, 1},
			expected: 2,
			result:   []int{1, 1},
			desc:     "邊界情況：兩個相同的元素",
		},
		{
			name:     "All Same Elements",
			input:    []int{1, 1, 1, 1, 1},
			expected: 2,
			result:   []int{1, 1},
			desc:     "特殊情況：所有元素都相同",
		},
		{
			name:     "No Duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
			result:   []int{1, 2, 3, 4, 5},
			desc:     "特殊情況：沒有重複元素",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			got := removeDuplicates(input)
			if got != tt.expected {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.expected)
			}

			for i := 0; i < tt.expected; i++ {
				if input[i] != tt.result[i] {
					t.Errorf("Expected nums[%d] to be %v, but got %v", i, tt.result[i], input[i])
				}
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	// 基準測試的輸入
	input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5}

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 創建輸入的副本以避免修改原始輸入
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)
			removeDuplicatesBruteForce(inputCopy)
		}
	})

	b.Run("雙指針解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 創建輸入的副本以避免修改原始輸入
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)
			removeDuplicatesTwoPointers(inputCopy)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 創建輸入的副本以避免修改原始輸入
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)
			removeDuplicates(inputCopy)
		}
	})
}
