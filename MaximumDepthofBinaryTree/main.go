package MaximumDepthofBinaryTree

// TreeNode 定義二元樹節點
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 使用 DFS 遞迴法計算二元樹的最大深度
// 時間複雜度: O(n)，其中 n 是節點數
// 空間複雜度: O(h)，其中 h 是樹的高度（遞迴呼叫堆疊）
func maxDepth(root *TreeNode) int {
	// 基礎情況：空節點深度為 0
	if root == nil {
		return 0
	}

	// 遞迴計算左子樹和右子樹的深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 返回較大的深度 + 1（當前節點）
	return 1 + max(leftDepth, rightDepth)
}

// maxDepthBFS 使用 BFS 層序遍歷計算二元樹的最大深度
// 時間複雜度: O(n)，需要訪問每個節點一次
// 空間複雜度: O(w)，其中 w 是樹的最大寬度
func maxDepthBFS(root *TreeNode) int {
	// 處理空樹
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}

	// 層序遍歷
	for len(queue) > 0 {
		// 當前層的節點數量
		levelSize := len(queue)

		// 處理當前層的所有節點
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 將子節點加入佇列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 處理完一層，深度加 1
		depth++
	}

	return depth
}

// maxDepthDFSIterative 使用 DFS 迭代法計算二元樹的最大深度
// 時間複雜度: O(n)，需要訪問每個節點一次
// 空間複雜度: O(h)，堆疊的最大深度
func maxDepthDFSIterative(root *TreeNode) int {
	// 處理空樹
	if root == nil {
		return 0
	}

	// 定義堆疊元素：(節點, 當前深度)
	type StackItem struct {
		node  *TreeNode
		depth int
	}

	maxD := 0
	stack := []StackItem{{root, 1}}

	// DFS 遍歷
	for len(stack) > 0 {
		// 彈出堆疊頂部元素
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		currentDepth := item.depth

		// 更新最大深度
		if currentDepth > maxD {
			maxD = currentDepth
		}

		// 將子節點推入堆疊，深度 +1
		if node.Left != nil {
			stack = append(stack, StackItem{node.Left, currentDepth + 1})
		}
		if node.Right != nil {
			stack = append(stack, StackItem{node.Right, currentDepth + 1})
		}
	}

	return maxD
}

// max 返回兩個整數中的較大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
