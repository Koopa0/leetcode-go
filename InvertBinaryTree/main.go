package main

// TreeNode 定義二元樹節點
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 使用 DFS 遞迴翻轉二元樹
// 翻轉二元樹意味著交換每個節點的左右子樹
// 時間複雜度: O(n)，其中 n 是節點數量
// 空間複雜度: O(h)，其中 h 是樹的高度（遞迴堆疊）
func invertTree(root *TreeNode) *TreeNode {
	// 基本情況：空節點或葉子節點
	if root == nil {
		return nil
	}

	// 交換左右子樹
	root.Left, root.Right = root.Right, root.Left

	// 遞迴翻轉左右子樹
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// invertTreePostOrder 使用後序遍歷翻轉二元樹
// 先遞迴處理子樹，再交換當前節點的左右子樹
// 時間複雜度: O(n)
// 空間複雜度: O(h)
func invertTreePostOrder(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 先遞迴翻轉左右子樹
	left := invertTreePostOrder(root.Left)
	right := invertTreePostOrder(root.Right)

	// 再交換當前節點的左右子樹
	root.Left = right
	root.Right = left

	return root
}

// invertTreeBFS 使用 BFS 廣度優先搜尋翻轉二元樹
// 使用佇列逐層處理節點，交換每個節點的左右子樹
// 時間複雜度: O(n)
// 空間複雜度: O(w)，其中 w 是樹的最大寬度
func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 初始化佇列
	queue := []*TreeNode{root}

	// 層序遍歷
	for len(queue) > 0 {
		// 取出佇列首個節點
		node := queue[0]
		queue = queue[1:]

		// 交換左右子樹
		node.Left, node.Right = node.Right, node.Left

		// 將子節點加入佇列
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// invertTreeIterative 使用 DFS 迭代（堆疊）翻轉二元樹
// 使用堆疊模擬遞迴過程
// 時間複雜度: O(n)
// 空間複雜度: O(h)
func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 初始化堆疊
	stack := []*TreeNode{root}

	// 迭代處理
	for len(stack) > 0 {
		// 彈出堆疊頂部節點
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 交換左右子樹
		node.Left, node.Right = node.Right, node.Left

		// 將子節點壓入堆疊
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return root
}

// 輔助函數：層序遍歷輸出樹（用於驗證）
func levelOrderPrint(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

func main() {
	// 範例 1: [4,2,7,1,3,6,9]
	//       4
	//      / \
	//     2   7
	//    / \ / \
	//   1  3 6  9
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 7}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 9}

	println("範例 1: 翻轉前 [4,2,7,1,3,6,9]")
	before := levelOrderPrint(root1)
	print("層序遍歷: ")
	for _, level := range before {
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

	invertTree(root1)
	println("\n翻轉後:")
	after := levelOrderPrint(root1)
	print("層序遍歷: ")
	for _, level := range after {
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

	// 範例 2: [2,1,3]
	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}

	println("\n範例 2: 翻轉前 [2,1,3]")
	before2 := levelOrderPrint(root2)
	print("層序遍歷: ")
	for _, level := range before2 {
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

	invertTree(root2)
	println("\n翻轉後:")
	after2 := levelOrderPrint(root2)
	print("層序遍歷: ")
	for _, level := range after2 {
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
}
