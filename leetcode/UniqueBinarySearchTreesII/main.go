package UniqueBinarySearchTreesII

import "fmt"

// TreeNode 定義樹節點結構
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遞迴分治法解決方案
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesHelper(1, n)
}

func generateTreesHelper(start, end int) []*TreeNode {
	result := []*TreeNode{}

	// 基本情況：如果範圍為空，返回一個包含 null 的列表
	if start > end {
		result = append(result, nil)
		return result
	}

	// 嘗試以每個數字作為根節點
	for i := start; i <= end; i++ {
		// 生成所有可能的左子樹
		leftTrees := generateTreesHelper(start, i-1)

		// 生成所有可能的右子樹
		rightTrees := generateTreesHelper(i+1, end)

		// 組合所有可能的左右子樹
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				// 建立根節點
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right

				// 將新樹加入結果列表
				result = append(result, root)
			}
		}
	}

	return result
}

func generateTreesMemo(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	// 建立記憶化映射表
	memo := make(map[string][]*TreeNode)
	return generateTreesWithMemo(1, n, memo)
}

func generateTreesWithMemo(start, end int, memo map[string][]*TreeNode) []*TreeNode {
	result := []*TreeNode{}

	// 基本情況
	if start > end {
		result = append(result, nil)
		return result
	}

	// 建立唯一鍵
	key := fmt.Sprintf("%d-%d", start, end)

	// 檢查是否已經計算過這個範圍
	if cached, ok := memo[key]; ok {
		return cached
	}

	// 嘗試以每個數字作為根節點
	for i := start; i <= end; i++ {
		// 生成所有可能的左子樹
		leftTrees := generateTreesWithMemo(start, i-1, memo)

		// 生成所有可能的右子樹
		rightTrees := generateTreesWithMemo(i+1, end, memo)

		// 組合所有可能的左右子樹
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				// 建立根節點
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right

				// 將新樹加入結果列表
				result = append(result, root)
			}
		}
	}

	// 將結果存入記憶化映射表
	memo[key] = result
	return result
}

func generateTreesDP(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	// 動態規劃表
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil} // 空樹

	// 對於長度從 1 到 n 的 BST
	for len := 1; len <= n; len++ {
		dp[len] = []*TreeNode{}

		// 嘗試每個可能的根節點位置
		for root := 1; root <= len; root++ {
			leftSize := root - 1    // 左子樹的節點數
			rightSize := len - root // 右子樹的節點數

			// 獲取所有可能的左子樹和右子樹
			for _, leftTree := range dp[leftSize] {
				for _, rightTree := range dp[rightSize] {
					// 建立當前根節點
					rootNode := &TreeNode{Val: root}

					// 複製並調整左子樹的值
					rootNode.Left = copyTree(leftTree, 0)

					// 複製並調整右子樹的值
					rootNode.Right = copyTree(rightTree, root)

					// 加入到結果中
					dp[len] = append(dp[len], rootNode)
				}
			}
		}
	}

	// 調整最終樹的值範圍為 1 到 n
	result := []*TreeNode{}
	for _, tree := range dp[n] {
		result = append(result, adjustTreeValues(tree, 0, 1))
	}

	return result
}

// 複製樹並調整節點值
func copyTree(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return nil
	}

	newNode := &TreeNode{Val: root.Val + offset}
	newNode.Left = copyTree(root.Left, offset)
	newNode.Right = copyTree(root.Right, offset)

	return newNode
}

// 調整樹的值範圍
func adjustTreeValues(root *TreeNode, offset int, start int) *TreeNode {
	if root == nil {
		return nil
	}

	newNode := &TreeNode{Val: root.Val - offset + start}
	newNode.Left = adjustTreeValues(root.Left, offset, start)
	newNode.Right = adjustTreeValues(root.Right, offset, start)

	return newNode
}
