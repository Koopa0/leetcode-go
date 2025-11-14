package AddBinary

import "strconv"

// 暴力解法 - 使用 Go 內建函數進行進位轉換
func addBinaryBruteForce(a string, b string) string {
	// 1. 將二進位字串轉換為十進位整數
	numA, _ := strconv.ParseInt(a, 2, 64)
	numB, _ := strconv.ParseInt(b, 2, 64)

	// 2. 計算總和
	sum := numA + numB

	// 3. 將總和轉換回二進位字串
	return strconv.FormatInt(sum, 2)
}

// 優化解法 - 模擬加法過程
func addBinarySimulation(a string, b string) string {
	// 1. 初始化結果字串和進位
	result := ""
	carry := 0

	// 2. 獲取兩個字串的長度
	lenA, lenB := len(a), len(b)

	// 3. 從最右側開始遍歷
	for i, j := lenA-1, lenB-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		// 計算當前位的總和
		sum := carry

		// 添加 a 的當前位
		if i >= 0 {
			sum += int(a[i] - '0')
		}

		// 添加 b 的當前位
		if j >= 0 {
			sum += int(b[j] - '0')
		}

		// 計算新的結果和進位
		result = string('0'+byte(sum%2)) + result
		carry = sum / 2
	}

	return result
}

// 最佳解法 - 使用位元組陣列
func addBinary(a string, b string) string {
	// 1. 獲取兩個字串的長度
	lenA, lenB := len(a), len(b)

	// 2. 確定結果的最大長度（較長輸入的長度加1）
	maxLen := max(lenA, lenB)
	maxLen++ // 考慮可能的進位

	// 3. 創建結果陣列
	result := make([]byte, maxLen)

	// 4. 從最右側開始計算
	carry := 0
	i, j, k := lenA-1, lenB-1, maxLen-1

	for ; i >= 0 || j >= 0 || carry > 0; i, j, k = i-1, j-1, k-1 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
		}

		if j >= 0 {
			sum += int(b[j] - '0')
		}

		result[k] = byte(sum%2) + '0'
		carry = sum / 2
	}

	// 5. 處理前導零（如果有的話）
	result = result[k+1:]

	// 6. 處理空結果的情況
	if len(result) == 0 {
		return "0"
	}

	return string(result)
}
