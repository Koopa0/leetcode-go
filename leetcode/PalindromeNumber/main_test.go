package PalindromeNumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	// 使用表格驅動測試
	tests := []struct {
		name     string
		input    int
		expected bool
		desc     string
	}{
		// 基本功能測試
		{"TC1", 121, true, "基本回文數測試"},
		{"TC2", 123, false, "非回文數測試"},
		{"TC3", 12321, true, "奇數位數回文測試"},
		{"TC4", 1221, true, "偶數位數回文測試"},

		// 邊界情況測試
		{"TC5", 0, true, "零是回文數"},
		{"TC6", -121, false, "負數不是回文數"},
		{"TC7", 10, false, "末尾為0的數不是回文數（除了0）"},

		// 極端情況測試
		{"TC8", 1, true, "單個數字是回文數"},
		{"TC9", 1000000001, true, "大數回文測試"},
		{"TC10", 2147483647, false, "最大32位整數（非回文）"},
		{"TC11", -2147483648, false, "最小32位整數（負數非回文）"},

		// 特殊情況測試
		{"TC12", 123454321, true, "大回文數測試"},
		{"TC13", 22, true, "兩位數相同回文測試"},
		{"TC14", 1000021, false, "含有前導0的非回文測試"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("測試 %s 失敗：輸入 %d，期望 %v，得到 %v, 描述：%s",
					tc.name, tc.input, tc.expected, result, tc.desc)
			}
		})
	}
}
