package PalindromeNumber

func isPalindrome(x int) bool {
	// 特殊情況：
	// 1. 所有負數都不是回文數
	// 2. 末尾為0的數字不可能是回文數（除非是0本身）
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 反轉後半部分的數字
	reversed := 0

	// 當原數字大於或等於反轉數字時繼續循環
	// 這確保我們只處理數字的一半
	for x > reversed {
		// 提取原數字的最後一位，加到反轉數字的末尾
		reversed = reversed*10 + x%10
		// 去掉原數字的最後一位
		x /= 10
	}

	// 對於偶數位數，如1221，最終x==reversed
	// 對於奇數位數，如12321，需要去掉中間的數字，所以檢查x==reversed/10
	return x == reversed || x == reversed/10
}
