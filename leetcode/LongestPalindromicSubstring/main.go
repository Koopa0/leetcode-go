package LongestPalindromicSubstring

// 最長回文子串
// 時間複雜度: O(n²)，其中n是字符串的長度
// 空間複雜度: O(1)
func longestPalindrome(s string) string {
	// 處理邊界情況，如果字符串長度小於2，直接返回原字符串
	if len(s) < 2 {
		return s
	}

	// 初始化起始索引和最大長度
	start, maxLength := 0, 1

	// 定義擴展函數，從中心點向兩側擴展尋找回文
	expandAroundCenter := func(left, right int) {
		// 當左右指針在有效範圍內且指向的字符相同時，繼續擴展
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// 計算當前回文的長度
			currentLength := right - left + 1

			// 如果當前回文長度大於已知最大長度，更新結果
			if currentLength > maxLength {
				maxLength = currentLength
				start = left
			}

			// 向兩側擴展
			left--
			right++
		}
	}

	// 遍歷所有可能的中心點
	for i := 0; i < len(s); i++ {
		// 擴展奇數長度回文（單字符中心）
		expandAroundCenter(i, i)

		// 擴展偶數長度回文（相鄰字符對中心）
		expandAroundCenter(i, i+1)
	}

	// 返回最長回文子串
	return s[start : start+maxLength]
}
