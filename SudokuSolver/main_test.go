package SudokuSolver

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name          string
		board         [][]byte
		expectedBoard [][]byte
		description   string
	}{
		{
			name: "標準9x9數獨",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expectedBoard: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
			description: "LeetCode提供的標準測試用例",
		},
		{
			name: "幾乎填滿的數獨",
			board: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '.'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
			expectedBoard: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
			description: "只有一個空格的邊緣情況",
		},
		{
			name: "空數獨",
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			expectedBoard: nil, // 不檢查具體解，只驗證解題器不會崩潰
			description:   "測試完全空白的數獨",
		},
		{
			name: "困難數獨",
			board: [][]byte{
				{'.', '.', '5', '3', '.', '.', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
				{'.', '7', '.', '.', '1', '.', '5', '.', '.'},
				{'4', '.', '.', '.', '.', '5', '3', '.', '.'},
				{'.', '1', '.', '.', '7', '.', '.', '.', '6'},
				{'.', '.', '3', '2', '.', '.', '.', '8', '.'},
				{'.', '6', '.', '5', '.', '.', '.', '.', '9'},
				{'.', '.', '4', '.', '.', '.', '.', '3', '.'},
				{'.', '.', '.', '.', '.', '9', '7', '.', '.'},
			},
			expectedBoard: nil, // 不檢查具體解，只驗證解題器不會崩潰
			description:   "測試一個公認困難的數獨",
		},
		{
			name: "性能測試 - 難數獨",
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '3', '.', '8', '5'},
				{'.', '.', '1', '.', '2', '.', '.', '.', '.'},
				{'.', '.', '.', '5', '.', '7', '.', '.', '.'},
				{'.', '.', '4', '.', '.', '.', '1', '.', '.'},
				{'.', '9', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '7', '3'},
				{'.', '.', '2', '.', '1', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '4', '.', '.', '.', '9'},
			},
			expectedBoard: nil, // 不檢查具體解，只驗證解題器不會崩潰
			description:   "測試算法對於較難數獨的性能",
		},
	}

	// 運行每個測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製輸入板，以免修改測試數據
			inputCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				inputCopy[i] = make([]byte, len(tt.board[i]))
				copy(inputCopy[i], tt.board[i])
			}

			// 調用解題函數
			solveSudoku(inputCopy)

			// 如果期望值為nil，說明我們不檢查具體結果，只驗證解題器不會崩潰
			if tt.expectedBoard != nil {
				if !reflect.DeepEqual(inputCopy, tt.expectedBoard) {
					t.Errorf("solveSudoku() = %v, want %v", formatBoard(inputCopy), formatBoard(tt.expectedBoard))
				}
			}

			// 對於所有情況，檢查結果是否是有效的數獨解
			if isValidSudoku(inputCopy) {
				// 檢查是否所有單元格都已填充
				allFilled := true
				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						if inputCopy[i][j] == '.' {
							allFilled = false
							break
						}
					}
					if !allFilled {
						break
					}
				}

				if !allFilled {
					t.Logf("數獨未完全填充，但結果有效")
				}
			} else {
				t.Errorf("解出的數獨無效")
			}
		})
	}
}

// 輔助函數：檢查數獨解是否有效
func isValidSudoku(board [][]byte) bool {
	// 檢查行
	for i := 0; i < 9; i++ {
		used := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if used[board[i][j]] {
				return false
			}
			used[board[i][j]] = true
		}
	}

	// 檢查列
	for j := 0; j < 9; j++ {
		used := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			if used[board[i][j]] {
				return false
			}
			used[board[i][j]] = true
		}
	}

	// 檢查3x3子網格
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			used := make(map[byte]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					row, col := boxRow*3+i, boxCol*3+j
					if board[row][col] == '.' {
						continue
					}
					if used[board[row][col]] {
						return false
					}
					used[board[row][col]] = true
				}
			}
		}
	}

	return true
}

// 輔助函數：格式化數獨網格以便於打印
func formatBoard(board [][]byte) string {
	result := "[\n"
	for i := 0; i < 9; i++ {
		result += "  ["
		for j := 0; j < 9; j++ {
			if j > 0 {
				result += ","
			}
			result += string(board[i][j])
		}
		result += "],\n"
	}
	result += "]"
	return result
}

// 性能測試
func BenchmarkSolveSudoku(b *testing.B) {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '8', '5'},
		{'.', '.', '1', '.', '2', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '7', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '1', '.', '.'},
		{'.', '9', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '7', '3'},
		{'.', '.', '2', '.', '1', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '4', '.', '.', '.', '9'},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次測試前複製原始數獨
		boardCopy := make([][]byte, len(board))
		for i := range board {
			boardCopy[i] = make([]byte, len(board[i]))
			copy(boardCopy[i], board[i])
		}

		solveSudoku(boardCopy)
	}
}

// 測試極端情況的空格分布
func TestSolveSudokuWithDifferentEmptyCellDistributions(t *testing.T) {
	testCases := []struct {
		name        string
		createBoard func() [][]byte
	}{
		{
			name: "左上角集中空格",
			createBoard: func() [][]byte {
				// 創建一個大部分已填充，但左上角3x3子網格全為空的數獨
				board := create9x9Board()
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						board[i][j] = '.'
					}
				}
				return board
			},
		},
		{
			name: "對角線空格",
			createBoard: func() [][]byte {
				// 創建一個只有對角線為空的數獨
				board := create9x9Board()
				for i := 0; i < 9; i++ {
					board[i][i] = '.'
				}
				return board
			},
		},
		{
			name: "棋盤狀空格",
			createBoard: func() [][]byte {
				// 創建一個棋盤狀的空格分布
				board := create9x9Board()
				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						if (i+j)%2 == 0 {
							board[i][j] = '.'
						}
					}
				}
				return board
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			board := tc.createBoard()
			solveSudoku(board)

			if !isValidSudoku(board) {
				t.Errorf("%s: 解出的數獨無效", tc.name)
			}

			// 檢查是否所有單元格都已填充
			allFilled := true
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if board[i][j] == '.' {
						allFilled = false
						break
					}
				}
				if !allFilled {
					break
				}
			}

			if !allFilled {
				t.Errorf("%s: 數獨未完全填充", tc.name)
			}
		})
	}
}

// 創建一個完全填充的有效9x9數獨（用於測試目的）
func create9x9Board() [][]byte {
	return [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}
}
