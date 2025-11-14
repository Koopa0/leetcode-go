package main

import (
	"reflect"
	"testing"
)

// 輔助函數：複製切片
func copySlice(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	return result
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "範例1：零在開頭和中間",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "範例2：只有一個零",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "沒有零",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "零在開頭",
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
		{
			name:     "零在末尾",
			nums:     []int{1, 2, 0, 0},
			expected: []int{1, 2, 0, 0},
		},
		{
			name:     "全是零",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "交替零和非零",
			nums:     []int{1, 0, 2, 0, 3, 0, 4},
			expected: []int{1, 2, 3, 4, 0, 0, 0},
		},
		{
			name:     "單一非零元素",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "空陣列",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "大量零在前",
			nums:     []int{0, 0, 0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0, 0, 0},
		},
		{
			name:     "負數和零",
			nums:     []int{-1, 0, -2, 0, -3},
			expected: []int{-1, -2, -3, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := copySlice(tt.nums)
			moveZeroes(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("moveZeroes() = %v, want %v", numsCopy, tt.expected)
			}
		})
	}
}

func TestMoveZeroesOptimized(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "範例1：零在開頭和中間",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "範例2：只有一個零",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "沒有零",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "交替零和非零",
			nums:     []int{1, 0, 2, 0, 3, 0, 4},
			expected: []int{1, 2, 3, 4, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := copySlice(tt.nums)
			moveZeroesOptimized(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("moveZeroesOptimized() = %v, want %v", numsCopy, tt.expected)
			}
		})
	}
}

func TestMoveZeroesBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "範例1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "範例2",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "交替零和非零",
			nums:     []int{1, 0, 2, 0, 3, 0, 4},
			expected: []int{1, 2, 3, 4, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := copySlice(tt.nums)
			moveZeroesBruteForce(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("moveZeroesBruteForce() = %v, want %v", numsCopy, tt.expected)
			}
		})
	}
}

func TestMoveZeroesSnowball(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "範例1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "範例2",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "零在開頭",
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
		{
			name:     "交替零和非零",
			nums:     []int{1, 0, 2, 0, 3, 0, 4},
			expected: []int{1, 2, 3, 4, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := copySlice(tt.nums)
			moveZeroesSnowball(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("moveZeroesSnowball() = %v, want %v", numsCopy, tt.expected)
			}
		})
	}
}

// 基準測試：比較四種解法的效能
func BenchmarkMoveZeroes(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{
			name: "小陣列_10個元素_50%零",
			nums: []int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0},
		},
		{
			name: "中陣列_100個元素_50%零",
			nums: func() []int {
				nums := make([]int, 100)
				for i := 0; i < 100; i++ {
					if i%2 == 0 {
						nums[i] = i + 1
					} else {
						nums[i] = 0
					}
				}
				return nums
			}(),
		},
		{
			name: "大陣列_10000個元素_20%零",
			nums: func() []int {
				nums := make([]int, 10000)
				for i := 0; i < 10000; i++ {
					if i%5 == 0 {
						nums[i] = 0
					} else {
						nums[i] = i + 1
					}
				}
				return nums
			}(),
		},
		{
			name: "全是零_1000個元素",
			nums: make([]int, 1000),
		},
		{
			name: "沒有零_1000個元素",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := 0; i < 1000; i++ {
					nums[i] = i + 1
				}
				return nums
			}(),
		},
	}

	for _, bm := range benchmarks {
		b.Run("標準雙指標_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := copySlice(bm.nums)
				moveZeroes(numsCopy)
			}
		})

		b.Run("優化版_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := copySlice(bm.nums)
				moveZeroesOptimized(numsCopy)
			}
		})

		b.Run("暴力法_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := copySlice(bm.nums)
				moveZeroesBruteForce(numsCopy)
			}
		})

		b.Run("雪球法_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := copySlice(bm.nums)
				moveZeroesSnowball(numsCopy)
			}
		})
	}
}
