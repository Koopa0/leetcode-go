package MergeSortedArray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
		desc     string // 測試目的描述
	}{
		{
			name:     "基本測試",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
			desc:     "測試兩個非空有序陣列的合併",
		},
		{
			name:     "nums2 為空",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
			desc:     "測試 nums2 為空的情況",
		},
		{
			name:     "nums1 為空",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
			desc:     "測試 nums1 為空的情況",
		},
		{
			name:     "包含相同元素",
			nums1:    []int{1, 2, 2, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 1, 2, 2, 2, 3},
			desc:     "測試包含相同元素的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製原始 nums1 用於測試
			originalNums1 := make([]int, len(tt.nums1))
			copy(originalNums1, tt.nums1)

			merge(tt.nums1, tt.m, tt.nums2, tt.n)

			if !reflect.DeepEqual(tt.nums1, tt.expected) {
				t.Errorf("merge() = %v, want %v, original nums1: %v", tt.nums1, tt.expected, originalNums1)
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	// 創建足夠大的測試資料
	createTestArrays := func(size int) ([]int, int, []int, int) {
		nums1 := make([]int, size*2)
		nums2 := make([]int, size)

		// 填充測試資料
		for i := 0; i < size; i++ {
			nums1[i] = i * 2
			nums2[i] = i*2 + 1
		}

		return nums1, size, nums2, size
	}

	// 小型資料集測試
	b.Run("SmallDataset", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums1, m, nums2, n := createTestArrays(10)
			merge(nums1, m, nums2, n)
		}
	})

	// 中型資料集測試
	b.Run("MediumDataset", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums1, m, nums2, n := createTestArrays(100)
			merge(nums1, m, nums2, n)
		}
	})

	// 大型資料集測試
	b.Run("LargeDataset", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums1, m, nums2, n := createTestArrays(1000)
			merge(nums1, m, nums2, n)
		}
	})
}
