package GroupAnagrams

import (
	"sort"
	"strings"
)

// 將字符串排序（用於創建字母異位詞的唯一標識）
func sortString(s string) string {
	// 將字符串轉換為字符切片
	chars := strings.Split(s, "")
	// 對字符進行排序
	sort.Strings(chars)
	// 將排序後的字符重新連接為字符串
	return strings.Join(chars, "")
}

// 主函數：對字母異位詞進行分組
func groupAnagrams(strs []string) [][]string {
	// 創建映射：排序後的字符串 -> 原始字符串列表
	anagramMap := make(map[string][]string)

	// 遍歷每個輸入字符串
	for _, str := range strs {
		// 獲取排序後的字符串作為鍵
		sortedStr := sortString(str)

		// 將原始字符串添加到對應的列表中
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// 創建結果數組
	result := make([][]string, 0, len(anagramMap))

	// 收集所有的字符串組
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
