package ValidPalindrome

import (
	"testing"
)

// 測試 isPalindrome 函數
func TestIsPalindrome(t *testing.T) {
	// 定義測試案例結構
	type testCase struct {
		name     string
		input    string
		expected bool
	}

	// 創建測試案例表
	tests := []testCase{
		{
			name:     "正常回文句子",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "非回文句子",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "空字串",
			input:    "",
			expected: true,
		},
		{
			name:     "只有標點符號",
			input:    ".,!?",
			expected: true,
		},
		{
			name:     "只有一個字符",
			input:    "a",
			expected: true,
		},
		{
			name:     "只有數字的回文",
			input:    "123321",
			expected: true,
		},
		{
			name:     "只有數字的非回文",
			input:    "123456",
			expected: false,
		},
		{
			name:     "混合大小寫的回文",
			input:    "AbBa",
			expected: true,
		},
		{
			name:     "混合字母數字的回文",
			input:    "a1b2b1a",
			expected: true,
		},
		{
			name:     "混合字母數字的非回文",
			input:    "a1b2c3",
			expected: false,
		},
		{
			name:     "很長的回文",
			input:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			expected: true,
		},
		{
			name:     "包含空格的回文",
			input:    "never odd or even",
			expected: true,
		},
		{
			name:     "包含大量標點的回文",
			input:    "A Santa at NASA!",
			expected: true,
		},
	}

	// 運行所有測試案例
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 調用被測試的函數
			result := isPalindrome(tc.input)

			// 驗證結果
			if result != tc.expected {
				t.Errorf("isPalindrome(%q) = %v; 期望 %v", tc.input, result, tc.expected)
			}
		})
	}
}

// 測試輔助函數 isAlphanumeric
func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		name     string
		char     byte
		expected bool
	}{
		{"小寫字母", 'a', true},
		{"大寫字母", 'Z', true},
		{"數字", '9', true},
		{"空格", ' ', false},
		{"標點符號", '.', false},
		{"特殊字符", '*', false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isAlphanumeric(tc.char)
			if result != tc.expected {
				t.Errorf("isAlphanumeric(%q) = %v; 期望 %v", tc.char, result, tc.expected)
			}
		})
	}
}

// 測試輔助函數 toLower
func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		char     byte
		expected byte
	}{
		{"小寫字母", 'a', 'a'},
		{"大寫字母", 'Z', 'z'},
		{"數字", '9', '9'},
		{"空格", ' ', ' '},
		{"標點符號", '.', '.'},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := toLower(tc.char)
			if result != tc.expected {
				t.Errorf("toLower(%q) = %q; 期望 %q", tc.char, result, tc.expected)
			}
		})
	}
}

// 如果使用了 isDigit 函數，也應該測試它
func TestIsDigit(t *testing.T) {
	tests := []struct {
		name     string
		char     byte
		expected bool
	}{
		{"數字0", '0', true},
		{"數字9", '9', true},
		{"小寫字母", 'a', false},
		{"大寫字母", 'Z', false},
		{"空格", ' ', false},
		{"標點符號", '.', false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isDigit(tc.char)
			if result != tc.expected {
				t.Errorf("isDigit(%q) = %v; 期望 %v", tc.char, result, tc.expected)
			}
		})
	}
}
