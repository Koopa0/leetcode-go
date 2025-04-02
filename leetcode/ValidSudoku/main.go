package ValidSudoku

func isValidSudoku(board [][]byte) bool {
	// 初始化映射來跟踪已見數字
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// 遍歷數獨板的每個單元格
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// 跳過空格
			if board[row][col] == '.' {
				continue
			}

			digit := board[row][col]

			// 計算3x3子方格的索引（0-8）
			boxIndex := (row/3)*3 + (col / 3)

			// 檢查數字是否已經在行、列或子方格中出現過
			if rows[row][digit] || cols[col][digit] || boxes[boxIndex][digit] {
				return false // 無效的數獨
			}

			// 將數字添加到對應的映射中
			rows[row][digit] = true
			cols[col][digit] = true
			boxes[boxIndex][digit] = true
		}
	}

	// 如果遍歷完整個數獨板沒有發現重複，則數獨有效
	return true
}
