package Permutations

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
		length   int // 期望的排列數量
	}{
		{
			name:     "空數組",
			input:    []int{},
			expected: [][]int{{}},
			length:   1,
		},
		{
			name:     "單個元素",
			input:    []int{1},
			expected: [][]int{{1}},
			length:   1,
		},
		{
			name:     "兩個元素",
			input:    []int{1, 2},
			expected: [][]int{{1, 2}, {2, 1}},
			length:   2,
		},
		{
			name:  "三個元素",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
			length: 6,
		},
		{
			name:   "包含負數",
			input:  []int{-1, 0, 1},
			length: 6, // 不檢查具體排列，只檢查數量
		},
		{
			name:   "檢查性能 - 較大輸入",
			input:  []int{1, 2, 3, 4, 5},
			length: 120, // 5! = 120
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := permute(test.input)

			// 檢查排列數量
			if len(result) != test.length {
				t.Errorf("期望 %d 個排列，得到 %d 個", test.length, len(result))
			}

			// 如果預期結果不為空，檢查具體排列
			if test.expected != nil && len(test.expected) > 0 {
				// 創建一個 map 來檢查所有預期的排列都存在
				expected := make(map[string]bool)
				for _, perm := range test.expected {
					key := fmt.Sprintf("%v", perm)
					expected[key] = true
				}

				// 檢查結果中的每個排列
				for _, perm := range result {
					key := fmt.Sprintf("%v", perm)
					if !expected[key] {
						t.Errorf("找到了意外的排列: %v", perm)
					}
					delete(expected, key) // 標記為已找到
				}

				// 檢查是否有預期但未找到的排列
				if len(expected) > 0 {
					for key := range expected {
						t.Errorf("未找到預期的排列: %s", key)
					}
				}
			}
		})
	}
}
