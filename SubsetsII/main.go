package SubsetsII

import (
	"fmt"
	"sort"
)

// 暴力解法
func subsetsWithDupBruteForce(nums []int) [][]int {
	// 先排序，方便後續字串轉換比較
	sort.Ints(nums)

	var result [][]int
	var current []int
	var backtrack func(index int)

	// 回溯函數
	backtrack = func(index int) {
		// 創建當前組合的副本並加入結果集
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)

		// 遍歷剩餘的元素
		for i := index; i < len(nums); i++ {
			// 將當前元素加入組合
			current = append(current, nums[i])
			// 遞迴生成包含當前元素的所有子集
			backtrack(i + 1)
			// 回溯：移除當前元素
			current = current[:len(current)-1]
		}
	}

	backtrack(0)

	// 使用雜湊表去除重複子集
	seen := make(map[string]bool)
	var uniqueResult [][]int

	for _, subset := range result {
		// 將子集轉換為字串以進行雜湊
		key := fmt.Sprintf("%v", subset)
		if !seen[key] {
			seen[key] = true
			uniqueResult = append(uniqueResult, subset)
		}
	}

	return uniqueResult
}

// 優化解法
func subsetsWithDupOptimized(nums []int) [][]int {
	// 先排序，確保相同元素相鄰
	sort.Ints(nums)

	var result [][]int
	var current []int
	var backtrack func(index int)

	// 回溯函數
	backtrack = func(index int) {
		// 創建當前組合的副本並加入結果集
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)

		// 遍歷剩餘的元素
		for i := index; i < len(nums); i++ {
			// 在同一層決策中跳過重複元素
			if i > index && nums[i] == nums[i-1] {
				continue
			}

			// 將當前元素加入組合
			current = append(current, nums[i])
			// 遞迴生成包含當前元素的所有子集
			backtrack(i + 1)
			// 回溯：移除當前元素
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

// 最佳解法 - 迭代方法
func subsetsWithDup(nums []int) [][]int {
	// 先排序，確保相同元素相鄰
	sort.Ints(nums)

	// 初始化結果集，包含空集合
	result := [][]int{[]int{}}

	start, end := 0, 0

	// 遍歷所有元素
	for i := 0; i < len(nums); i++ {
		start = 0

		// 如果當前元素與前一個元素相同，只從上一輪新增的子集開始
		if i > 0 && nums[i] == nums[i-1] {
			start = end + 1
		}

		end = len(result) - 1
		size := len(result)

		// 將當前元素添加到指定範圍的子集中形成新的子集
		for j := start; j < size; j++ {
			// 創建新的子集
			newSubset := make([]int, len(result[j]))
			copy(newSubset, result[j])
			newSubset = append(newSubset, nums[i])
			result = append(result, newSubset)
		}
	}

	return result
}
