package SubstringwithConcatenationofAllWords

func findSubstring(s string, words []string) []int {

	result := []int{}

	// 處理邊界情況
	if len(s) == 0 || len(words) == 0 {
		return result
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount

	if len(s) < totalLen {
		return result
	}

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	for i := 0; i < wordLen; i++ {
		// 提前終止條件：剩餘字符串長度不足
		if i+totalLen > len(s) {
			break
		}

		left, right := i, i
		currFreq := make(map[string]int)
		count := 0

		for right+wordLen <= len(s) {
			// 提前終止條件：剩餘長度不足
			if right+totalLen-(right-left) > len(s) {
				break
			}

			currWord := s[right : right+wordLen]
			right += wordLen

			if _, exists := wordFreq[currWord]; !exists {
				currFreq = make(map[string]int)
				count = 0
				left = right
				continue
			}

			currFreq[currWord]++
			count++

			for currFreq[currWord] > wordFreq[currWord] {
				leftWord := s[left : left+wordLen]
				currFreq[leftWord]--
				count--
				left += wordLen
			}

			if count == wordCount {
				result = append(result, left)

				leftWord := s[left : left+wordLen]
				currFreq[leftWord]--
				count--
				left += wordLen
			}
		}
	}

	return result
}
