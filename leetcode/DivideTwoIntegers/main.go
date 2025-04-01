package DivideTwoIntegers

import "math"

// divide 不使用乘法、除法和取餘運算符實現兩整數相除
func divide(dividend int, divisor int) int {
	// 處理除數為零的情況
	if divisor == 0 {
		return math.MaxInt32 // 在實際場景中應該返回錯誤，但LeetCode期望返回一個整數
	}

	// 處理被除數為零的情況
	if dividend == 0 {
		return 0
	}

	// 處理溢出情況：當被除數為最小整數且除數為-1時
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32 // 結果將是2^31，超出範圍，返回2^31-1
	}

	// 處理除數為1或-1的簡單情況
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend // 這裡不會再有溢出，因為前面已處理了特殊情況
	}

	// 確定結果的符號
	negative := (dividend < 0) != (divisor < 0)

	// 轉換為正數進行計算
	// 注意：math.MinInt32的絕對值會溢出，需要特殊處理
	absDividend := abs(dividend)
	absDivisor := abs(divisor)

	// 執行二進制長除法
	result := 0
	for absDividend >= absDivisor {
		// 找到最大的k，使得divisor * 2^k <= absDividend
		temp := absDivisor
		multiple := 1

		// 當temp的兩倍仍小於等於absDividend時，繼續左移
		for absDividend >= (temp << 1) {
			temp <<= 1     // temp = temp * 2
			multiple <<= 1 // multiple = multiple * 2
		}

		// 從被除數中減去temp，並將對應的2^k加到結果中
		absDividend -= temp
		result += multiple
	}

	// 應用符號
	if negative {
		result = -result
	}

	return result
}

// abs 返回整數的絕對值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
