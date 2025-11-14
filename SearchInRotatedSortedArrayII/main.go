package SearchInRotatedSortedArrayII

// 暴力解法
func bruteForceSearch(nums []int, target int) bool {
	// 1. 初始化
	n := len(nums)

	// 2. 核心暴力邏輯：遍歷整個陣列
	for i := 0; i < n; i++ {
		if nums[i] == target {
			return true
		}
	}

	// 3. 結果處理與返回：未找到目標值
	return false
}

// 優化解法：修改的二分搜尋
func optimizedSearch(nums []int, target int) bool {
	// 1. 初始化
	left, right := 0, len(nums)-1

	// 2. 核心優化邏輯：二分搜尋的變體
	for left <= right {
		mid := left + (right-left)/2

		// 找到目標值
		if nums[mid] == target {
			return true
		}

		// 處理左邊界與中間元素相同的特殊情況
		// 這是重複元素引入的複雜性
		if nums[left] == nums[mid] {
			left++
			continue
		}

		// 判斷中間元素在哪個排序子陣列中
		if nums[left] <= nums[mid] {
			// 中間元素在左側排序子陣列
			if nums[left] <= target && target < nums[mid] {
				// 目標在左側子陣列
				right = mid - 1
			} else {
				// 目標在右側子陣列
				left = mid + 1
			}
		} else {
			// 中間元素在右側排序子陣列
			if nums[mid] < target && target <= nums[right] {
				// 目標在右側子陣列
				left = mid + 1
			} else {
				// 目標在左側子陣列
				right = mid - 1
			}
		}
	}

	// 3. 結果處理與返回：未找到目標值
	return false
}

// 最佳解法：高度優化的二分搜尋
func search(nums []int, target int) bool {
	// 1. 初始化，使用最佳策略
	left, right := 0, len(nums)-1

	// 2. 核心最佳邏輯
	for left <= right {
		mid := left + (right-left)/2

		// 找到目標值
		if nums[mid] == target {
			return true
		}

		// 處理左邊界、中間元素、右邊界三者相等的特殊情況
		// 這種情況下無法確定搜尋方向，只能逐步縮小範圍
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		// 判斷中間元素在哪個排序子陣列中
		if nums[left] <= nums[mid] {
			// 中間元素在左側排序子陣列
			if nums[left] <= target && target < nums[mid] {
				// 目標在左側子陣列
				right = mid - 1
			} else {
				// 目標在右側子陣列
				left = mid + 1
			}
		} else {
			// 中間元素在右側排序子陣列
			if nums[mid] < target && target <= nums[right] {
				// 目標在右側子陣列
				left = mid + 1
			} else {
				// 目標在左側子陣列
				right = mid - 1
			}
		}
	}

	// 3. 結果處理與返回
	return false
}
