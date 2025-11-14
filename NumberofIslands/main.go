package main

// numIslands 使用 DFS 深度優先搜尋計算島嶼數量
// 遍歷網格，每次遇到 '1' 時進行 DFS，將相連的陸地全部標記為已訪問
// 每次 DFS 代表找到一個新島嶼，計數器加 1
// 時間複雜度: O(m × n)，其中 m 和 n 是網格的行數和列數
// 空間複雜度: O(m × n)，遞迴堆疊在最壞情況下可能達到網格大小
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// 遍歷整個網格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 發現陸地，進行 DFS
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

// dfs 深度優先搜尋，將相連的陸地標記為已訪問
func dfs(grid [][]byte, i, j int) {
	rows := len(grid)
	cols := len(grid[0])

	// 邊界檢查和陸地檢查
	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != '1' {
		return
	}

	// 標記為已訪問（使用 '0' 或其他字元）
	grid[i][j] = '0'

	// 遞迴訪問四個方向：上、下、左、右
	dfs(grid, i-1, j) // 上
	dfs(grid, i+1, j) // 下
	dfs(grid, i, j-1) // 左
	dfs(grid, i, j+1) // 右
}

// numIslandsBFS 使用 BFS 廣度優先搜尋計算島嶼數量
// 時間複雜度: O(m × n)
// 空間複雜度: O(min(m, n))，佇列在最壞情況下的大小
func numIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// 方向陣列：上、下、左、右
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				// BFS
				queue := [][]int{{i, j}}
				grid[i][j] = '0'

				for len(queue) > 0 {
					curr := queue[0]
					queue = queue[1:]
					row, col := curr[0], curr[1]

					// 檢查四個方向
					for _, dir := range directions {
						newRow := row + dir[0]
						newCol := col + dir[1]

						// 邊界檢查和陸地檢查
						if newRow >= 0 && newRow < rows &&
						   newCol >= 0 && newCol < cols &&
						   grid[newRow][newCol] == '1' {
							queue = append(queue, []int{newRow, newCol})
							grid[newRow][newCol] = '0' // 標記為已訪問
						}
					}
				}
			}
		}
	}

	return count
}

// numIslandsDFSIterative 使用迭代 DFS（堆疊）計算島嶼數量
// 時間複雜度: O(m × n)
// 空間複雜度: O(m × n)，堆疊在最壞情況下的大小
func numIslandsDFSIterative(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// 方向陣列：上、下、左、右
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				// 使用堆疊進行迭代 DFS
				stack := [][]int{{i, j}}
				grid[i][j] = '0'

				for len(stack) > 0 {
					// 彈出堆疊頂部
					curr := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					row, col := curr[0], curr[1]

					// 檢查四個方向
					for _, dir := range directions {
						newRow := row + dir[0]
						newCol := col + dir[1]

						if newRow >= 0 && newRow < rows &&
						   newCol >= 0 && newCol < cols &&
						   grid[newRow][newCol] == '1' {
							stack = append(stack, []int{newRow, newCol})
							grid[newRow][newCol] = '0'
						}
					}
				}
			}
		}
	}

	return count
}

func main() {
	// 範例 1
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	println("範例 1:")
	println("島嶼數量:", numIslands(grid1)) // 1

	// 範例 2（需要重新建立，因為 grid 會被修改）
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	println("\n範例 2:")
	println("島嶼數量:", numIslands(grid2)) // 3
}
