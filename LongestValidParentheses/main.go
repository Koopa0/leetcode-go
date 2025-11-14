package LongestValidParentheses

func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0

	// 從左到右掃描
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	// 重置計數
	left, right = 0, 0

	// 從右到左掃描
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLength
}
