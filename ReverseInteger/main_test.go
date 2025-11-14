package ReverseInteger

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    int
		description string
	}{
		// 基本測試案例
		{"positive_number", 123, 321, "反轉正數"},
		{"negative_number", -123, -321, "反轉負數"},
		{"zero", 0, 0, "反轉零"},

		// 末尾零測試
		{"trailing_zeros", 120, 21, "末尾有零的數字"},
		{"trailing_zeros_negative", -120, -21, "末尾有零的負數"},

		// 邊界值測試
		{"single_digit", 7, 7, "單個數字"},
		{"single_digit_negative", -8, -8, "單個負數數字"},

		// 溢出測試
		{"overflow_positive", 1534236469, 0, "反轉後正溢出"},
		{"overflow_negative", -1563847412, 0, "反轉後負溢出"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.input)
			if result != tt.expected {
				t.Errorf("reverse(%d) = %d; want %d (%s)",
					tt.input, result, tt.expected, tt.description)
			}
		})
	}
}
