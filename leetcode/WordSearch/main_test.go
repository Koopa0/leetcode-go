package WordSearch

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
		desc     string
	}{
		{
			name: "基本測試1",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
			desc:     "測試基本路徑查找",
		},
		{
			name: "基本測試2",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
			desc:     "測試另一個有效路徑",
		},
		{
			name: "基本測試3",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
			desc:     "測試不允許重複使用同一個格子",
		},
		{
			name:     "空網格",
			board:    [][]byte{},
			word:     "A",
			expected: false,
			desc:     "測試空網格的處理",
		},
		{
			name: "單格網格",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
			desc:     "測試單一格子的情況",
		},
		{
			name: "長單字",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCESCFSADEE",
			expected: true,
			desc:     "測試長單字路徑",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 深拷貝 board 以避免修改原始測試資料
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}

			got := exist(boardCopy, tt.word)
			if got != tt.expected {
				t.Errorf("exist() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試資料
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 深拷貝 board
			boardCopy := make([][]byte, len(board))
			for j := range board {
				boardCopy[j] = make([]byte, len(board[j]))
				copy(boardCopy[j], board[j])
			}
			bruteForceExist(boardCopy, word)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 深拷貝 board
			boardCopy := make([][]byte, len(board))
			for j := range board {
				boardCopy[j] = make([]byte, len(board[j]))
				copy(boardCopy[j], board[j])
			}
			optimizedExist(boardCopy, word)
		}
	})

	b.Run("Optimal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 深拷貝 board
			boardCopy := make([][]byte, len(board))
			for j := range board {
				boardCopy[j] = make([]byte, len(board[j]))
				copy(boardCopy[j], board[j])
			}
			exist(boardCopy, word)
		}
	})
}
