package LengthLastWord

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		desc     string // 測試目的描述
	}{
		{
			name:     "基本功能測試 1",
			input:    "Hello World",
			expected: 5,
			desc:     "測試基本情況，兩個單詞間有一個空格",
		},
		{
			name:     "基本功能測試 2",
			input:    "   fly me   to   the moon  ",
			expected: 4,
			desc:     "測試前後有空格，單詞間有多個空格的情況",
		},
		{
			name:     "基本功能測試 3",
			input:    "luffy is still joyboy",
			expected: 6,
			desc:     "測試多個單詞的情況",
		},
		{
			name:     "邊界情況測試 1",
			input:    "a",
			expected: 1,
			desc:     "測試只有一個字元的情況",
		},
		{
			name:     "邊界情況測試 2",
			input:    "a ",
			expected: 1,
			desc:     "測試尾部有空格的情況",
		},
		{
			name:     "邊界情況測試 3",
			input:    " a",
			expected: 1,
			desc:     "測試開頭有空格的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLastWord(tt.input)
			if got != tt.expected {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 用於基準測試的範例輸入
	input := "   fly me   to   the moon  "

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLastWordBruteForce(input)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLastWordOptimized(input)
		}
	})

	b.Run("Optimal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLastWord(input)
		}
	})
}
