package SubsetsII

import (
	"fmt"
	"sort"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
		desc     string // 測試目的描述
	}{
		{
			name:     "基本測試1",
			input:    []int{1, 2, 2},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
			desc:     "測試包含重複元素的基本情況",
		},
		{
			name:     "基本測試2",
			input:    []int{0},
			expected: [][]int{{}, {0}},
			desc:     "測試只有一個元素的情況",
		},
		{
			name:     "空陣列測試",
			input:    []int{},
			expected: [][]int{{}},
			desc:     "測試空陣列應返回一個包含空集合的陣列",
		},
		{
			name:     "多重複元素測試",
			input:    []int{1, 2, 2, 3, 3},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 2, 3, 3}, {1, 2, 3}, {1, 2, 3, 3}, {1, 3}, {1, 3, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 2, 3, 3}, {2, 3}, {2, 3, 3}, {3}, {3, 3}},
			desc:     "測試多個重複元素的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.input)
			// 需要特殊比較函數，因為子集順序可能不同
			if !compareSubsets(got, tt.expected) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 比較兩個子集集合是否相等（忽略順序）
func compareSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 將每個子集轉換為字串以便比較
	aMap := make(map[string]bool)
	bMap := make(map[string]bool)

	for _, subset := range a {
		sort.Ints(subset) // 確保子集內部順序一致
		aMap[fmt.Sprintf("%v", subset)] = true
	}

	for _, subset := range b {
		sort.Ints(subset) // 確保子集內部順序一致
		bMap[fmt.Sprintf("%v", subset)] = true
	}

	// 檢查兩個 map 是否相等
	for k := range aMap {
		if !bMap[k] {
			return false
		}
	}

	for k := range bMap {
		if !aMap[k] {
			return false
		}
	}

	return true
}

func BenchmarkSubsetsWithDup(b *testing.B) {
	// 測試用例
	input := []int{1, 2, 2, 3, 3, 4}

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsWithDupBruteForce(input)
		}
	})

	b.Run("優化回溯解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsWithDupOptimized(input)
		}
	})

	b.Run("最佳迭代解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsWithDup(input)
		}
	})
}
