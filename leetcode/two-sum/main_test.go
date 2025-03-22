package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "基本案例",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "有重複元素",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "負數元素",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "較大",
			nums:   []int{1, 3, 5, 7, 9, 11, 13, 16, 17},
			target: 16,
			want:   []int{3, 4},
		},
		{
			name:   "無解案例",
			nums:   []int{1, 2, 3, 4},
			target: 8,
			want:   nil,
		},
	}

	// 運行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
