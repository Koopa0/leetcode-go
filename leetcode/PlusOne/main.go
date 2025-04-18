package PlusOne

// 暴力解決方案
func bruteForcePlusOne(digits []int) []int {
	// 從最後一位開始
	n := len(digits)

	// 加 1 操作
	digits[n-1]++

	// 處理進位
	for i := n - 1; i > 0; i-- {
		if digits[i] < 10 {
			break // 如果沒有進位，提前結束
		}
		digits[i] = 0 // 當前位設為 0
		digits[i-1]++ // 向前進位
	}

	// 處理最高位的進位情況
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...) // 在最前面新增一個 1
	}

	return digits
}

// 優化解決方案
func optimizedPlusOne(digits []int) []int {
	n := len(digits)

	// 從最低位開始遍歷
	for i := n - 1; i >= 0; i-- {
		// 如果當前位小於 9，加 1 後直接返回（無需進位）
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		// 如果當前位是 9，設為 0 並繼續處理進位
		digits[i] = 0
	}

	// 如果代碼執行到這裡，說明所有位都是 9（例如 999）
	// 需要在最前面新增一個 1
	newDigits := make([]int, n+1)
	newDigits[0] = 1

	return newDigits
}

// 最佳解決方案
func plusOne(digits []int) []int {
	// 從後向前遍歷
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果當前位小於 9，加 1 後直接返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		// 當前位是 9，設為 0
		digits[i] = 0
	}

	// 如果所有位都是 9，則需要新增一位
	result := make([]int, len(digits)+1)
	result[0] = 1 // 首位設為 1，其餘位默認為 0

	return result
}
