package RomanToInteger

// romanToInt 將羅馬數字轉換為整數
func romanToInt(s string) int {
	// 建立羅馬數字到整數值的映射
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	lastValue := 0

	// 從右到左遍歷羅馬數字字串
	for i := len(s) - 1; i >= 0; i-- {
		// 獲取當前字的值
		curValue := romanValues[s[i]]

		// 判斷是否為減法情況
		if curValue >= lastValue {
			// 當前值大於等於上一個處理的值，執行加法
			result += curValue
		} else {
			// 當前值小於上一個處理的值，執行減法
			result -= curValue
		}

		// 更新最後處理的值
		lastValue = curValue
	}

	return result
}
