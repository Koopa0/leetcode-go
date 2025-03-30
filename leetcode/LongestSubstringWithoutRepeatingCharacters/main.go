package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {

	// 假設字符集僅為 ASCII
	chart := [128]int{}

	// 遍歷初始化每個都是-1 假設都沒被找到
	for i := range chart {
		chart[i] = -1
	}

	maxlength := 0
	left := 0

	for right := 0; right < len(s); right++ {

		// chart["a"] > left 代表被找到了 所以 left 要移到被找到的重複的字的後一個位置
		if chart[s[right]] >= left {
			left = chart[s[right]] + 1
		}

		chart[s[right]] = right
		// 因為是左右指針 且index 為０ 開始 所以如果直接相減會少1
		// 所以正確長度為左右指針相減 + 1
		currentLength := right - left + 1
		// 假設當前找到的長度大於儲存的最長長度 則儲存
		if currentLength > maxlength {
			maxlength = currentLength
		}
	}

	return maxlength
}
