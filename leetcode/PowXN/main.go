package PowXN

import "math"

func myPow(x float64, n int) float64 {
	// 處理基本情況
	if n == 0 {
		return 1
	}

	// 處理整數邊界情況 INT_MIN
	if n == math.MinInt32 {
		// 將 INT_MIN 分解為 -2^30 * 2 = -2^31
		// 計算 x^(-2^30) 然後再平方，最後取倒數
		return 1 / (myPow(x, -(n/2)) * myPow(x, -(n/2)))
	}

	// 處理負指數
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	// 二進制分解計算冪
	result := 1.0
	current_product := x

	// 遍歷 n 的二進制表示
	for n > 0 {
		// 如果當前位是 1，將當前的冪乘到結果中
		if n&1 == 1 {
			result *= current_product
		}

		// 將 current_product 平方，為下一位準備
		current_product *= current_product

		// 右移一位，檢查下一個二進制位
		n >>= 1
	}

	return result
}
