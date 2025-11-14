package DecodeWays

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		desc     string
	}{
		{
			name:     "基本功能測試1",
			input:    "12",
			expected: 2,
			desc:     "可以解碼為 'AB' (1, 2) 或 'L' (12)",
		},
		{
			name:     "基本功能測試2",
			input:    "226",
			expected: 3,
			desc:     "可以解碼為 'BZ' (2, 26), 'VF' (22, 6), 或 'BBF' (2, 2, 6)",
		},
		{
			name:     "邊界情況測試1",
			input:    "0",
			expected: 0,
			desc:     "單個0無法解碼",
		},
		{
			name:     "邊界情況測試2",
			input:    "06",
			expected: 0,
			desc:     "以0開頭且不組成有效的兩位數",
		},
		{
			name:     "特殊情況測試1",
			input:    "10",
			expected: 1,
			desc:     "只能解碼為 'J' (10)",
		},
		{
			name:     "特殊情況測試2",
			input:    "101",
			expected: 1,
			desc:     "只能解碼為 'JA' (10, 1)",
		},
		{
			name:     "連續0測試",
			input:    "100",
			expected: 0,
			desc:     "無法解碼，因為00不是有效的編碼",
		},
		{
			name:     "大輸入測試",
			input:    "11111111111111111111",
			expected: 10946,
			desc:     "斐波那契數列F(21)，因為每個位置可以單獨解碼或與前一個組合",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optimizedDpNumDecodings(tt.input)
			if got != tt.expected {
				t.Errorf("numDecodings(%q) = %v, want %v - %s", tt.input, got, tt.expected, tt.desc)
			}
		})
	}
}

func BenchmarkNumDecodings(b *testing.B) {
	// 準備測試輸入
	inputs := []string{"12", "226", "11111111111111111111"}

	b.Run("暴力遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, input := range inputs {
				bruteForceNumDecodings(input)
			}
		}
	})

	b.Run("記憶化遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, input := range inputs {
				memoizedNumDecodings(input)
			}
		}
	})

	b.Run("動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, input := range inputs {
				dpNumDecodings(input)
			}
		}
	})

	b.Run("空間優化動態規劃", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, input := range inputs {
				optimizedDpNumDecodings(input)
			}
		}
	})
}
