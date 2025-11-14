package UniquePathsII

// 暴力解法 (遞迴)
func bruteForceSolution(obstacleGrid [][]int) int {
	// 如果起點或終點是障礙物，則無法到達終點
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 使用記憶化遞迴優化
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示尚未計算
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 越界或遇到障礙物
		if i >= m || j >= n || obstacleGrid[i][j] == 1 {
			return 0
		}

		// 到達終點
		if i == m-1 && j == n-1 {
			return 1
		}

		// 已經計算過的子問題
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// 向下和向右移動
		memo[i][j] = dfs(i+1, j) + dfs(i, j+1)
		return memo[i][j]
	}

	return dfs(0, 0)
}

// 優化解法 (自下而上的動態規劃)
func optimizedSolution(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 創建動態規劃表格
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起點
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}

	// 填充第一行
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}

	// 填充第一列
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && dp[0][j-1] == 1 {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	// 填充其餘部分
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[m-1][n-1]
}

// 最佳解法 (空間優化的動態規劃)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 創建一維 DP 陣列
	dp := make([]int, n)

	// 初始化起點
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	// 填充第一行
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && dp[j-1] > 0 {
			dp[j] = dp[j-1]
		}
	}

	// 處理剩餘行
	for i := 1; i < m; i++ {
		// 更新每一行的第一個元素
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}

		// 更新該行的其餘元素
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}
