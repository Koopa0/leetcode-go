package UniquePaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
		desc     string
	}{
		{
			name:     "3x7 Grid",
			m:        3,
			n:        7,
			expected: 28,
			desc:     "LeetCode 範例 1",
		},
		{
			name:     "3x2 Grid",
			m:        3,
			n:        2,
			expected: 3,
			desc:     "LeetCode 範例 2",
		},
		{
			name:     "1x1 Grid",
			m:        1,
			n:        1,
			expected: 1,
			desc:     "邊界情況：只有一個格子",
		},
		{
			name:     "1x5 Grid",
			m:        1,
			n:        5,
			expected: 1,
			desc:     "邊界情況：只有一列",
		},
		{
			name:     "5x1 Grid",
			m:        5,
			n:        1,
			expected: 1,
			desc:     "邊界情況：只有一行",
		},
		{
			name:     "10x10 Grid",
			m:        10,
			n:        10,
			expected: 48620,
			desc:     "較大的輸入",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 測試所有實現
			got1 := uniquePathsMemoization(tt.m, tt.n)
			if got1 != tt.expected {
				t.Errorf("uniquePathsMemoization() = %v, want %v", got1, tt.expected)
			}

			got2 := uniquePathsDP(tt.m, tt.n)
			if got2 != tt.expected {
				t.Errorf("uniquePathsDP() = %v, want %v", got2, tt.expected)
			}

			got3 := uniquePathsOptimized(tt.m, tt.n)
			if got3 != tt.expected {
				t.Errorf("uniquePathsOptimized() = %v, want %v", got3, tt.expected)
			}

			got4 := uniquePathsCombination(tt.m, tt.n)
			if got4 != tt.expected {
				t.Errorf("uniquePathsCombination() = %v, want %v", got4, tt.expected)
			}
		})
	}
}

func BenchmarkUniquePaths(b *testing.B) {
	// 中等規模的輸入
	m, n := 10, 10

	b.Run("Memoization", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uniquePathsMemoization(m, n)
		}
	})

	b.Run("DynamicProgramming", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uniquePathsDP(m, n)
		}
	})

	b.Run("OptimizedDP", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uniquePathsOptimized(m, n)
		}
	})

	b.Run("Combination", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uniquePathsCombination(m, n)
		}
	})
}
