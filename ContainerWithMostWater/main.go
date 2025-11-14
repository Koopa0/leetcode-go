package ContainerWithMostWater

func maxArea(height []int) int {
	// 初始化變量
	maxWater := 0
	left := 0
	right := len(height) - 1

	// 使用雙指針技術
	for left < right {
		// 計算當前容器的寬度
		width := right - left

		// 容器的高度由較短的線決定
		h := min(height[left], height[right])

		// 計算當前容器的容量並更新最大容量
		currentArea := width * h
		maxWater = max(maxWater, currentArea)

		// 移動指向較短線的指針
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
