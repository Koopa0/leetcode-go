package TrappingRainWater

func trap(height []int) int {
	// 邊界情況處理
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	totalWater := 0

	for left < right {
		// 決定移動哪個指針
		if leftMax < rightMax {
			left++
			// 更新leftMax或計算捕獲的水
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
		} else {
			right--
			// 更新rightMax或計算捕獲的水
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
		}
	}

	return totalWater
}
