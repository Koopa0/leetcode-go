package DivideTwoIntegers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string // 測試名稱
		dividend int    // 輸入：被除數
		divisor  int    // 輸入：除數
		want     int    // 期望輸出
		desc     string // 測試描述
	}{
		// 基本功能測試
		{
			name:     "正數除法",
			dividend: 10,
			divisor:  3,
			want:     3,
			desc:     "測試基本的正數除法，結果向零截斷",
		},
		{
			name:     "負數除數",
			dividend: 7,
			divisor:  -3,
			want:     -2,
			desc:     "測試除數為負數的情況",
		},
		{
			name:     "負數被除數",
			dividend: -7,
			divisor:  3,
			want:     -2,
			desc:     "測試被除數為負數的情況",
		},
		{
			name:     "兩數皆為負",
			dividend: -7,
			divisor:  -3,
			want:     2,
			desc:     "測試被除數和除數都為負數的情況",
		},

		// 邊界情況測試
		{
			name:     "被除數為零",
			dividend: 0,
			divisor:  5,
			want:     0,
			desc:     "測試被除數為零的情況，結果應為零",
		},
		{
			name:     "除數為1",
			dividend: 10,
			divisor:  1,
			want:     10,
			desc:     "測試除數為1的情況，結果應與被除數相同",
		},
		{
			name:     "除數為-1",
			dividend: 10,
			divisor:  -1,
			want:     -10,
			desc:     "測試除數為-1的情況，結果應為被除數的相反數",
		},
		{
			name:     "溢出情況",
			dividend: math.MinInt32,
			divisor:  -1,
			want:     math.MaxInt32,
			desc:     "測試特殊溢出情況：最小整數除以-1",
		},

		// 大數測試
		{
			name:     "大數除法",
			dividend: 1000000000,
			divisor:  3,
			want:     333333333,
			desc:     "測試大數除法的效率和精度",
		},
		{
			name:     "最小整數除以2",
			dividend: math.MinInt32,
			divisor:  2,
			want:     math.MinInt32 / 2, // -1073741824
			desc:     "測試最小整數除以2的情況",
		},

		// 精確截斷測試
		{
			name:     "精確截斷測試1",
			dividend: 5,
			divisor:  2,
			want:     2,
			desc:     "測試正數除法的精確截斷",
		},
		{
			name:     "精確截斷測試2",
			dividend: -5,
			divisor:  2,
			want:     -2,
			desc:     "測試負數除法的精確截斷",
		},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divide(tt.dividend, tt.divisor)
			if got != tt.want {
				t.Errorf("divide(%v, %v) = %v, 期望 %v, %s",
					tt.dividend, tt.divisor, got, tt.want, tt.desc)
			}
		})
	}
}

// 測試除數為0的情況（在實際環境中應該拋出錯誤）
func TestDivideByZero(t *testing.T) {
	result := divide(10, 0)
	if result != math.MaxInt32 {
		t.Errorf("divide(10, 0) = %v, 期望 %v 或拋出錯誤",
			result, math.MaxInt32)
	}
}

// 性能測試
func BenchmarkDivide(b *testing.B) {
	// 大數除法性能測試
	for i := 0; i < b.N; i++ {
		divide(math.MaxInt32, 3)
	}
}

// 比較不同除法實現的性能
func BenchmarkDivideImplementations(b *testing.B) {
	b.Run("標準實現", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			divide(math.MaxInt32, 3)
		}
	})

	// 如果有替代實現，可以添加其他性能測試
	// b.Run("替代實現", func(b *testing.B) {
	//     for i := 0; i < b.N; i++ {
	//         divideAlternative(math.MaxInt32, 3)
	//     }
	// })
}
