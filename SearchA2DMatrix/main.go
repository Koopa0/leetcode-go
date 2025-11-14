package SearchA2DMatrix

// 暴力解法
func bruteForceSearchMatrix(matrix [][]int, target int) bool {
	// 獲取矩陣的行數和列數
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	// 遍歷整個矩陣
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	// 未找到目標值
	return false
}

// 兩次二分搜尋解法
func twoBinarySearchMatrix(matrix [][]int, target int) bool {
	// 獲取矩陣的行數和列數
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	// 第一次二分搜尋：找到可能包含目標值的列
	left, right := 0, m-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid][0] <= target && target <= matrix[mid][n-1] {
			// 目標值可能在這一列
			return binarySearch(matrix[mid], target)
		} else if matrix[mid][0] > target {
			// 目標值可能在更上面的列
			right = mid - 1
		} else {
			// 目標值可能在更下面的列
			left = mid + 1
		}
	}

	// 未找到可能包含目標值的列
	return false
}

// 輔助函數：在一維數組中進行二分搜尋
func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// 單次二分搜尋解法（最優解）
func searchMatrix(matrix [][]int, target int) bool {
	// 獲取矩陣的行數和列數
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	// 在虛擬的一維數組上執行二分搜尋
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2

		// 將一維索引轉換為二維索引
		row, col := mid/n, mid%n

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 未找到目標值
	return false
}
