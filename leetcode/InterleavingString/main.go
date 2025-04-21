package InterleavingString

// 暴力解法 - 遞迴回溯
func isInterleaveBruteForce(s1 string, s2 string, s3 string) bool {
	// 如果長度不匹配，直接返回false
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 遞迴輔助函數
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 基本情況：所有字符都已匹配
		if k == len(s3) {
			return i == len(s1) && j == len(s2)
		}

		// 檢查是否可以從s1取字符
		var result bool = false
		if i < len(s1) && s1[i] == s3[k] {
			result = dfs(i+1, j, k+1)
		}

		// 如果從s1取字符不行，嘗試從s2取字符
		if !result && j < len(s2) && s2[j] == s3[k] {
			result = dfs(i, j+1, k+1)
		}

		return result
	}

	return dfs(0, 0, 0)
}

// 優化解法 - 記憶化遞迴
func isInterleaveOptimized(s1 string, s2 string, s3 string) bool {
	// 如果長度不匹配，直接返回false
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 創建記憶化數組
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1表示尚未計算
		}
	}

	// 遞迴輔助函數
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 基本情況：所有字符都已匹配
		if k == len(s3) {
			return i == len(s1) && j == len(s2)
		}

		// 檢查記憶化數組
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		// 初始化結果
		result := false

		// 嘗試從s1取字符
		if i < len(s1) && s1[i] == s3[k] {
			result = dfs(i+1, j, k+1)
		}

		// 如果從s1取字符不行，嘗試從s2取字符
		if !result && j < len(s2) && s2[j] == s3[k] {
			result = dfs(i, j+1, k+1)
		}

		// 更新記憶化數組
		if result {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}

		return result
	}

	return dfs(0, 0, 0)
}

// 最優解法 - 動態規劃
func isInterleaveDP(s1 string, s2 string, s3 string) bool {
	// 如果長度不匹配，直接返回false
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 創建動態規劃表格
	m, n := len(s1), len(s2)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 初始化：空字符串可以交錯形成空字符串
	dp[0][0] = true

	// 填充第一行：s1的前i個字符是否能交錯形成s3的前i個字符
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	// 填充第一列：s2的前j個字符是否能交錯形成s3的前j個字符
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	// 填充其餘單元格
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 當前位置的s3字符索引
			k := i + j - 1

			// 從s1取字符或從s2取字符
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[k]) ||
				(dp[i][j-1] && s2[j-1] == s3[k])
		}
	}

	return dp[m][n]
}
