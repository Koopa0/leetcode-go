package SameTree

/**
 * 二元樹節點的定義
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遞迴解決方案
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 如果兩個節點都為空，則視為相同
	if p == nil && q == nil {
		return true
	}

	// 如果一個節點為空而另一個不為空，則不同
	if p == nil || q == nil {
		return false
	}

	// 如果節點值不同，則不同
	if p.Val != q.Val {
		return false
	}

	// 遞迴檢查左子樹和右子樹是否都相同
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 迭代解決方案
func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	// 使用堆疊存儲待比較的節點對
	type NodePair struct {
		p, q *TreeNode
	}

	stack := []NodePair{{p, q}}

	for len(stack) > 0 {
		// 彈出堆疊頂部的節點對
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 取出節點對
		node1, node2 := pair.p, pair.q

		// 如果兩個節點都為空，繼續檢查下一對
		if node1 == nil && node2 == nil {
			continue
		}

		// 如果一個節點為空或節點值不同，則不同
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		// 將右子樹和左子樹對推入堆疊
		stack = append(stack, NodePair{node1.Right, node2.Right})
		stack = append(stack, NodePair{node1.Left, node2.Left})
	}

	// 如果所有節點對都相同，則返回 true
	return true
}

// BFS 解決方案
func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	// 使用佇列存儲待比較的節點對
	type NodePair struct {
		p, q *TreeNode
	}

	queue := []NodePair{{p, q}}

	for len(queue) > 0 {
		// 取出佇列首部的節點對
		pair := queue[0]
		queue = queue[1:]

		// 取出節點對
		node1, node2 := pair.p, pair.q

		// 如果兩個節點都為空，繼續檢查下一對
		if node1 == nil && node2 == nil {
			continue
		}

		// 如果一個節點為空或節點值不同，則不同
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		// 將左子樹和右子樹對加入佇列
		queue = append(queue, NodePair{node1.Left, node2.Left})
		queue = append(queue, NodePair{node1.Right, node2.Right})
	}

	// 如果所有節點對都相同，則返回 true
	return true
}
