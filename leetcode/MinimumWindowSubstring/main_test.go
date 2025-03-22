package MinimumWindowSubstring

import (
	"strings"
	"testing"
)

func TestMinWindow(t *testing.T) {
	// 定義測試用例
	testCases := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{
			name:     "範例 1：基本情況",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			name:     "範例 2：單個字符",
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			name:     "範例 3：不可能的情況",
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			name:     "邊界情況：空字符串",
			s:        "",
			t:        "a",
			expected: "",
		},
		{
			name:     "邊界情況：空目標",
			s:        "a",
			t:        "",
			expected: "",
		},
		{
			name:     "邊界情況：相同字符串",
			s:        "abc",
			t:        "abc",
			expected: "abc",
		},
		{
			name:     "特殊情況：重複字符",
			s:        "aaaaaaaaaaaabbbbbcdd",
			t:        "abcdd",
			expected: "abbbbbcdd",
		},
		{
			name:     "性能測試：長字符串",
			s:        strings.Repeat("abcdefghijklmnopqrstuvwxyz", 1000),
			t:        "xyz",
			expected: "xyz",
		},
	}

	// 運行測試用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minWindow(tc.s, tc.t)
			if result != tc.expected {
				t.Errorf("預期 %q，得到 %q", tc.expected, result)
			}
		})
	}
}
