package PlusOne

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		desc     string // 測試目的描述
	}{
		{
			name:     "基本案例",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
			desc:     "測試基本功能，最後一位加 1",
		},
		{
			name:     "需進位案例",
			input:    []int{1, 2, 9},
			expected: []int{1, 3, 0},
			desc:     "測試需要進位的情況",
		},
		{
			name:     "全部進位案例",
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
			desc:     "測試全部數字都需要進位的情況",
		},
		{
			name:     "單一數字案例",
			input:    []int{9},
			expected: []int{1, 0},
			desc:     "測試只有一個數字且需要進位的情況",
		},
		{
			name:     "邊緣案例-零",
			input:    []int{0},
			expected: []int{1},
			desc:     "測試輸入為 [0] 的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("plusOne() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試案例
	input := []int{9, 9, 9, 9, 9, 9, 9}

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製輸入以避免修改原始數據
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)

			// 執行暴力解法
			bruteForcePlusOne(inputCopy)
		}
	})

	b.Run("優化解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製輸入以避免修改原始數據
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)

			// 執行優化解法
			optimizedPlusOne(inputCopy)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 複製輸入以避免修改原始數據
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)

			// 執行最佳解法
			plusOne(inputCopy)
		}
	})
}
