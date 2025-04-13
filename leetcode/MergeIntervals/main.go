package MergeIntervals

import "sort"

// 暴力解法
func mergeIntervalsBruteForce(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. 初始化
	result := make([][]int, len(intervals))
	copy(result, intervals)

	// 2. 核心暴力邏輯
	changed := true
	for changed {
		changed = false
		temp := [][]int{}

		for i := 0; i < len(result); i++ {
			merged := false

			for j := 0; j < len(temp); j++ {
				// 檢查是否可以合併
				if canMerge(temp[j], result[i]) {
					// 合併區間
					temp[j] = []int{
						min(temp[j][0], result[i][0]),
						max(temp[j][1], result[i][1]),
					}
					merged = true
					changed = true
					break
				}
			}

			if !merged {
				temp = append(temp, result[i])
			}
		}

		result = temp
	}

	// 3. 結果處理與返回
	return result
}

// 檢查兩個區間是否可以合併
func canMerge(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[1], b[1])
}

// 排序優化解法
func mergeIntervalsSorted(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. 初始化並排序
	// 按照區間的起始值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 核心優化邏輯
	result := [][]int{intervals[0]} // 從第一個區間開始

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastIdx := len(result) - 1
		last := result[lastIdx]

		// 檢查當前區間是否與結果集合中的最後一個區間重疊
		if current[0] <= last[1] {
			// 合併重疊的區間
			result[lastIdx][1] = max(last[1], current[1])
		} else {
			// 不重疊，將當前區間添加到結果中
			result = append(result, current)
		}
	}

	// 3. 結果處理與返回
	return result
}

// 最佳解法
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. 初始化並排序
	// 按照區間的起始值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 核心最佳邏輯
	result := [][]int{}

	// 從第一個區間開始
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		// 如果當前區間與當前合併區間重疊
		if intervals[i][0] <= current[1] {
			// 更新當前合併區間的結束值
			current[1] = max(current[1], intervals[i][1])
		} else {
			// 不重疊，將當前合併區間添加到結果中
			result = append(result, current)
			// 開始新的合併區間
			current = intervals[i]
		}
	}

	// 將最後的合併區間添加到結果中
	result = append(result, current)

	// 3. 結果處理與返回
	return result
}
