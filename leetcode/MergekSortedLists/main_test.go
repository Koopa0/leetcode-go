package MergekSortedLists

import (
	"reflect"
	"testing"
)

// 構建鏈表輔助函數
func buildList(values []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, val := range values {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return dummy.Next
}

// 將鏈表轉換為數組輔助函數
func listToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeKLists(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		lists    []*ListNode
		expected []int
		desc     string
	}{
		{
			name: "基本測試",
			lists: []*ListNode{
				buildList([]int{1, 4, 5}),
				buildList([]int{1, 3, 4}),
				buildList([]int{2, 6}),
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
			desc:     "測試多個非空鏈表的基本合併",
		},
		{
			name:     "空輸入",
			lists:    []*ListNode{},
			expected: nil,
			desc:     "測試空鏈表數組",
		},
		{
			name: "單個空鏈表",
			lists: []*ListNode{
				nil,
			},
			expected: nil,
			desc:     "測試只有一個空鏈表的情況",
		},
		{
			name: "多個空鏈表",
			lists: []*ListNode{
				nil,
				nil,
				nil,
			},
			expected: nil,
			desc:     "測試多個空鏈表的情況",
		},
		{
			name: "混合空和非空鏈表",
			lists: []*ListNode{
				buildList([]int{1, 3, 5}),
				nil,
				buildList([]int{2, 4, 6}),
				nil,
			},
			expected: []int{1, 2, 3, 4, 5, 6},
			desc:     "測試空和非空鏈表混合的情況",
		},
		{
			name: "單個鏈表",
			lists: []*ListNode{
				buildList([]int{1, 2, 3, 4, 5}),
			},
			expected: []int{1, 2, 3, 4, 5},
			desc:     "測試只有一個非空鏈表的情況",
		},
		{
			name: "完全相同的元素",
			lists: []*ListNode{
				buildList([]int{1, 1, 1}),
				buildList([]int{1, 1, 1}),
			},
			expected: []int{1, 1, 1, 1, 1, 1},
			desc:     "測試鏈表包含完全相同的元素",
		},
		{
			name: "負數值",
			lists: []*ListNode{
				buildList([]int{-3, -1, 1, 3}),
				buildList([]int{-4, -2, 0, 2}),
			},
			expected: []int{-4, -3, -2, -1, 0, 1, 2, 3},
			desc:     "測試包含負數的情況",
		},
		{
			name: "長度不同的鏈表",
			lists: []*ListNode{
				buildList([]int{1}),
				buildList([]int{2, 3, 4, 5, 6}),
				buildList([]int{7, 8}),
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
			desc:     "測試長度差異很大的鏈表",
		},
		{
			name: "極端大值和小值",
			lists: []*ListNode{
				buildList([]int{-10000, -5000}),
				buildList([]int{5000, 10000}),
			},
			expected: []int{-10000, -5000, 5000, 10000},
			desc:     "測試包含極端大小值的情況",
		},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 深拷貝輸入鏈表，避免測試間相互影響
			var listsCopy []*ListNode
			for _, list := range tt.lists {
				if list == nil {
					listsCopy = append(listsCopy, nil)
				} else {
					listsCopy = append(listsCopy, buildList(listToArray(list)))
				}
			}

			result := MergeKLists(listsCopy)
			resultArray := listToArray(result)

			if !reflect.DeepEqual(resultArray, tt.expected) {
				t.Errorf("MergeKLists() = %v, want %v\n%s", resultArray, tt.expected, tt.desc)
			}
		})
	}
}

// 性能測試
func BenchmarkMergeKLists(b *testing.B) {
	// 生成測試數據
	lists := []*ListNode{
		buildList([]int{1, 4, 7, 10, 13, 16, 19}),
		buildList([]int{2, 5, 8, 11, 14, 17, 20}),
		buildList([]int{3, 6, 9, 12, 15, 18, 21}),
	}

	// 執行基準測試
	for i := 0; i < b.N; i++ {
		// 重新構建鏈表，避免修改原始數據
		var listsCopy []*ListNode
		for _, list := range lists {
			listsCopy = append(listsCopy, buildList(listToArray(list)))
		}

		MergeKLists(listsCopy)
	}
}

// 比較不同算法的性能
func BenchmarkCompareAlgorithms(b *testing.B) {
	// 生成測試數據
	lists := []*ListNode{
		buildList([]int{1, 4, 7, 10}),
		buildList([]int{2, 5, 8, 11}),
		buildList([]int{3, 6, 9, 12}),
	}

	b.Run("MinHeap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 重新構建鏈表
			var listsCopy []*ListNode
			for _, list := range lists {
				listsCopy = append(listsCopy, buildList(listToArray(list)))
			}

			MergeKLists(listsCopy)
		}
	})

	b.Run("Linear", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 重新構建鏈表
			var listsCopy []*ListNode
			for _, list := range lists {
				listsCopy = append(listsCopy, buildList(listToArray(list)))
			}

			MergeKListsLinear(listsCopy)
		}
	})

	b.Run("DivideConquer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 重新構建鏈表
			var listsCopy []*ListNode
			for _, list := range lists {
				listsCopy = append(listsCopy, buildList(listToArray(list)))
			}

			MergeKListsDivideConquer(listsCopy)
		}
	})
}
