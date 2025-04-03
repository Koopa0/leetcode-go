package MultiplyStrings

import "strings"

func multiply(num1 string, num2 string) string {
	// 特殊情況處理
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 初始化結果數組，長度為兩個數字的長度之和
	length := len(num1) + len(num2)
	result := make([]int, length)

	// 從右到左遍歷兩個數字
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			// 計算當前位置的乘積
			mul := int(num1[i]-'0') * int(num2[j]-'0')

			// 結果位置
			p1 := i + j + 1 // 個位位置
			p2 := i + j     // 進位位置

			// 將乘積加到對應位置
			sum := mul + result[p1]
			result[p1] = sum % 10
			result[p2] += sum / 10
		}
	}

	// 構建結果字符串
	var sb strings.Builder
	for _, digit := range result {
		// 跳過前導零
		if sb.Len() == 0 && digit == 0 {
			continue
		}
		sb.WriteByte(byte(digit) + '0')
	}

	// 如果結果為空，返回 "0"
	if sb.Len() == 0 {
		return "0"
	}

	return sb.String()
}
