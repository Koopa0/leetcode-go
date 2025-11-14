package SymmetricTree

// TreeNode 定義
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遞迴解法
func isSymmetric(root *TreeNode) bool {
	// 處理空樹的情況
	if root == nil {
		return true
	}

	// 檢查左右子樹是否互為鏡像
	return isMirror(root.Left, root.Right)
}

// 輔助函數：檢查兩個樹是否互為鏡像
func isMirror(left, right *TreeNode) bool {
	// 如果兩個節點都為空，認為是對稱的
	if left == nil && right == nil {
		return true
	}

	// 如果一個節點為空而另一個不為空，不對稱
	if left == nil || right == nil {
		return false
	}

	// 檢查當前節點值是否相等
	if left.Val != right.Val {
		return false
	}

	// 遞迴檢查：左的左與右的右，左的右與右的左
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// 迭代解法（使用佇列）
func isSymmetricIterative(root *TreeNode) bool {
	// 處理空樹的情況
	if root == nil {
		return true
	}

	// 初始化佇列
	queue := []*TreeNode{root.Left, root.Right}

	// 循環處理佇列中的節點
	for len(queue) > 0 {
		// 每次取出兩個節點進行比較
		left := queue[0]
		right := queue[1]
		queue = queue[2:] // 移除已處理的節點

		// 如果兩個節點都為空，繼續檢查下一對
		if left == nil && right == nil {
			continue
		}

		// 如果一個節點為空而另一個不為空，不對稱
		if left == nil || right == nil {
			return false
		}

		// 如果節點值不相等，不對稱
		if left.Val != right.Val {
			return false
		}

		// 將下一層的節點按照對應關係入佇列
		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}

	// 所有節點都檢查完畢且都符合對稱條件
	return true
}

// 層序遍歷解法
func isSymmetricByLevel(root *TreeNode) bool {
	// 處理空樹的情況
	if root == nil {
		return true
	}

	// 初始化當前層的節點列表
	currentLevel := []*TreeNode{root}

	// 逐層處理
	for len(currentLevel) > 0 {
		// 建立下一層的節點列表
		nextLevel := []*TreeNode{}
		// 建立當前層的值列表，包括 nil 節點
		levelValues := []interface{}{}

		// 處理當前層的每個節點
		for _, node := range currentLevel {
			if node == nil {
				// 用特殊值表示 nil 節點
				levelValues = append(levelValues, nil)
			} else {
				// 記錄節點值
				levelValues = append(levelValues, node.Val)
				// 將下一層的節點加入列表，即使是 nil 也加入
				nextLevel = append(nextLevel, node.Left, node.Right)
			}
		}

		// 檢查當前層是否對稱（回文序列）
		if !isPalindrome(levelValues) {
			return false
		}

		// 更新當前層為下一層
		currentLevel = nextLevel
	}

	return true
}

// 輔助函數：檢查切片是否為回文序列
func isPalindrome(arr []interface{}) bool {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		// 比較兩端的元素是否相等
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}
