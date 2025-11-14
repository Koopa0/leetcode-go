package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	// 初始化左右指針
	left, right := 0, len(nums)-1

	// 當搜索範圍有效時繼續搜索
	for left <= right {
		// 計算中間位置（避免整數溢出）
		mid := left + (right-left)/2

		// 找到目標值
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 目標在右半部分
			left = mid + 1
		} else {
			// 目標在左半部分
			right = mid - 1
		}
	}

	// 如果沒有找到目標值，left 指針指向正確的插入位置
	return left
}
