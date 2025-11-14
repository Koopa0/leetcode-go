package SymmetricTree

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
		desc     string
	}{
		{
			name: "對稱樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			expected: true,
			desc:     "標準對稱樹測試",
		},
		{
			name: "非對稱樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: false,
			desc:     "標準非對稱樹測試",
		},
		{
			name:     "空樹",
			root:     nil,
			expected: true,
			desc:     "空樹應被視為對稱",
		},
		{
			name:     "只有根節點",
			root:     &TreeNode{Val: 1},
			expected: true,
			desc:     "只有一個節點的樹應被視為對稱",
		},
		// 可以添加更多測試案例
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.root)
			if got != tt.expected {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.expected)
			}

			// 測試迭代版本
			got = isSymmetricIterative(tt.root)
			if got != tt.expected {
				t.Errorf("isSymmetricIterative() = %v, want %v", got, tt.expected)
			}

			// 測試層序遍歷版本
			got = isSymmetricByLevel(tt.root)
			if got != tt.expected {
				t.Errorf("isSymmetricByLevel() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSymmetricTree(b *testing.B) {
	// 建立一個大型對稱樹用於測試
	root := createLargeSymmetricTree(10) // 深度為10的對稱樹

	b.Run("遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSymmetric(root)
		}
	})

	b.Run("迭代", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSymmetricIterative(root)
		}
	})

	b.Run("層序", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSymmetricByLevel(root)
		}
	})
}

// 輔助函數：建立大型對稱樹
func createLargeSymmetricTree(depth int) *TreeNode {
	if depth == 0 {
		return nil
	}

	root := &TreeNode{Val: 1}
	buildSymmetricTree(root, depth-1)
	return root
}

func buildSymmetricTree(node *TreeNode, depthRemaining int) {
	if depthRemaining == 0 {
		return
	}

	node.Left = &TreeNode{Val: depthRemaining}
	node.Right = &TreeNode{Val: depthRemaining}

	buildSymmetricTree(node.Left, depthRemaining-1)
	buildSymmetricTree(node.Right, depthRemaining-1)
}
