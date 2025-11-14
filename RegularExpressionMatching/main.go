package RegularExpressionMatching

// 正則表達式匹配

// 實現正則表達式匹配
// 支持 '.' 匹配任何單個字
// 支持 '*' 匹配零個或多個前面的元素
func isMatch(s string, p string) bool {
	// 獲取字串和模式的長度
	m, n := len(s), len(p)

	// 創建一個二維動態規劃數組
	// dp[i][j] 表示 s 的前 i 個字是否匹配 p 的前 j 個字
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 基本情況: 空字串匹配空模式
	dp[0][0] = true

	// 處理模式 p 可能匹配空字串的情況（例如: a*b*c*）
	for j := 1; j <= n; j++ {
		// 只有當前字是 '*' 且前面的模式匹配空字串時，當前模式才能匹配空字串
		if p[j-1] == '*' && j >= 2 {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 填充 dp 數組
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 情況 1: 不使用 '*' 匹配（即匹配零次前面的元素）
				if j >= 2 {
					dp[i][j] = dp[i][j-2]
				}

				// 情況 2: 使用 '*' 匹配（即匹配一次或多次前面的元素）
				if j >= 2 && (p[j-2] == '.' || p[j-2] == s[i-1]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				// 當前字是否匹配
				match := p[j-1] == '.' || p[j-1] == s[i-1]
				dp[i][j] = match && dp[i-1][j-1]
			}
		}
	}

	// 最終結果
	return dp[m][n]
}
