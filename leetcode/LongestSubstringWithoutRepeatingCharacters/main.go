package LongestSubstringWithoutRepeatingCharacters

// lengthOfLongestSubstring 函數返回字符串 s 中不含重複字符的最長子串的長度
func lengthOfLongestSubstring(s string) int {
	// 如果字符串為空，直接返回 0
	if len(s) == 0 {
		return 0
	}

	// 使用映射來記錄每個字符最後一次出現的位置
	charIndexMap := make(map[byte]int)

	// 初始化變量
	maxLength := 0 // 最長子串的長度
	left := 0      // 窗口的左邊界

	// 遍歷字符串中的每個字符
	for right := 0; right < len(s); right++ {
		// 當前字符
		currentChar := s[right]

		// 如果當前字符已經在窗口中出現過
		if lastIndex, found := charIndexMap[currentChar]; found && lastIndex >= left {
			// 更新窗口的左邊界為該字符上一次出現的位置的下一個位置
			left = lastIndex + 1
		}

		// 計算當前窗口大小並更新最大長度
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		// 更新當前字符的位置
		charIndexMap[currentChar] = right
	}

	return maxLength
}
