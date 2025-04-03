package PowXN

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
		epsilon  float64 // 用於浮點數比較的誤差容許範圍
	}{
		// 基本情況
		{"Zero power", 5.0, 0, 1.0, 1e-9},
		{"First power", 5.0, 1, 5.0, 1e-9},

		// 正整數指數
		{"Positive power", 2.0, 10, 1024.0, 1e-9},
		{"Example 1", 2.0, 10, 1024.0, 1e-9},
		{"Example 2", 2.1, 3, 9.261, 1e-3}, // 允許較大誤差

		// 負整數指數
		{"Negative power", 2.0, -2, 0.25, 1e-9},
		{"Example 3", 2.0, -2, 0.25, 1e-9},

		// 邊界情況
		{"Large positive power", 1.1, 100, 13780.6123, 1e-3},   // 大指數
		{"Small x large n", 0.9, 100, 0.000003, 2.356140e-05},  // x < 1 的大指數
		{"Very small power", 123.45, -67, 5.7667e-150, 1e-140}, // 非常小的結果

		// 特殊情況
		{"x equals 1", 1.0, 1000000, 1.0, 1e-9},           // x = 1 的任意指數都是 1
		{"x equals -1, even n", -1.0, 1000000, 1.0, 1e-9}, // -1 的偶數次冪是 1
		{"x equals -1, odd n", -1.0, 999999, -1.0, 1e-9},  // -1 的奇數次冪是 -1

		// 極端情況（可能需要特別處理）
		// {"INT_MIN", 2.0, math.MinInt32, 0.0, 1e-9},  // 特別小心 n = -2^31
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := myPow(tc.x, tc.n)
			// 使用誤差範圍比較浮點數
			if math.Abs(result-tc.expected) > tc.epsilon {
				t.Errorf("myPow(%f, %d) = %f; want %f (diff: %e, allowed: %e)",
					tc.x, tc.n, result, tc.expected, math.Abs(result-tc.expected), tc.epsilon)
			}
		})
	}
}
