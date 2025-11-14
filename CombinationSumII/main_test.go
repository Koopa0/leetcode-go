package CombinationSumII

import (
	"reflect"
	"sort"
	"testing"
)

// 用於對二維切片進行排序以便比較
func sortResult(nums [][]int) {
	// 首先對每個內部切片進行排序
	for i := range nums {
		sort.Ints(nums[i])
	}

	// 然後對整個二維切片進行排序
	sort.Slice(nums, func(i, j int) bool {
		// 比較兩個切片的長度
		if len(nums[i]) != len(nums[j]) {
			return len(nums[i]) < len(nums[j])
		}

		// 如果長度相同，依次比較每個元素
		for k := 0; k < len(nums[i]); k++ {
			if nums[i][k] != nums[j][k] {
				return nums[i][k] < nums[j][k]
			}
		}

		return false // 如果完全相同，則保持原順序
	})
}

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want:       [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want:       [][]int{{1, 2, 2}, {5}},
		},
		{
			name:       "No Solution",
			candidates: []int{1, 2, 3},
			target:     10,
			want:       [][]int{},
		},
		{
			name:       "Single Element Equal to Target",
			candidates: []int{7, 2, 5, 8},
			target:     8,
			want:       [][]int{{8}},
		},
		{
			name:       "Multiple Duplicate Elements",
			candidates: []int{1, 1, 1, 1, 1, 1, 1},
			target:     4,
			want:       [][]int{{1, 1, 1, 1}},
		},
		{
			name:       "Minimum Input Size",
			candidates: []int{1},
			target:     1,
			want:       [][]int{{1}},
		},
		{
			name:       "All Elements Larger Than Target",
			candidates: []int{10, 20, 30},
			target:     5,
			want:       [][]int{},
		},
		{
			name:       "Maximum Target Value",
			candidates: []int{10, 10, 10},
			target:     30,
			want:       [][]int{{10, 10, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)

			// 排序結果以進行比較
			sortResult(got)
			sortResult(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
