package LongestPalindromicSubstring

import "strings"

// Manacher 算法：在線性時間內找出字串中最長的回文子串
func longestPalindrome(s string) string {

	// 步驟 1: 預處理原始字串，插入特殊字符
	// 在字串的每個字符之間插入特殊字符 '#'，並在開頭和結尾添加哨兵字符 '^' 和 '$'
	// 例如: "babad" 變成 "^#b#a#b#a#d#$"
	// 這樣做的目的是:
	// 1. 統一處理奇數和偶數長度的回文
	// 2. 避免邊界檢查，簡化代碼
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)

	// 步驟 2: 創建回文半徑數組 P
	// P[i] 表示以 i 為中心的回文子串的半徑（不包括中心點自身）
	// 例如: P[i] = 3 表示以 i 為中心，向左右各擴展 3 個字符形成的子串是回文
	P := make([]int, n)

	// C: 當前找到的回文中心
	// R: 當前右邊界，即 C + P[C]，表示當前找到的回文的右邊界位置
	C, R := 0, 0

	// 步驟 3: 計算每個位置的回文半徑
	// 注意：我們跳過首尾的哨兵字符，所以範圍是 [1, n-1)
	for i := 1; i < n-1; i++ {
		// 情況 1: i 在 R 的範圍內，可以利用對稱性質
		if R > i {
			// i 關於 C 的對稱點是 2*C - i
			// 利用回文的對稱性，P[i] 至少等於 P[2*C-i] 或 R-i 中的較小值
			// 為什麼取兩者的較小值？
			// - 如果 P[2*C-i] < R-i，則 P[i] = P[2*C-i]
			// - 如果 P[2*C-i] >= R-i，則 P[i] 至少為 R-i，然後再擴展
			P[i] = min(R-i, P[2*C-i])
		}
		// 否則，P[i] 默認為 0

		// 步驟 4: 中心擴展
		// 無論是否利用了對稱性，都嘗試向外擴展
		// 檢查 T[i-1-P[i]] 和 T[i+1+P[i]] 是否相等，如果相等則繼續擴展
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		// 步驟 5: 更新 C 和 R
		// 如果當前回文的右邊界超過了 R，則更新中心點 C 和右邊界 R
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}

	// 步驟 6: 找出最長回文子串
	// 遍歷回文半徑數組 P，找出最大值及其對應的中心點
	maxLen := 0      // 最大回文半徑
	centerIndex := 0 // 最大回文的中心點
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}

	// 步驟 7: 將處理後的字串映射回原始字串
	// 在預處理字串中，回文的範圍是 [centerIndex-maxLen, centerIndex+maxLen]
	// 對應到原始字串的起始位置是 (centerIndex-maxLen)/2
	// 結束位置是 (centerIndex+maxLen)/2
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}
