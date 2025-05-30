package Combinations

// 回溯法基本實現
func combineBacktrack(n int, k int) [][]int {
	result := [][]int{}

	// 回溯函數
	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		// 若當前組合已有 k 個元素，則加入結果列表
		if len(current) == k {
			// 創建一個新的切片以避免引用問題
			combination := make([]int, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// 嘗試從 start 到 n 的每個數字
		for i := start; i <= n; i++ {
			// 選擇當前數字
			current = append(current, i)

			// 遞迴處理下一個位置
			backtrack(i+1, current)

			// 回溯：撤銷選擇
			current = current[:len(current)-1]
		}
	}

	// 從空組合開始回溯
	backtrack(1, []int{})

	return result
}

// 優化的回溯法實現
func combineBacktrackOptimized(n int, k int) [][]int {
	result := [][]int{}

	// 回溯函數
	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		// 若當前組合已有 k 個元素，則加入結果列表
		if len(current) == k {
			// 創建一個新的切片以避免引用問題
			combination := make([]int, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// 剪枝：計算剩餘所需元素數量
		remain := k - len(current)

		// 只遍歷可能產生有效組合的數字
		// i 最大可以是 n-(k-len(current))+1
		for i := start; i <= n-remain+1; i++ {
			// 選擇當前數字
			current = append(current, i)

			// 遞迴處理下一個位置
			backtrack(i+1, current)

			// 回溯：撤銷選擇
			current = current[:len(current)-1]
		}
	}

	// 從空組合開始回溯
	backtrack(1, []int{})

	return result
}

// 字典序迭代實現
func combineIterative(n int, k int) [][]int {
	// 處理邊界情況
	if k == 0 || k > n {
		return [][]int{}
	}

	result := [][]int{}

	// 初始化第一個組合：[1,2,...,k]
	comb := make([]int, k)
	for i := 0; i < k; i++ {
		comb[i] = i + 1 // 填入 1,2,...,k
	}

	// 將第一個組合加入結果
	first := make([]int, k)
	copy(first, comb)
	result = append(result, first)

	// 不斷生成下一個組合，直到沒有為止
	for {
		// 從右往左，找到第一個可以增加的位置
		i := k - 1 // 從最右邊開始

		// 如果當前位置的數字已達到最大可能值，則向左移動
		// 對於位置i，最大可能值為 n-(k-1-i)
		// 例如，對於n=4, k=2：
		// - 位置0的最大值是 4-(2-1-0)=3
		// - 位置1的最大值是 4-(2-1-1)=4
		for i >= 0 && comb[i] == n-(k-1-i) {
			i--
		}

		// 如果沒有位置可以增加，則已生成所有組合
		if i < 0 {
			break
		}

		// 將該位置的數字加1
		comb[i]++

		// 重設該位置右邊的所有數字
		// 每個數字應該是前一個數字加1
		for j := i + 1; j < k; j++ {
			comb[j] = comb[i] + (j - i)
		}

		// 將新組合加入結果
		next := make([]int, k)
		copy(next, comb)
		result = append(result, next)
	}

	return result
}
