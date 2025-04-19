package AddBinary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
		desc     string // 測試目的描述
	}{
		{
			name:     "基本測試1",
			a:        "11",
			b:        "1",
			expected: "100",
			desc:     "測試基本二進位加法",
		},
		{
			name:     "基本測試2",
			a:        "1010",
			b:        "1011",
			expected: "10101",
			desc:     "測試多位二進位加法",
		},
		{
			name:     "僅一個輸入為0",
			a:        "0",
			b:        "1",
			expected: "1",
			desc:     "測試一個輸入為0的情況",
		},
		{
			name:     "兩個輸入都為0",
			a:        "0",
			b:        "0",
			expected: "0",
			desc:     "測試兩個輸入都為0的情況",
		},
		{
			name:     "長度不同的輸入",
			a:        "1",
			b:        "111",
			expected: "1000",
			desc:     "測試輸入長度差異大的情況",
		},
		{
			name:     "較大數值測試",
			a:        "11111111",
			b:        "11111111",
			expected: "111111110",
			desc:     "測試較大數值的加法",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("addBinary() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 用於基準測試的輸入
	a := "10101010101010101010"
	bVeryLarge := "11111111111111111111"

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 注意：對於大數輸入，這可能會失敗
			addBinaryBruteForce(a, bVeryLarge)
		}
	})

	b.Run("模擬加法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addBinarySimulation(a, bVeryLarge)
		}
	})

	b.Run("優化的迭代解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addBinary(a, bVeryLarge)
		}
	})
}
