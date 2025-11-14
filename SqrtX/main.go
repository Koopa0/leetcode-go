package SqrtX

// 暴力解法
func bruteForceSqrt(x int) int {
	// 特殊情況處理
	if x == 0 {
		return 0
	}

	// 從1開始逐一檢查
	i := 1
	for i*i <= x {
		i++
	}

	// 返回最後一個平方不超過x的整數
	return i - 1
}

// 二分搜尋解法
func binarySearchSqrt(x int) int {
	// 特殊情況處理
	if x == 0 {
		return 0
	}

	// 設置搜尋範圍
	left, right := 1, x

	// 二分搜尋
	for left <= right {
		mid := left + (right-left)/2

		// 避免整數溢出的方式比較
		if mid <= x/mid && (mid+1) > x/(mid+1) {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

// 牛頓法解法
func newtonSqrt(x int) int {
	// 特殊情況處理
	if x == 0 {
		return 0
	}

	// 選擇初始猜測值
	r := x

	// 牛頓迭代
	for r > x/r {
		r = (r + x/r) / 2
	}

	// 返回向下取整的結果
	return r
}
