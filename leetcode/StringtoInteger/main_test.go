package StringtoInteger

import "testing"

func TestMyAtoi(t *testing.T) {
	// 定義測試用例
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		// 基本功能測試
		{"basic_positive", "42", 42},
		{"basic_negative", "-42", -42},
		{"with_whitespace", "   -42", -42},
		{"with_words", "4193 with words", 4193},

		// 邊緣情況
		{"empty_string", "", 0},
		{"only_whitespace", "     ", 0},
		{"only_sign_plus", "+", 0},
		{"only_sign_minus", "-", 0},
		{"sign_with_whitespace", "    +    ", 0},

		// 溢出測試
		{"overflow_positive", "2147483648", 2147483647},   // INT_MAX
		{"overflow_negative", "-2147483649", -2147483648}, // INT_MIN
		{"large_overflow_positive", "9999999999", 2147483647},
		{"large_overflow_negative", "-9999999999", -2147483648},

		// 特殊字
		{"with_dots", "3.14159", 3},
		{"with_plus_and_spaces", "  +  413", 0},
		{"with_multiple_signs", "+-12", 0},
		{"with_letters_before", "words and 987", 0},

		// 零值測試
		{"zero", "0", 0},
		{"multiple_zeros", "000", 0},
		{"negative_zero", "-0", 0},

		// 前導零測試
		{"leading_zeros", "0042", 42},
		{"negative_leading_zeros", "-0042", -42},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := myAtoi(tt.input)
			if actual != tt.expected {
				t.Errorf("myAtoi(%q) = %d; expected %d", tt.input, actual, tt.expected)
			}
		})
	}
}
