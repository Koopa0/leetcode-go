package LetterCombinationsofaPhoneNumber

func letterCombinations(digits string) []string {
	// 如果輸入為空，返回空列表
	if len(digits) == 0 {
		return []string{}
	}

	// 定義數字到字母的映射
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	// 回溯函數
	var backtrack func(index int, current string)
	backtrack = func(index int, current string) {
		// 如果當前組合長度等於輸入長度，表示已經處理完所有數字
		if len(current) == len(digits) {
			result = append(result, current)
			return
		}

		// 獲取當前數字對應的字母
		// e.g. digits = "23"，index = 0，letters = "abc" digit = '2'
		// e.g. digits = "23"，index = 1，letters = "def"
		// e.g. digits = "23"，index = 2，letters = ""
		digit := digits[index]
		letters := phoneMap[digit]

		// 嘗試每個字母
		for i := 0; i < len(letters); i++ {
			// 添加當前字母到組合中
			backtrack(index+1, current+string(letters[i]))
		}
	}

	// 從第一個數字開始回溯
	backtrack(0, "")

	return result
}
