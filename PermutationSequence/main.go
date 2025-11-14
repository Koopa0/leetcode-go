package PermutationSequence

import (
	"container/list"
	"strings"
)

// 暴力解法
func bruteForceSolution(n int, k int) string {
	// 初始化數字集合
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// 生成所有排列
	var permutations []string
	generatePermutations(nums, 0, &permutations)

	// 返回第 k 個排列
	return permutations[k-1]
}

// 生成所有排列的輔助函數
func generatePermutations(nums []int, index int, result *[]string) {
	if index == len(nums) {
		// 將當前排列轉換為字串
		var sb strings.Builder
		for _, num := range nums {
			sb.WriteRune(rune('0' + num))
		}
		*result = append(*result, sb.String())
		return
	}

	for i := index; i < len(nums); i++ {
		// 交換當前元素和 index 位置的元素
		nums[index], nums[i] = nums[i], nums[index]

		// 遞迴生成剩餘元素的排列
		generatePermutations(nums, index+1, result)

		// 回溯，恢復原始順序
		nums[index], nums[i] = nums[i], nums[index]
	}
}

// 數學方法直接計算
func mathematicalSolution(n int, k int) string {
	// 創建 1 到 n 的數字集合
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// 計算 (n-1)!
	factorial := 1
	for i := 1; i < n; i++ {
		factorial *= i
	}

	// 調整 k 從 0 開始計算
	k = k - 1

	var result strings.Builder
	for i := n; i > 0; i-- {
		// 計算當前位置的數字索引
		index := k / factorial

		// 將對應的數字加入結果
		result.WriteRune(rune('0' + nums[index]))

		// 從數字集合中移除已使用的數字
		nums = append(nums[:index], nums[index+1:]...)

		// 更新 k 和階乘
		k = k % factorial
		if i > 1 {
			factorial = factorial / (i - 1)
		}
	}

	return result.String()
}

// 最佳解法
func getPermutation(n int, k int) string {
	// 創建數字鏈表
	numList := list.New()
	for i := 1; i <= n; i++ {
		numList.PushBack(i)
	}

	// 計算 (n-1)!
	factorial := 1
	for i := 1; i < n; i++ {
		factorial *= i
	}

	// 調整 k 從 0 開始計算
	k = k - 1

	var result strings.Builder
	for i := n; i > 0; i-- {
		// 計算當前位置的數字索引
		index := k / factorial

		// 從鏈表中找到對應的數字
		e := numList.Front()
		for j := 0; j < index; j++ {
			e = e.Next()
		}

		// 將數字加入結果
		result.WriteRune(rune('0' + e.Value.(int)))

		// 從鏈表中移除已使用的數字
		numList.Remove(e)

		// 更新 k 和階乘
		k = k % factorial
		if i > 1 {
			factorial = factorial / (i - 1)
		}
	}

	return result.String()
}
