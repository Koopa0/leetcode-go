package LengthLastWord

import "strings"

// 暴力解法 - 分割字串
func lengthOfLastWordBruteForce(s string) int {
	// 使用空格分割字串
	words := strings.Fields(s)

	// 如果沒有單詞，返回 0（根據題目，這種情況不會發生）
	if len(words) == 0 {
		return 0
	}

	// 返回最後一個單詞的長度
	return len(words[len(words)-1])
}

// 優化解法 - 從後向前遍歷
func lengthOfLastWordOptimized(s string) int {
	// 獲取字串長度
	length := len(s)

	// 計數器，用於記錄最後一個單詞的長度
	count := 0

	// 從尾部開始向前遍歷
	for i := length - 1; i >= 0; i-- {
		// 如果遇到非空格字元
		if s[i] != ' ' {
			count++
		} else if count > 0 {
			// 如果已經計數了一些字元，並且現在遇到空格，說明已經找到最後一個單詞
			break
		}
		// 如果是尾部空格（count==0），則繼續向前遍歷
	}

	return count
}

// 最佳解法 - 單次遍歷的改進版
func lengthOfLastWord(s string) int {
	// 獲取字串長度
	length := len(s)

	// 從尾部跳過空格，找到最後一個單詞的結束位置
	end := length - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}

	// 如果字串全是空格（根據題目，這種情況不會發生）
	if end < 0 {
		return 0
	}

	// 從最後一個單詞的結束位置向前找到其開始位置
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	// 返回最後一個單詞的長度
	return end - start
}
