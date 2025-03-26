package JumpGame

func canJump(nums []int) bool {
	// 數組的長度
	n := len(nums)

	// 如果數組只有一個元素，已經在最後一個位置
	if n == 1 {
		return true
	}

	// 初始化能到達的最遠位置
	maxReach := 0

	// 遍歷數組中的每個元素
	for i := 0; i < n; i++ {
		// 如果當前位置已經超出了能到達的最遠位置，返回 false
		if i > maxReach {
			return false
		}

		// 更新能到達的最遠位置
		maxReach = max(maxReach, i+nums[i])

		// 如果能到達的最遠位置已經超過或等於最後一個索引，返回 true
		if maxReach >= n-1 {
			return true
		}
	}

	// 如果遍歷完數組後仍未能到達最後一個索引，返回 false
	return false
}
