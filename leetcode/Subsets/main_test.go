package Subsets

import (
	"fmt"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
		desc     string
	}{
		{
			name:     "空陣列",
			input:    []int{},
			expected: [][]int{{}},
			desc:     "測試空陣列",
		},
		{
			name:     "單一元素",
			input:    []int{1},
			expected: [][]int{{}, {1}},
			desc:     "測試只有一個元素的陣列",
		},
		{
			name:     "三個元素",
			input:    []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
			desc:     "測試標準情況",
		},
		{
			name:     "負數元素",
			input:    []int{-1, -2},
			expected: [][]int{{}, {-1}, {-2}, {-1, -2}},
			desc:     "測試包含負數的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsIterative(tt.input)
			// 需要處理結果順序可能不同的情況
			if !subsetsEqual(got, tt.expected) {
				t.Errorf("subsets() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 比較兩個子集集合是否相等（不考慮順序）
func subsetsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 將每個子集轉換為字串以便比較
	mapA := make(map[string]bool)
	for _, subset := range a {
		sort.Ints(subset) // 首先排序每個子集
		key := fmt.Sprintf("%v", subset)
		mapA[key] = true
	}

	// 檢查 b 中的每個子集是否在 mapA 中
	for _, subset := range b {
		sort.Ints(subset)
		key := fmt.Sprintf("%v", subset)
		if !mapA[key] {
			return false
		}
	}

	return true
}

func BenchmarkSolutions(b *testing.B) {
	// 測試用輸入
	input := []int{1, 2, 3, 4, 5}

	b.Run("回溯法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsBacktracking(input)
		}
	})

	b.Run("位元操作", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsBitmask(input)
		}
	})

	b.Run("迭代增量法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsIterative(input)
		}
	})
}
