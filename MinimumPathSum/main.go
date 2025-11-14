package MinimumPathSum

// 暴力遞迴解法
func minPathSumBruteForce(grid [][]int) int {
	return minPathSumHelper(grid, 0, 0)
}

func minPathSumHelper(grid [][]int, i, j int) int {
	// 獲取網格尺寸
	m, n := len(grid), len(grid[0])

	// 基本情況：到達終點
	if i == m-1 && j == n-1 {
		return grid[i][j]
	}

	// 如果在最後一列，只能向右移動
	if i == m-1 {
		return grid[i][j] + minPathSumHelper(grid, i, j+1)
	}

	// 如果在最後一行，只能向下移動
	if j == n-1 {
		return grid[i][j] + minPathSumHelper(grid, i+1, j)
	}

	// 可以向右或向下移動，選擇路徑和較小的方向
	rightPath := minPathSumHelper(grid, i, j+1)
	downPath := minPathSumHelper(grid, i+1, j)

	return grid[i][j] + min(rightPath, downPath)
}

// 記憶化遞迴解法
func minPathSumMemoization(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 初始化記憶化數組，使用 -1 表示尚未計算
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return memoHelper(grid, 0, 0, memo)
}

func memoHelper(grid [][]int, i, j int, memo [][]int) int {
	m, n := len(grid), len(grid[0])

	// 如果結果已經計算過，直接返回
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// 基本情況：到達終點
	if i == m-1 && j == n-1 {
		memo[i][j] = grid[i][j]
		return memo[i][j]
	}

	// 如果在最後一列，只能向右移動
	if i == m-1 {
		memo[i][j] = grid[i][j] + memoHelper(grid, i, j+1, memo)
		return memo[i][j]
	}

	// 如果在最後一行，只能向下移動
	if j == n-1 {
		memo[i][j] = grid[i][j] + memoHelper(grid, i+1, j, memo)
		return memo[i][j]
	}

	// 可以向右或向下移動，選擇路徑和較小的方向
	rightPath := memoHelper(grid, i, j+1, memo)
	downPath := memoHelper(grid, i+1, j, memo)

	memo[i][j] = grid[i][j] + min(rightPath, downPath)
	return memo[i][j]
}

// 動態規劃解法
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// 創建 DP 表格
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起點
	dp[0][0] = grid[0][0]

	// 填充第一行
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 填充第一列
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 填充其餘位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

// 空間優化的動態規劃解法
func minPathSumOptimized(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// 只使用一維數組
	dp := make([]int, n)

	// 初始化第一個元素
	dp[0] = grid[0][0]

	// 填充第一行
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// 處理剩餘行
	for i := 1; i < m; i++ {
		// 更新行的第一個元素
		dp[0] = dp[0] + grid[i][0]

		// 更新行的其餘元素
		for j := 1; j < n; j++ {
			dp[j] = grid[i][j] + min(dp[j], dp[j-1])
		}
	}

	return dp[n-1]
}
