package RecoverBinarySearchTree

import (
	"fmt"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		name     string
		input    *TreeNode
		expected *TreeNode
		desc     string
	}{
		{
			name: "Example 1",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 2},
				},
			},
			expected: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
			},
			desc: "測試根節點與其右子節點的右子節點被交換的情況",
		},
		{
			name: "Example 2",
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 2},
				},
			},
			expected: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 3},
				},
			},
			desc: "測試根節點與其右子節點的左子節點被交換的情況",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製輸入樹以便進行比較
			inputCopy := copyTree(tt.input)

			// 執行恢復操作
			morrisRecoverTree(tt.input)

			// 比較結果
			if !isSameTree(tt.input, tt.expected) {
				t.Errorf("recoverTree() = %v, want %v", treeToString(tt.input), treeToString(tt.expected))
				t.Errorf("Original input: %v", treeToString(inputCopy))
			}
		})
	}
}

// 輔助函數：複製樹
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

// 輔助函數：比較兩棵樹
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 輔助函數：將樹轉換為字串（用於錯誤信息）
func treeToString(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return fmt.Sprintf("{%d, %s, %s}", root.Val, treeToString(root.Left), treeToString(root.Right))
}

func BenchmarkSolutions(b *testing.B) {
	// 建立一個具有一定規模的測試樹
	root := generateLargeTestTree(1000)

	// 暴力解法
	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// 創建樹的副本
			treeCopy := copyTree(root)
			bruteForceRecoverTree(treeCopy)
		}
	})

	// 優化解法
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			treeCopy := copyTree(root)
			optimizedRecoverTree(treeCopy)
		}
	})

	// 最佳解法 (Morris)
	b.Run("Morris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			treeCopy := copyTree(root)
			morrisRecoverTree(treeCopy)
		}
	})
}

// 生成大型測試樹
func generateLargeTestTree(size int) *TreeNode {
	// 生成一個平衡的 BST
	nodes := make([]int, size)
	for i := 0; i < size; i++ {
		nodes[i] = i
	}

	// 交換兩個節點來模擬錯誤
	idx1, idx2 := size/3, 2*size/3
	nodes[idx1], nodes[idx2] = nodes[idx2], nodes[idx1]

	// 從有序數組構建 BST
	return buildBST(nodes, 0, len(nodes)-1)
}

// 從有序數組構建 BST
func buildBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = buildBST(nums, start, mid-1)
	root.Right = buildBST(nums, mid+1, end)

	return root
}
