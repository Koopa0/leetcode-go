package RestoreIPAddresses

import "strconv"

// 回溯法解決方案
func restoreIpAddresses(s string) []string {
	// 初始化結果列表
	result := []string{}

	// 邊界檢查
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	// 回溯函數
	backtrack(s, 0, 0, "", &result)

	return result
}

// 回溯函數
func backtrack(s string, index int, count int, current string, result *[]string) {
	// 如果已經找到 4 個部分且用完所有字
	if count == 4 && index == len(s) {
		// 去掉最後多餘的點
		*result = append(*result, current[:len(current)-1])
		return
	}

	// 如果已經有 4 個部分但還有剩餘字，或字用完但還未形成 4 個部分
	if count == 4 || index == len(s) {
		return
	}

	// 嘗試從當前位置取 1、2 或 3 個字
	for i := 1; i <= 3 && index+i <= len(s); i++ {
		// 取出長度為 i 的部分
		part := s[index : index+i]

		// 檢查是否有效：
		// 1. 如果長度超過 1，則不能以 0 開頭
		// 2. 數值必須在 0-255 之間
		val, _ := strconv.Atoi(part)
		if (len(part) > 1 && part[0] == '0') || val > 255 {
			continue
		}

		// 將有效部分添加到當前 IP 中，並遞迴處理剩餘部分
		backtrack(s, index+i, count+1, current+part+".", result)
	}
}

// 迭代法解決方案
func restoreIpAddressesIterative(s string) []string {
	result := []string{}

	// 邊界檢查
	if len(s) < 4 || len(s) > 12 {
		return result
	}

	// 迭代嘗試所有可能的分割位置
	for a := 1; a <= 3 && a < len(s); a++ {
		for b := a + 1; b <= a+3 && b < len(s); b++ {
			for c := b + 1; c <= b+3 && c < len(s); c++ {
				// 檢查四個部分是否都有效
				s1 := s[0:a]
				s2 := s[a:b]
				s3 := s[b:c]
				s4 := s[c:]

				if isValidPart(s1) && isValidPart(s2) && isValidPart(s3) && isValidPart(s4) {
					ip := s1 + "." + s2 + "." + s3 + "." + s4
					result = append(result, ip)
				}
			}
		}
	}

	return result
}

// 檢查 IP 地址的一部分是否有效
func isValidPart(s string) bool {
	// 長度超過 3 的部分無效
	if len(s) > 3 || len(s) == 0 {
		return false
	}

	// 如果長度超過 1，則不能以 0 開頭
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	// 檢查數值是否在 0-255 之間
	val, _ := strconv.Atoi(s)
	return val <= 255
}
