package N_Queens

import "testing"

func TestSolveNQueens(t *testing.T) {
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
			desc:     "1x1 棋盤應該有 1 個解法",
		},
		{
			name:     "n=4",
			n:        4,
			expected: 2,
			desc:     "4x4 棋盤應該有 2 個解法",
		},
		{
			name:     "n=8",
			n:        8,
			expected: 92,
			desc:     "8x8 棋盤應該有 92 個解法",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solutions := solveNQueens(tt.n)
			if len(solutions) != tt.expected {
				t.Errorf("solveNQueens(%d) = %v solutions, want %v", tt.n, len(solutions), tt.expected)
			}
		})
	}
}
