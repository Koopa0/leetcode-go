package fourSum

import (
	"reflect"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:     "標準測試用例1",
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:     "標準測試用例2",
			nums:     []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
		{
			name:     "空數組",
			nums:     []int{},
			target:   0,
			expected: [][]int{},
		},
		{
			name:     "長度不足4的數組",
			nums:     []int{1, 2, 3},
			target:   6,
			expected: [][]int{},
		},
		{
			name:     "無解的情況",
			nums:     []int{1, 2, 3, 4, 5},
			target:   100,
			expected: [][]int{},
		},
		{
			name:     "有大數的情況",
			nums:     []int{1000000000, 1000000000, 1000000000, 1000000000},
			target:   -294967296,
			expected: [][]int{},
		},
		{
			name:     "有重複元素的情況",
			nums:     []int{0, 0, 0, 0, 0, 0},
			target:   0,
			expected: [][]int{{0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fourSum(tt.nums, tt.target)

			// 由於結果順序可能不同，需要進行排序後比較
			sortNestedArrays := func(arr [][]int) {
				sort.Slice(arr, func(i, j int) bool {
					for k := range arr[i] {
						if arr[i][k] != arr[j][k] {
							return arr[i][k] < arr[j][k]
						}
					}
					return false
				})
			}

			sortNestedArrays(result)
			sortNestedArrays(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("fourSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
