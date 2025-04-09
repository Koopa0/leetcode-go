package N_Queens

// 位元操作優化版本的 solveNQueens 函數
func solveNQueens(n int) [][]string {
	result := [][]string{}

	// 創建一個 n×n 的棋盤，初始化為全部 '.'
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// 位元掩碼：(1 << n) - 1 表示低 n 位全部為 1
	// 例如，n=4 時，掩碼為 0b1111
	mask := (1 << n) - 1

	// 回溯函數，使用位元操作優化
	var backtrack func(row, cols, posDiag, negDiag int)
	backtrack = func(row, cols, posDiag, negDiag int) {
		// 如果已經放置了 n 個皇后，將當前棋盤配置添加到結果中
		if row == n {
			solution := make([]string, n)
			for i := range board {
				solution[i] = string(board[i])
			}
			result = append(result, solution)
			return
		}

		// 計算所有可以放置皇后的位置
		// 1 表示可放置，0 表示不可放置
		availablePositions := ^(cols | posDiag | negDiag) & mask

		// 逐一嘗試所有可放置的位置
		for availablePositions != 0 {
			// 獲取最右邊的 1（最右邊可放置的位置）
			position := availablePositions & -availablePositions

			// 計算對應的列索引
			col := countTrailingZeros(position)

			// 放置皇后
			board[row][col] = 'Q'

			// 更新約束條件並遞迴到下一行
			// 列：將當前位置標記為已使用
			// 正對角線：左移一位（因為下一行 row+1，所以 row+col 增加 1）
			// 反對角線：右移一位（因為下一行 row+1，所以 row-col 減少 1）
			backtrack(row+1,
				cols|position,
				(posDiag|position)<<1,
				(negDiag|position)>>1)

			// 回溯：移除皇后
			board[row][col] = '.'

			// 移除當前位置，嘗試下一個位置
			availablePositions &= ^position
		}
	}

	// 從第一行開始回溯，初始沒有任何列或對角線被占用
	backtrack(0, 0, 0, 0)
	return result
}

// countTrailingZeros 計算一個整數二進制表示中尾部 0 的數量
// 例如，8 (1000) 有 3 個尾部 0
func countTrailingZeros(num int) int {
	if num == 0 {
		return 0
	}

	count := 0
	for (num & 1) == 0 {
		count++
		num >>= 1
	}
	return count
}
