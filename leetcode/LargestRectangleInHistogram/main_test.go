package LargestRectangleInHistogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
		desc     string
	}{
		{
			name:     "Example 1",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
			desc:     "測試基本功能",
		},
		{
			name:     "Example 2",
			heights:  []int{2, 4},
			expected: 4,
			desc:     "測試只有兩個元素的情況",
		},
		{
			name:     "Empty Array",
			heights:  []int{},
			expected: 0,
			desc:     "測試空陣列",
		},
		{
			name:     "Single Element",
			heights:  []int{5},
			expected: 5,
			desc:     "測試只有一個元素的情況",
		},
		{
			name:     "All Same Height",
			heights:  []int{3, 3, 3, 3, 3},
			expected: 15,
			desc:     "測試所有柱子高度相同的情況",
		},
		{
			name:     "Increasing Heights",
			heights:  []int{1, 2, 3, 4, 5},
			expected: 9,
			desc:     "測試高度遞增的情況",
		},
		{
			name:     "Decreasing Heights",
			heights:  []int{5, 4, 3, 2, 1},
			expected: 9,
			desc:     "測試高度遞減的情況",
		},
		{
			name:     "Large Input",
			heights:  generateLargeInput(1000),
			expected: 1000,
			desc:     "測試大型輸入",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.heights)
			if got != tt.expected {
				t.Errorf("largestRectangleArea() = %v, 預期 %v", got, tt.expected)
			}
		})
	}
}

func generateLargeInput(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = 1
	}
	return result
}

func BenchmarkSolutions(b *testing.B) {
	// 用於基準測試的範例輸入
	input := []int{2, 1, 5, 6, 2, 3}

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			largestRectangleAreaBruteForce(input)
		}
	})

	b.Run("DivideConquer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			largestRectangleAreaDivideConquer(input)
		}
	})

	b.Run("MonotonicStack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			largestRectangleArea(input)
		}
	})
}
