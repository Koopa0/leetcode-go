package CombinationSumII

import "sort"

// combinationSum2 返回所有可能的組合，使得組合中的數字和等於目標值
// 每個數字在每個組合中只能使用一次，且解集不能包含重複的組合
func combinationSum2(candidates []int, target int) [][]int {
	// 結果切片，用於存儲所有有效的組合
	result := [][]int{}

	// 對候選數組進行排序，使相同的數字相鄰，便於去重
	sort.Ints(candidates)

	// 定義用於回溯的輔助函數
	var backtrack func(start int, remain int, path []int)
	backtrack = func(start int, remain int, path []int) {
		// 基礎情況：如果剩餘目標值為0，表示找到一個有效組合
		if remain == 0 {
			// 注意：需要創建path的副本，因為path會在後續回溯過程中被修改
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		// 如果剩餘目標值小於0，或者已經處理完所有候選數字，則返回
		if remain < 0 || start >= len(candidates) {
			return
		}

		// 嘗試從start開始的每個候選數字
		for i := start; i < len(candidates); i++ {
			// 跳過重複數字以避免重複組合
			// 只有當i大於start（即不是該層的第一個數字）且當前數字等於前一個數字時才跳過
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// 當前數字大於剩餘目標值，由於數組已排序，後續數字也一定大於剩餘目標值，可提前退出
			if candidates[i] > remain {
				break
			}

			// 選擇當前數字
			path = append(path, candidates[i])

			// 遞迴調用，注意索引為i+1（因為每個數字只能使用一次）
			// 並減少剩餘目標值
			backtrack(i+1, remain-candidates[i], path)

			// 回溯：移除最後添加的數字，嘗試下一個選擇
			path = path[:len(path)-1]
		}
	}

	// 使用空路徑和完整目標值開始回溯
	backtrack(0, target, []int{})

	return result
}
