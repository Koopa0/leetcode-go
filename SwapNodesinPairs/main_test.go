package SwapNodesinPairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	// 定義測試用例表
	tests := []struct {
		name     string // 測試用例名稱
		input    []int  // 輸入鏈表的值
		expected []int  // 期望輸出鏈表的值
		desc     string // 測試用例描述
	}{
		{
			name:     "正常情況-偶數節點",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
			desc:     "測試包含偶數個節點的鏈表",
		},
		{
			name:     "正常情況-奇數節點",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 1, 4, 3, 5},
			desc:     "測試包含奇數個節點的鏈表",
		},
		{
			name:     "邊界情況-空鏈表",
			input:    []int{},
			expected: []int{},
			desc:     "測試空鏈表",
		},
		{
			name:     "邊界情況-單節點",
			input:    []int{1},
			expected: []int{1},
			desc:     "測試只有一個節點的鏈表",
		},
		{
			name:     "邊界情況-兩個節點",
			input:    []int{1, 2},
			expected: []int{2, 1},
			desc:     "測試只有兩個節點的鏈表",
		},
		{
			name:     "極限情況-最大節點數",
			input:    make([]int, 100), // 創建包含100個節點的鏈表
			expected: nil,              // 這裡需要根據預期結果動態生成
			desc:     "測試最大節點數的鏈表",
		},
	}

	// 為極限情況生成預期結果
	input := make([]int, 100)
	expected := make([]int, 100)
	for i := 0; i < 100; i++ {
		input[i] = i
		if i%2 == 0 && i+1 < 100 {
			expected[i] = i + 1
			expected[i+1] = i
		}
	}
	tests[5].input = input
	tests[5].expected = expected

	// 執行測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 創建輸入鏈表
			head := createLinkedList(tt.input)

			// 執行函數
			result := swapPairs(head)

			// 將結果轉換回切片進行比較
			resultSlice := linkedListToSlice(result)

			if len(tt.expected) == 0 && len(resultSlice) == 0 {
				// 兩者都是空切片，測試通過
				return
			}

			// 驗證結果
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("swapPairs() = %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}

// 輔助函數：創建鏈表
func createLinkedList(values []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// 輔助函數：將鏈表轉換為切片
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}
