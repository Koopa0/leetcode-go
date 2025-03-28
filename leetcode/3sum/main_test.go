package threesum

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
		desc     string
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
			desc:     "基本測試用例，有多個解",
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
			desc:     "無解的情況",
		},
		{
			name:     "Example 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
			desc:     "全零數組",
		},
		{
			name:     "Contains duplicates",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
			desc:     "包含重複元素的情況",
		},
		{
			name:     "All negative",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: [][]int{},
			desc:     "全為負數，無解",
		},
		{
			name:     "All positive",
			nums:     []int{1, 2, 3, 4, 5},
			expected: [][]int{},
			desc:     "全為正數，無解",
		},
		{
			name:     "Large input",
			nums:     generateLargeInput(1000), // 假設有一個生成大輸入的函數
			expected: nil,                      // 此處不關注具體輸出，主要測試性能
			desc:     "大規模輸入測試",
		},
		{
			name:     "Minimum length",
			nums:     []int{1, 2, 3},
			expected: [][]int{},
			desc:     "最小長度為3的數組",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)

			// 由於結果的順序不重要，需要先排序再比較
			sortResults := func(res [][]int) {
				for i := range res {
					sort.Ints(res[i])
				}
				sort.Slice(res, func(i, j int) bool {
					for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
						if res[i][k] != res[j][k] {
							return res[i][k] < res[j][k]
						}
					}
					return len(res[i]) < len(res[j])
				})
			}

			if tt.name != "Large input" { // 對於大輸入，不檢查具體結果
				sortResults(got)
				sortResults(tt.expected)

				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.expected)
				}
			}
		})
	}
}

// 生成大規模測試數據
func generateLargeInput(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(200) - 100 // 生成 [-100, 99] 範圍內的隨機數
	}
	return result
}
