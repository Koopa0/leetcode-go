package ScrambleString

// 暴力解法 (遞迴)
func isScrambleBruteForce(s1 string, s2 string) bool {
	// 字串長度不同，直接返回 false
	if len(s1) != len(s2) {
		return false
	}

	// 兩個字串相同，直接返回 true
	if s1 == s2 {
		return true
	}

	// 檢查字元組成是否相同
	if !hasSameChars(s1, s2) {
		return false
	}

	n := len(s1)
	// 嘗試所有可能的分割點
	for i := 1; i < n; i++ {
		// 不交換的情況
		if isScrambleBruteForce(s1[:i], s2[:i]) && isScrambleBruteForce(s1[i:], s2[i:]) {
			return true
		}

		// 交換的情況
		if isScrambleBruteForce(s1[:i], s2[n-i:]) && isScrambleBruteForce(s1[i:], s2[:n-i]) {
			return true
		}
	}

	return false
}

// 輔助函數：檢查兩個字串是否包含相同的字元
func hasSameChars(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)

	for _, ch := range s1 {
		counts[ch]++
	}

	for _, ch := range s2 {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}

// 優化解法 (記憶化搜索)
func isScrambleMemoization(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	n := len(s1)
	// 初始化記憶化數組，-1 表示未計算，0 表示 false，1 表示 true
	memo := make([][][]int8, n)
	for i := range memo {
		memo[i] = make([][]int8, n)
		for j := range memo[i] {
			memo[i][j] = make([]int8, n+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

	var checkScramble func(i1, i2, length int) bool
	checkScramble = func(i1, i2, length int) bool {
		// 檢查記憶化數組
		if memo[i1][i2][length] != -1 {
			return memo[i1][i2][length] == 1
		}

		// 基本情況：長度為 1
		if length == 1 {
			result := s1[i1] == s2[i2]
			if result {
				memo[i1][i2][length] = 1
			} else {
				memo[i1][i2][length] = 0
			}
			return result
		}

		// 檢查字元組成是否相同
		if !hasSameCharsWithIndices(s1, s2, i1, i2, length) {
			memo[i1][i2][length] = 0
			return false
		}

		// 嘗試所有可能的分割點
		for i := 1; i < length; i++ {
			// 不交換的情況
			if checkScramble(i1, i2, i) && checkScramble(i1+i, i2+i, length-i) {
				memo[i1][i2][length] = 1
				return true
			}

			// 交換的情況
			if checkScramble(i1, i2+length-i, i) && checkScramble(i1+i, i2, length-i) {
				memo[i1][i2][length] = 1
				return true
			}
		}

		memo[i1][i2][length] = 0
		return false
	}

	return checkScramble(0, 0, n)
}

// 輔助函數：檢查指定索引開始的子字串是否包含相同的字元
func hasSameCharsWithIndices(s1, s2 string, i1, i2, length int) bool {
	counts := make(map[rune]int)

	for i := 0; i < length; i++ {
		counts[rune(s1[i1+i])]++
		counts[rune(s2[i2+i])]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

// 最佳解法 (動態規劃)
func isScrambleDP(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	n := len(s1)

	// 檢查字元組成是否相同
	if !hasSameChars(s1, s2) {
		return false
	}

	// 初始化 DP 數組
	// dp[len][i][j] 表示 s1 從索引 i 開始，s2 從索引 j 開始，長度為 len 的子字串是否可以互相擾亂
	dp := make([][][]bool, n+1)
	for len := 0; len <= n; len++ {
		dp[len] = make([][]bool, n)
		for i := 0; i < n; i++ {
			dp[len][i] = make([]bool, n)
		}
	}

	// 基本情況：長度為 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}

	// 填充 DP 數組
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			for j := 0; j <= n-length; j++ {
				for k := 1; k < length; k++ {
					// 不交換的情況
					if dp[k][i][j] && dp[length-k][i+k][j+k] {
						dp[length][i][j] = true
						break
					}

					// 交換的情況
					if dp[k][i][j+length-k] && dp[length-k][i+k][j] {
						dp[length][i][j] = true
						break
					}
				}
			}
		}
	}

	return dp[n][0][0]
}

// 增強性能的混合方法
func isScramble(s1 string, s2 string) bool {
	// 基本條件檢查
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	// 字元頻率快速檢查
	var freq [26]int
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
		freq[s2[i]-'a']--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}

	// 使用字串對作為鍵的記憶化映射
	memo := make(map[string]bool)

	var scrambleCheck func(a, b string) bool
	scrambleCheck = func(a, b string) bool {
		// 基本情況檢查
		if a == b {
			return true
		}

		if len(a) == 1 {
			return a == b
		}

		// 檢查記憶化結果
		key := a + ":" + b
		if result, exists := memo[key]; exists {
			return result
		}

		n := len(a)

		// 增量字元頻率檢查
		for i := 1; i < n; i++ {
			// 不交換的情況
			var leftEqual, rightEqual bool = true, true

			// 快速檢查前綴是否可能匹配
			var charCount [26]int
			for j := 0; j < i; j++ {
				charCount[a[j]-'a']++
				charCount[b[j]-'a']--
			}

			for _, count := range charCount {
				if count != 0 {
					leftEqual = false
					break
				}
			}

			// 只有當字元頻率匹配時才遞迴調用
			if leftEqual {
				// 檢查後綴是否可能匹配
				var charCount2 [26]int
				for j := i; j < n; j++ {
					charCount2[a[j]-'a']++
					charCount2[b[j]-'a']--
				}

				for _, count := range charCount2 {
					if count != 0 {
						rightEqual = false
						break
					}
				}

				if rightEqual && scrambleCheck(a[:i], b[:i]) && scrambleCheck(a[i:], b[i:]) {
					memo[key] = true
					return true
				}
			}

			// 交換的情況
			leftEqual, rightEqual = true, true

			// 重置字元計數數組
			var charCount3 [26]int
			for j := 0; j < i; j++ {
				charCount3[a[j]-'a']++
				charCount3[b[n-i+j]-'a']--
			}

			for _, count := range charCount3 {
				if count != 0 {
					leftEqual = false
					break
				}
			}

			if leftEqual {
				var charCount4 [26]int
				for j := i; j < n; j++ {
					charCount4[a[j]-'a']++
					charCount4[b[j-i]-'a']--
				}

				for _, count := range charCount4 {
					if count != 0 {
						rightEqual = false
						break
					}
				}

				if rightEqual && scrambleCheck(a[:i], b[n-i:]) && scrambleCheck(a[i:], b[:n-i]) {
					memo[key] = true
					return true
				}
			}
		}

		// 所有分割點都檢查完後才設置為 false
		memo[key] = false
		return false
	}

	return scrambleCheck(s1, s2)
}
