package Subsets

// 回溯法解決方案
func subsetsBacktracking(nums []int) [][]int {
	// 初始化結果陣列
	result := [][]int{}

	// 定義回溯函數
	var backtrack func(index int, current []int)
	backtrack = func(index int, current []int) {
		// 如果已處理完所有元素，將當前子集添加到結果中
		if index == len(nums) {
			// 創建當前子集的副本
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// 選擇1：包含當前元素
		current = append(current, nums[index])
		backtrack(index+1, current)

		// 選擇2：不包含當前元素（回溯）
		current = current[:len(current)-1]
		backtrack(index+1, current)
	}

	// 啟動回溯過程
	backtrack(0, []int{})

	return result
}

// 位元操作解決方案
func subsetsBitmask(nums []int) [][]int {
	n := len(nums)
	// 子集總數為 2^n
	totalSubsets := 1 << n // 相當於 2^n

	result := make([][]int, 0, totalSubsets)

	// 遍歷從 0 到 2^n-1 的每個數字
	for i := 0; i < totalSubsets; i++ {
		// 創建當前子集
		subset := []int{}

		// 檢查每一位是否為 1
		for j := 0; j < n; j++ {
			// 如果第 j 位是 1，則包含 nums[j]
			if (i & (1 << j)) != 0 {
				subset = append(subset, nums[j])
			}
		}

		// 將當前子集加入結果
		result = append(result, subset)
	}

	return result
}

// 迭代增量法解決方案
func subsetsIterative(nums []int) [][]int {
	// 初始化結果為只包含空集的列表
	result := [][]int{{}}

	// 對於 nums 中的每個元素
	for _, num := range nums {
		// 創建存放新子集的列表
		newSubsets := [][]int{}

		// 對於結果中的每個現有子集
		for _, curr := range result {
			// 創建當前子集的副本
			newSubset := make([]int, len(curr), len(curr)+1)
			copy(newSubset, curr)

			// 將當前元素加入副本，形成新的子集
			newSubset = append(newSubset, num)

			// 將新子集加入 newSubsets
			newSubsets = append(newSubsets, newSubset)
		}

		// 將所有新子集加入結果
		result = append(result, newSubsets...)
	}

	return result
}
