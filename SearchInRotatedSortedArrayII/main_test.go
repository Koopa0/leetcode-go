package SearchInRotatedSortedArrayII

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
		desc     string
	}{
		{
			name:     "標準旋轉陣列找到目標",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
			desc:     "測試目標值存在於旋轉陣列中",
		},
		{
			name:     "標準旋轉陣列未找到目標",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
			desc:     "測試目標值不存在於旋轉陣列中",
		},
		{
			name:     "全部相同元素陣列",
			nums:     []int{1, 1, 1, 1, 1, 1},
			target:   1,
			expected: true,
			desc:     "測試全部元素相同的邊緣情況",
		},
		{
			name:     "全部相同元素陣列未找到目標",
			nums:     []int{1, 1, 1, 1, 1, 1},
			target:   2,
			expected: false,
			desc:     "測試全部元素相同且目標不存在",
		},
		{
			name:     "只有一個元素的陣列",
			nums:     []int{5},
			target:   5,
			expected: true,
			desc:     "測試單元素陣列",
		},
		{
			name:     "沒有旋轉的已排序陣列",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
			desc:     "測試標準二分搜尋的情況",
		},
		{
			name:     "包含大量重複的旋轉陣列",
			nums:     []int{2, 2, 2, 3, 2, 2, 2},
			target:   3,
			expected: true,
			desc:     "測試最壞情況：大量重複元素中的單一不同值",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("search() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 準備測試資料
	nums := []int{4, 5, 6, 7, 0, 1, 2, 3, 3, 3, 3}
	target := 0

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bruteForceSearch(nums, target)
		}
	})

	b.Run("優化解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			optimizedSearch(nums, target)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search(nums, target)
		}
	})
}
