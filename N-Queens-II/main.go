package N_Queens_II

func totalNQueens(n int) int {
	count := 0 // 解法計數器

	var backtrack func(int, int, int, int)
	backtrack = func(row, cols, posDiag, negDiag int) {
		if row == n {
			count++
			return
		}

		// 獲取所有可能放置皇后的位置
		// ~(cols | posDiag | negDiag) 得到所有非攻擊位置
		// & ((1 << n) - 1) 保留僅 n 個位元
		availablePositions := ^(cols | posDiag | negDiag) & ((1 << n) - 1)

		// 遍歷所有可能的位置
		for availablePositions > 0 {
			// 獲取最低位的 1（代表一個可放置的位置）
			position := availablePositions & -availablePositions

			// 更新 availablePositions，移除已處理的位置
			availablePositions -= position

			// 計算列位置
			col := 0
			for p := position; p > 1; p >>= 1 {
				col++
			}

			// 遞迴下一行
			backtrack(
				row+1,
				cols|position,
				(posDiag|position)<<1,
				(negDiag|position)>>1,
			)
			// 不需要顯式回溯，因為我們傳遞的是值而非引用
		}
	}

	backtrack(0, 0, 0, 0)
	return count
}
