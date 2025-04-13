package MergeIntervals

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
		desc     string // 測試目的描述
	}{
		{
			name:     "基本功能測試 1",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
			desc:     "合併重疊區間 [1,3] 和 [2,6]",
		},
		{
			name:     "基本功能測試 2",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
			desc:     "邊界接觸視為重疊",
		},
		{
			name:     "邊界情況測試 1",
			input:    [][]int{{1, 4}},
			expected: [][]int{{1, 4}},
			desc:     "單個區間不需要合併",
		},
		{
			name:     "邊界情況測試 2",
			input:    [][]int{},
			expected: [][]int{},
			desc:     "空輸入應返回空結果",
		},
		{
			name:     "特殊情況測試",
			input:    [][]int{{1, 4}, {0, 4}},
			expected: [][]int{{0, 4}},
			desc:     "未排序的區間應正確合併",
		},
		{
			name:     "大型輸入測試",
			input:    generateLargeInput(1000),
			expected: [][]int{{0, 999}},
			desc:     "處理大量重疊區間",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeIntervals(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 生成大型輸入
func generateLargeInput(size int) [][]int {
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = []int{i, i + 1}
	}
	return result
}

func BenchmarkMergeIntervals(b *testing.B) {
	// 準備測試資料
	small := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	medium := generateRandomIntervals(100)
	large := generateRandomIntervals(1000)

	// 暴力解法
	b.Run("BruteForce-Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervalsBruteForce(small)
		}
	})

	b.Run("BruteForce-Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervalsBruteForce(medium)
		}
	})

	// 排序優化解法
	b.Run("Sorted-Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervalsSorted(small)
		}
	})

	b.Run("Sorted-Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervalsSorted(medium)
		}
	})

	b.Run("Sorted-Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervalsSorted(large)
		}
	})

	// 最佳解法
	b.Run("Optimal-Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervals(small)
		}
	})

	b.Run("Optimal-Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervals(medium)
		}
	})

	b.Run("Optimal-Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeIntervals(large)
		}
	})
}

// 生成隨機區間
func generateRandomIntervals(size int) [][]int {
	result := make([][]int, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		start := rand.Intn(1000)
		end := start + rand.Intn(20) + 1
		result[i] = []int{start, end}
	}

	return result
}
