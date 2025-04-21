package BinaryTreeInorderTraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	tests := []struct {
		name     string
		input    *TreeNode
		expected []int
		desc     string // 測試目的描述
	}{
		{
			name:     "空樹",
			input:    nil,
			expected: []int{},
			desc:     "測試空樹的情況",
		},
		{
			name:     "只有根節點",
			input:    &TreeNode{Val: 1},
			expected: []int{1},
			desc:     "測試只有一個節點的情況",
		},
		{
			name: "標準二元樹",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 3, 2},
			desc:     "測試LeetCode範例1中的標準二元樹",
		},
		{
			name: "平衡二元樹",
			input: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: []int{2, 1, 3},
			desc:     "測試平衡二元樹",
		},
		{
			name: "左偏樹",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{3, 2, 1},
			desc:     "測試左偏樹（類似鏈表）",
		},
		{
			name: "右偏樹",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 2, 3},
			desc:     "測試右偏樹（類似鏈表）",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversalMorris(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	// 建立一個有100個節點的平衡二元樹
	root := buildBalancedTree(100)

	b.Run("遞迴解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalRecursive(root)
		}
	})

	b.Run("迭代解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalIterative(root)
		}
	})

	b.Run("莫里斯遍歷", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalMorris(root)
		}
	})
}

// 輔助函數，用於建立平衡二元樹
func buildBalancedTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}

	root := &TreeNode{Val: n / 2}
	root.Left = buildBalancedTree(n / 2)
	root.Right = buildBalancedTree(n - n/2 - 1)

	return root
}
