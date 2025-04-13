package PermutationSequence

import "testing"

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected string
		desc     string // 測試目的描述
	}{
		{
			name:     "基本測試1",
			n:        3,
			k:        3,
			expected: "213",
			desc:     "測試 n=3, k=3 的基本情況",
		},
		{
			name:     "基本測試2",
			n:        4,
			k:        9,
			expected: "2314",
			desc:     "測試 n=4, k=9 的基本情況",
		},
		{
			name:     "邊緣情況1",
			n:        1,
			k:        1,
			expected: "1",
			desc:     "測試 n=1 的邊緣情況",
		},
		{
			name:     "邊緣情況2",
			n:        9,
			k:        362880,
			expected: "987654321",
			desc:     "測試 n=9, k=9! 的邊緣情況（最後一個排列）",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPermutation(tt.n, tt.k)
			if got != tt.expected {
				t.Errorf("getPermutation(%v, %v) = %v, want %v", tt.n, tt.k, got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 基準測試的範例輸入
	n, k := 9, 12345

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bruteForceSolution(n, k)
		}
	})

	b.Run("數學方法直接計算", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathematicalSolution(n, k)
		}
	})

	b.Run("使用鏈表的最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getPermutation(n, k)
		}
	})
}
