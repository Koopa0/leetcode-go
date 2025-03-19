package ContainerWithMostWater

import (
	"testing"
)

// TestMaxArea 使用表驅動測試方法測試 maxArea 函數
func TestMaxArea(t *testing.T) {
	// 定義測試用例表
	tests := []struct {
		name     string // 測試用例名稱
		height   []int  // 輸入的高度數組
		expected int    // 期望的輸出結果
	}{
		{
			name:     "LeetCode 範例",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "最小數組長度",
			height:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "所有元素相同",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20, // 5 * (5-1) = 20
		},
		{
			name:     "遞增數組",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6, // min(1,5) * 4 = 4 或 min(2,5) * 3 = 6
		},
		{
			name:     "遞減數組",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6, // min(5,1) * 4 = 4 或 min(5,2) * 3 = 6
		},
		{
			name:     "中間高兩側低",
			height:   []int{1, 3, 6, 7, 9, 4, 2},
			expected: 12, // min(1,2) * 6 = 6 或 min(3,4) * 4 = 12 或 min(6,4) * 2 = 8 等
		},
		{
			name:     "兩側高中間低",
			height:   []int{9, 6, 3, 2, 5, 7, 10},
			expected: 54, // min(9,10) * 6 = 54
		},
		{
			name:     "含零高度",
			height:   []int{0, 7, 0, 8, 0},
			expected: 14, // min(7,8) * 2 = 14
		},
		{
			name:     "最大高度測試",
			height:   []int{10000, 1, 10000},
			expected: 20000, // min(10000,10000) * 2 = 20000
		},
	}

	// 執行測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.expected {
				t.Errorf("maxArea(%v) = %d; 期望結果 %d", tt.height, got, tt.expected)
			}
		})
	}
}

// TestMaxAreaEmpty 測試空數組或長度為1的數組
func TestMaxAreaInvalid(t *testing.T) {
	// 空數組應該返回0
	result := maxArea([]int{})
	if result != 0 {
		t.Errorf("maxArea([]): 期望 0, 得到 %d", result)
	}

	// 長度為1的數組無法形成容器，應該返回0
	result = maxArea([]int{5})
	if result != 0 {
		t.Errorf("maxArea([5]): 期望 0, 得到 %d", result)
	}
}

// TestMaxAreaLargeInput 測試大規模輸入
func TestMaxAreaLargeInput(t *testing.T) {
	if testing.Short() {
		t.Skip("跳過大規模測試")
	}

	// 創建一個大型數組
	const size = 10000
	largeArray := make([]int, size)
	for i := 0; i < size; i++ {
		largeArray[i] = i % 100 // 使用模運算來限制高度值
	}

	// 執行測試並記錄時間
	result := maxArea(largeArray)

	// 檢查結果是否大於0（實際值取決於生成的數組）
	if result <= 0 {
		t.Errorf("大規模測試失敗: 得到的結果為 %d, 應該大於0", result)
	}
}

// BenchmarkMaxArea 基準測試來測量性能
func BenchmarkMaxArea(b *testing.B) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}

// BenchmarkMaxAreaLarge 大規模數據的基準測試
func BenchmarkMaxAreaLarge(b *testing.B) {
	// 創建一個大型數組
	const size = 10000
	largeArray := make([]int, size)
	for i := 0; i < size; i++ {
		largeArray[i] = i % 100
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(largeArray)
	}
}
