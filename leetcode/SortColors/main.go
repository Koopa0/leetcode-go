package SortColors

// 計數排序解法
func countingSortSortColors(nums []int) {
	// 1. 初始化計數器
	count := [3]int{}

	// 2. 統計每個顏色的數量
	for _, num := range nums {
		count[num]++
	}

	// 3. 根據計數重建陣列
	index := 0

	// 放入所有的紅色 (0)
	for i := 0; i < count[0]; i++ {
		nums[index] = 0
		index++
	}

	// 放入所有的白色 (1)
	for i := 0; i < count[1]; i++ {
		nums[index] = 1
		index++
	}

	// 放入所有的藍色 (2)
	for i := 0; i < count[2]; i++ {
		nums[index] = 2
		index++
	}
}

// 三路快排解法
func threeWayPartitioningSortColors(nums []int) {
	// 1. 初始化三個指針
	left := 0              // 0（紅色）區域的右邊界
	current := 0           // 當前處理的元素
	right := len(nums) - 1 // 2（藍色）區域的左邊界

	// 2. 遍歷陣列
	for current <= right {
		if nums[current] == 0 {
			// 當前元素是紅色，放到左邊
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		} else if nums[current] == 1 {
			// 當前元素是白色，保持在中間
			current++
		} else {
			// 當前元素是藍色，放到右邊
			nums[current], nums[right] = nums[right], nums[current]
			right--
			// 注意：current 不遞增，因為交換後的元素還需要檢查
		}
	}
}

// 雙指針解法
func twoPointersSortColors(nums []int) {
	// 1. 初始化雙指針
	p0 := 0 // 下一個 0 應該放的位置
	p1 := 0 // 下一個 1 應該放的位置

	// 2. 遍歷陣列
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// 遇到 0，與 p0 位置交換
			nums[i], nums[p0] = nums[p0], nums[i]

			// 如果 p0 < p1，說明 p0 位置原本是 1
			// 交換後 1 被移到了 i 位置，需要再把它移到 p1 位置
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}

			// 更新指針
			p0++
			p1++
		} else if nums[i] == 1 {
			// 遇到 1，與 p1 位置交換
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
		// 遇到 2 時不做任何處理，讓它留在後面
	}
}
