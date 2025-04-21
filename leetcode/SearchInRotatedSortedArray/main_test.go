package SearchInRotatedSortedArray

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "基本示例1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "基本示例2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "單元素數組，目標存在",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "單元素數組，目標不存在",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "未旋轉數組",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "目標在旋轉點左側",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "目標在旋轉點右側",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: 5,
		},
		{
			name:     "目標為數組最小值",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "目標為數組最大值",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   7,
			expected: 3,
		},
		{
			name:     "極端旋轉：旋轉1次",
			nums:     []int{2, 1},
			target:   1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %d, 期望 %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
