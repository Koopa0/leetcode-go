package UniqueBinarySearchTrees

// 暴力遞迴解法
func numTreesRecursive(n int) int {
	// 基本情況：空樹或單節點樹
	if n == 0 || n == 1 {
		return 1
	}

	total := 0

	// 嘗試每個可能的根節點
	for i := 1; i <= n; i++ {
		// 左子樹的結構數量
		leftStructures := numTreesRecursive(i - 1)
		// 右子樹的結構數量
		rightStructures := numTreesRecursive(n - i)
		// 當根節點為 i 時的總數 = 左子樹數量 × 右子樹數量
		total += leftStructures * rightStructures
	}

	return total
}

// 記憶化遞迴解法
func numTreesMemo(n int) int {
	// 創建記憶化陣列
	memo := make([]int, n+1)
	return numTreesHelper(n, memo)
}

func numTreesHelper(n int, memo []int) int {
	// 如果已經計算過，直接返回
	if memo[n] > 0 {
		return memo[n]
	}

	// 基本情況：空樹或單節點樹
	if n == 0 || n == 1 {
		memo[n] = 1
		return 1
	}

	total := 0

	// 嘗試每個可能的根節點
	for i := 1; i <= n; i++ {
		// 左子樹的結構數量
		leftStructures := numTreesHelper(i-1, memo)
		// 右子樹的結構數量
		rightStructures := numTreesHelper(n-i, memo)
		// 當根節點為 i 時的總數 = 左子樹數量 × 右子樹數量
		total += leftStructures * rightStructures
	}

	// 存儲結果以供後續使用
	memo[n] = total
	return total
}

// 動態規劃解法
func numTrees(n int) int {
	// 創建 DP 陣列
	dp := make([]int, n+1)

	// 初始化基本情況
	dp[0] = 1 // 空樹
	dp[1] = 1 // 單節點樹

	// 自底向上填充 DP 表
	for i := 2; i <= n; i++ {
		// 嘗試每個可能的根節點
		for j := 1; j <= i; j++ {
			// 左子樹的節點數量為 j-1
			// 右子樹的節點數量為 i-j
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
