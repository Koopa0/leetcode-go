package SetMatrixZeroes

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
		desc   string
	}{
		{
			name:   "基本功能測試 1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
			desc:   "矩陣中間有一個 0",
		},
		{
			name:   "基本功能測試 2",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
			desc:   "第一行有兩個 0",
		},
		{
			name:   "邊界情況測試 1",
			matrix: [][]int{{0}},
			want:   [][]int{{0}},
			desc:   "單一元素矩陣",
		},
		{
			name:   "邊界情況測試 2",
			matrix: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			want:   [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			desc:   "沒有 0 的矩陣",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製輸入矩陣，避免修改原始測試數據
			input := make([][]int, len(tt.matrix))
			for i, row := range tt.matrix {
				input[i] = make([]int, len(row))
				copy(input[i], row)
			}

			setZeroes(input)

			// 比較結果
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", input, tt.want)
			}
		})
	}
}

func BenchmarkSetZeroes(b *testing.B) {
	// 創建一個較大的測試矩陣
	m, n := 50, 50
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			// 隨機生成 0 或非 0 值
			if rand.Intn(10) == 0 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = rand.Intn(100) + 1
			}
		}
	}

	// 備份原矩陣用於各個測試
	copyMatrix := func() [][]int {
		c := make([][]int, m)
		for i := range matrix {
			c[i] = make([]int, n)
			for j := range matrix[i] {
				c[i][j] = matrix[i][j]
			}
		}
		return c
	}

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cp := copyMatrix()
			setZeroesBruteForce(cp)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cp := copyMatrix()
			setZeroesOptimized(cp)
		}
	})

	b.Run("Optimal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cp := copyMatrix()
			setZeroes(cp)
		}
	})
}
