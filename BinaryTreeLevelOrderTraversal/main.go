package main

// TreeNode 定義二元樹節點
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 使用 BFS 廣度優先搜尋進行二元樹層序遍歷
// 使用佇列逐層處理節點，記錄每一層的節點值
// 時間複雜度: O(n)，其中 n 是節點數量
// 空間複雜度: O(n)，佇列最多存儲一層的節點
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	// 空樹處理
	if root == nil {
		return result
	}

	// 初始化佇列，加入根節點
	queue := []*TreeNode{root}

	// 當佇列不為空時持續處理
	for len(queue) > 0 {
		// 記錄當前層的大小
		levelSize := len(queue)
		currentLevel := []int{}

		// 處理當前層的所有節點
		for i := 0; i < levelSize; i++ {
			// 取出佇列首個節點
			node := queue[0]
			queue = queue[1:]

			// 記錄當前節點的值
			currentLevel = append(currentLevel, node.Val)

			// 將子節點加入佇列（下一層）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 將當前層的結果加入最終結果
		result = append(result, currentLevel)
	}

	return result
}

// levelOrderDFS 使用 DFS 深度優先搜尋進行層序遍歷
// 透過遞迴和層級參數記錄每一層的節點值
// 時間複雜度: O(n)
// 空間複雜度: O(h)，其中 h 是樹的高度（遞迴堆疊）
func levelOrderDFS(root *TreeNode) [][]int {
	result := [][]int{}
	dfsHelper(root, 0, &result)
	return result
}

// dfsHelper 輔助函數：遞迴遍歷並記錄每一層的節點
func dfsHelper(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// 如果當前層還沒有記錄，建立新的層
	if level >= len(*result) {
		*result = append(*result, []int{})
	}

	// 將當前節點值加入對應的層
	(*result)[level] = append((*result)[level], node.Val)

	// 遞迴處理左右子樹（層級加 1）
	dfsHelper(node.Left, level+1, result)
	dfsHelper(node.Right, level+1, result)
}

// levelOrderIterative 使用迭代方式的層序遍歷（優化版）
// 使用雙指標技巧減少記憶體分配
// 時間複雜度: O(n)
// 空間複雜度: O(n)
func levelOrderIterative(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		levelValues := make([]int, len(currentLevel))
		nextLevel := []*TreeNode{}

		// 處理當前層的所有節點
		for i, node := range currentLevel {
			levelValues[i] = node.Val

			// 收集下一層的節點
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		result = append(result, levelValues)
		currentLevel = nextLevel
	}

	return result
}

func main() {
	// 範例 1: [3,9,20,null,null,15,7]
	//     3
	//    / \
	//   9  20
	//     /  \
	//    15   7
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	println("範例 1: [3,9,20,null,null,15,7]")
	result1 := levelOrder(root1)
	print("層序遍歷: ")
	for _, level := range result1 {
		print("[")
		for i, val := range level {
			if i > 0 {
				print(",")
			}
			print(val)
		}
		print("] ")
	}
	println()

	// 範例 2: [1]
	root2 := &TreeNode{Val: 1}
	println("\n範例 2: [1]")
	result2 := levelOrder(root2)
	print("層序遍歷: ")
	for _, level := range result2 {
		print("[")
		for i, val := range level {
			if i > 0 {
				print(",")
			}
			print(val)
		}
		print("] ")
	}
	println()

	// 範例 3: []
	println("\n範例 3: []")
	result3 := levelOrder(nil)
	println("層序遍歷:", len(result3), "層")
}
