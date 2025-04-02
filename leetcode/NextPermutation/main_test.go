package NextPermutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	// 定義表格驅動測試
	tests := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{
			name:     "基本測試 - 簡單遞增數組",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
			desc:     "驗證簡單的遞增數組的下一個排列",
		},
		{
			name:     "基本測試 - 完全降序數組",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
			desc:     "當數組完全降序時，應返回升序排列（最小排列）",
		},
		{
			name:     "基本測試 - 包含重複元素",
			input:    []int{1, 1, 5},
			expected: []int{1, 5, 1},
			desc:     "驗證處理重複元素的能力",
		},
		{
			name:     "單元素數組",
			input:    []int{1},
			expected: []int{1},
			desc:     "單元素數組沒有其他排列，應保持不變",
		},
		{
			name:     "相同元素數組",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
			desc:     "全部元素相同的數組沒有其他排列，應保持不變",
		},
		{
			name:     "複雜測試 - 部分遞增部分遞減",
			input:    []int{1, 3, 5, 4, 2},
			expected: []int{1, 4, 2, 3, 5},
			desc:     "測試更複雜的部分遞增部分遞減數組",
		},
		{
			name:     "包含零的數組",
			input:    []int{0, 1, 2, 5, 3, 3, 0},
			expected: []int{0, 1, 3, 0, 2, 3, 5},
			desc:     "驗證包含零和其他重複元素的處理",
		},
		{
			name:     "兩個元素 - 遞增",
			input:    []int{1, 2},
			expected: []int{2, 1},
			desc:     "兩個元素的遞增數組",
		},
		{
			name:     "兩個元素 - 遞減",
			input:    []int{2, 1},
			expected: []int{1, 2},
			desc:     "兩個元素的遞減數組",
		},
		{
			name:     "性能測試 - 較大數組",
			input:    generateLargeArray(100), // 生成一個較大的測試數組
			expected: nil,                     // 這裡我們不比較具體結果，只測試不會崩潰
			desc:     "測試處理較大數組的性能和穩定性",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製輸入以避免修改原測試數據
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// 執行函數
			nextPermutation(input)

			// 對於性能測試，我們不檢查具體結果
			if tt.name == "性能測試 - 較大數組" {
				return
			}

			// 檢查結果是否符合預期
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("期望 %v，但得到 %v。%s", tt.expected, input, tt.desc)
			}
		})
	}
}

// 生成用於性能測試的大數組
func generateLargeArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = size - i
	}
	return result
}
