package LongestSubstringWithoutRepeatingCharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	// 定義測試用例表
	tests := []struct {
		name     string // 測試用例名稱
		input    string // 輸入字符串
		expected int    // 預期輸出
		desc     string // 測試用例描述
	}{
		{
			name:     "Example 1",
			input:    "abcabcbb",
			expected: 3,
			desc:     "測試普通情況，最長子串為 'abc'",
		},
		{
			name:     "Example 2",
			input:    "bbbbb",
			expected: 1,
			desc:     "測試所有字符相同的情況",
		},
		{
			name:     "Example 3",
			input:    "pwwkew",
			expected: 3,
			desc:     "測試最長子串在中間的情況",
		},
		{
			name:     "Empty String",
			input:    "",
			expected: 0,
			desc:     "測試空字符串",
		},
		{
			name:     "Single Character",
			input:    "a",
			expected: 1,
			desc:     "測試單個字符",
		},
		{
			name:     "All Unique Characters",
			input:    "abcdefg",
			expected: 7,
			desc:     "測試所有字符都不重複的情況",
		},
		{
			name:     "Repeated Patterns",
			input:    "abcdabcd",
			expected: 4,
			desc:     "測試重複模式",
		},
		{
			name:     "Complex String",
			input:    "!@#$%^&*()_+",
			expected: 12,
			desc:     "測試特殊字符",
		},
		{
			name:     "Mixed Characters",
			input:    "a1b2c3d4e5",
			expected: 10,
			desc:     "測試混合字符",
		},
		{
			name:     "Long String",
			input:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			expected: 62,
			desc:     "測試長字符串，所有字符都不重複",
		},
		{
			name:     "Long String with Repeats",
			input:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expected: 26,
			desc:     "測試長字符串，有重複字符",
		},
	}

	// 執行測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 調用被測試的函數
			got := lengthOfLongestSubstring(tt.input)

			// 驗證結果
			if got != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%s) = %d, want %d\n%s", tt.input, got, tt.expected, tt.desc)
			}
		})
	}
}

// 基準測試，用於測量函數的性能
func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	// 重置計時器
	b.ResetTimer()

	// 執行 b.N 次迭代
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")
	}
}
