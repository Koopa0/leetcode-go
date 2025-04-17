package MinimumPathSum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
		desc     string
	}{
		{
			name:     "基本測試 1",
			grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
			desc:     "LeetCode 提供的範例 1",
		},
		{
			name:     "基本測試 2",
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
			desc:     "LeetCode 提供的範例 2",
		},
		{
			name:     "單行測試",
			grid:     [][]int{{1, 2, 3, 4, 5}},
			expected: 15,
			desc:     "單行網格，僅有一條路徑",
		},
		{
			name:     "單列測試",
			grid:     [][]int{{1}, {2}, {3}, {4}, {5}},
			expected: 15,
			desc:     "單列網格，僅有一條路徑",
		},
		{
			name:     "最小網格",
			grid:     [][]int{{5}},
			expected: 5,
			desc:     "1x1 網格，僅有一個元素",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPathSum(tt.grid)
			if got != tt.expected {
				t.Errorf("minPathSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkMinPathSum(b *testing.B) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	b.Run("暴力遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minPathSumBruteForce(grid)
		}
	})

	b.Run("記憶化遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minPathSumMemoization(grid)
		}
	})

	b.Run("動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minPathSum(grid)
		}
	})

	b.Run("空間優化動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minPathSumOptimized(grid)
		}
	})
}
