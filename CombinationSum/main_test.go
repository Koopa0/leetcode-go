package CombinationSum

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	// 定義測試用例表
	testCases := []struct {
		name        string
		candidates  []int
		target      int
		expected    [][]int
		description string
	}{
		{
			name:        "Example 1",
			candidates:  []int{2, 3, 6, 7},
			target:      7,
			expected:    [][]int{{2, 2, 3}, {7}},
			description: "Basic example with multiple solutions",
		},
		{
			name:        "Example 2",
			candidates:  []int{2, 3, 5},
			target:      8,
			expected:    [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
			description: "Example with multiple different combinations",
		},
		{
			name:        "Example 3",
			candidates:  []int{2},
			target:      1,
			expected:    [][]int{},
			description: "No solution possible",
		},
		{
			name:        "Empty candidates",
			candidates:  []int{},
			target:      5,
			expected:    [][]int{},
			description: "Edge case: empty candidates array",
		},
		{
			name:        "Zero target",
			candidates:  []int{1, 2, 3},
			target:      0,
			expected:    [][]int{{}},
			description: "Edge case: target is zero (one empty combination)",
		},
	}

	// 執行測試用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := combinationSum(tc.candidates, tc.target)

			// 比較結果
			if !compareNestedSlices(result, tc.expected) {
				t.Errorf("Test case %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

// 比較兩個嵌套切片是否相等（因順序可能不同需要自定義比較函數）
func compareNestedSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 將每個組合轉換為字串，用於比較
	mapA := make(map[string]bool)
	for _, combo := range a {
		sort.Ints(combo) // 確保順序一致
		key := sliceToString(combo)
		mapA[key] = true
	}

	// 檢查b中的每個組合是否都在mapA中
	for _, combo := range b {
		sort.Ints(combo)
		key := sliceToString(combo)
		if !mapA[key] {
			return false
		}
	}

	return true
}

// 將整數切片轉換為字串
func sliceToString(slice []int) string {
	var builder strings.Builder
	for _, num := range slice {
		builder.WriteString(fmt.Sprintf("%d,", num))
	}
	return builder.String()
}
