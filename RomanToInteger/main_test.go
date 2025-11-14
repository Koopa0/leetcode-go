package RomanToInteger

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		desc     string
	}{
		{
			name:     "基本測試 - III",
			input:    "III",
			expected: 3,
			desc:     "測試基本加法：I + I + I = 3",
		},
		{
			name:     "減法規則 - IV",
			input:    "IV",
			expected: 4,
			desc:     "測試基本減法規則：IV = 4",
		},
		{
			name:     "減法規則 - IX",
			input:    "IX",
			expected: 9,
			desc:     "測試基本減法規則：IX = 9",
		},
		{
			name:     "混合情況 - LVIII",
			input:    "LVIII",
			expected: 58,
			desc:     "測試混合情況：L + V + I + I + I = 58",
		},
		{
			name:     "複雜情況 - MCMXCIV",
			input:    "MCMXCIV",
			expected: 1994,
			desc:     "測試複雜情況：M + CM + XC + IV = 1000 + 900 + 90 + 4 = 1994",
		},
		{
			name:     "邊界情況 - 最小值 I",
			input:    "I",
			expected: 1,
			desc:     "測試最小值：I = 1",
		},
		{
			name:     "邊界情況 - 最大值 MMMCMXCIX",
			input:    "MMMCMXCIX",
			expected: 3999,
			desc:     "測試最大值：MMM + CM + XC + IX = 3000 + 900 + 90 + 9 = 3999",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := romanToInt(test.input)
			if result != test.expected {
				t.Errorf("romanToInt(%s) = %d; 期望 %d - %s",
					test.input, result, test.expected, test.desc)
			}
		})
	}
}
