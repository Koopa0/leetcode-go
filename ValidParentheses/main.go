package ValidParentheses

func isValid(s string) bool {
	// 創建一個堆疊來存儲開放括號
	stack := []rune{}

	// 創建閉合括號到開放括號的映射
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍歷字串中的每個字
	for _, char := range s {
		// 檢查當前字是否為閉合括號
		if openBracket, exists := bracketMap[char]; exists {
			// 如果堆疊為空或頂部不匹配，則字串無效
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// 彈出匹配的開放括號
			stack = stack[:len(stack)-1]
		} else {
			// 當前字是開放括號，入堆疊
			stack = append(stack, char)
		}
	}

	// 如果堆疊為空，則所有括號都已正確匹配
	return len(stack) == 0
}
