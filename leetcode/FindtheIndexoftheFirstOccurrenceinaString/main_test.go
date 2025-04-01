package FindtheIndexoftheFirstOccurrenceinaString

import (
	"strings"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		// 基本測試用例
		{"Example 1", "sadbutsad", "sad", 0},
		{"Example 2", "leetcode", "leeto", -1},

		// 邊界情況
		{"Empty haystack", "", "a", -1},
		{"Empty needle", "abc", "", 0}, // 根據定義，空 needle 應該返回 0
		{"Identical strings", "abc", "abc", 0},

		// 特殊情況
		{"Needle at end", "abcabc", "bc", 1},
		{"Multiple occurrences", "abababa", "aba", 0},
		{"No match", "abcdef", "xyz", -1},

		// 長字符串
		{"Long strings", "a" + strings.Repeat("b", 10000) + "c", "bc", 10000},

		// 重複字符
		{"Repeating characters", "aaaaa", "aa", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
