package MinimumWindowSubstring

import "math"

func minWindow(s string, t string) string {
	// 邊界情況：如果 t 為空，返回空字符串
	if len(t) == 0 {
		return ""
	}

	// 創建 t 中字符的頻率數組（僅 ASCII）
	tCount := make([]int, 128)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// 計算所需的總字符數
	totalCharsNeeded := len(t)

	// 初始化變量
	windowCount := make([]int, 128)
	charsFound := 0
	left, right := 0, 0
	minLen := math.MaxInt32
	minStart, minEnd := 0, 0

	// 通過移動右指針擴展窗口
	for right < len(s) {
		// 將當前字符添加到窗口計數中
		rightChar := s[right]
		windowCount[rightChar]++

		// 如果這個字符是必需的，並且我們還沒有超過所需的計數
		if tCount[rightChar] > 0 && windowCount[rightChar] <= tCount[rightChar] {
			charsFound++
		}

		// 當所有字符都找到時，嘗試從左側收縮窗口
		for charsFound == totalCharsNeeded {
			// 更新最小窗口
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
				minEnd = right
			}

			// 從窗口中移除最左側的字符
			leftChar := s[left]
			windowCount[leftChar]--

			// 如果此移除影響了我們的字符計數
			if tCount[leftChar] > 0 && windowCount[leftChar] < tCount[leftChar] {
				charsFound--
			}

			// 從左側收縮窗口
			left++
		}

		// 從右側擴展窗口
		right++
	}

	// 返回最小窗口子串，如果沒有找到則返回空字符串
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minStart : minEnd+1]
}
