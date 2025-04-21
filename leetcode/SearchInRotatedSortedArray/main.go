package SearchInRotatedSortedArray

func search(nums []int, target int) int {
	// 初始化左右指針
	left, right := 0, len(nums)-1

	for left <= right {
		// 計算中間索引
		mid := left + (right-left)/2 // 避免整數溢出

		// 找到目標值
		if nums[mid] == target {
			return mid
		}

		// 判斷左半部分是否有序
		if nums[left] <= nums[mid] {
			// 目標值在左半部分的有序區間內
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // 在左半部分搜索
			} else {
				left = mid + 1 // 在右半部分搜索
			}
		} else { // 右半部分有序
			// 目標值在右半部分的有序區間內
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // 在右半部分搜索
			} else {
				right = mid - 1 // 在左半部分搜索
			}
		}
	}

	// 未找到目標值
	return -1
}
