package ValidNumber

import "regexp"

// 有限狀態機解決方案
func isNumberFSM(s string) bool {
	// 定義狀態轉換表
	// 狀態: 0-初始, 1-符號後, 2-整數部分, 3-小數點後無數字, 4-小數點後有數字, 5-e後, 6-e後符號, 7-e後數字
	transition := [][]int{
		{1, 2, 3, -1, -1},   // 0: 初始狀態
		{-1, 2, 3, -1, -1},  // 1: 符號後
		{-1, 2, 4, 5, -1},   // 2: 整數部分
		{-1, 4, -1, -1, -1}, // 3: 小數點後無數字
		{-1, 4, -1, 5, -1},  // 4: 小數點後有數字
		{6, 7, -1, -1, -1},  // 5: e後
		{-1, 7, -1, -1, -1}, // 6: e後符號
		{-1, 7, -1, -1, -1}, // 7: e後數字
	}

	currentState := 0 // 起始狀態

	for _, ch := range s {
		var targetState int

		switch ch {
		case '+', '-':
			targetState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			targetState = 1
		case '.':
			targetState = 2
		case 'e', 'E':
			targetState = 3
		default:
			return false // 不合法字
		}

		// 獲取下一個狀態
		nextState := transition[currentState][targetState]
		if nextState == -1 {
			return false // 無效轉換
		}

		currentState = nextState
	}

	// 檢查最終狀態是否為有效終止狀態
	return currentState == 2 || currentState == 4 || currentState == 7
}

// 線性掃描解決方案
func isNumberLinear(s string) bool {
	// 跟踪關鍵元素的出現
	seenDigit := false
	seenDot := false
	seenE := false

	for i, ch := range s {
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			seenDigit = true
		case '.':
			if seenDot || seenE {
				return false // 已經有小數點或在指數部分
			}
			seenDot = true
		case 'e', 'E':
			if seenE || !seenDigit {
				return false // 已經有e或沒有前導數字
			}
			seenE = true
			seenDigit = false // 重置數字標誌，指數部分需要至少一個數字
		case '+', '-':
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false // 符號只能出現在開頭或e/E後
			}
		default:
			return false // 不合法字
		}
	}

	return seenDigit // 必須至少有一個數字
}

// 正則表達式解決方案
func isNumber(s string) bool {
	// 定義有效數字的正則表達式模式
	pattern := `^[+-]?(\d+\.?|\d*\.\d+)([eE][+-]?\d+)?$`

	// 編譯正則表達式
	re := regexp.MustCompile(pattern)

	// 匹配字串
	return re.MatchString(s)
}
