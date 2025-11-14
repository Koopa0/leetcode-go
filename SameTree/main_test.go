package SameTree

import "testing"

func TestIsSameTree(t *testing.T) {
	// 輔助函數：從陣列構建樹
	buildTree := func(nums []interface{}) *TreeNode {
		if len(nums) == 0 || nums[0] == nil {
			return nil
		}

		root := &TreeNode{Val: nums[0].(int)}
		queue := []*TreeNode{root}
		i := 1

		for len(queue) > 0 && i < len(nums) {
			node := queue[0]
			queue = queue[1:]

			if i < len(nums) && nums[i] != nil {
				node.Left = &TreeNode{Val: nums[i].(int)}
				queue = append(queue, node.Left)
			}
			i++

			if i < len(nums) && nums[i] != nil {
				node.Right = &TreeNode{Val: nums[i].(int)}
				queue = append(queue, node.Right)
			}
			i++
		}

		return root
	}

	tests := []struct {
		name     string
		p        []interface{}
		q        []interface{}
		expected bool
		desc     string
	}{
		{
			name:     "相同的樹",
			p:        []interface{}{1, 2, 3},
			q:        []interface{}{1, 2, 3},
			expected: true,
			desc:     "兩棵樹結構和值都相同",
		},
		{
			name:     "不同結構的樹",
			p:        []interface{}{1, 2},
			q:        []interface{}{1, nil, 2},
			expected: false,
			desc:     "節點值相同但結構不同",
		},
		{
			name:     "不同值的樹",
			p:        []interface{}{1, 2, 1},
			q:        []interface{}{1, 1, 2},
			expected: false,
			desc:     "結構相同但節點值不同",
		},
		{
			name:     "空樹",
			p:        []interface{}{},
			q:        []interface{}{},
			expected: true,
			desc:     "兩個空樹應該相同",
		},
		{
			name:     "一棵空樹",
			p:        []interface{}{1},
			q:        []interface{}{},
			expected: false,
			desc:     "一棵樹為空而另一棵不為空",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pTree := buildTree(tt.p)
			qTree := buildTree(tt.q)
			got := isSameTreeBFS(pTree, qTree)
			if got != tt.expected {
				t.Errorf("isSameTree() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	// 建立測試樹
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	b.Run("遞迴", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTree(p, q)
		}
	})

	b.Run("迭代", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTreeIterative(p, q)
		}
	})

	b.Run("BFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTreeBFS(p, q)
		}
	})
}
