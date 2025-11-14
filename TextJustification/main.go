package TextJustification

import "strings"

// 蠻力解法
func fullJustifyForce(words []string, maxWidth int) []string {
	// 初始化結果陣列和當前行資訊
	result := []string{}
	line := []string{}
	lineLength := 0

	// 逐個處理單詞
	for i := 0; i < len(words); i++ {
		word := words[i]

		// 檢查添加當前單詞後是否超出最大寬度
		// 需要考慮單詞間至少一個空格
		if lineLength+len(word)+len(line) > maxWidth {
			// 格式化當前行並添加到結果中
			result = append(result, formatLine(line, lineLength, maxWidth, false))

			// 重置當前行
			line = []string{word}
			lineLength = len(word)
		} else {
			// 將單詞添加到當前行
			line = append(line, word)
			lineLength += len(word)
		}
	}

	// 處理最後一行
	if len(line) > 0 {
		result = append(result, formatLine(line, lineLength, maxWidth, true))
	}

	return result
}

// 格式化行
func formatLine(line []string, lineLength, maxWidth int, isLastLine bool) string {
	// 當前行只有一個單詞或是最後一行，左對齊
	if len(line) == 1 || isLastLine {
		return leftJustify(line, maxWidth)
	}

	// 計算需要插入的空格總數
	totalSpaces := maxWidth - lineLength
	// 計算單詞間隙數
	slots := len(line) - 1
	// 計算每個間隙的基本空格數
	spacesPerSlot := totalSpaces / slots
	// 計算需要額外加一個空格的間隙數
	extraSpaces := totalSpaces % slots

	var sb strings.Builder

	// 構建對齊的行
	for i, word := range line {
		sb.WriteString(word)

		// 如果不是最後一個單詞，添加空格
		if i < len(line)-1 {
			spaces := spacesPerSlot
			// 左側間隙獲得額外空格
			if i < extraSpaces {
				spaces++
			}
			sb.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return sb.String()
}

// 左對齊行
func leftJustify(line []string, maxWidth int) string {
	var sb strings.Builder
	sb.WriteString(line[0])

	// 添加單詞，每個單詞間一個空格
	for i := 1; i < len(line); i++ {
		sb.WriteString(" ")
		sb.WriteString(line[i])
	}

	// 填充剩餘空格
	remainingSpaces := maxWidth - sb.Len()
	sb.WriteString(strings.Repeat(" ", remainingSpaces))

	return sb.String()
}

// 優化解法
func fullJustifyOptimized(words []string, maxWidth int) []string {
	result := []string{}

	index := 0
	for index < len(words) {
		// 找出當前行可以包含的單詞
		count := 1
		lineLength := len(words[index])
		lastWord := index

		for lastWord+1 < len(words) && lineLength+1+len(words[lastWord+1]) <= maxWidth {
			lastWord++
			lineLength += 1 + len(words[lastWord])
			count++
		}

		// 創建當前行
		var currentLine strings.Builder

		// 處理只有一個單詞或最後一行的情況
		if count == 1 || lastWord == len(words)-1 {
			// 左對齊
			currentLine.WriteString(words[index])

			for i := index + 1; i <= lastWord; i++ {
				currentLine.WriteString(" ")
				currentLine.WriteString(words[i])
			}

			// 填充剩餘空格
			currentLine.WriteString(strings.Repeat(" ", maxWidth-lineLength))
		} else {
			// 計算空格分配
			spaces := maxWidth - (lineLength - (count - 1)) // 減去已計算的單詞間空格
			spacesPerSlot := spaces / (count - 1)
			extraSpaces := spaces % (count - 1)

			currentLine.WriteString(words[index])

			for i := index + 1; i <= lastWord; i++ {
				// 計算當前間隙應有的空格數
				spaceCount := spacesPerSlot
				if i-index-1 < extraSpaces {
					spaceCount++
				}

				currentLine.WriteString(strings.Repeat(" ", spaceCount))
				currentLine.WriteString(words[i])
			}
		}

		result = append(result, currentLine.String())
		index = lastWord + 1
	}

	return result
}

// 最佳解法
func fullJustify(words []string, maxWidth int) []string {
	result := []string{}

	i := 0
	for i < len(words) {
		// 決定當前行包含的單詞範圍 [i, j]
		j := i
		lineLength := len(words[i])

		// 貪婪地添加盡可能多的單詞
		for j+1 < len(words) && lineLength+1+len(words[j+1]) <= maxWidth {
			j++
			lineLength += 1 + len(words[j])
		}

		// 創建當前行
		builder := strings.Builder{}

		// 特殊情況：最後一行或只有一個單詞
		if j == len(words)-1 || i == j {
			builder.WriteString(words[i])

			for k := i + 1; k <= j; k++ {
				builder.WriteByte(' ')
				builder.WriteString(words[k])
			}

			// 填充尾部空格
			remainingSpaces := maxWidth - builder.Len()
			builder.WriteString(strings.Repeat(" ", remainingSpaces))
		} else {
			// 計算空格
			wordCount := j - i + 1
			totalSpaces := maxWidth - (lineLength - (wordCount - 1))

			if wordCount == 1 {
				// 單詞後填充空格
				builder.WriteString(words[i])
				builder.WriteString(strings.Repeat(" ", totalSpaces))
			} else {
				// 計算間隙和每個間隙的空格數
				gaps := wordCount - 1
				spacesPerGap := totalSpaces / gaps
				extraSpaces := totalSpaces % gaps

				builder.WriteString(words[i])

				for k := i + 1; k <= j; k++ {
					// 計算當前間隙的空格數
					currentSpaces := spacesPerGap
					if k-i-1 < extraSpaces {
						currentSpaces++
					}

					builder.WriteString(strings.Repeat(" ", currentSpaces))
					builder.WriteString(words[k])
				}
			}
		}

		result = append(result, builder.String())
		i = j + 1
	}

	return result
}
