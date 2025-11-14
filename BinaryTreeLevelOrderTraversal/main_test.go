package main

import (
	"reflect"
	"testing"
)

// 輔助函數：比較兩個二維陣列是否相等
func equal2DSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "範例1：標準二元樹",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "範例2：單一節點",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "範例3：空樹",
			root: nil,
			want: [][]int{},
		},
		{
			name: "只有左子樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "只有右子樹",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{Val: 3},
				},
			},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "完全二元樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "不平衡樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 4},
			},
			want: [][]int{{1}, {2, 4}, {3}},
		},
		{
			name: "兩層樹",
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 8},
			},
			want: [][]int{{5}, {3, 8}},
		},
		{
			name: "深度較大的樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			want: [][]int{{1}, {2}, {3}, {4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)
			if !equal2DSlice(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrderDFS(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "範例1：標準二元樹",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "範例2：單一節點",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "範例3：空樹",
			root: nil,
			want: [][]int{},
		},
		{
			name: "完全二元樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrderDFS(tt.root)
			if !equal2DSlice(got, tt.want) {
				t.Errorf("levelOrderDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrderIterative(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "範例1：標準二元樹",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "範例2：單一節點",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "範例3：空樹",
			root: nil,
			want: [][]int{},
		},
		{
			name: "完全二元樹",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrderIterative(tt.root)
			if !equal2DSlice(got, tt.want) {
				t.Errorf("levelOrderIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基準測試：比較三種解法的效能
func BenchmarkLevelOrder(b *testing.B) {
	// 建立測試用的樹
	trees := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "小樹_3層",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "完全二元樹_3層",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "大樹_7層完全二元樹",
			root: func() *TreeNode {
				// 建立一個 7 層的完全二元樹
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
		b.Run("BFS_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				levelOrder(tree.root)
			}
		})

		b.Run("DFS_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				levelOrderDFS(tree.root)
			}
		})

		b.Run("迭代_"+tree.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				levelOrderIterative(tree.root)
			}
		})
	}
}
