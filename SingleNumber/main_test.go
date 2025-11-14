package main

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "範例1：兩個重複，一個單獨",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "範例2：五個數字",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "範例3：單一元素",
			nums: []int{1},
			want: 1,
		},
		{
			name: "單獨數字在開頭",
			nums: []int{5, 3, 3, 7, 7},
			want: 5,
		},
		{
			name: "單獨數字在中間",
			nums: []int{1, 1, 6, 2, 2},
			want: 6,
		},
		{
			name: "單獨數字在末尾",
			nums: []int{3, 3, 9, 9, 8},
			want: 8,
		},
		{
			name: "負數",
			nums: []int{-1, -1, -2},
			want: -2,
		},
		{
			name: "混合正負數",
			nums: []int{-1, 5, -1, 5, 3},
			want: 3,
		},
		{
			name: "零與其他數字",
			nums: []int{0, 1, 0},
			want: 1,
		},
		{
			name: "單獨數字為零",
			nums: []int{1, 1, 0},
			want: 0,
		},
		{
			name: "大陣列",
			nums: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 6, 7, 8, 9, 11, 10, 11},
			want: 22, // 實際上沒有單獨的，但這裡用於測試邏輯
		},
	}

	// 修正最後一個測試案例
	tests[len(tests)-1].nums = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 6, 7, 8, 9, 10, 11}
	tests[len(tests)-1].want = 11

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製切片避免修改原始資料
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := singleNumber(numsCopy)
			if got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleNumberWithMap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "範例1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "範例2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "範例3",
			nums: []int{1},
			want: 1,
		},
		{
			name: "負數",
			nums: []int{-1, -1, -2},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumberWithMap(tt.nums)
			if got != tt.want {
				t.Errorf("singleNumberWithMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleNumberWithSet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "範例1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "範例2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "範例3",
			nums: []int{1},
			want: 1,
		},
		{
			name: "負數",
			nums: []int{-1, -1, -2},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumberWithSet(tt.nums)
			if got != tt.want {
				t.Errorf("singleNumberWithSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基準測試：比較三種解法的效能
func BenchmarkSingleNumber(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{
			name: "小陣列_3個元素",
			nums: []int{2, 2, 1},
		},
		{
			name: "中陣列_11個元素",
			nums: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "大陣列_10001個元素",
			nums: func() []int {
				nums := make([]int, 10001)
				for i := 0; i < 5000; i++ {
					nums[i*2] = i
					nums[i*2+1] = i
				}
				nums[10000] = 9999 // 單獨的數字
				return nums
			}(),
		},
	}

	for _, bm := range benchmarks {
		b.Run("XOR_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				singleNumber(bm.nums)
			}
		})

		b.Run("Map_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				singleNumberWithMap(bm.nums)
			}
		})

		b.Run("Set_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				singleNumberWithSet(bm.nums)
			}
		})
	}
}
