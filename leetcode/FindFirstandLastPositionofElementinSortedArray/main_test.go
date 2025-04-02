package FindFirstandLastPositionofElementinSortedArray

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "標準情況 - 目標值存在多次",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "目標值不存在",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "空數組",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "只有一個元素且匹配",
			nums:     []int{5},
			target:   5,
			expected: []int{0, 0},
		},
		{
			name:     "只有一個元素但不匹配",
			nums:     []int{5},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "目標值在數組開頭",
			nums:     []int{8, 8, 9, 10},
			target:   8,
			expected: []int{0, 1},
		},
		{
			name:     "目標值在數組結尾",
			nums:     []int{5, 6, 7, 8, 8},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "所有元素都是目標值",
			nums:     []int{8, 8, 8, 8, 8},
			target:   8,
			expected: []int{0, 4},
		},
		{
			name:     "大型數組 - 目標值存在",
			nums:     generateSortedArray(1000, 500, 520),
			target:   510,
			expected: []int{10, 10},
		},
		{
			name:     "大型數組 - 目標值多次出現",
			nums:     generateSortedArray(1000, 500, 505),
			target:   502,
			expected: []int{2, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := searchRange(test.nums, test.target)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("對於 %v 和目標值 %d，期望 %v，但得到 %v",
					test.nums, test.target, test.expected, result)
			}
		})
	}
}

// 生成測試用的排序數組
func generateSortedArray(size, start, duplicateVal int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		if i == duplicateVal {
			arr[i] = duplicateVal
		} else {
			arr[i] = start + i
		}
	}
	return arr
}
