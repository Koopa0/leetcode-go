package EditDistance

// 暴力解法 - 遞迴
func minDistanceRecursive(word1 string, word2 string) int {
	return editDistRecursive(word1, word2, len(word1), len(word2))
}

func editDistRecursive(s1 string, s2 string, m int, n int) int {
	// 基本情況
	if m == 0 {
		return n // 如果第一個字串為空，需要 n 次插入操作
	}
	if n == 0 {
		return m // 如果第二個字串為空，需要 m 次刪除操作
	}

	// 如果最後一個字元相同，不需要額外操作
	if s1[m-1] == s2[n-1] {
		return editDistRecursive(s1, s2, m-1, n-1)
	}

	// 否則，取三種操作中的最小值
	// 1. 插入：將字元插入到 s1
	// 2. 刪除：從 s1 刪除字元
	// 3. 替換：將 s1 中的字元替換為 s2 中的字元
	insert := editDistRecursive(s1, s2, m, n-1)
	d := editDistRecursive(s1, s2, m-1, n)
	replace := editDistRecursive(s1, s2, m-1, n-1)

	return 1 + min(insert, min(d, replace))
}

// 優化解法 - 記憶化遞迴
func minDistanceMemoized(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 建立備忘錄，初始值為 -1 表示未計算
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return editDistMemoized(word1, word2, m, n, memo)
}

func editDistMemoized(s1 string, s2 string, m int, n int, memo [][]int) int {
	// 如果已經計算過，直接返回結果
	if memo[m][n] != -1 {
		return memo[m][n]
	}

	// 基本情況
	if m == 0 {
		memo[m][n] = n
		return n
	}
	if n == 0 {
		memo[m][n] = m
		return m
	}

	// 如果最後一個字元相同
	if s1[m-1] == s2[n-1] {
		memo[m][n] = editDistMemoized(s1, s2, m-1, n-1, memo)
		return memo[m][n]
	}

	// 取三種操作的最小值
	insert := editDistMemoized(s1, s2, m, n-1, memo)
	d := editDistMemoized(s1, s2, m-1, n, memo)
	replace := editDistMemoized(s1, s2, m-1, n-1, memo)

	memo[m][n] = 1 + min(insert, min(d, replace))
	return memo[m][n]
}

// 最佳解法 - 動態規劃（自底向上）
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// 建立 dp 表格
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化第一行和第一列
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 從 word1 的前 i 個字元轉換到空字串需要 i 次刪除
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 從空字串轉換到 word2 的前 j 個字元需要 j 次插入
	}

	// 填充 dp 表格
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果字元相同，不需要額外操作
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 取三種操作的最小值
				insert := dp[i][j-1]    // 插入操作
				d := dp[i-1][j]         // 刪除操作
				replace := dp[i-1][j-1] // 替換操作
				dp[i][j] = 1 + min(insert, min(d, replace))
			}
		}
	}

	return dp[m][n]
}
