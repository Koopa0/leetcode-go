package RemoveDuplicatesfromSortedArray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		result   []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 1, 2},
			expected: 2,
			result:   []int{1, 2},
		},
		{
			name:     "Example 2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
			result:   []int{0, 1, 2, 3, 4},
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
			result:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All same elements",
			input:    []int{1, 1, 1, 1, 1},
			expected: 1,
			result:   []int{1},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: 1,
			result:   []int{1},
		},
		{
			name:     "Edge values",
			input:    []int{-100, -100, 0, 0, 100, 100},
			expected: 3,
			result:   []int{-100, 0, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 創建輸入數組的副本以避免修改原始測試數據
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)

			// 調用函數
			k := removeDuplicates(nums)

			// 檢查返回值
			if k != tt.expected {
				t.Errorf("expected %d unique elements, got %d", tt.expected, k)
			}

			// 檢查前 k 個元素是否符合預期
			for i := 0; i < tt.expected; i++ {
				if i < len(nums) && nums[i] != tt.result[i] {
					t.Errorf("at index %d: expected %d, got %d", i, tt.result[i], nums[i])
				}
			}
		})
	}
}
