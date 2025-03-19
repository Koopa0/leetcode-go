package ContainerWithMostWater

func maxArea(height []int) int {

	left, right := 0, len(height)-1

	maxArea := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
