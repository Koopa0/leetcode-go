package ReverseInteger

func reverse(x int) int {
	// 常數定義，避免硬編碼
	const IntMax = 1<<31 - 1 // 2^31 - 1
	const IntMin = -1 << 31  // -2^31

	// 初始化結果
	result := 0

	// 一位一位地提取並反轉
	for x != 0 {
		// 提取最後一位數字
		digit := x % 10
		x /= 10

		// 檢查溢出
		// 如果 result > INT_MAX/10 或者 (result == INT_MAX/10 且 digit > INT_MAX%10)
		if result > IntMax/10 || (result == IntMax/10 && digit > IntMax%10) {
			return 0
		}
		// 如果 result < INT_MIN/10 或者 (result == INT_MIN/10 且 digit < INT_MIN%10)
		if result < IntMin/10 || (result == IntMin/10 && digit < IntMin%10) {
			return 0
		}

		// 更新結果
		result = result*10 + digit
	}

	return result
}
