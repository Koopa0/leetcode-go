package MaximalRectangle

// 暴力解法
func maximalRectangleForce(matrix [][]byte) int {
	// 1. 初始化
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	maxArea := 0

	// 2. 核心暴力邏輯
	// 遍歷所有可能的左上角
	for r1 := 0; r1 < rows; r1++ {
		for c1 := 0; c1 < cols; c1++ {
			// 只考慮以 '1' 開始的左上角
			if matrix[r1][c1] == '1' {
				// 遍歷所有可能的右下角
				for r2 := r1; r2 < rows; r2++ {
					for c2 := c1; c2 < cols; c2++ {
						// 檢查矩形是否全為 '1'
						isValid := true
						for r := r1; r <= r2; r++ {
							for c := c1; c <= c2; c++ {
								if matrix[r][c] == '0' {
									isValid = false
									break
								}
							}
							if !isValid {
								break
							}
						}

						// 計算面積並更新最大值
						if isValid {
							area := (r2 - r1 + 1) * (c2 - c1 + 1)
							if area > maxArea {
								maxArea = area
							}
						}
					}
				}
			}
		}
	}

	// 3. 返回結果
	return maxArea
}

// 優化解法 - 使用高度直方圖
func maximalRectangleStack(matrix [][]byte) int {
	// 1. 初始化與邊界情況處理
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols) // 每列的高度
	maxArea := 0

	// 2. 核心優化邏輯
	for r := 0; r < rows; r++ {
		// 更新高度直方圖
		for c := 0; c < cols; c++ {
			if matrix[r][c] == '1' {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}

		// 計算當前行的最大矩形面積
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	// 3. 返回結果
	return maxArea
}

// 計算直方圖中的最大矩形面積 (LeetCode 84 的解法)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{} // 存儲索引的單調遞增堆疊
	maxArea := 0

	for i := 0; i <= n; i++ {
		var h int
		if i == n {
			h = 0 // 哨兵值，確保堆疊被清空
		} else {
			h = heights[i]
		}

		// 當前高度小於堆疊頂部高度時，計算堆疊頂部元素為高度的矩形面積
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // 彈出堆疊頂部

			var width int
			if len(stack) == 0 {
				width = i // 左邊界是 -1，右邊界是 i
			} else {
				width = i - stack[len(stack)-1] - 1 // 左邊界是新的堆疊頂部，右邊界是 i
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i) // 將當前索引推入堆疊
	}

	return maxArea
}

// 最佳解法 - 動態規劃優化
func maximalRectangle(matrix [][]byte) int {
	// 1. 初始化與邊界情況處理
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols) // 每列的高度
	left := make([]int, cols)    // 左邊界數組
	right := make([]int, cols)   // 右邊界數組

	// 初始化右邊界為列數
	for i := 0; i < cols; i++ {
		right[i] = cols
	}

	maxArea := 0

	// 2. 核心最佳邏輯
	for r := 0; r < rows; r++ {
		// 更新高度
		for c := 0; c < cols; c++ {
			if matrix[r][c] == '1' {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}

		// 計算左邊界
		leftBoundary := 0
		for c := 0; c < cols; c++ {
			if matrix[r][c] == '1' {
				// 左邊界是當前leftBoundary和上一行左邊界的最大值
				left[c] = max(left[c], leftBoundary)
			} else {
				// 重置左邊界和leftBoundary
				left[c] = 0
				leftBoundary = c + 1
			}
		}

		// 計算右邊界
		rightBoundary := cols
		for c := cols - 1; c >= 0; c-- {
			if matrix[r][c] == '1' {
				// 右邊界是當前rightBoundary和上一行右邊界的最小值
				right[c] = min(right[c], rightBoundary)
			} else {
				// 重置右邊界和rightBoundary
				right[c] = cols
				rightBoundary = c
			}
		}

		// 計算最大面積
		for c := 0; c < cols; c++ {
			area := heights[c] * (right[c] - left[c])
			maxArea = max(maxArea, area)
		}
	}

	// 3. 返回結果
	return maxArea
}
