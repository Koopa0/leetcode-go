package FindtheIndexoftheFirstOccurrenceinaString

func strStr(haystack string, needle string) int {
	// 獲取兩個字串的長度
	n, m := len(haystack), len(needle)

	// 如果 needle 為空，返回 0（題目約束已確保 needle 不為空）
	if m == 0 {
		return 0
	}

	// 如果 needle 長度大於 haystack 長度，不可能找到匹配
	if m > n {
		return -1
	}

	// 構建部分匹配表（PMT）
	// next[i] 表示 needle[0...i] 的最長相L同前綴後綴長度
	next := make([]int, m)

	// 初始化
	//next[0] = 0

	// 填充 next 數組
	for i, j := 1, 0; i < m; {
		if needle[i] == needle[j] {
			j++
			next[i] = j
			i++
		} else if j > 0 {
			j = next[j-1]
		} else {
			next[i] = 0
			i++
		}
	}

	// 使用 KMP 算法在 haystack 中搜索 needle
	for i, j := 0, 0; i < n; {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == m {
				// 找到完全匹配，返回起始索引
				return i - j
			}
		} else if j > 0 {
			// 部分匹配失敗，使用 next 數組回退
			j = next[j-1]
		} else {
			// 第一個字就不匹配，直接往前移動
			i++
		}
	}

	// 如果沒有找到匹配，返回 -1
	return -1
}
