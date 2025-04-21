package RemoveDuplicatesfromSortedArray

func removeDuplicates(nums []int) int {
	// 處理邊界情況
	if len(nums) <= 1 {
		return len(nums)
	}

	// 初始化慢指針和快指針
	slow := 0

	// 使用快指針遍歷數組
	for fast := 1; fast < len(nums); fast++ {
		// 當找到一個與慢指針指向的元素不同的元素時
		if nums[fast] != nums[slow] {
			// 慢指針前進
			slow++
			// 將不同的元素放到慢指針的位置
			nums[slow] = nums[fast]
		}
		// 注意：如果元素相同，我們只移動快指針，慢指針不動
	}

	// 返回唯一元素的數量（慢指針位置 + 1）
	return slow + 1
}
