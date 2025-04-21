package WordSearch

// 暴力解法（帶回溯的 DFS）
func bruteForceExist(board [][]byte, word string) bool {
	// 1. 初始化
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])
	// 方向陣列，代表上、右、下、左四個方向
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// 2. 創建訪問標記矩陣
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 3. DFS 輔助函數
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// 成功條件：已匹配整個單字
		if index == len(word) {
			return true
		}

		// 檢查當前位置是否有效且匹配當前字元
		if i < 0 || i >= rows || j < 0 || j >= cols ||
			visited[i][j] || board[i][j] != word[index] {
			return false
		}

		// 標記當前位置為已訪問
		visited[i][j] = true

		// 嘗試四個方向
		for _, dir := range directions {
			newI, newJ := i+dir[0], j+dir[1]
			if dfs(newI, newJ, index+1) {
				return true
			}
		}

		// 回溯：取消標記
		visited[i][j] = false

		return false
	}

	// 4. 從每個可能的起始位置開始 DFS
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// 優化解法（直接在原網格上標記）
func optimizedExist(board [][]byte, word string) bool {
	// 1. 初始化
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])
	// 方向陣列，代表上、右、下、左四個方向
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// 2. DFS 輔助函數
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// 成功條件：已匹配整個單字
		if index == len(word) {
			return true
		}

		// 檢查當前位置是否有效且匹配當前字元
		if i < 0 || i >= rows || j < 0 || j >= cols ||
			board[i][j] != word[index] {
			return false
		}

		// 臨時保存當前值
		temp := board[i][j]
		// 標記為已訪問（使用特殊字元）
		board[i][j] = '#'

		// 嘗試四個方向
		for _, dir := range directions {
			newI, newJ := i+dir[0], j+dir[1]
			if dfs(newI, newJ, index+1) {
				return true
			}
		}

		// 回溯：恢復原值
		board[i][j] = temp

		return false
	}

	// 3. 從每個可能的起始位置開始 DFS
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// 最佳化解法（提前剪枝與方向優化）
func exist(board [][]byte, word string) bool {
	// 1. 初始化
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])

	// 2. 前處理：檢查字母頻率
	// 如果網格中的某個字母數量少於單字中需要的數量，直接返回 false
	boardFreq := make(map[byte]int)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			boardFreq[board[i][j]]++
		}
	}

	wordFreq := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		wordFreq[word[i]]++
	}

	for ch, count := range wordFreq {
		if boardFreq[ch] < count {
			return false
		}
	}

	// 3. 方向陣列，代表上、右、下、左四個方向
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// 4. DFS 輔助函數
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// 成功條件：已匹配整個單字
		if index == len(word) {
			return true
		}

		// 檢查當前位置是否有效且匹配當前字元
		if i < 0 || i >= rows || j < 0 || j >= cols ||
			board[i][j] != word[index] {
			return false
		}

		// 臨時保存當前值
		temp := board[i][j]
		// 標記為已訪問
		board[i][j] = '#'

		// 嘗試四個方向
		for _, dir := range directions {
			newI, newJ := i+dir[0], j+dir[1]
			if dfs(newI, newJ, index+1) {
				return true
			}
		}

		// 回溯：恢復原值
		board[i][j] = temp

		return false
	}

	// 5. 從每個可能的起始位置開始 DFS
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
