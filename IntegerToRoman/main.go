package IntegerToRoman

import "strings"

// intToRoman 使用貪心算法將整數轉換為羅馬數字
func intToRoman(num int) string {
	// 定義所有符號及其值，按值從大到小排序
	symbols := []struct {
		symbol string
		value  int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	// 結果字串
	var result strings.Builder

	// 貪心算法：從最大的符號開始
	for _, s := range symbols {
		// 當前數字大於或等於當前符號值時
		for num >= s.value {
			// 將符號添加到結果中
			result.WriteString(s.symbol)
			// 從數字中減去該符號值
			num -= s.value
		}

		// 如果數字已經為 0，提前結束
		if num == 0 {
			break
		}
	}

	return result.String()
}
