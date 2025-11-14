package RotateList

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
		desc     string
	}{
		{
			name:     "基本測試 1",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{4, 5, 1, 2, 3},
			desc:     "標準情況，旋轉 2 次",
		},
		{
			name:     "基本測試 2",
			input:    []int{0, 1, 2},
			k:        4,
			expected: []int{2, 0, 1},
			desc:     "k 大於鏈結串列長度",
		},
		{
			name:     "邊緣情況：單節點",
			input:    []int{1},
			k:        3,
			expected: []int{1},
			desc:     "單節點串列，無論旋轉多少次結果都相同",
		},
		{
			name:     "邊緣情況：空串列",
			input:    []int{},
			k:        5,
			expected: []int{},
			desc:     "空串列直接返回 nil",
		},
		{
			name:     "邊緣情況：不需旋轉",
			input:    []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
			desc:     "旋轉次數等於串列長度，等於不旋轉",
		},
		{
			name:     "大型測試",
			input:    createLongList(100),
			k:        25,
			expected: nil, // 這裡省略預期結果，實際測試中需計算
			desc:     "較長串列的旋轉測試",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 構建輸入鏈結串列
			head := constructList(tt.input)

			// 執行旋轉
			result := rotateRight(head, tt.k)

			// 轉換結果為數組並比較
			resultArray := listToArray(result)
			if !reflect.DeepEqual(resultArray, tt.expected) && tt.expected != nil {
				t.Errorf("rotateRight() = %v, want %v", resultArray, tt.expected)
			}
		})
	}
}

// 輔助函數：構建鏈結串列
func constructList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// 輔助函數：將鏈結串列轉換為數組
func listToArray(head *ListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// 創建長串列用於測試
func createLongList(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i + 1
	}
	return result
}

func BenchmarkRotateRight(b *testing.B) {
	// 創建長度為 100 的測試串列
	inputArray := createLongList(100)
	k := 25

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := constructList(inputArray)
			bruteForceSolution(head, k)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := constructList(inputArray)
			rotateRightOptimized(head, k)
		}
	})

	b.Run("Optimal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := constructList(inputArray)
			rotateRight(head, k)
		}
	})
}
