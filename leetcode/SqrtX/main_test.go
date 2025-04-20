package SqrtX

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		desc     string // 測試目的描述
	}{
		// 基本功能測試
		{"基本測試-完全平方數", 4, 2, "對於完全平方數應該返回其平方根"},
		{"基本測試-非完全平方數", 8, 2, "對於非完全平方數應該返回向下取整的平方根"},

		// 邊界情況測試
		{"邊界測試-零", 0, 0, "對於0應該返回0"},
		{"邊界測試-一", 1, 1, "對於1應該返回1"},

		// 特殊情況測試
		{"特殊測試-最大整數", 2147483647, 46340, "對於最大整數應該能正確計算"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newtonSqrt(tt.input)
			if got != tt.expected {
				t.Errorf("mySqrt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 用於基準測試的輸入
	input := 2147483647

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bruteForceSqrt(input)
		}
	})

	b.Run("二分搜尋解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			binarySearchSqrt(input)
		}
	})

	b.Run("牛頓法解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			newtonSqrt(input)
		}
	})
}
