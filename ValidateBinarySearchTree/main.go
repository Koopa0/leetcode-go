package ValidateBinarySearchTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力解法（錯誤方法）
func wrongIsValidBST(root *TreeNode) bool {
	// 空樹被認為是有效的BST
	if root == nil {
		return true
	}

	// 檢查左子節點是否小於當前節點
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	// 檢查右子節點是否大於當前節點
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	// 遞迴檢查左右子樹
	return wrongIsValidBST(root.Left) && wrongIsValidBST(root.Right)
}

// 優化解法：遞迴與值域限制
func isValidBSTWithRange(root *TreeNode) bool {
	// 使用helper函數，初始範圍為無限大
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(node *TreeNode, min int, max int) bool {
	// 空節點被認為是有效的BST
	if node == nil {
		return true
	}

	// 檢查當前節點的值是否在有效範圍內
	if node.Val <= min || node.Val >= max {
		return false
	}

	// 遞迴檢查左子樹（更新上界）和右子樹（更新下界）
	return isValidBSTHelper(node.Left, min, node.Val) &&
		isValidBSTHelper(node.Right, node.Val, max)
}

// 最佳解法：中序遍歷法
func isValidBST(root *TreeNode) bool {
	// 前一個節點的值，初始設為最小值
	var prev *TreeNode = nil
	// 使用閉包記住前一個節點
	var valid = true

	// 中序遍歷函數
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || !valid {
			return
		}

		// 遍歷左子樹
		inOrder(node.Left)

		// 檢查當前節點的值是否大於前一個節點的值
		if prev != nil && node.Val <= prev.Val {
			valid = false
			return
		}

		// 更新前一個節點
		prev = node

		// 遍歷右子樹
		inOrder(node.Right)
	}

	// 開始中序遍歷
	inOrder(root)
	return valid
}
