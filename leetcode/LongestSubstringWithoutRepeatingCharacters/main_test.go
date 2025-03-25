package LongestSubstringWithoutRepeatingCharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		desc     string
	}{
		{
			name:     "Example 1",
			input:    "abcabcbb",
			expected: 3,
			desc:     "測試基本情況，最長子字串為 'abc'",
		},
		{
			name:     "Example 2",
			input:    "bbbbb",
			expected: 1,
			desc:     "測試所有字符都相同的情況",
		},
		{
			name:     "Example 3",
			input:    "pwwkew",
			expected: 3,
			desc:     "測試子字串 'wke'，而非子序列 'pwke'",
		},
		{
			name:     "Empty String",
			input:    "",
			expected: 0,
			desc:     "測試空字串",
		},
		{
			name:     "Single Character",
			input:    "a",
			expected: 1,
			desc:     "測試只有一個字符的字串",
		},
		{
			name:     "All Unique Characters",
			input:    "abcdefg",
			expected: 7,
			desc:     "測試所有字符都不重複的情況",
		},
		{
			name:     "With Spaces",
			input:    "hello world",
			expected: 6,
			desc:     "測試包含空格的字串，最長子字串為 ' world'",
		},
		{
			name:     "With Symbols",
			input:    "a!b@c#d$",
			expected: 8,
			desc:     "測試包含特殊符號的字串",
		},
		{
			name:     "Long String",
			input:    "abcdefghijklmnopqrstuvwxy",
			expected: 25,
			desc:     "測試較長的字串，包含所有英文字母和數字",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.input); got != tt.expected {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v\n%s", got, tt.expected, tt.desc)
			}
		})
	}
}
