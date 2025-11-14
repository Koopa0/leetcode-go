package ValidPalindrome

func isPalindrome(s string) bool {
	// 初始化雙指針
	left, right := 0, len(s)-1

	for left < right {
		// 跳過左側非字母數字
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// 跳過右側非字母數字
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// 比較字（忽略大小寫）
		if left < right {
			leftChar, rightChar := s[left], s[right]

			// 如果都是數字，直接比較
			if isDigit(leftChar) && isDigit(rightChar) {
				if leftChar != rightChar {
					return false
				}
			} else {
				// 如果是字母，轉換為小寫後比較
				if toLower(leftChar) != toLower(rightChar) {
					return false
				}
			}

			left++
			right--
		}
	}

	return true
}

// 檢查字是否為字母或數字
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// 檢查字是否為數字
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// 將字轉為小寫
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
