package StringtoInteger

import "testing"

func TestMyAtoi(t *testing.T) {
	// 定義測試用例表
	tests := []struct {
		name     string // 測試名稱
		input    string // 輸入字串
		expected int    // 期望輸出
		desc     string // 測試案例說明
	}{
		// 基本功能測試案例
		{
			name:     "Basic positive number",
			input:    "42",
			expected: 42,
			desc:     "測試基本的正整數轉換",
		},
		{
			name:     "Number with leading spaces",
			input:    "   42",
			expected: 42,
			desc:     "測試帶有前導空格的正整數",
		},
		{
			name:     "Negative number with leading spaces",
			input:    "   -42",
			expected: -42,
			desc:     "測試帶有前導空格的負整數",
		},
		{
			name:     "Number with trailing words",
			input:    "4193 with words",
			expected: 4193,
			desc:     "測試帶有尾隨單詞的整數",
		},
		{
			name:     "Positive sign with number",
			input:    "+123",
			expected: 123,
			desc:     "測試帶有正號的整數",
		},

		// 邊界情況測試案例
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
			desc:     "測試空字串",
		},
		{
			name:     "Only spaces",
			input:    "     ",
			expected: 0,
			desc:     "測試只包含空格的字串",
		},
		{
			name:     "Only sign without number",
			input:    "-",
			expected: 0,
			desc:     "測試只有符號而沒有數字的情況",
		},
		{
			name:     "Invalid start character",
			input:    "words and 987",
			expected: 0,
			desc:     "測試以非法字元開始的字串",
		},
		{
			name:     "Zero value",
			input:    "0",
			expected: 0,
			desc:     "測試零值",
		},
		{
			name:     "Zero with leading zeros",
			input:    "000",
			expected: 0,
			desc:     "測試帶有前導零的零值",
		},

		// 溢出測試案例
		{
			name:     "Positive overflow",
			input:    "2147483648", // MAX_INT32 + 1
			expected: 2147483647,   // MAX_INT32
			desc:     "測試正整數溢出情況",
		},
		{
			name:     "Large positive overflow",
			input:    "9999999999",
			expected: 2147483647, // MAX_INT32
			desc:     "測試大正整數溢出情況",
		},
		{
			name:     "Negative overflow",
			input:    "-2147483649", // MIN_INT32 - 1
			expected: -2147483648,   // MIN_INT32
			desc:     "測試負整數溢出情況",
		},
		{
			name:     "Large negative overflow",
			input:    "-9999999999",
			expected: -2147483648, // MIN_INT32
			desc:     "測試大負整數溢出情況",
		},

		// 特殊情況測試案例
		{
			name:     "Multiple signs",
			input:    "+-42",
			expected: 0,
			desc:     "測試多個符號的情況",
		},
		{
			name:     "Decimal number",
			input:    "3.14159",
			expected: 3,
			desc:     "測試帶小數點的數字",
		},
		{
			name:     "Number with spaces in the middle",
			input:    "123 456",
			expected: 123,
			desc:     "測試中間帶空格的數字",
		},
		{
			name:     "Number with sign in the middle",
			input:    "123-456",
			expected: 123,
			desc:     "測試中間帶符號的數字",
		},
		{
			name:     "Number at the maximum edge",
			input:    "2147483647", // MAX_INT32
			expected: 2147483647,
			desc:     "測試最大 32 位整數值",
		},
		{
			name:     "Number at the minimum edge",
			input:    "-2147483648", // MIN_INT32
			expected: -2147483648,
			desc:     "測試最小 32 位整數值",
		},
	}

	// 執行所有測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.input)
			if got != tt.expected {
				t.Errorf("myAtoi(%q) = %d; want %d - %s",
					tt.input, got, tt.expected, tt.desc)
			}
		})
	}
}

// 效能測試
func BenchmarkMyAtoi(b *testing.B) {
	// 基準測試用例
	testCases := []string{
		"42",          // 簡單正數
		"-42",         // 簡單負數
		"   42",       // 帶前導空格
		"2147483647",  // 最大 32 位整數
		"-2147483648", // 最小 32 位整數
		"9999999999",  // 溢出案例
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			myAtoi(tc)
		}
	}
}
