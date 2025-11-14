package MaximalRectangle

import (
	"math/rand"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
		desc     string
	}{
		{
			name: "範例1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
			desc:     "典型測試情況",
		},
		{
			name:     "空矩陣",
			matrix:   [][]byte{},
			expected: 0,
			desc:     "處理空矩陣",
		},
		{
			name:     "單個0",
			matrix:   [][]byte{{'0'}},
			expected: 0,
			desc:     "一個元素為0的矩陣",
		},
		{
			name:     "單個1",
			matrix:   [][]byte{{'1'}},
			expected: 1,
			desc:     "一個元素為1的矩陣",
		},
		{
			name: "全為1",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 9,
			desc:     "測試全部為1的矩陣",
		},
		{
			name: "全為0",
			matrix: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
			desc:     "測試全部為0的矩陣",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximalRectangle(tt.matrix)
			if got != tt.expected {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkMaximalRectangle(b *testing.B) {
	// 創建一個較大的測試矩陣
	rows, cols := 100, 100
	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
		for j := range matrix[i] {
			if rand.Intn(2) == 1 {
				matrix[i][j] = '1'
			} else {
				matrix[i][j] = '0'
			}
		}
	}

	// 暴力解法太慢，不適合基準測試

	b.Run("優化解法（堆疊）", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 使用方法2的實現
			maximalRectangleStack(matrix)
		}
	})

	b.Run("最佳解法（動態規劃）", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 使用方法3的實現
			maximalRectangle(matrix)
		}
	})
}
