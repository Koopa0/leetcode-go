package LargestRectangleInHistogram

// 暴力解法
func largestRectangleAreaBruteForce(heights []int) int {
	maxArea := 0
	n := len(heights)

	// 1. 初始化

	// 2. 雙層迴圈進行暴力搜尋
	for i := 0; i < n; i++ {
		minHeight := heights[i]
		for j := i; j < n; j++ {
			// 找出 i 到 j 之間的最小高度
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			// 計算面積並更新最大面積
			width := j - i + 1
			area := minHeight * width
			if area > maxArea {
				maxArea = area
			}
		}
	}

	// 3. 返回結果
	return maxArea
}

// 分治法解決方案
func largestRectangleAreaDivideConquer(heights []int) int {
	return divideConquer(heights, 0, len(heights)-1)
}

func divideConquer(heights []int, start, end int) int {
	// 1. 基本情況
	if start > end {
		return 0
	}
	if start == end {
		return heights[start]
	}

	// 2. 找出最小高度的索引
	minIndex := start
	for i := start; i <= end; i++ {
		if heights[i] < heights[minIndex] {
			minIndex = i
		}
	}

	// 3. 計算三種情況下的最大面積
	// 3.1 使用最小高度作為高度的矩形
	currentArea := heights[minIndex] * (end - start + 1)

	// 3.2 遞迴計算左側子陣列的最大矩形
	leftArea := divideConquer(heights, start, minIndex-1)

	// 3.3 遞迴計算右側子陣列的最大矩形
	rightArea := divideConquer(heights, minIndex+1, end)

	// 4. 返回三者中的最大值
	return max(currentArea, max(leftArea, rightArea))
}

// 單調堆疊最佳解法
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// 1. 初始化
	stack := make([]int, 0) // 存儲索引的單調遞增堆疊
	maxArea := 0

	// 2. 處理每個柱子
	for i := 0; i <= n; i++ {
		// 當到達陣列末尾或當前高度小於堆疊頂部元素的高度時
		var h int
		if i == n {
			h = 0 // 使用高度 0 來清空堆疊
		} else {
			h = heights[i]
		}

		// 當堆疊非空且當前高度小於堆疊頂部元素的高度
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // 彈出堆疊頂部

			// 計算寬度：右界(i) - 左界(stack頂部的右側)
			var width int
			if len(stack) == 0 {
				width = i // 左側沒有更小的元素，寬度為 i
			} else {
				width = i - stack[len(stack)-1] - 1 // 左界是新的堆疊頂部右側
			}

			// 計算面積並更新最大面積
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		// 將當前索引入堆疊
		stack = append(stack, i)
	}

	// 3. 返回結果
	return maxArea
}
