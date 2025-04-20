package SetMatrixZeroes

// 暴力解法
func setZeroesBruteForce(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	// 1. 創建矩陣的副本
	cp := make([][]int, m)
	for i := range cp {
		cp[i] = make([]int, n)
		for j := range matrix[i] {
			cp[i][j] = matrix[i][j]
		}
	}

	// 2. 根據副本中的 0 修改原矩陣
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cp[i][j] == 0 {
				// 將該行設為 0
				for col := 0; col < n; col++ {
					matrix[i][col] = 0
				}
				// 將該列設為 0
				for row := 0; row < m; row++ {
					matrix[row][j] = 0
				}
			}
		}
	}
}

// 優化解法
func setZeroesOptimized(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	// 記錄哪些行和列需要設為 0
	rowZero := make([]bool, m)
	colZero := make([]bool, n)

	// 標記含有 0 的行和列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}

	// 根據標記設置 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowZero[i] || colZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 最佳解法
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	firstRowHasZero := false
	firstColHasZero := false

	// 檢查第一行是否有 0
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// 檢查第一列是否有 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// 用第一行和第一列作為標記
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 標記該行需要設為 0
				matrix[0][j] = 0 // 標記該列需要設為 0
			}
		}
	}

	// 根據第一行和第一列的標記更新矩陣
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 如果第一行原本有 0，將整行設為 0
	if firstRowHasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 如果第一列原本有 0，將整列設為 0
	if firstColHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
