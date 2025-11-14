package ProductofArrayExceptSelf

// productExceptSelf 計算數組中除了當前元素外所有元素的乘積
// 輸入: nums - 整數數組
// 輸出: 一個新數組，其中 answer[i] 是除了 nums[i] 外所有元素的乘積
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// 創建結果數組
	answer := make([]int, n)

	// 第一步: 從左到右遍歷，計算當前元素左側所有元素的乘積
	// 初始化第一個元素的左側乘積為 1（因為左側沒有元素）
	answer[0] = 1
	for i := 1; i < n; i++ {
		// 當前位置的左側乘積 = 前一個位置的左側乘積 * 前一個位置的元素值
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 第二步: 從右到左遍歷，計算當前元素右側所有元素的乘積
	// 並與左側乘積相乘得到最終結果
	rightProduct := 1 // 初始化右側乘積為 1
	for i := n - 1; i >= 0; i-- {
		// 將當前位置的左側乘積與右側乘積相乘
		answer[i] *= rightProduct
		// 更新右側乘積，為下一個元素準備
		rightProduct *= nums[i]
	}

	return answer
}
