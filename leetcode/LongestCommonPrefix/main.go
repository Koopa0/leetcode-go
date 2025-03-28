package LongestCommonPrefix

import "strings"

// longestCommonPrefix 查找字符串數組中的最長公共前綴
// 如果不存在公共前綴，返回空字符串 ""
func longestCommonPrefix(strs []string) string {
	// 處理邊緣情況: 空數組或包含空字符串
	if len(strs) == 0 {
		return ""
	}

	// 選擇第一個字符串作為初始前綴
	prefix := strs[0]

	// 遍歷其餘的字符串
	for i := 1; i < len(strs); i++ {
		// 當當前字符串不以當前前綴開頭時
		for !strings.HasPrefix(strs[i], prefix) {
			// 縮短前綴，移除最後一個字符
			prefix = prefix[:len(prefix)-1]

			// 如果前綴變為空，則表示沒有公共前綴
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
