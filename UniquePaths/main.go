package UniquePaths

// 暴力遞迴解法
func uniquePathsBruteForce(m int, n int) int {
	return uniquePathsRecursive(0, 0, m, n)
}

func uniquePathsRecursive(i, j, m, n int) int {
	// 已達終點
	if i == m-1 && j == n-1 {
		return 1
	}

	// 越界
	if i >= m || j >= n {
		return 0
	}

	// 向右移動 + 向下移動
	return uniquePathsRecursive(i, j+1, m, n) + uniquePathsRecursive(i+1, j, m, n)
}

// 記憶化遞迴解法
func uniquePathsMemoization(m int, n int) int {
	// 初始化備忘錄
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	return uniquePathsMemo(0, 0, m, n, memo)
}

func uniquePathsMemo(i, j, m, n int, memo [][]int) int {
	// 已達終點
	if i == m-1 && j == n-1 {
		return 1
	}

	// 越界
	if i >= m || j >= n {
		return 0
	}

	// 如果已經計算過，直接返回儲存的結果
	if memo[i][j] > 0 {
		return memo[i][j]
	}

	// 計算向右和向下的路徑數，並儲存結果
	memo[i][j] = uniquePathsMemo(i, j+1, m, n, memo) + uniquePathsMemo(i+1, j, m, n, memo)
	return memo[i][j]
}

// 動態規劃解法（二維）
func uniquePathsDP(m int, n int) int {
	// 初始化 dp 表格
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 第一列都是 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1 // 第一行都是 1
	}

	// 填充 dp 表格
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 空間優化的動態規劃解法
func uniquePathsOptimized(m int, n int) int {
	// 只使用一維陣列
	dp := make([]int, n)

	// 初始化：第一行全部為 1
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// 從第二行開始填充
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1] // dp[j] 是上一行的值，dp[j-1] 是同一行前一列的值
		}
	}

	return dp[n-1]
}

// 數學組合解法
func uniquePathsCombination(m int, n int) int {
	// 計算組合數 C(m+n-2, m-1)

	// 為避免溢出，選擇較小的數作為 k
	k := m - 1
	if n-1 < k {
		k = n - 1
	}

	result := 1
	// 計算 C(m+n-2, k) = (m+n-2)! / (k! * (m+n-2-k)!)
	for i := 1; i <= k; i++ {
		result = result * (m + n - 1 - i) / i
	}

	return result
}
