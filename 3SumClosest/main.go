package threeSumClosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	// 對數組進行排序
	sort.Ints(nums)

	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2] // 初始最接近和

	// 固定第一個數，然後使用雙指針尋找另外兩個數
	for i := 0; i < n-2; i++ {
		// 如果當前元素與前一個元素相同，跳過以避免重複計算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			// 計算當前三數之和
			sum := nums[i] + nums[left] + nums[right]

			// 如果找到完全匹配的和，直接返回
			if sum == target {
				return sum
			}

			// 更新最接近的和
			if abs(sum-target) < abs(closestSum-target) {
				closestSum = sum
			}

			if sum < target {
				// 和太小，左指針右移
				left++
			} else {
				// 和太大，右指針左移
				right--
			}
		}
	}

	return closestSum
}

// 計算絕對值的輔助函數
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
