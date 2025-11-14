package ScrambleString

import "testing"

func TestIsScramble(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
		desc     string // 測試目的描述
	}{
		{
			name:     "Example 1",
			s1:       "great",
			s2:       "rgeat",
			expected: true,
			desc:     "標準測試案例 - 可擾亂",
		},
		{
			name:     "Example 2",
			s1:       "abcde",
			s2:       "caebd",
			expected: false,
			desc:     "標準測試案例 - 不可擾亂",
		},
		{
			name:     "Example 3",
			s1:       "a",
			s2:       "a",
			expected: true,
			desc:     "邊界情況 - 單一字元",
		},
		{
			name:     "Same Strings",
			s1:       "abc",
			s2:       "abc",
			expected: true,
			desc:     "特殊情況 - 相同字串",
		},
		{
			name:     "Different Character Sets",
			s1:       "abc",
			s2:       "def",
			expected: false,
			desc:     "特殊情況 - 不同字元集",
		},
		{
			name:     "Long String",
			s1:       "abcdefghijklmn",
			s2:       "efghijklmnabc",
			expected: false,
			desc:     "效能測試 - 較長字串",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isScramble(tt.s1, tt.s2)
			if got != tt.expected {
				t.Errorf("isScramble() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkScrambleSolutions(b *testing.B) {
	// 實例輸入
	s1 := "great"
	s2 := "rgeat"

	b.Run("Brute Force", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isScrambleBruteForce(s1, s2)
		}
	})

	b.Run("Memoization", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isScrambleMemoization(s1, s2)
		}
	})

	b.Run("Dynamic Programming", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isScrambleDP(s1, s2)
		}
	})

	b.Run("Optimized DP", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isScramble(s1, s2)
		}
	})
}
