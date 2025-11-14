package LongestValidParentheses

import (
	"strings"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "示例 1：普通情況",
			input:    "(()",
			expected: 2,
		},
		{
			name:     "示例 2：連續有效括號",
			input:    ")()())",
			expected: 4,
		},
		{
			name:     "空字串",
			input:    "",
			expected: 0,
		},
		{
			name:     "全部無效括號",
			input:    ")))(((",
			expected: 0,
		},
		{
			name:     "全部有效括號",
			input:    "()()()",
			expected: 6,
		},
		{
			name:     "嵌套括號",
			input:    "((()))",
			expected: 6,
		},
		{
			name:     "混合括號",
			input:    "(())())",
			expected: 6,
		},
		{
			name:     "長字串",
			input:    strings.Repeat("()", 15000),
			expected: 30000,
		},
		{
			name:     "只有左括號",
			input:    "(((((",
			expected: 0,
		},
		{
			name:     "只有右括號",
			input:    ")))))",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestValidParentheses(tc.input)
			if result != tc.expected {
				t.Errorf("預期 %d，得到 %d", tc.expected, result)
			}
		})
	}
}
