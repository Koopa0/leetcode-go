package MaximumSubarray

// 暴力解法
func maxSubArrayBruteForce(nums []int) int {
	// 1. 初始化最大和為第一個元素
	maxSum := nums[0]
	n := len(nums)

	// 2. 遍歷所有可能的起始位置
	for i := 0; i < n; i++ {
		currentSum := 0
		// 3. 遍歷所有可能的結束位置
		for j := i; j < n; j++ {
			// 4. 累加當前元素
			currentSum += nums[j]
			// 5. 更新最大和
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	// 6. 返回結果
	return maxSum
}

// 動態規劃解法
func maxSubArrayDP(nums []int) int {
	// 1. 初始化動態規劃陣列
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	maxSum := dp[0]

	// 2. 填充動態規劃陣列
	for i := 1; i < n; i++ {
		// 3. 應用狀態轉移方程
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		// 4. 更新全局最大和
		maxSum = max(maxSum, dp[i])
	}

	// 5. 返回結果
	return maxSum
}

// Kadane's 演算法 (最佳解法)
func maxSubArray(nums []int) int {
	// 1. 初始化當前最大和與全局最大和
	currentMax := nums[0]
	globalMax := nums[0]

	// 2. 遍歷陣列的其餘部分
	for i := 1; i < len(nums); i++ {
		// 3. 更新當前最大和
		currentMax = max(nums[i], currentMax+nums[i])
		// 4. 更新全局最大和
		globalMax = max(globalMax, currentMax)
	}

	// 5. 返回結果
	return globalMax
}
