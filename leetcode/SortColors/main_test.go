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
		desc     string // 測試目的描述
	}{
		{
			name:     "基本情況 1",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
			desc:     "標準測試案例",
		},
		{
			name:     "基本情況 2",
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
			desc:     "最小測試案例",
		},
		{
			name:     "單一顏色",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
			desc:     "只有一種顏色的情況",
		},
		{
			name:     "已排序陣列",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
			desc:     "已經排序好的陣列",
		},
		{
			name:     "逆序陣列",
			input:    []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
			desc:     "完全倒序的陣列",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製輸入以免修改原始數據
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// 執行排序
			twoPointersSortColors(input)

			// 檢查結果
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("sortColors() = %v, want %v", input, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試輸入
	input := []int{2, 0, 2, 1, 1, 0}

	b.Run("計數排序", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製輸入以免修改原始數據
			testInput := make([]int, len(input))
			copy(testInput, input)

			countingSortSortColors(testInput)
		}
	})

	b.Run("三路快排", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製輸入以免修改原始數據
			testInput := make([]int, len(input))
			copy(testInput, input)

			threeWayPartitioningSortColors(testInput)
		}
	})

	b.Run("雙指針", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製輸入以免修改原始數據
			testInput := make([]int, len(input))
			copy(testInput, input)

			twoPointersSortColors(testInput)
		}
	})
}
