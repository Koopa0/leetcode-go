package ReverseLinkedList

import (
	"reflect"
	"testing"
)

// 輔助函數：從陣列建立鏈結串列
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	curr := head

	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}

	return head
}

// 輔助函數：將鏈結串列轉換為陣列
func linkedListToArray(head *ListNode) []int {
	var result []int
	curr := head

	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}

	return result
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "空鏈結串列",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "單節點",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "兩個節點",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "多個節點",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "相同值的節點",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "負數值",
			input:    []int{-5, -3, 0, 3, 5},
			expected: []int{5, 3, 0, -3, -5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (迭代法)", func(t *testing.T) {
			head := createLinkedList(tt.input)
			result := reverseList(head)
			got := linkedListToArray(result)

			// 處理空切片的比較
			if len(got) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList() = %v, want %v", got, tt.expected)
			}
		})

		t.Run(tt.name+" (遞迴法)", func(t *testing.T) {
			head := createLinkedList(tt.input)
			result := reverseListRecursive(head)
			got := linkedListToArray(result)

			// 處理空切片的比較
			if len(got) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseListRecursive() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 測試鏈結串列的連結正確性
func TestReverseListLinkage(t *testing.T) {
	// 建立鏈結串列 1 -> 2 -> 3
	head := createLinkedList([]int{1, 2, 3})

	// 反轉
	reversed := reverseList(head)

	// 檢查連結
	if reversed.Val != 3 {
		t.Errorf("頭節點應為 3，得到 %d", reversed.Val)
	}
	if reversed.Next.Val != 2 {
		t.Errorf("第二個節點應為 2，得到 %d", reversed.Next.Val)
	}
	if reversed.Next.Next.Val != 1 {
		t.Errorf("第三個節點應為 1，得到 %d", reversed.Next.Next.Val)
	}
	if reversed.Next.Next.Next != nil {
		t.Errorf("尾節點的 Next 應為 nil")
	}
}

func BenchmarkReverseList(b *testing.B) {
	// 建立測試用的鏈結串列
	values := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = i
	}

	b.Run("迭代法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := createLinkedList(values)
			reverseList(head)
		}
	})

	b.Run("遞迴法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := createLinkedList(values)
			reverseListRecursive(head)
		}
	})
}

// 大型鏈結串列測試
func TestReverseListLarge(t *testing.T) {
	size := 5000
	values := make([]int, size)
	expected := make([]int, size)

	for i := 0; i < size; i++ {
		values[i] = i
		expected[size-1-i] = i
	}

	t.Run("大型鏈結串列 (迭代法)", func(t *testing.T) {
		head := createLinkedList(values)
		result := reverseList(head)
		got := linkedListToArray(result)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("大型鏈結串列反轉失敗")
		}
	})

	t.Run("大型鏈結串列 (遞迴法)", func(t *testing.T) {
		head := createLinkedList(values)
		result := reverseListRecursive(head)
		got := linkedListToArray(result)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("大型鏈結串列反轉失敗")
		}
	})
}
