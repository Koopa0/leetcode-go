package EditDistance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
		desc     string
	}{
		{
			name:     "基本測試1",
			word1:    "horse",
			word2:    "ros",
			expected: 3,
			desc:     "基本的編輯距離計算",
		},
		{
			name:     "基本測試2",
			word1:    "intention",
			word2:    "execution",
			expected: 5,
			desc:     "較長字串的編輯距離",
		},
		{
			name:     "相同字串",
			word1:    "same",
			word2:    "same",
			expected: 0,
			desc:     "相同字串的編輯距離應該為0",
		},
		{
			name:     "一個空字串",
			word1:    "",
			word2:    "abc",
			expected: 3,
			desc:     "一個字串為空時的編輯距離",
		},
		{
			name:     "兩個空字串",
			word1:    "",
			word2:    "",
			expected: 0,
			desc:     "兩個空字串的編輯距離應該為0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if got != tt.expected {
				t.Errorf("minDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 測試用例
	word1 := "intention"
	word2 := "execution"

	b.Run("暴力遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minDistanceRecursive(word1, word2)
		}
	})

	b.Run("記憶化遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minDistanceMemoized(word1, word2)
		}
	})

	b.Run("動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			minDistance(word1, word2)
		}
	})
}
