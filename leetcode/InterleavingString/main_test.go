package InterleavingString

import (
	"strings"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		s3       string
		expected bool
		desc     string // 測試目的說明
	}{
		{
			name:     "Example1",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
			desc:     "測試基本交錯情況",
		},
		{
			name:     "Example2",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbbaccc",
			expected: false,
			desc:     "測試無法交錯的情況",
		},
		{
			name:     "EmptyStrings",
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
			desc:     "測試空字串",
		},
		{
			name:     "OneEmptyString",
			s1:       "",
			s2:       "abc",
			s3:       "abc",
			expected: true,
			desc:     "測試一個空字串的情況",
		},
		{
			name:     "LengthMismatch",
			s1:       "abc",
			s2:       "def",
			s3:       "abcde",
			expected: false,
			desc:     "測試長度不匹配的情況",
		},
		{
			name:     "LargeStrings",
			s1:       strings.Repeat("a", 50),
			s2:       strings.Repeat("b", 50),
			s3:       strings.Repeat("ab", 50),
			expected: true,
			desc:     "測試較大字串的性能",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 測試暴力解法（小輸入）
			if len(tt.s1) < 20 && len(tt.s2) < 20 {
				got := isInterleaveBruteForce(tt.s1, tt.s2, tt.s3)
				if got != tt.expected {
					t.Errorf("isInterleaveBruteForce() = %v, want %v", got, tt.expected)
				}
			}

			// 測試記憶化解法
			got := isInterleaveOptimized(tt.s1, tt.s2, tt.s3)
			if got != tt.expected {
				t.Errorf("isInterleaveOptimized() = %v, want %v", got, tt.expected)
			}

			// 測試動態規劃解法
			got = isInterleaveDP(tt.s1, tt.s2, tt.s3)
			if got != tt.expected {
				t.Errorf("isInterleaveDP() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	s1 := "aabccaabccaabcc"
	s2 := "dbbcadbbcadbbca"
	s3 := "aadbbcbcacaadbbcbcacaadbbcbcac"

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveBruteForce(s1, s2, s3)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveOptimized(s1, s2, s3)
		}
	})

	b.Run("DP", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveDP(s1, s2, s3)
		}
	})
}
