package BinaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遞迴解法
func inorderTraversalRecursive(root *TreeNode) []int {
	// 結果陣列
	var result []int

	// 定義輔助遞迴函數
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		// 基本情況：如果節點為空，直接返回
		if node == nil {
			return
		}

		// 遞迴遍歷左子樹
		inorder(node.Left)

		// 訪問當前節點（加入結果陣列）
		result = append(result, node.Val)

		// 遞迴遍歷右子樹
		inorder(node.Right)
	}

	// 調用輔助函數
	inorder(root)

	return result
}

// 遞迴解法
func inorderTraversal(root *TreeNode) []int {
	// 結果陣列
	var result []int

	// 定義輔助遞迴函數
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		// 基本情況：如果節點為空，直接返回
		if node == nil {
			return
		}

		// 遞迴遍歷左子樹
		inorder(node.Left)

		// 訪問當前節點（加入結果陣列）
		result = append(result, node.Val)

		// 遞迴遍歷右子樹
		inorder(node.Right)
	}

	// 調用輔助函數
	inorder(root)

	return result
}

// 迭代解法（使用堆疊）
func inorderTraversalIterative(root *TreeNode) []int {
	// 結果陣列
	var result []int

	// 使用切片作為堆疊
	var stack []*TreeNode

	// 當前正在處理的節點
	current := root

	// 當堆疊非空或當前節點非空時繼續循環
	for current != nil || len(stack) > 0 {
		// 盡可能向左走，將所有左側節點壓入堆疊
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 當無法繼續向左時，彈出堆疊頂部的節點
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 訪問當前節點
		result = append(result, current.Val)

		// 移動到右子節點
		current = current.Right
	}

	return result
}

// 莫里斯遍歷
func inorderTraversalMorris(root *TreeNode) []int {
	// 結果陣列
	result := []int{}

	// 當前正在處理的節點
	current := root

	// 當當前節點非空時繼續循環
	for current != nil {
		// 如果當前節點沒有左子節點
		if current.Left == nil {
			// 訪問當前節點
			result = append(result, current.Val)
			// 移動到右子節點
			current = current.Right
		} else {
			// 找到當前節點在中序遍歷中的前驅（左子樹的最右節點）
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			// 如果前驅的右指針為空，建立臨時鏈接
			if predecessor.Right == nil {
				predecessor.Right = current
				current = current.Left
			} else {
				// 如果前驅的右指針已指向當前節點，恢復樹的原始結構
				predecessor.Right = nil
				// 訪問當前節點
				result = append(result, current.Val)
				// 移動到右子節點
				current = current.Right
			}
		}
	}

	return result
}
