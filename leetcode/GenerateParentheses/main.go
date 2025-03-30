package GenerateParentheses

// generateParenthesis 生成所有可能的有效括號組合
func generateParenthesis(n int) []string {
	// 存儲結果的切片
	result := []string{}

	// 調用回溯函數開始生成
	backtrack(&result, "", 0, 0, n)

	return result
}

// backtrack 使用回溯法生成有效括號組合
// result: 存儲最終結果的切片的指針
// current: 當前構建的字符串
// open: 當前已使用的左括號數量
// close: 當前已使用的右括號數量
// max: 括號對數 n
func backtrack(result *[]string, current string, open, close, max int) {
	// 基本情況：如果當前字符串長度為 2*max，說明我們已經構建了一個完整的有效組合
	if len(current) == 2*max {
		*result = append(*result, current) // 將當前組合添加到結果中
		return
	}

	// 情況1：如果已使用的左括號數量小於 max，可以添加左括號
	if open < max {
		// 添加左括號並遞歸
		backtrack(result, current+"(", open+1, close, max)
	}

	// 情況2：如果已使用的右括號數量小於已使用的左括號數量，可以添加右括號
	if close < open {
		// 添加右括號並遞歸
		backtrack(result, current+")", open, close+1, max)
	}

	// 注意：當 open == max 且 close == open 時，兩個條件都不滿足，遞歸自然終止
}
