package InsertInterval

// 暴力解法
func insertBruteForce(intervals [][]int, newInterval []int) [][]int {
	// 1. 找到正確的插入位置並插入新區間
	n := len(intervals)
	var position int
	for position = 0; position < n; position++ {
		if intervals[position][0] > newInterval[0] {
			break
		}
	}

	// 插入新區間
	result := make([][]int, 0, n+1)
	result = append(result, intervals[:position]...)
	result = append(result, newInterval)
	result = append(result, intervals[position:]...)

	// 2. 合併重疊區間
	return merge(result)
}

// 合併重疊區間的函數（類似 LeetCode 56 題）
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := result[len(result)-1]
		next := intervals[i]

		// 檢查是否重疊
		if current[1] >= next[0] {
			// 重疊，更新結束點
			if current[1] < next[1] {
				current[1] = next[1]
			}
		} else {
			// 不重疊，加入新區間
			result = append(result, next)
		}
	}

	return result
}

// 優化解法
func insertOptimized(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	n := len(intervals)
	i := 0

	// 將所有在新區間之前的區間加入結果
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 合併所有與新區間重疊的區間
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	// 加入合併後的新區間
	result = append(result, newInterval)

	// 將所有在新區間之後的區間加入結果
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// 最佳解法
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	inserted := false

	for _, interval := range intervals {
		// 如果當前區間在新區間之後
		if interval[0] > newInterval[1] {
			// 如果新區間尚未插入，先插入新區間
			if !inserted {
				result = append(result, newInterval)
				inserted = true
			}
			result = append(result, interval)
		} else if interval[1] < newInterval[0] {
			// 當前區間在新區間之前
			result = append(result, interval)
		} else {
			// 當前區間與新區間重疊，更新新區間
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	// 如果新區間尚未插入，將其加入結果
	if !inserted {
		result = append(result, newInterval)
	}

	return result
}
