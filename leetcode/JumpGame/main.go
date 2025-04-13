package JumpGame

// 暴力解法
func canJumpBruteForce(nums []int) bool {
	// 輔助函數：檢查從某位置是否可達終點
	var canJumpFromPosition func(position int) bool
	canJumpFromPosition = func(position int) bool {
		// 已到達最後位置
		if position == len(nums)-1 {
			return true
		}

		// 計算能跳到的最遠距離（不超過陣列長度）
		furthestJump := min(position+nums[position], len(nums)-1)

		// 嘗試所有可能的跳躍
		for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
			if canJumpFromPosition(nextPosition) {
				return true
			}
		}

		// 所有可能的跳躍都無法到達終點
		return false
	}

	// 從位置 0 開始
	return canJumpFromPosition(0)
}

// 動態規劃解法（帶記憶化）
func canJumpDP(nums []int) bool {
	// 定義狀態
	const (
		UNKNOWN = 0
		GOOD    = 1
		BAD     = 2
	)

	// 記憶表，初始化為未知狀態
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = UNKNOWN
	}
	// 終點總是可達的
	memo[len(nums)-1] = GOOD

	// 從某個位置是否可達終點
	var canJumpFromPosition func(position int) bool
	canJumpFromPosition = func(position int) bool {
		// 如果已計算過該位置的狀態，直接返回
		if memo[position] != UNKNOWN {
			return memo[position] == GOOD
		}

		// 計算能跳到的最遠距離
		furthestJump := min(position+nums[position], len(nums)-1)

		// 嘗試所有可能的跳躍
		for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
			if canJumpFromPosition(nextPosition) {
				memo[position] = GOOD
				return true
			}
		}

		// 所有可能的跳躍都無法到達終點
		memo[position] = BAD
		return false
	}

	// 從位置 0 開始
	return canJumpFromPosition(0)
}

// 貪婪解法
func canJumpGreedy(nums []int) bool {
	// 最遠可達位置
	maxReach := 0

	// 遍歷陣列中的每個位置
	for i := 0; i < len(nums); i++ {
		// 如果當前位置已經超過了能到達的最遠位置，表示無法到達該位置
		if i > maxReach {
			return false
		}

		// 更新最遠可達位置
		maxReach = max(maxReach, i+nums[i])

		// 如果最遠可達位置已經覆蓋到終點，返回成功
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	// 理論上不會執行到這裡，因為在循環中會返回結果
	return false
}
