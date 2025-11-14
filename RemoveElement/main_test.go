package RemoveElement

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	// 定義測試用例表格
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		result   []int
	}{
		{
			name:     "示例1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			result:   []int{2, 2},
		},
		{
			name:     "示例2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			result:   []int{0, 1, 3, 0, 4},
		},
		{
			name:     "空數組",
			nums:     []int{},
			val:      1,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "全部元素等於val",
			nums:     []int{2, 2, 2, 2},
			val:      2,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "沒有元素等於val",
			nums:     []int{1, 3, 5, 7},
			val:      2,
			expected: 4,
			result:   []int{1, 3, 5, 7},
		},
		{
			name:     "單一元素等於val",
			nums:     []int{1},
			val:      1,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "單一元素不等於val",
			nums:     []int{1},
			val:      2,
			expected: 1,
			result:   []int{1},
		},
	}

	// 遍歷測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 創建原數組的副本，因為函數會修改數組
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			// 調用被測函數
			got := removeElement(numsCopy, tt.val)

			// 檢查返回值
			if got != tt.expected {
				t.Errorf("removeElement() 返回值 = %v, 期望 %v", got, tt.expected)
			}

			// 檢查數組前 k 個元素（需要先排序）
			sort.Ints(numsCopy[:got])
			resultCopy := make([]int, len(tt.result))
			copy(resultCopy, tt.result)
			sort.Ints(resultCopy)

			if !reflect.DeepEqual(numsCopy[:got], resultCopy) {
				t.Errorf("removeElement() 處理後數組前 %d 個元素 = %v, 期望 %v", got, numsCopy[:got], resultCopy)
			}
		})
	}
}
