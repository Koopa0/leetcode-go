package main

import (
	"reflect"
	"testing"
)

// 輔助函數：比較兩棵樹是否相同
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 輔助函數：複製樹
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newNode := &TreeNode{Val: root.Val}
	newNode.Left = copyTree(root.Left)
	newNode.Right = copyTree(root.Right)
	return newNode
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "範例1：完整二元樹",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "範例2：小樹",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name:     "範例3：空樹",
			root:     nil,
			expected: nil,
		},
		{
			name: "單一節點",
			root: &TreeNode{Val: 1},
			expected: &TreeNode{Val: 1},
		},
		{
			name: "只有左子樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
		},
		{
			name: "只有右子樹",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
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
			expected: &TreeNode{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製原始樹，因為 invertTree 會修改原樹
			rootCopy := copyTree(tt.root)
			got := invertTree(rootCopy)
			if !isSameTree(got, tt.expected) {
				t.Errorf("invertTree() failed for %s", tt.name)
			}
		})
	}
}

func TestInvertTreePostOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "範例1：完整二元樹",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name:     "空樹",
			root:     nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCopy := copyTree(tt.root)
			got := invertTreePostOrder(rootCopy)
			if !isSameTree(got, tt.expected) {
				t.Errorf("invertTreePostOrder() failed for %s", tt.name)
			}
		})
	}
}

func TestInvertTreeBFS(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "範例1：完整二元樹",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "範例2：小樹",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name:     "空樹",
			root:     nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCopy := copyTree(tt.root)
			got := invertTreeBFS(rootCopy)
			if !isSameTree(got, tt.expected) {
				t.Errorf("invertTreeBFS() failed for %s", tt.name)
			}
		})
	}
}

func TestInvertTreeIterative(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "範例1：完整二元樹",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name:     "空樹",
			root:     nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCopy := copyTree(tt.root)
			got := invertTreeIterative(rootCopy)
			if !isSameTree(got, tt.expected) {
				t.Errorf("invertTreeIterative() failed for %s", tt.name)
			}
		})
	}
}

// 測試層序遍歷輸出
func TestLevelOrderPrint(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	expected := [][]int{{4}, {2, 7}, {1, 3, 6, 9}}
	got := levelOrderPrint(root)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("levelOrderPrint() = %v, want %v", got, expected)
	}
}

// 基準測試：比較四種解法的效能
func BenchmarkInvertTree(b *testing.B) {
	// 建立測試用的樹
	trees := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "小樹_3層",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
		},
		{
			name: "大樹_7層完全二元樹",
			root: func() *TreeNode {
				root := &TreeNode{Val: 1}
				queue := []*TreeNode{root}
				val := 2

				for len(queue) > 0 && val < 128 {
					node := queue[0]
					queue = queue[1:]

					node.Left = &TreeNode{Val: val}
					val++
					queue = append(queue, node.Left)

					if val < 128 {
						node.Right = &TreeNode{Val: val}
						val++
						queue = append(queue, node.Right)
					}
				}
				return root
			}(),
		},
	}

	for _, tree := range trees {
		b.Run("DFS遞迴_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				treeCopy := copyTree(tree.root)
				invertTree(treeCopy)
			}
		})

		b.Run("DFS後序_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				treeCopy := copyTree(tree.root)
				invertTreePostOrder(treeCopy)
			}
		})

		b.Run("BFS_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				treeCopy := copyTree(tree.root)
				invertTreeBFS(treeCopy)
			}
		})

		b.Run("DFS迭代_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				treeCopy := copyTree(tree.root)
				invertTreeIterative(treeCopy)
			}
		})
	}
}
