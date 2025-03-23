package ProductofArrayExceptSelf

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	// 定義測試表格
	tests := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{
			name:     "範例1",
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
			desc:     "測試基本功能 - 所有元素均為正數",
		},
		{
			name:     "範例2",
			input:    []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
			desc:     "測試包含零的情況",
		},
		{
			name:     "最小輸入",
			input:    []int{1, 2},
			expected: []int{2, 1},
			desc:     "測試最小長度數組",
		},
		{
			name:     "包含多個零",
			input:    []int{0, 0, 1, 2},
			expected: []int{0, 0, 0, 0},
			desc:     "測試包含多個零的情況",
		},
		{
			name:     "所有元素相同",
			input:    []int{2, 2, 2, 2},
			expected: []int{8, 8, 8, 8},
			desc:     "測試所有元素相同的情況",
		},
		{
			name:     "負數元素",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-24, -12, -8, -6},
			desc:     "測試所有元素為負數的情況",
		},
		{
			name:     "混合正負元素",
			input:    []int{-1, 2, -3, 4},
			expected: []int{-24, 12, -8, 6},
			desc:     "測試正負元素混合的情況",
		},
		{
			name:     "元素為1的情況",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
			desc:     "測試所有元素為1的特殊情況",
		},
	}

	// 執行測試表格中的每個測試案例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v, 期望 %v, %s", tt.input, result, tt.expected, tt.desc)
			}
		})
	}
}

// 性能測試
func BenchmarkProductExceptSelf(b *testing.B) {
	// 創建一個較大的測試數組
	largeInput := make([]int, 1000)
	for i := range largeInput {
		largeInput[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productExceptSelf(largeInput)
	}
}
