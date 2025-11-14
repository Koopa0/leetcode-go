package SearchA2DMatrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
		desc     string
	}{
		{
			name:     "基本測試-目標存在",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
			desc:     "測試目標值在矩陣中",
		},
		{
			name:     "基本測試-目標不存在",
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
			desc:     "測試目標值不在矩陣中",
		},
		{
			name:     "邊界情況-空矩陣",
			matrix:   [][]int{},
			target:   1,
			expected: false,
			desc:     "測試空矩陣",
		},
		{
			name:     "邊界情況-單元素矩陣-目標存在",
			matrix:   [][]int{{1}},
			target:   1,
			expected: true,
			desc:     "測試單元素矩陣，目標存在",
		},
		{
			name:     "邊界情況-單元素矩陣-目標不存在",
			matrix:   [][]int{{1}},
			target:   2,
			expected: false,
			desc:     "測試單元素矩陣，目標不存在",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.matrix, tt.target)
			if got != tt.expected {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 設置測試用例
	testCases := []struct {
		name   string
		matrix [][]int
		target int
	}{
		{
			name:   "小矩陣-目標存在",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
		},
		{
			name:   "小矩陣-目標不存在",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
		},
		{
			name:   "大矩陣-目標存在",
			matrix: generateLargeMatrix(50, 50), // 假設有一個函數生成較大的測試矩陣
			target: 42,
		},
	}

	for _, tc := range testCases {
		// 測試暴力解法
		b.Run(fmt.Sprintf("BruteForce-%s", tc.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bruteForceSearchMatrix(tc.matrix, tc.target)
			}
		})

		// 測試兩次二分搜尋解法
		b.Run(fmt.Sprintf("TwoBinarySearch-%s", tc.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				twoBinarySearchMatrix(tc.matrix, tc.target)
			}
		})

		// 測試單次二分搜尋解法（最優解）
		b.Run(fmt.Sprintf("OptimalBinarySearch-%s", tc.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				searchMatrix(tc.matrix, tc.target)
			}
		})
	}
}

// 生成大型測試矩陣的輔助函數
func generateLargeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	value := 1

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = value
			value++
		}
	}

	return matrix
}
