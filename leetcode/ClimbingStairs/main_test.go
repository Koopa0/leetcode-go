package ClimbingStairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		desc     string // 測試目的描述
	}{
		{
			name:     "基本測試 - 2 階",
			input:    2,
			expected: 2,
			desc:     "測試基本情況 n=2",
		},
		{
			name:     "基本測試 - 3 階",
			input:    3,
			expected: 3,
			desc:     "測試基本情況 n=3",
		},
		{
			name:     "中等測試 - 5 階",
			input:    5,
			expected: 8,
			desc:     "測試中等情況 n=5",
		},
		{
			name:     "中等測試 - 10 階",
			input:    10,
			expected: 89,
			desc:     "測試中等情況 n=10",
		},
		{
			name:     "極限測試 - 45 階",
			input:    45,
			expected: 1836311903,
			desc:     "測試限制條件下的極限情況 n=45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairsOptimized(tt.input)
			if got != tt.expected {
				t.Errorf("climbStairsOptimized() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	// 以 n=30 為例進行基準測試
	n := 30

	b.Run("暴力遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairsBruteForce(n)
		}
	})

	b.Run("記憶化遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairsWithMemo(n)
		}
	})

	b.Run("動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairsDP(n)
		}
	})

	b.Run("空間最佳化動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairsOptimized(n)
		}
	})
}
