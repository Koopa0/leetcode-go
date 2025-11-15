package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   bool
	}{
		{
			name:   "範例1：回文_偶數長度",
			values: []int{1, 2, 2, 1},
			want:   true,
		},
		{
			name:   "範例2：不是回文",
			values: []int{1, 2},
			want:   false,
		},
		{
			name:   "範例3：單一節點",
			values: []int{1},
			want:   true,
		},
		{
			name:   "空鏈結串列",
			values: []int{},
			want:   true,
		},
		{
			name:   "兩個相同節點",
			values: []int{1, 1},
			want:   true,
		},
		{
			name:   "兩個不同節點",
			values: []int{1, 2},
			want:   false,
		},
		{
			name:   "回文_奇數長度",
			values: []int{1, 2, 3, 2, 1},
			want:   true,
		},
		{
			name:   "不是回文_奇數長度",
			values: []int{1, 2, 3, 4, 5},
			want:   false,
		},
		{
			name:   "所有相同值",
			values: []int{1, 1, 1, 1},
			want:   true,
		},
		{
			name:   "長回文",
			values: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want:   true,
		},
		{
			name:   "長回文_偶數",
			values: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want:   true,
		},
		{
			name:   "幾乎回文_最後不同",
			values: []int{1, 2, 3, 2, 2},
			want:   false,
		},
		{
			name:   "負數回文",
			values: []int{-1, -2, -2, -1},
			want:   true,
		},
		{
			name:   "混合正負數回文",
			values: []int{-1, 0, 1, 0, -1},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			got := isPalindrome(head)
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeArray(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   bool
	}{
		{
			name:   "範例1：回文",
			values: []int{1, 2, 2, 1},
			want:   true,
		},
		{
			name:   "範例2：不是回文",
			values: []int{1, 2},
			want:   false,
		},
		{
			name:   "單一節點",
			values: []int{1},
			want:   true,
		},
		{
			name:   "回文_奇數長度",
			values: []int{1, 2, 3, 2, 1},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			got := isPalindromeArray(head)
			if got != tt.want {
				t.Errorf("isPalindromeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeRecursive(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   bool
	}{
		{
			name:   "範例1：回文",
			values: []int{1, 2, 2, 1},
			want:   true,
		},
		{
			name:   "範例2：不是回文",
			values: []int{1, 2},
			want:   false,
		},
		{
			name:   "回文_奇數長度",
			values: []int{1, 2, 3, 2, 1},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			got := isPalindromeRecursive(head)
			if got != tt.want {
				t.Errorf("isPalindromeRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基準測試：比較三種解法的效能
func BenchmarkIsPalindrome(b *testing.B) {
	benchmarks := []struct {
		name   string
		values []int
	}{
		{
			name:   "短鏈結串列_4個節點_回文",
			values: []int{1, 2, 2, 1},
		},
		{
			name:   "中鏈結串列_10個節點_回文",
			values: []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
		},
		{
			name:   "長鏈結串列_1000個節點_回文",
			values: func() []int {
				values := make([]int, 1000)
				for i := 0; i < 500; i++ {
					values[i] = i
					values[999-i] = i
				}
				return values
			}(),
		},
		{
			name:   "不是回文_提前退出",
			values: []int{1, 2, 3, 4},
		},
		{
			name:   "單一節點",
			values: []int{1},
		},
	}

	for _, bm := range benchmarks {
		b.Run("快慢指標_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := createList(bm.values)
				isPalindrome(head)
			}
		})

		b.Run("陣列法_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := createList(bm.values)
				isPalindromeArray(head)
			}
		})

		b.Run("遞迴法_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := createList(bm.values)
				isPalindromeRecursive(head)
			}
		})
	}
}
