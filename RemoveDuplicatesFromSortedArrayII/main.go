package RemoveDuplicatesFromSortedArrayII

// 暴力解法
func removeDuplicatesBruteForce(nums []int) int {
	// 處理邊界情況
	if len(nums) <= 2 {
		return len(nums)
	}

	// 創建一個額外的陣列來存儲結果
	result := []int{}
	count := 1

	// 將第一個元素加入結果
	result = append(result, nums[0])

	// 遍歷剩餘的陣列
	for i := 1; i < len(nums); i++ {
		// 如果當前元素與前一個元素不同，重置計數器
		if nums[i] != nums[i-1] {
			result = append(result, nums[i])
			count = 1
		} else if count < 2 { // 如果當前元素與前一個元素相同且計數 < 2
			result = append(result, nums[i])
			count++
		}
		// 如果計數已達到 2，跳過該元素
	}

	// 將結果複製回原陣列
	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	return len(result)
}

// 最佳解法 - 雙指針
func removeDuplicatesTwoPointers(nums []int) int {
	// 處理邊界情況
	n := len(nums)
	if n <= 2 {
		return n
	}

	// slow 指向要寫入結果的位置
	// fast 用於遍歷原陣列
	slow := 2

	// 從索引 2 開始遍歷
	for fast := 2; fast < n; fast++ {
		// 如果當前元素不是第三個重複的元素
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		// 否則跳過該元素
	}

	return slow
}

func removeDuplicates(nums []int) int {

	index := 1
	occurance := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			occurance++
		} else {
			occurance = 1
		}

		if occurance <= 2 {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
