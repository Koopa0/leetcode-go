package Combinations

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
		desc     string
	}{
		{
			name:     "範例1",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			desc:     "測試基本功能",
		},
		{
			name:     "範例2",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
			desc:     "測試邊界情況：n=k=1",
		},
		{
			name:     "空組合",
			n:        5,
			k:        0,
			expected: [][]int{},
			desc:     "測試邊界情況：k=0",
		},
		{
			name:     "k大於n",
			n:        3,
			k:        4,
			expected: [][]int{},
			desc:     "測試邊界情況：k>n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combineIterative(tt.n, tt.k)
			// 排序結果以便比較，因為順序可能不同
			sortResults := func(res [][]int) {
				for i := range res {
					sort.Ints(res[i])
				}
				sort.Slice(res, func(i, j int) bool {
					for k := 0; k < len(res[i]); k++ {
						if k >= len(res[j]) {
							return false
						}
						if res[i][k] != res[j][k] {
							return res[i][k] < res[j][k]
						}
					}
					return len(res[i]) < len(res[j])
				})
			}

			sortResults(got)
			sortResults(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("combine() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkCombine(b *testing.B) {
	// 基準測試使用的輸入
	n, k := 10, 5

	b.Run("基本回溯法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			combineBacktrack(n, k)
		}
	})

	b.Run("優化回溯法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			combineBacktrackOptimized(n, k)
		}
	})

	b.Run("字典序迭代實現", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			combineIterative(n, k)
		}
	})
}
