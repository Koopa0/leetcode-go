package MinimumWindowSubstring

import (
	"strings"
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected string
		desc     string // 測試目的描述
	}{
		{
			name:     "基本功能測試 1",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
			desc:     "測試基本滑動視窗功能",
		},
		{
			name:     "基本功能測試 2",
			s:        "a",
			t:        "a",
			expected: "a",
			desc:     "測試單字元匹配",
		},
		{
			name:     "邊界情況測試 1",
			s:        "a",
			t:        "aa",
			expected: "",
			desc:     "測試無法匹配的情況",
		},
		{
			name:     "邊界情況測試 2",
			s:        "",
			t:        "a",
			expected: "",
			desc:     "測試空字串輸入",
		},
		{
			name:     "邊界情況測試 3",
			s:        "ab",
			t:        "",
			expected: "ab",
			desc:     "測試空目標字串",
		},
		{
			name:     "特殊情況測試",
			s:        "AAAAAAAABC",
			t:        "ABC",
			expected: "ABC",
			desc:     "測試連續重複字元",
		},
		{
			name:     "大型輸入測試",
			s:        strings.Repeat("A", 10000) + "BC",
			t:        "ABC",
			expected: "ABC",
			desc:     "測試大型輸入處理效率",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.expected {
				t.Errorf("minWindow() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 用於效能測試的輸入
	s := "ADOBECODEBANC"
	t := "ABC"

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minWindowBruteForce(s, t)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minWindowOptimized(s, t)
		}
	})

	b.Run("Optimal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minWindow(s, t)
		}
	})
}
