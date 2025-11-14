package MaximumDepthofBinaryTree

import (
	"testing"
)

// 輔助函數：從陣列建立二元樹（層序遍歷）
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// 處理左子節點
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 處理右子節點
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected int
	}{
		{
			name:     "空樹",
			input:    []interface{}{},
			expected: 0,
		},
		{
			name:     "單節點",
			input:    []interface{}{1},
			expected: 1,
		},
		{
			name:     "範例 1 - 平衡樹",
			input:    []interface{}{3, 9, 20, nil, nil, 15, 7},
			expected: 3,
		},
		{
			name:     "範例 2 - 偏右樹",
			input:    []interface{}{1, nil, 2},
			expected: 2,
		},
		{
			name:     "完全二元樹",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: 3,
		},
		{
			name:     "偏左的不平衡樹",
			input:    []interface{}{1, 2, nil, 3, nil, 4},
			expected: 4,
		},
		{
			name:     "偏右的不平衡樹",
			input:    []interface{}{1, nil, 2, nil, 3, nil, 4},
			expected: 4,
		},
		{
			name:     "兩層完全樹",
			input:    []interface{}{1, 2, 3},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (DFS 遞迴)", func(t *testing.T) {
			root := createTree(tt.input)
			got := maxDepth(root)
			if got != tt.expected {
				t.Errorf("maxDepth() = %v, want %v", got, tt.expected)
			}
		})

		t.Run(tt.name+" (BFS)", func(t *testing.T) {
			root := createTree(tt.input)
			got := maxDepthBFS(root)
			if got != tt.expected {
				t.Errorf("maxDepthBFS() = %v, want %v", got, tt.expected)
			}
		})

		t.Run(tt.name+" (DFS 迭代)", func(t *testing.T) {
			root := createTree(tt.input)
			got := maxDepthDFSIterative(root)
			if got != tt.expected {
				t.Errorf("maxDepthDFSIterative() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 測試深度計算的正確性
func TestMaxDepthDeepTree(t *testing.T) {
	// 建立一個深度為 10 的偏左樹
	// 1 -> 2 -> 3 -> ... -> 10
	root := &TreeNode{Val: 1}
	current := root
	for i := 2; i <= 10; i++ {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}

	expected := 10

	t.Run("深度為 10 的樹 (DFS 遞迴)", func(t *testing.T) {
		got := maxDepth(root)
		if got != expected {
			t.Errorf("maxDepth() = %v, want %v", got, expected)
		}
	})

	t.Run("深度為 10 的樹 (BFS)", func(t *testing.T) {
		got := maxDepthBFS(root)
		if got != expected {
			t.Errorf("maxDepthBFS() = %v, want %v", got, expected)
		}
	})

	t.Run("深度為 10 的樹 (DFS 迭代)", func(t *testing.T) {
		got := maxDepthDFSIterative(root)
		if got != expected {
			t.Errorf("maxDepthDFSIterative() = %v, want %v", got, expected)
		}
	})
}

func BenchmarkMaxDepth(b *testing.B) {
	// 建立測試用的完全二元樹（深度 10）
	values := make([]interface{}, 1023) // 2^10 - 1 個節點
	for i := 0; i < 1023; i++ {
		values[i] = i
	}
	root := createTree(values)

	b.Run("DFS 遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxDepth(root)
		}
	})

	b.Run("BFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxDepthBFS(root)
		}
	})

	b.Run("DFS 迭代", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxDepthDFSIterative(root)
		}
	})
}

// 測試極端情況：非常不平衡的樹
func TestMaxDepthUnbalanced(b *testing.T) {
	// 建立一個深度為 100 的偏左樹
	root := &TreeNode{Val: 1}
	current := root
	for i := 2; i <= 100; i++ {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}

	expected := 100

	tests := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{"DFS 遞迴", maxDepth},
		{"BFS", maxDepthBFS},
		{"DFS 迭代", maxDepthDFSIterative},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(t *testing.T) {
			got := tt.fn(root)
			if got != expected {
				t.Errorf("%s() = %v, want %v", tt.name, got, expected)
			}
		})
	}
}
