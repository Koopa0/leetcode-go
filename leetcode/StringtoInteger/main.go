package StringtoInteger

import "math"

func myAtoi(s string) int {
	// 初始化變數
	result := 0
	sign := 1
	i := 0

	// 32 位整數的範圍
	const MaxInt = math.MaxInt32 // 2147483647
	const MinInt = math.MinInt32 // -2147483648

	// 跳過前導空格
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// 處理符號
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 處理數字
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// 檢查溢出 (分別處理正數和負數的情況)
		if sign == 1 {
			if result > MaxInt/10 || (result == MaxInt/10 && digit > MaxInt%10) {
				return MaxInt
			}
		} else {
			if -result < MinInt/10 || (-result == MinInt/10 && digit > -MinInt%10) {
				return MinInt
			}
		}

		result = result*10 + digit
		i++
	}

	// 應用符號
	return sign * result
}
