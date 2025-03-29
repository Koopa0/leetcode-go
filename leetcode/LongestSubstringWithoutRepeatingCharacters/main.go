package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	// 假設字符集僅為 ASCII
	charIndex := [128]int{} // 初始化為 0
	for i := range charIndex {
		charIndex[i] = -1 // 初始化為 -1，表示字符尚未出現
	}

	maxLength := 0
	left := 0

	// 遍歷字符串
	// e.g. s = "abcabcbb"，right = 0，s[right] = "a"，
	// charIndex[s[right]] = -1，left = 0，currentLength = 1，maxLength = 1
	// e.g. s = "abcabcbb"，right = 1，s[right] = "b"，
	// charIndex[s[right]] = -1，left = 0，currentLength = 2，maxLength = 2
	for right := 0; right < len(s); right++ {
		if charIndex[s[right]] >= left {
			left = charIndex[s[right]] + 1
		}
		charIndex[s[right]] = right
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
