package StringtoInteger

func myAtoi(s string) int {
	// 定義 32 位有符號整數的範圍限制
	const IntMax = 1<<31 - 1 // 2^31 - 1
	const IntMin = -1 << 31  // -2^31

	// 初始化結果和符號
	result := 0
	sign := 1 // 1 表示正數，-1 表示負數
	i := 0
	n := len(s)

	// 步驟 1：跳過前導空白
	for i < n && s[i] == ' ' {
		i++
	}

	// 處理邊緣情況：字符串全是空白
	if i == n {
		return 0
	}

	// 步驟 2：處理符號
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 步驟 3：讀取數字並轉換
	for i < n && s[i] >= '0' && s[i] <= '9' {
		// 獲取當前數字
		digit := int(s[i] - '0')

		// 步驟 4：檢查溢出
		// 對於正數，如果 result > INT_MAX/10 或者 result == INT_MAX/10 且 digit > 7，將溢出
		// 對於負數，如果 result > INT_MAX/10 或者 result == INT_MAX/10 且 digit > 8，將溢出
		if result > IntMax/10 || (result == IntMax/10 && digit > IntMax%10) {
			if sign == 1 {
				return IntMax
			} else {
				return IntMin
			}
		}

		// 更新結果
		result = result*10 + digit
		i++
	}

	// 應用符號
	return sign * result
}
