package MaximumSubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		desc     string // 測試目的描述
	}{
		{
			name:     "Example 1",
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
			desc:     "標準測試案例",
		},
		{
			name:     "Example 2",
			input:    []int{1},
			expected: 1,
			desc:     "只有一個元素的陣列",
		},
		{
			name:     "Example 3",
			input:    []int{5, 4, -1, 7, 8},
			expected: 23,
			desc:     "全部元素之和就是最大子陣列",
		},
		{
			name:     "All Negative",
			input:    []int{-1, -2, -3, -4},
			expected: -1,
			desc:     "全部是負數的情況",
		},
		{
			name:     "Alternating",
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
			desc:     "正負交替的陣列",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.input)
			if got != tt.expected {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試用例
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxSubArrayBruteForce(input)
		}
	})

	b.Run("DynamicProgramming", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxSubArrayDP(input)
		}
	})

	b.Run("Kadane", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxSubArray(input)
		}
	})
}
