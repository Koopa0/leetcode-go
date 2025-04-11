package N_Queens

import "math/bits"

// solveNQueens 函數解決 N-Queens 問題，返回所有可能的解法
// 參數 n：棋盤大小和皇后數量
// 返回值：包含所有解法的二維字串陣列，每個解法表示一種有效的皇后放置方式
func solveNQueens(n int) [][]string {
	// 初始化結果陣列，用於存儲所有找到的解法
	// 這個結構將收集所有有效的棋盤配置
	var result [][]string

	// 創建棋盤，使用二維 byte 陣列表示
	// 選擇 byte 而非 string 是為了提高修改效率，因為在回溯過程中需要頻繁地放置和移除皇后
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		// 初始化棋盤，所有位置都設為 '.'（空格）
		// 這代表初始狀態下棋盤是空的，沒有皇后
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// constructResult 是一個輔助函數，用於將當前棋盤狀態轉換為題目要求的輸出格式
	// 當找到一個解法時，我們將棋盤的當前狀態轉換為字串陣列並添加到結果集中
	// 這個函數僅在找到完整解法時被調用
	constructResult := func() []string {
		solution := make([]string, n)
		for i := range board {
			// 將每一行的 byte 陣列轉換為單個字串
			// 這樣做是因為題目要求返回字串形式的棋盤表示
			solution[i] = string(board[i])
		}
		return solution
	}

	// dfs 是主要的遞迴回溯函數，用於逐行放置皇后
	// 使用深度優先搜索遍歷所有可能的皇后放置方式
	// 參數說明：
	// - row: 當前處理的行（從 0 開始）
	// - cols: 使用位運算表示已被占用的列（第 i 位為 1 表示第 i 列已被占用）
	// - diag1: 使用位運算表示已被占用的主對角線（↗↙方向，從左上到右下）
	// - diag2: 使用位運算表示已被占用的副對角線（↖↘方向，從右上到左下）
	var dfs func(row, cols, diag1, diag2 int)
	dfs = func(row, cols, diag1, diag2 int) {
		// 基本情況：如果已經處理完所有行（放置了 n 個皇后）
		// 意味著我們找到了一個有效解法
		// 當 row == n 時，表示已經成功處理了 0 到 n-1 的所有行
		if row == n {
			// 將當前棋盤狀態轉換為需要的格式，添加到結果集
			// 這時棋盤上已經有了 n 個皇后，且彼此不衝突
			result = append(result, constructResult())
			return
		}

		// 計算當前行所有可以放置皇后的列位置
		// 這是整個算法的核心部分，使用位運算高效計算
		// 步驟解析：
		// 1. cols | diag1 | diag2：合併所有被攻擊的位置（位為1表示被攻擊）
		//    - cols 記錄了哪些列已有皇后（第 i 位為 1 表示第 i 列已被占用）
		//    - diag1 記錄了哪些主對角線已有皇后
		//    - diag2 記錄了哪些副對角線已有皇后
		// 2. ^(...)：取反，將被攻擊的位置變為0，安全位置變為1
		//    這樣，位為 1 的位置就是可以放置皇后的安全位置
		// 3. & ((1 << n) - 1)：應用掩碼，只保留棋盤範圍內的位（低n位）
		//    例如：當 n=8 時，掩碼是 11111111，確保我們只考慮有效的列位置
		availablePositions := ^(cols | diag1 | diag2) & ((1 << n) - 1)

		// 遍歷所有可用的位置（為1的位表示可用）
		// 只要還有可用位置（availablePositions > 0），就繼續嘗試
		// 這個迴圈會嘗試當前行的每一個安全位置
		for availablePositions > 0 {
			// 從可用位置中選取最右邊的1（最低有效位）
			// 技巧：x & -x 可以提取整數 x 中的最右邊的1
			// 原理：-x 是 x 的二補數（取反加1），與原數進行 AND 操作後只保留最右邊的1
			// 例如：若 availablePositions = 00101100，則 position = 00000100
			position := availablePositions & -availablePositions

			// 計算選中位置對應的列索引
			// bits.TrailingZeros 返回整數中最低位1的位置
			// 例如：0100 中最低位1在位置2（從0開始計數）
			// 這樣我們就知道要在第 col 列放置皇后
			col := bits.TrailingZeros(uint(position))

			// 在選定位置放置皇后
			// 將棋盤上對應位置標記為 'Q'，表示放置了皇后
			board[row][col] = 'Q'

			// 遞迴處理下一行，更新三種約束條件
			// 1. cols | position：標記當前列為已占用
			//    將新放置的皇后所在列添加到被占用的列集合中
			// 2. (diag1 | position) << 1：更新主對角線約束並左移1位
			//    - 左移原因：在下一行，相同主對角線的列索引會減少1
			//    - 主對角線特性：r+c 值相同的格子位於同一主對角線
			//    - 例如：位置(r,c)的主對角線同時包含(r+1,c-1)
			// 3. (diag2 | position) >> 1：更新副對角線約束並右移1位
			//    - 右移原因：在下一行，相同副對角線的列索引會增加1
			//    - 副對角線特性：r-c 值相同的格子位於同一副對角線
			//    - 例如：位置(r,c)的副對角線同時包含(r+1,c+1)
			dfs(row+1,
				cols|position,       // 更新列約束
				(diag1|position)<<1, // 更新主對角線約束
				(diag2|position)>>1) // 更新副對角線約束

			// 回溯：撤銷當前選擇，恢復棋盤狀態
			// 這是回溯算法的關鍵步驟，允許我們嘗試所有可能的配置
			// 當遞迴返回時，意味著基於當前選擇的所有可能性都已探索完畢
			// 現在需要撤銷這個選擇，嘗試下一個可能的位置
			board[row][col] = '.'

			// 從可用位置中移除剛才嘗試過的位置（清除最右邊的1）
			// 技巧：x & (x-1) 可以清除整數 x 中的最右邊的1
			// 原理：當計算 x-1 時，最右邊的1變成0，其右側的0變成1
			//      然後與原數進行 AND 操作，結果就是清除了最右邊的1
			// 例如：00101100 & 00101011 = 00101000
			availablePositions &= availablePositions - 1
		}
	}

	// 啟動遞迴搜索，從第0行開始，初始沒有任何約束
	// 參數分別是：當前行、已占用的列、已占用的主對角線、已占用的副對角線
	// 初始值都是0，表示一開始棋盤上沒有任何皇后，沒有任何位置被攻擊
	dfs(0, 0, 0, 0)

	// 返回所有找到的解法
	// 此時 result 中包含了所有符合條件的棋盤配置
	return result
}
