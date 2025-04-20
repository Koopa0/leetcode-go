package SimplifyPath

import "strings"

// 暴力解法 使用堆疊
func simplifyPath(path string) string {
	// 將路徑以斜線分割
	dirs := strings.Split(path, "/")

	// 建立一個堆疊來存儲目錄名稱
	stack := []string{}

	// 遍歷所有目錄名稱
	for _, dir := range dirs {
		// 如果是空字串或 '.'，跳過
		if dir == "" || dir == "." {
			continue
		}

		// 如果是 '..'，且堆疊不為空，彈出頂部元素
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// 否則將目錄名稱推入堆疊
			stack = append(stack, dir)
		}
	}

	// 組合最終路徑
	result := "/" + strings.Join(stack, "/")

	return result
}

// 優化解法
func simplifyPathOptimized(path string) string {
	stack := []string{}

	// 使用更簡潔的方式遍歷分割後的路徑
	for _, dir := range strings.Split(path, "/") {
		switch dir {
		case "", ".":
			// 跳過空字串和當前目錄
			continue
		case "..":
			// 返回上一層目錄
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			// 添加有效目錄名稱
			stack = append(stack, dir)
		}
	}

	// 處理邊界情況：空堆疊
	if len(stack) == 0 {
		return "/"
	}

	// 組合最終路徑
	return "/" + strings.Join(stack, "/")
}

// 最佳解法
func simplifyPathOptimal(path string) string {
	// 創建堆疊
	stack := []string{}

	if len(path) == 0 {
		return "/"
	}

	if len(path) == 1 && path[0] == '/' {
		return path
	}

	// 確保路徑以斜線開頭和結尾，方便處理
	if path[len(path)-1] != '/' {
		path += "/"
	}

	// 遍歷路徑
	start := 0
	for i := 1; i < len(path); i++ {
		// 找到目錄分隔符
		if path[i] == '/' {
			// 提取目錄名稱
			dir := path[start+1 : i]
			start = i

			// 處理不同類型的目錄名稱
			if dir == "" || dir == "." {
				continue
			} else if dir == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, dir)
			}
		}
	}

	// 組合最終路徑
	if len(stack) == 0 {
		return "/"
	}

	result := ""
	for _, dir := range stack {
		result += "/" + dir
	}

	return result
}
