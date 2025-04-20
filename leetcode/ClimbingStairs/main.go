package ClimbingStairs

// 暴力遞迴解法
func climbStairsBruteForce(n int) int {
	// 基本情況
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// 遞迴計算 f(n) = f(n-1) + f(n-2)
	return climbStairsBruteForce(n-1) + climbStairsBruteForce(n-2)
}

// 記憶化遞迴解法
func climbStairsWithMemo(n int) int {
	// 建立記憶化陣列
	memo := make([]int, n+1)
	return climbStairsHelper(n, memo)
}

func climbStairsHelper(n int, memo []int) int {
	// 基本情況
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// 檢查是否已經計算過這個值
	if memo[n] != 0 {
		return memo[n]
	}

	// 計算並儲存結果
	memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	return memo[n]
}

// 動態規劃（自底向上）解法
func climbStairsDP(n int) int {
	// 處理基本情況
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// 建立 dp 陣列
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	// 自底向上填充 dp 陣列
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 空間最佳化的動態規劃解法
func climbStairsOptimized(n int) int {
	// 處理基本情況
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// 初始化前兩個狀態
	prev := 1 // f(1)
	curr := 2 // f(2)

	// 迭代計算 f(3) 到 f(n)
	for i := 3; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}

	return curr
}
