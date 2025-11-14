package ZigzagConversion

import "strings"

func convert(s string, numRows int) string {
	// 處理特殊情況
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	// 使用 strings.Builder 來提高效率
	builders := make([]strings.Builder, numRows)

	// 預估每個 builder 的容量
	estCapacity := (len(s) / numRows) + 1
	for i := range builders {
		builders[i].Grow(estCapacity)
	}

	// 計算曲折模式的週期長度
	cycleLen := 2 * (numRows - 1)

	// 遍歷原始字串的每個字
	for i := 0; i < len(s); i++ {
		position := i % cycleLen
		if position < numRows {
			builders[position].WriteByte(s[i])
		} else {
			builders[cycleLen-position].WriteByte(s[i])
		}
	}

	// 連接所有行的字串
	result := strings.Builder{}
	result.Grow(len(s))
	for _, builder := range builders {
		result.WriteString(builder.String())
	}

	return result.String()
}
