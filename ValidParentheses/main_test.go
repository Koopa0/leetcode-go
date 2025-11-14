package ValidParentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		desc     string
	}{
		{
			name:     "Example 1 - Simple Valid",
			input:    "()",
			expected: true,
			desc:     "測試單對括號",
		},
		{
			name:     "Example 2 - Multiple Valid",
			input:    "()[]{}",
			expected: true,
			desc:     "測試多對不同類型的括號",
		},
		{
			name:     "Example 3 - Invalid Type Matching",
			input:    "(]",
			expected: false,
			desc:     "測試不匹配的括號類型",
		},
		{
			name:     "Nested Valid",
			input:    "([{}])",
			expected: true,
			desc:     "測試嵌套的括號結構",
		},
		{
			name:     "Incorrect Order",
			input:    "([)]",
			expected: false,
			desc:     "測試錯誤順序的括號",
		},
		{
			name:     "Unbalanced Open",
			input:    "(((",
			expected: false,
			desc:     "測試只有開放括號的情況",
		},
		{
			name:     "Unbalanced Close",
			input:    ")))",
			expected: false,
			desc:     "測試只有閉合括號的情況",
		},
		{
			name:     "Single Character",
			input:    "(",
			expected: false,
			desc:     "測試單個開放括號",
		},
		{
			name:     "Complex Nested Valid",
			input:    "{[()]}([]){()}",
			expected: true,
			desc:     "測試複雜的嵌套和連續括號結構",
		},
		{
			name:     "Complex Nested Invalid",
			input:    "{[(])}",
			expected: false,
			desc:     "測試複雜的無效嵌套結構",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			if got != tt.expected {
				t.Errorf("%s: isValid(%q) = %v, want %v, %s",
					tt.name, tt.input, got, tt.expected, tt.desc)
			}
		})
	}
}
