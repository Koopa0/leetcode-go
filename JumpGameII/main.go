package JumpGameII

func jump(nums []int) int {
	// 如果數組長度為 1，直接返回 0（不需要跳躍）
	if len(nums) == 1 {
		return 0
	}

	jumps := 0    // 跳躍次數
	currEnd := 0  // 當前能到達的最遠位置
	farthest := 0 // 下一步能到達的最遠位置

	// 遍歷數組（注意：不需要遍歷最後一個元素）
	for i := 0; i < len(nums)-1; i++ {
		// 更新下一步能到達的最遠位置
		farthest = max(farthest, i+nums[i])

		// 如果已經到達當前跳躍能到達的邊界
		if i == currEnd {
			// 進行一次跳躍
			jumps++
			// 更新當前能到達的最遠位置
			currEnd = farthest

			// 優化：如果已經能到達終點，提前結束
			if currEnd >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}
