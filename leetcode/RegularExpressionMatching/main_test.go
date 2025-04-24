package RegularExpressionMatching

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	// 定義測試表格
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
		desc     string
	}{
		// 基本功能測試
		{
			name:     "基本匹配 - 無特殊字",
			s:        "abc",
			p:        "abc",
			expected: true,
			desc:     "測試基本字串匹配，無特殊字",
		},
		{
			name:     "基本不匹配 - 無特殊字",
			s:        "abc",
			p:        "abd",
			expected: false,
			desc:     "測試基本字串不匹配",
		},

		// '.' 字測試
		{
			name:     "點匹配單個字",
			s:        "abc",
			p:        "a.c",
			expected: true,
			desc:     "測試 '.' 匹配任意單個字",
		},
		{
			name:     "多個點匹配",
			s:        "abc",
			p:        "...",
			expected: true,
			desc:     "測試多個 '.' 匹配多個字",
		},

		// '*' 字測試
		{
			name:     "*匹配零次",
			s:        "abc",
			p:        "abd*c",
			expected: true,
			desc:     "測試 '*' 匹配零次前面的元素",
		},
		{
			name:     "*匹配一次",
			s:        "abc",
			p:        "ab*c",
			expected: true,
			desc:     "測試 '*' 匹配一次前面的元素",
		},
		{
			name:     "*匹配多次",
			s:        "abbbc",
			p:        "ab*c",
			expected: true,
			desc:     "測試 '*' 匹配多次前面的元素",
		},

		// 組合測試 (. 和 *)
		{
			name:     ".*匹配",
			s:        "abcdefg",
			p:        "a.*g",
			expected: true,
			desc:     "測試 '.*' 組合匹配任意字序列",
		},
		{
			name:     "複雜組合測試1",
			s:        "aab",
			p:        "c*a*b",
			expected: true,
			desc:     "測試複雜的 '*' 組合，包括匹配零次",
		},
		{
			name:     "複雜組合測試2",
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
			desc:     "測試更複雜的模式",
		},

		// 邊界情況
		{
			name:     "空字串匹配",
			s:        "",
			p:        "",
			expected: true,
			desc:     "測試空字串匹配空模式",
		},
		{
			name:     "空字串匹配星號模式",
			s:        "",
			p:        "a*",
			expected: true,
			desc:     "測試空字串匹配包含 '*' 的模式",
		},
		{
			name:     "空字串匹配複雜星號模式",
			s:        "",
			p:        "a*b*c*",
			expected: true,
			desc:     "測試空字串匹配複雜的 '*' 模式",
		},
		{
			name:     "空模式匹配非空字串",
			s:        "abc",
			p:        "",
			expected: false,
			desc:     "測試空模式匹配非空字串",
		},

		// 極端情況
		{
			name:     "長字串匹配",
			s:        "aaaaaaaaaaaaaaaaaaab",
			p:        "a*a*a*a*a*a*a*a*a*a*c",
			expected: false,
			desc:     "測試長字串匹配長模式",
		},
		{
			name:     "以星號結尾",
			s:        "aaa",
			p:        "a*",
			expected: true,
			desc:     "測試以 '*' 結尾的模式匹配",
		},
	}

	// 運行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 測試動態規劃版本
			result := isMatch(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatch(%q, %q) = %v, 期望 %v, 描述: %s",
					tt.s, tt.p, result, tt.expected, tt.desc)
			}
		})
	}
}

// 性能測試
func BenchmarkIsMatch(b *testing.B) {
	// 基準測試例子
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{"短字串", "abc", "a.c"},
		{"中等字串", "aabcbcbcbc", "a*b*c*b*c*"},
		{"長字串", "aaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*b"},
	}

	// 運行基準測試
	for _, tc := range testCases {
		b.Run("DP_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isMatch(tc.s, tc.p)
			}
		})
	}
}
