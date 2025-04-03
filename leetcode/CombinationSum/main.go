package CombinationSum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	// 排序候選數組（非必須，但有利於優化）
	sort.Ints(candidates)

	// 用於存儲結果
	var result [][]int

	// 用於回溯的內部函數
	var backtrack func(start int, remaining int, current []int)

	backtrack = func(start int, remaining int, current []int) {
		// 成功找到組合
		if remaining == 0 {
			// 創建當前組合的副本並加入結果集
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// 嘗試所有可能的選擇
		for i := start; i < len(candidates); i++ {
			// 剪枝：如果當前候選數已經大於剩餘目標，後面的數更大，無需繼續
			if candidates[i] > remaining {
				break
			}

			// 選擇當前候選數
			current = append(current, candidates[i])

			// 遞歸：注意索引仍為 i，因為可以重複使用同一元素
			backtrack(i, remaining-candidates[i], current)

			// 回溯：移除最後加入的元素
			current = current[:len(current)-1]
		}
	}

	// 從空組合開始回溯
	backtrack(0, target, []int{})

	return result
}
