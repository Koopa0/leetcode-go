package LongestCommonPrefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		input    []string
		expected string
		desc     string
	}{
		{
			name:     "Example 1",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
			desc:     "基本情況：三個字符串具有公共前綴 'fl'",
		},
		{
			name:     "Example 2",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
			desc:     "沒有公共前綴",
		},
		{
			name:     "Single string",
			input:    []string{"alone"},
			expected: "alone",
			desc:     "只有一個字符串時，前綴就是該字符串本身",
		},
		{
			name:     "Empty array",
			input:    []string{},
			expected: "",
			desc:     "空數組應返回空字符串",
		},
		{
			name:     "Contains empty string",
			input:    []string{"test", "", "testing"},
			expected: "",
			desc:     "包含空字符串時應返回空字符串",
		},
		{
			name:     "All identical",
			input:    []string{"abc", "abc", "abc"},
			expected: "abc",
			desc:     "所有字符串相同時，前綴就是該字符串",
		},
		{
			name:     "Long strings with short prefix",
			input:    []string{"aaaaaaaaaaaaaaaaaaabc", "aaaaaaaaaaaaaaaaaaade", "aaaaaaaaaaaaaaaaaaaxy"},
			expected: "aaaaaaaaaaaaaaaaaaa",
			desc:     "長字符串具有較長的公共前綴",
		},
		{
			name:     "Different string lengths",
			input:    []string{"ab", "abcd", "abc"},
			expected: "ab",
			desc:     "不同長度的字符串",
		},
		{
			name:     "Unicode characters",
			input:    []string{"你好世界", "你好朋友", "你好家人"},
			expected: "你好",
			desc:     "測試Unicode字符",
		},
	}

	// 執行測試表中的所有測試案例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %v; want %v\n%s",
					tt.input, result, tt.expected, tt.desc)
			}
		})
	}
}
