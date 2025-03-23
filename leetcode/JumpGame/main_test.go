package JumpGame

import "testing"

func TestCanJump(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		nums     []int
		expected bool
		desc     string
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
			desc:     "能夠從索引 0 跳到索引 1，再從索引 1 跳到索引 4",
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
			desc:     "總會到達索引 3，但該位置的跳躍長度為 0，無法到達索引 4",
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
			desc:     "只有一個元素時，已經在最後一個位置",
		},
		{
			name:     "Start with zero",
			nums:     []int{0, 1},
			expected: false,
			desc:     "起始位置為 0，無法跳躍，且數組長度大於 1",
		},
		{
			name:     "Exactly reachable",
			nums:     []int{1, 2, 3},
			expected: true,
			desc:     "每步都能剛好到達下一個位置",
		},
		{
			name:     "Large jumps",
			nums:     []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			expected: true,
			desc:     "包含大跳躍距離的情況",
		},
		{
			name:     "Trapped by zeros",
			nums:     []int{1, 0, 1, 0},
			expected: false,
			desc:     "會被中間的 0 困住",
		},
		{
			name:     "Maximum input size",
			nums:     make([]int, 10000),
			expected: false,
			desc:     "測試最大輸入大小，所有元素為 0，無法到達最後一個位置",
		},
	}

	// 特殊設置大數組測試
	bigArray := make([]int, 10000)
	for i := range bigArray {
		bigArray[i] = 1 // 每個位置都可以跳 1 步
	}
	tests = append(tests, struct {
		name     string
		nums     []int
		expected bool
		desc     string
	}{
		name:     "Maximum input size - all ones",
		nums:     bigArray,
		expected: true,
		desc:     "測試最大輸入大小，所有元素為 1，可以一步步到達最後一個位置",
	})

	// 運行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			if result != tt.expected {
				t.Errorf("canJump(%v) = %v, want %v, desc: %s", tt.nums, result, tt.expected, tt.desc)
			}
		})
	}
}
