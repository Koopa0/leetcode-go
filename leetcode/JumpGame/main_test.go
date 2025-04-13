package JumpGame

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
		desc     string
	}{
		{
			name:     "基本成功情況",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
			desc:     "可以從位置 0 跳到位置 1，再從位置 1 跳到終點",
		},
		{
			name:     "基本失敗情況",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
			desc:     "無論如何都會被困在位置 3",
		},
		{
			name:     "邊緣情況 - 只有一個元素",
			nums:     []int{0},
			expected: true,
			desc:     "已在終點，無需跳躍",
		},
		{
			name:     "邊緣情況 - 第一個為 0",
			nums:     []int{0, 1},
			expected: false,
			desc:     "無法從起點移動",
		},
		{
			name:     "大跳躍",
			nums:     []int{5, 0, 0, 0, 0, 0},
			expected: true,
			desc:     "可以直接從位置 0 跳到終點",
		},
		{
			name:     "剛好達到終點",
			nums:     []int{2, 0, 0},
			expected: true,
			desc:     "可以剛好跳到終點",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJumpGreedy(tt.nums)
			if got != tt.expected {
				t.Errorf("canJumpGreedy() = %v, want %v, 描述: %v", got, tt.expected, tt.desc)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試案例
	input := []int{2, 3, 1, 1, 4}

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			canJumpBruteForce(input)
		}
	})

	b.Run("DynamicProgramming", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			canJumpDP(input)
		}
	})

	b.Run("Greedy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			canJumpGreedy(input)
		}
	})
}
