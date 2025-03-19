package TrappingRainWater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "LeetCode 示例",
			heights:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "空陣列",
			heights:  []int{},
			expected: 0,
		},
		{
			name:     "單元素陣列",
			heights:  []int{5},
			expected: 0,
		},
		{
			name:     "兩個元素的陣列",
			heights:  []int{5, 3},
			expected: 0,
		},
		{
			name:     "單調遞增",
			heights:  []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "單調遞減",
			heights:  []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "平坦地形",
			heights:  []int{3, 3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "U形地形",
			heights:  []int{3, 0, 3},
			expected: 3,
		},
		{
			name:     "W形地形",
			heights:  []int{3, 0, 1, 0, 3},
			expected: 8,
		},
		{
			name:     "解釋範例",
			heights:  []int{3, 1, 0, 2, 4},
			expected: 6,
		},
		{
			name:     "極端高度差",
			heights:  []int{10000, 0, 10000},
			expected: 10000,
		},
		{
			name:     "多個凹槽",
			heights:  []int{5, 2, 1, 2, 1, 5},
			expected: 14,
		},
		{
			name:     "兩端低中間高",
			heights:  []int{1, 4, 2, 5, 1, 2, 3},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.heights)

			if got != tt.expected {
				t.Errorf("trap(%v) = %d, 期望 %d", tt.heights, got, tt.expected)
			}
		})
	}
}

func BenchmarkTrap(b *testing.B) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 5, 2, 4, 1, 0, 3}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trap(heights)
	}
}

func BenchmarkTrapTableDriven(b *testing.B) {
	benchmarks := []struct {
		name    string
		heights []int
	}{
		{"小輸入", []int{3, 0, 3}},
		{"中輸入", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{"大輸入", make([]int, 1000)},
	}

	for i := range benchmarks[2].heights {
		benchmarks[2].heights[i] = i % 10
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				trap(bm.heights)
			}
		})
	}
}

func TestTrapWithExamples(t *testing.T) {
	heights := []int{3, 1, 0, 2, 4}
	expected := 6

	got := trap(heights)

	if got != expected {
		t.Errorf("trap(%v) = %d, 期望 %d", heights, got, expected)
	}

	positionWater := []int{0, 2, 3, 1, 0}
	total := 0

	for i, water := range positionWater {
		total += water
		t.Logf("位置 %d (高度 %d): 水量 = %d", i, heights[i], water)
	}

	t.Logf("總水量: %d", total)

	if total != expected {
		t.Errorf("手動計算總水量 = %d, 期望 %d", total, expected)
	}
}
