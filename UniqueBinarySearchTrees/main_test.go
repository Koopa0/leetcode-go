package UniqueBinarySearchTrees

import "testing"

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		desc     string
	}{
		{
			name:     "單節點樹",
			input:    1,
			expected: 1,
			desc:     "測試單節點樹的情況",
		},
		{
			name:     "兩個節點",
			input:    2,
			expected: 2,
			desc:     "測試兩個節點的情況",
		},
		{
			name:     "三個節點",
			input:    3,
			expected: 5,
			desc:     "測試三個節點的情況",
		},
		{
			name:     "四個節點",
			input:    4,
			expected: 14,
			desc:     "測試四個節點的情況",
		},
		{
			name:     "大輸入",
			input:    10,
			expected: 16796,
			desc:     "測試較大輸入值",
		},
		{
			name:     "最大輸入",
			input:    19,
			expected: 1767263190,
			desc:     "測試問題約束的最大輸入",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numTrees(tt.input)
			if got != tt.expected {
				t.Errorf("numTrees() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試輸入
	input := 10

	b.Run("暴力遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesRecursive(input)
		}
	})

	b.Run("記憶化遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesMemo(input)
		}
	})

	b.Run("動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTrees(input)
		}
	})
}
