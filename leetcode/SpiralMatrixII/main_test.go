package SpiralMatrixII

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected [][]int
		desc     string
	}{
		{
			name:     "3x3 matrix",
			input:    3,
			expected: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			desc:     "測試 3x3 矩陣的螺旋填充",
		},
		{
			name:     "1x1 matrix",
			input:    1,
			expected: [][]int{{1}},
			desc:     "測試 1x1 矩陣的螺旋填充",
		},
		{
			name:     "2x2 matrix",
			input:    2,
			expected: [][]int{{1, 2}, {4, 3}},
			desc:     "測試 2x2 矩陣的螺旋填充",
		},
		{
			name:     "4x4 matrix",
			input:    4,
			expected: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
			desc:     "測試 4x4 矩陣的螺旋填充",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateMatrix(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 基準測試用的輸入
	input := 10

	b.Run("邊界收縮法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateMatrixBoundary(input)
		}
	})

	b.Run("方向陣列法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateMatrixDirection(input)
		}
	})

	b.Run("層次遍歷法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateMatrix(input)
		}
	})
}
