package JumpGameII

import "testing"

func TestJump(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "示例 2",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "單元素數組",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "已經在終點",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "每步只能跳 1",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "第一步可直達終點",
			nums:     []int{5, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "大型數組",
			nums:     []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := jump(tc.nums)
			if result != tc.expected {
				t.Errorf("jump(%v) = %d, 預期 %d", tc.nums, result, tc.expected)
			}
		})
	}
}
