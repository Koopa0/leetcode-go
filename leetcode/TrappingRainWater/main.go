package TrappingRainWater

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	totalWater := 0

	for left < right {

		if height[left] < height[right] {

			if height[left] >= leftMax {
				leftMax = height[left]
			} else {

				totalWater += leftMax - height[left]
			}
			left++
		} else {

			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}

	return totalWater
}
