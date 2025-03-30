package LongestPalindromicSubstring

import "strings"

func longestPalindrome(s string) string {

	// Manacher 所以要在頭尾加上任意特殊符號 以 ^ , $ 兩個符號定義
	// 然後在每個字串前加上同樣的符號, 確保他前後回文,
	// 加上符號的原因是 確保 s 為奇數 可以實現 Manacher 算法
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"

	// 找出 string 的長度
	n := len(T)

	// 儲存每個位置的回文數
	P := make([]int, n)

	// C: 當前找到的回文中心
	// R: 當前右邊界，即 C + P[C]，表示當前找到的回文的右邊界位置
	C, R := 0, 0

	// 步驟 3: 計算每個位置的回文半徑
	// 注意：我們跳過首尾的哨兵字符，所以範圍是 [1, n-1)
	for i := 1; i < len(T)-1; i++ {

		if R > i {
			// R > i 代表右邊界在當前的右邊 可以使用 Manacher 的重要觀念
			// 當前的 回文中的回文 , 會左右呼應 ， 所以當前檢查的數字的回文數
			// 至少會大於等於鏡像的回文數
			// 最小可能會是 右邊界到當前 index 的數量或 鏡像的回文數 取最小值
			P[i] = min(R-i, P[2*C-i])
		}

		// 步驟 4: 中心擴展
		// 無論是否利用了對稱性，都嘗試向外擴展
		// 檢查 T[i-1-P[i]] 和 T[i+1+P[i]] 是否相等，如果相等則繼續擴展

		// 假設 i = 6 他的回文數為 3
		// 如果要繼續檢查和比對 就要檢查 10 和 2
		// 10 就會是 6 + 1 + 3 (i + 擴展 + 回文數)
		// 2 就會是  6 - 1 - 3 (i - 擴展 - 回文數)
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		//步驟 5: 更新 C 和 R
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

	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}
