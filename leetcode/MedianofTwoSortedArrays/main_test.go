package MedianofTwoSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
		desc  string
	}{
		{
			name:  "示例 1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
			desc:  "基本情況，總長度為奇數",
		},
		{
			name:  "示例 2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
			desc:  "基本情況，總長度為偶數",
		},
		{
			name:  "空數組測試",
			nums1: []int{},
			nums2: []int{1},
			want:  1.0,
			desc:  "測試其中一個數組為空的情況",
		},
		{
			name:  "數組長度差異大",
			nums1: []int{1},
			nums2: []int{2, 3, 4, 5, 6, 7, 8},
			want:  4.5,
			desc:  "測試兩個數組長度差異很大的情況",
		},
		{
			name:  "重複元素",
			nums1: []int{1, 1, 3, 3},
			nums2: []int{2, 2},
			want:  2.0,
			desc:  "測試包含重複元素的情況",
		},
		{
			name:  "負數元素",
			nums1: []int{-5, -3, -1},
			nums2: []int{-6, -4, -2, 0},
			want:  -3.0,
			desc:  "測試包含負數的情況",
		},
		{
			name:  "極端情況：非常大的正負數",
			nums1: []int{-1000000, 1000000},
			nums2: []int{-999999, 999999},
			want:  0.0,
			desc:  "測試包含極端值的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
