package MinimumWindowSubstring

// 暴力解法
func minWindowBruteForce(s string, t string) string {
	// 1. 初始化
	if len(s) < len(t) {
		return ""
	}

	// 建立 t 的字元頻率表
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	minLen := len(s) + 1
	minStart := 0

	// 2. 暴力檢查所有子字串
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			// 檢查子字串 s[i:j+1] 是否包含 t 中所有字元
			subCount := make(map[byte]int)
			for k := i; k <= j; k++ {
				subCount[s[k]]++
			}

			// 檢查是否符合條件
			valid := true
			for c, count := range tCount {
				if subCount[c] < count {
					valid = false
					break
				}
			}

			// 更新最小長度子字串
			if valid && (j-i+1) < minLen {
				minLen = j - i + 1
				minStart = i
			}
		}
	}

	// 3. 結果處理與回傳
	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}

// 優化解法（滑動視窗）
func minWindowOptimized(s string, t string) string {
	// 1. 初始化，使用優化的資料結構
	if len(s) < len(t) {
		return ""
	}

	// 建立 t 的字元頻率表
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// 滑動視窗相關變數
	required := len(tCount)            // 需要匹配的不同字元數
	formed := 0                        // 已經匹配的不同字元數
	windowCounts := make(map[byte]int) // 目前視窗中的字元計數

	// 結果變數
	minLen := len(s) + 1
	minStart := 0

	// 2. 滑動視窗實作
	left, right := 0, 0

	for right < len(s) {
		// 擴展右邊界
		c := s[right]
		windowCounts[c]++

		// 檢查是否匹配一個字元
		if count, found := tCount[c]; found && windowCounts[c] == count {
			formed++
		}

		// 嘗試收縮左邊界
		for left <= right && formed == required {
			c := s[left]

			// 更新最小視窗
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			// 移出左側字元
			windowCounts[c]--
			if count, found := tCount[c]; found && windowCounts[c] < count {
				formed--
			}

			left++
		}

		right++
	}

	// 3. 結果處理與回傳
	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}

// 最佳解法（優化的滑動視窗）
func minWindow(s string, t string) string {
	// 1. 使用最佳策略初始化
	if len(s) < len(t) {
		return ""
	}

	if len(t) == 0 {
		return s
	}

	// 使用陣列代替雜湊表，適用於已知範圍的字元集
	tCount := [128]int{}
	sCount := [128]int{}

	// 記錄 t 中每個字元的需求數量
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// 記錄 t 中出現的不同字元數
	charTypes := 0
	for _, count := range tCount {
		if count > 0 {
			charTypes++
		}
	}

	// 結果變數
	minLen := len(s) + 1
	minStart := 0

	// 滑動視窗變數
	left, right := 0, 0
	formed := 0

	// 2. 優化的滑動視窗實作
	for right < len(s) {
		// 擴展右邊界
		sCount[s[right]]++

		// 檢查是否正好匹配某個字元
		if tCount[s[right]] > 0 && sCount[s[right]] == tCount[s[right]] {
			formed++
		}

		// 嘗試收縮左邊界
		for left <= right && formed == charTypes {
			// 更新最小視窗
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			// 移出左側字元
			if tCount[s[left]] > 0 && sCount[s[left]] == tCount[s[left]] {
				formed--
			}
			sCount[s[left]]--
			left++
		}

		right++
	}

	// 3. 結果處理與回傳
	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
