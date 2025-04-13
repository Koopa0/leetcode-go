package SpiralMatrixII

// 邊界收縮法實現
func generateMatrixBoundary(n int) [][]int {
	// 初始化 n×n 矩陣
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 定義邊界
	left, right := 0, n-1
	top, bottom := 0, n-1
	num := 1

	// 螺旋填充
	for num <= n*n {
		// 從左到右填充
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// 從上到下填充
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// 從右到左填充
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// 從下到上填充
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

// 方向陣列法實現
func generateMatrixDirection(n int) [][]int {
	// 初始化 n×n 矩陣
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 定義方向: 右、下、左、上
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col := 0, 0 // 起始位置
	dirIndex := 0    // 方向索引

	// 填充從 1 到 n²
	for num := 1; num <= n*n; num++ {
		// 在當前位置填入數字
		matrix[row][col] = num

		// 計算下一個位置
		nextRow := row + directions[dirIndex][0]
		nextCol := col + directions[dirIndex][1]

		// 如果下一個位置超出邊界或已填充，改變方向
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || matrix[nextRow][nextCol] != 0 {
			dirIndex = (dirIndex + 1) % 4
		}

		// 移動到下一個位置
		row += directions[dirIndex][0]
		col += directions[dirIndex][1]
	}

	return matrix
}

// 層次遍歷法實現
func generateMatrix(n int) [][]int {
	// 初始化 n×n 矩陣
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1              // 起始填充數字
	layers := (n + 1) / 2 // 層數

	for layer := 0; layer < layers; layer++ {
		// 計算當前層的大小
		len := n - 2*layer

		// 處理中心點的特殊情況
		if len == 1 {
			matrix[layer][layer] = num
			break
		}

		// 填充上邊
		for i := 0; i < len-1; i++ {
			matrix[layer][layer+i] = num
			num++
		}

		// 填充右邊
		for i := 0; i < len-1; i++ {
			matrix[layer+i][layer+len-1] = num
			num++
		}

		// 填充下邊
		for i := 0; i < len-1; i++ {
			matrix[layer+len-1][layer+len-1-i] = num
			num++
		}

		// 填充左邊
		for i := 0; i < len-1; i++ {
			matrix[layer+len-1-i][layer] = num
			num++
		}
	}

	return matrix
}
