package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "基本案例",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "全零案例",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "無解案例",
			nums: []int{1, 2, 3, 4},
			want: [][]int{},
		},
		{
			name: "重複數字案例",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)

			// 排序每個三元組，以便比較
			for _, triplet := range got {
				sort.Ints(triplet)
			}
			for _, triplet := range tt.want {
				sort.Ints(triplet)
			}

			// 排序結果，以便比較
			sort.Slice(got, func(i, j int) bool {
				if got[i][0] != got[j][0] {
					return got[i][0] < got[j][0]
				}
				if got[i][1] != got[j][1] {
					return got[i][1] < got[j][1]
				}
				return got[i][2] < got[j][2]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				if tt.want[i][0] != tt.want[j][0] {
					return tt.want[i][0] < tt.want[j][0]
				}
				if tt.want[i][1] != tt.want[j][1] {
					return tt.want[i][1] < tt.want[j][1]
				}
				return tt.want[i][2] < tt.want[j][2]
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
