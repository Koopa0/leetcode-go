package RecoverBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力解法
func bruteForceRecoverTree(root *TreeNode) {
	// 1. 初始化
	var nodes []*TreeNode

	// 2. 中序遍歷，將節點收集到陣列中
	inorder(root, &nodes)

	// 3. 尋找錯誤放置的節點
	var first, second *TreeNode
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			if first == nil {
				first = nodes[i]
			}
			second = nodes[i+1]
		}
	}

	// 4. 交換這兩個節點的值
	first.Val, second.Val = second.Val, first.Val
}

// 中序遍歷函數
func inorder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left, nodes)
	*nodes = append(*nodes, root)
	inorder(root.Right, nodes)
}

// 優化解法
func optimizedRecoverTree(root *TreeNode) {
	// 1. 初始化
	var first, second, prev *TreeNode

	// 2. 中序遍歷並找出錯誤節點
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 遍歷左子樹
		inorder(node.Left)

		// 處理當前節點
		if prev != nil && prev.Val > node.Val {
			// 找到違反規則的情況
			if first == nil {
				first = prev // 第一次發現錯誤時，記錄前一個節點
			}
			second = node // 無論是否是第一次發現錯誤，都更新 second
		}
		prev = node

		// 遍歷右子樹
		inorder(node.Right)
	}

	inorder(root)

	// 3. 交換錯誤節點的值
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// 最佳解法 - Morris Traversal
func morrisRecoverTree(root *TreeNode) {
	// 1. 初始化
	var first, second, prev *TreeNode

	// 2. Morris 中序遍歷
	current := root
	for current != nil {
		// 如果沒有左子樹，則訪問當前節點並轉向右子樹
		if current.Left == nil {
			// 處理當前節點
			if prev != nil && prev.Val > current.Val {
				if first == nil {
					first = prev
				}
				second = current
			}
			prev = current
			current = current.Right
		} else {
			// 有左子樹，找到前驅節點
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			// 第一次到達前驅節點
			if predecessor.Right == nil {
				predecessor.Right = current // 設置臨時鏈接
				current = current.Left
			} else {
				// 第二次到達前驅節點，說明左子樹已經遍歷完畢
				predecessor.Right = nil // 恢復樹結構

				// 處理當前節點
				if prev != nil && prev.Val > current.Val {
					if first == nil {
						first = prev
					}
					second = current
				}
				prev = current
				current = current.Right
			}
		}
	}

	// 3. 交換錯誤節點的值
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
