package AddTwoNumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// 輔助函數：將數組轉換為鏈表
	createList := func(nums []int) *ListNode {
		dummy := &ListNode{}
		curr := dummy
		for _, num := range nums {
			curr.Next = &ListNode{Val: num}
			curr = curr.Next
		}
		return dummy.Next
	}

	// 輔助函數：將鏈表轉換為數組（用於比較）
	listToSlice := func(head *ListNode) []int {
		var result []int
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	// 測試用例表
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
		desc     string
	}{
		{
			name:     "基本測試用例1",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
			desc:     "測試基本加法：342 + 465 = 807",
		},
		{
			name:     "基本測試用例2",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
			desc:     "測試兩個0相加",
		},
		{
			name:     "長度不等的鏈表",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
			desc:     "測試長度不等的鏈表相加，並有額外進位",
		},
		{
			name:     "一個數較大一個較小",
			l1:       []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			l2:       []int{5, 6, 4},
			expected: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0},
			desc:     "測試一個大數與小數相加",
		},
		{
			name:     "有進位的邊界情況",
			l1:       []int{9},
			l2:       []int{1},
			expected: []int{0, 1},
			desc:     "測試只有進位的情況",
		},
		{
			name:     "極端情況：長鏈表",
			l1:       make([]int, 100), // 100個0
			l2:       []int{1},
			expected: append([]int{1}, make([]int, 99)...),
			desc:     "測試100位與1位相加",
		},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.l1)
			l2 := createList(tt.l2)
			result := addTwoNumbers(l1, l2)
			resultSlice := listToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("加法結果不正確，期望 %v，得到 %v，測試描述：%s",
					tt.expected, resultSlice, tt.desc)
			}
		})
	}
}
