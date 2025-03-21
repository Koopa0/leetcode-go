package GroupAnagrams

import "strings"

// 時間複雜度: O(n * k)，其中 n 是字符串數量，k 是平均字符串長度
func groupAnagramsCount(strs []string) [][]string {
	// 創建一個哈希表，鍵是字符計數的表示，值是原始字符串的列表
	anagramMap := make(map[string][]string)

	// 遍歷所有輸入字符串
	for _, str := range strs {
		// 創建一個長度為26的數組，用於統計字符出現次數
		count := make([]int, 26)

		// 統計每個字符的出現次數
		for _, ch := range str {
			count[ch-'a']++
		}

		// 創建一個特殊的鍵，用於表示字符計數
		// 格式為: "字符:計數#字符:計數#..."
		var key strings.Builder
		for i, cnt := range count {
			if cnt > 0 {
				key.WriteByte(byte(i) + 'a')
				key.WriteByte(':')
				key.WriteString(string(rune(cnt + '0')))
				key.WriteByte('#')
			}
		}

		// 將原始字符串添加到對應的分組中
		anagramMap[key.String()] = append(anagramMap[key.String()], str)
	}

	// 初始化結果數組
	result := make([][]string, 0, len(anagramMap))

	// 將哈希表中的所有值（即所有分組）添加到結果數組中
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
