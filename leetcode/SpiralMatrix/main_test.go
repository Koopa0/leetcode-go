package SpiralMatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
		desc     string // 測試目的描述
	}{
		{
			name:     "3x3矩陣",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			desc:     "測試標準的3x3矩陣",
		},
		{
			name:     "3x4矩陣",
			input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
			desc:     "測試非正方形矩陣",
		},
		{
			name:     "1x1矩陣",
			input:    [][]int{{1}},
			expected: []int{1},
			desc:     "測試單元素矩陣",
		},
		{
			name:     "1x3矩陣",
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
			desc:     "測試單行矩陣",
		},
		{
			name:     "3x1矩陣",
			input:    [][]int{{1}, {2}, {3}},
			expected: []int{1, 2, 3},
			desc:     "測試單列矩陣",
		},
		{
			name:     "空矩陣",
			input:    [][]int{},
			expected: []int{},
			desc:     "測試空矩陣",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSpiralOrder(b *testing.B) {
	// 3x3 矩陣作為基準測試輸入
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			spiralOrderBruteForce(input)
		}
	})

	b.Run("優化解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			spiralOrderOptimized(input)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			spiralOrder(input)
		}
	})
}
