package IntegerToRoman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		// 基本測試用例
		{"單一符號 - 1", 1, "I"},
		{"單一符號 - 5", 5, "V"},
		{"單一符號 - 10", 10, "X"},
		{"單一符號 - 50", 50, "L"},
		{"單一符號 - 100", 100, "C"},
		{"單一符號 - 500", 500, "D"},
		{"單一符號 - 1000", 1000, "M"},

		// 重複符號測試
		{"重複符號 - 3", 3, "III"},
		{"重複符號 - 30", 30, "XXX"},
		{"重複符號 - 300", 300, "CCC"},
		{"重複符號 - 3000", 3000, "MMM"},

		// 減法規則測試
		{"減法規則 - 4", 4, "IV"},
		{"減法規則 - 9", 9, "IX"},
		{"減法規則 - 40", 40, "XL"},
		{"減法規則 - 90", 90, "XC"},
		{"減法規則 - 400", 400, "CD"},
		{"減法規則 - 900", 900, "CM"},

		// 組合測試
		{"組合 - 6", 6, "VI"},
		{"組合 - 7", 7, "VII"},
		{"組合 - 8", 8, "VIII"},
		{"組合 - 14", 14, "XIV"},
		{"組合 - 58", 58, "LVIII"},
		{"組合 - 1994", 1994, "MCMXCIV"},

		// 邊界測試
		{"邊界 - 最小", 1, "I"},
		{"邊界 - 最大", 3999, "MMMCMXCIX"},

		// LeetCode 提供的例子
		{"LeetCode 示例 1", 3, "III"},
		{"LeetCode 示例 2", 4, "IV"},
		{"LeetCode 示例 3", 9, "IX"},
		{"LeetCode 示例 4", 58, "LVIII"},
		{"LeetCode 示例 5", 1994, "MCMXCIV"},
	}

	// 測試貪心算法實現
	for _, tt := range tests {
		t.Run(tt.name+"_貪心算法", func(t *testing.T) {
			result := intToRoman(tt.input)
			if result != tt.expected {
				t.Errorf("intToRoman(%d) = %s, 期望 %s", tt.input, result, tt.expected)
			}
		})
	}
}

// 性能測試
func BenchmarkIntToRoman(b *testing.B) {
	// 測試用例
	testCases := []int{3, 4, 9, 58, 1994, 3999}

	b.Run("貪心算法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				intToRoman(tc)
			}
		}
	})
}
