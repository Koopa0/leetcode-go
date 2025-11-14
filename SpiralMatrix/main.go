package SpiralMatrix

func spiralOrderBruteForce(matrix [][]int) []int {
	// 處理邊界情況
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 1. 初始化
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*cols)

	// 定義四個邊界
	left, right := 0, cols-1
	top, bottom := 0, rows-1

	// 2. 核心暴力遍歷邏輯
	for left <= right && top <= bottom {
		// 向右遍歷
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// 向下遍歷
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 檢查邊界條件，避免重複遍歷
		if top > bottom || left > right {
			break
		}

		// 向左遍歷
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		// 向上遍歷
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	// 3. 結果處理與返回
	return result
}

// 優化解法 - 使用方向陣列
func spiralOrderOptimized(matrix [][]int) []int {
	// 處理邊界情況
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 1. 初始化帶優化的資料結構
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*cols)

	// 定義方向：右、下、左、上
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// 創建訪問標記矩陣
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 2. 核心優化邏輯
	row, col, dirIdx := 0, 0, 0 // 初始位置和方向
	totalElements := rows * cols

	for i := 0; i < totalElements; i++ {
		// 添加當前元素到結果
		result = append(result, matrix[row][col])
		visited[row][col] = true

		// 計算下一個位置
		nextRow := row + directions[dirIdx][0]
		nextCol := col + directions[dirIdx][1]

		// 檢查下一個位置是否有效，如果無效則改變方向
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			// 改變方向 (右->下->左->上->右...)
			dirIdx = (dirIdx + 1) % 4
			nextRow = row + directions[dirIdx][0]
			nextCol = col + directions[dirIdx][1]
		}

		// 更新當前位置
		row, col = nextRow, nextCol
	}

	// 3. 結果處理與返回
	return result
}

// 最佳解法 - 邊界收縮法
func spiralOrder(matrix [][]int) []int {
	// 處理邊界情況
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 1. 初始化帶最佳策略
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*cols)

	// 定義四個邊界
	left, right := 0, cols-1
	top, bottom := 0, rows-1

	// 2. 核心最佳邏輯
	for left <= right && top <= bottom {
		// 從左到右遍歷 top 行
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// 從上到下遍歷 right 列
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// 檢查是否已經遍歷完畢
		if top > bottom || left > right {
			break
		}

		// 從右到左遍歷 bottom 行
		for col := right; col >= left; col-- {
			result = append(result, matrix[bottom][col])
		}
		bottom--

		// 從下到上遍歷 left 列
		for row := bottom; row >= top; row-- {
			result = append(result, matrix[row][left])
		}
		left++
	}

	// 3. 結果處理與返回
	return result
}
