package SearchInsertPosition

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// 基本功能測試
		{"target in array", []int{1, 3, 5, 6}, 5, 2},
		{"target not in array", []int{1, 3, 5, 6}, 2, 1},

		// 邊界情況
		{"empty array", []int{}, 5, 0},
		{"single element, target smaller", []int{5}, 3, 0},
		{"single element, target equal", []int{5}, 5, 0},
		{"single element, target larger", []int{5}, 7, 1},

		// 特殊情況
		{"target smaller than all elements", []int{1, 3, 5, 6}, 0, 0},
		{"target larger than all elements", []int{1, 3, 5, 6}, 7, 4},
		{"duplicate target (not in problem, but good to test)", []int{1, 3, 5, 5, 6}, 5, 2}, // 找到第一個匹配

		// 性能測試案例
		{"large array", generateLargeArray(1000), 999, 999}, // 生成一個大數組進行測試
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("searchInsert(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

// 生成一個包含 n 個元素的大數組
func generateLargeArray(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}
