package ValidateBinarySearchTree

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name     string
		tree     *TreeNode
		expected bool
		desc     string
	}{
		{
			name: "範例1：有效BST",
			tree: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
			desc:     "基本的有效BST",
		},
		{
			name: "範例2：無效BST",
			tree: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			expected: false,
			desc:     "右子樹中有小於根節點的值",
		},
		{
			name:     "空樹",
			tree:     nil,
			expected: true,
			desc:     "空樹是有效的BST",
		},
		{
			name:     "單節點樹",
			tree:     &TreeNode{Val: 1},
			expected: true,
			desc:     "單節點樹是有效的BST",
		},
		{
			name: "具有重複值的樹",
			tree: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
			desc:     "有重複值，不是有效的BST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.tree)
			if got != tt.expected {
				t.Errorf("isValidBST() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	// 創建一個中等大小的BST用於基準測試
	tree := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   12,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 14},
		},
	}

	b.Run("遞迴與值域限制", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBSTWithRange(tree)
		}
	})

	b.Run("中序遍歷法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBST(tree)
		}
	})
}
