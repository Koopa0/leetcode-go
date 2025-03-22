package SortColors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{
			name:     "Example 1",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
			desc:     "測試範例1",
		},
		{
			name:     "Example 2",
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
			desc:     "測試範例2",
		},
		{
			name:     "Already Sorted",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
			desc:     "測試已排序的",
		},
		{
			name:     "Reverse Sorted",
			input:    []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
			desc:     "測試逆序",
		},
		{
			name:     "All Zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
			desc:     "測試全為0的",
		},
		{
			name:     "All Ones",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
			desc:     "測試全為1的",
		},
		{
			name:     "All Twos",
			input:    []int{2, 2, 2},
			expected: []int{2, 2, 2},
			desc:     "測試全為2的",
		},
		{
			name:     "Single Element",
			input:    []int{1},
			expected: []int{1},
			desc:     "測試只有一個元素的",
		},
		{
			name:     "No Zeros",
			input:    []int{1, 2, 1, 2},
			expected: []int{1, 1, 2, 2},
			desc:     "測試不含0的",
		},
		{
			name:     "No Ones",
			input:    []int{0, 2, 0, 2},
			expected: []int{0, 0, 2, 2},
			desc:     "測試不含1的",
		},
		{
			name:     "No Twos",
			input:    []int{0, 1, 0, 1},
			expected: []int{0, 0, 1, 1},
			desc:     "測試不含2的",
		},
		{
			name:     "Random Order",
			input:    []int{1, 0, 2, 1, 0, 2, 1, 0, 2},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
			desc:     "測試隨機順序的",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 創建輸入的副本
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// 執行排序
			sortColors(input)

			// 檢查結果
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("sortColors() = %v, want %v, desc: %s",
					input, tt.expected, tt.desc)
			}
		})
	}
}
