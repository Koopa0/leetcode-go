package DecodeWays

// 暴力遞迴解法
func bruteForceNumDecodings(s string) int {
	// 使用輔助函數從索引0開始遞迴
	return recursiveDecoding(s, 0)
}

func recursiveDecoding(s string, index int) int {
	// 基本情況：到達字串末尾，找到一種有效解碼
	if index == len(s) {
		return 1
	}

	// 如果當前數字是'0'，則無法獨立解碼
	if s[index] == '0' {
		return 0
	}

	// 嘗試解碼一個數字
	singleDigit := recursiveDecoding(s, index+1)

	// 嘗試解碼兩個數字（如果可能）
	doubleDigit := 0
	if index+1 < len(s) {
		// 檢查兩位數是否在有效範圍內
		twoDigits := (s[index]-'0')*10 + (s[index+1] - '0')
		if twoDigits >= 10 && twoDigits <= 26 {
			doubleDigit = recursiveDecoding(s, index+2)
		}
	}

	// 兩種可能性的總和
	return singleDigit + doubleDigit
}

// 帶記憶化的遞迴解法
func memoizedNumDecodings(s string) int {
	// 初始化記憶化數組，-1表示尚未計算
	memo := make([]int, len(s)+1)
	for i := range memo {
		memo[i] = -1
	}

	return memoizedDecoding(s, 0, memo)
}

func memoizedDecoding(s string, index int, memo []int) int {
	// 如果已經計算過這個子問題，直接返回結果
	if memo[index] != -1 {
		return memo[index]
	}

	// 基本情況：到達字串末尾，找到一種有效解碼
	if index == len(s) {
		return 1
	}

	// 如果當前數字是'0'，則無法獨立解碼
	if s[index] == '0' {
		memo[index] = 0
		return 0
	}

	// 嘗試解碼一個數字
	singleDigit := memoizedDecoding(s, index+1, memo)

	// 嘗試解碼兩個數字（如果可能）
	doubleDigit := 0
	if index+1 < len(s) {
		// 檢查兩位數是否在有效範圍內
		twoDigits := (s[index]-'0')*10 + (s[index+1] - '0')
		if twoDigits >= 10 && twoDigits <= 26 {
			doubleDigit = memoizedDecoding(s, index+2, memo)
		}
	}

	// 儲存並返回結果
	memo[index] = singleDigit + doubleDigit
	return memo[index]
}

// 自底向上的動態規劃解法
func dpNumDecodings(s string) int {
	n := len(s)

	// 特殊情況處理
	if n == 0 || s[0] == '0' {
		return 0
	}

	// dp[i] 表示 s 前 i 個字元的解碼方式數
	dp := make([]int, n+1)
	dp[0] = 1 // 空字串有1種解碼方式

	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}

	for i := 2; i <= n; i++ {
		// 考慮單個數字解碼
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// 考慮兩個數字解碼
		twoDigits := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// 空間優化版動態規劃解法
func optimizedDpNumDecodings(s string) int {
	n := len(s)

	// 特殊情況處理
	if n == 0 || s[0] == '0' {
		return 0
	}

	// 初始化前兩個狀態
	// dp_prev_prev 代表 dp[i-2]
	// dp_prev 代表 dp[i-1]
	dpPrevPrev := 1 // dp[0] = 1
	dpPrev := 1     // dp[1] = 1 (假設第一個字元不是 '0')

	if s[0] == '0' {
		dpPrev = 0
	}

	for i := 2; i <= n; i++ {
		current := 0

		// 考慮單個數字解碼
		if s[i-1] != '0' {
			current += dpPrev
		}

		// 考慮兩個數字解碼
		twoDigits := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigits >= 10 && twoDigits <= 26 {
			current += dpPrevPrev
		}

		// 更新狀態
		dpPrevPrev = dpPrev
		dpPrev = current
	}

	return dpPrev
}
