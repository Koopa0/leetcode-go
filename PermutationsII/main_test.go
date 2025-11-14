package PermutationsII

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "示例1：包含重複元素",
			nums:     []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name:     "示例2：不包含重複元素",
			nums:     []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name:     "單元素數組",
			nums:     []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "兩個相同元素",
			nums:     []int{1, 1},
			expected: [][]int{{1, 1}},
		},
		{
			name:     "包含負數",
			nums:     []int{-1, 1, -1},
			expected: [][]int{{-1, -1, 1}, {-1, 1, -1}, {1, -1, -1}},
		},
		{
			name: "多個重複元素",
			nums: []int{1, 1, 1, 2},
			expected: [][]int{
				{1, 1, 1, 2}, {1, 1, 2, 1}, {1, 2, 1, 1}, {2, 1, 1, 1},
			},
		},
		{
			name: "測試最大長度邊界",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			// expected 將包含 8! = 40320 個排列，太長無法列出
			expected: nil, // 在測試中使用長度檢查代替具體值比較
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permuteUnique(tt.nums)

			// 對於大型測試用例，只檢查結果長度
			if tt.name == "測試最大長度邊界" {
				// 計算預期的排列數量
				counts := make(map[int]int)
				for _, num := range tt.nums {
					counts[num]++
				}

				total := factorial(len(tt.nums))
				for _, count := range counts {
					total /= factorial(count)
				}

				if len(result) != total {
					t.Errorf("長度不符：預期 %d，實際 %d", total, len(result))
				}
				return
			}

			// 檢查結果包含所有預期的排列
			if !comparePermutations(result, tt.expected) {
				t.Errorf("預期 %v，獲得 %v", tt.expected, result)
			}
		})
	}
}

// 計算階乘
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 比較兩個排列集合是否相同（忽略順序）
func comparePermutations(actual, expected [][]int) bool {
	if len(actual) != len(expected) {
		return false
	}

	// 將排列轉換為字串以便比較
	actualMap := make(map[string]bool)
	for _, perm := range actual {
		key := fmt.Sprintf("%v", perm)
		actualMap[key] = true
	}

	// 檢查每個預期的排列是否在實際結果中
	for _, perm := range expected {
		key := fmt.Sprintf("%v", perm)
		if !actualMap[key] {
			return false
		}
	}

	return true
}
