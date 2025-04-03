package WildcardMatching

func isMatch(s string, p string) bool {
	m, n := len(p), len(s)

	// 當前行和前一行
	curr := make([]bool, n+1)
	prev := make([]bool, n+1)

	// 基本情況
	prev[0] = true

	for i := 1; i <= m; i++ {
		// 更新當前行的第一個元素
		if p[i-1] == '*' {
			curr[0] = prev[0]
		} else {
			curr[0] = false
		}

		for j := 1; j <= n; j++ {
			if p[i-1] == '*' {
				curr[j] = prev[j] || curr[j-1]
			} else if p[i-1] == '?' || p[i-1] == s[j-1] {
				curr[j] = prev[j-1]
			} else {
				curr[j] = false
			}
		}

		// 交換當前行和前一行
		prev, curr = curr, prev
	}

	return prev[n]
}
