package N_Queens_II

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
		desc     string
	}{
		{
			name:     "n=1",
			n:        1,
			expected: 1,
			desc:     "測試 n=1 的情況",
		},
		{
			name:     "n=4",
			n:        4,
			expected: 2,
			desc:     "測試 n=4 的情況",
		},
		{
			name:     "n=8",
			n:        8,
			expected: 92,
			desc:     "測試 n=8 的情況（經典國際象棋棋盤大小）",
		},
		{
			name:     "n=9",
			n:        9,
			expected: 352,
			desc:     "測試 n=9 的情況（最大允許輸入）",
		},
		{
			name:     "n=2",
			n:        2,
			expected: 0,
			desc:     "測試無解的情況",
		},
		{
			name:     "n=3",
			n:        3,
			expected: 0,
			desc:     "測試另一個無解的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalNQueens(tt.n)
			if got != tt.expected {
				t.Errorf("totalNQueens(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 選擇輸入 n=8 進行基準測試
	n := 8

	b.Run("BitOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			totalNQueens(n)
		}
	})
}
