package LongestPalindromicSubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// 定義測試表格，包含多種測試情況
	tests := []struct {
		name        string // 測試用例名稱
		input       string // 輸入字符串
		expected    string // 期望輸出
		description string // 測試目的描述
	}{
		{
			name:        "Example 1",
			input:       "babad",
			expected:    "bab", // 注意：返回 "aba" 也是正確的
			description: "測試基本用例，有多個相同長度的回文",
		},
		{
			name:        "Example 2",
			input:       "cbbd",
			expected:    "bb",
			description: "測試偶數長度回文",
		},
		{
			name:        "Single character",
			input:       "a",
			expected:    "a",
			description: "測試單個字符邊界情況",
		},
		{
			name:        "Two same characters",
			input:       "aa",
			expected:    "aa",
			description: "測試兩個相同字符的情況",
		},
		{
			name:        "Two different characters",
			input:       "ab",
			expected:    "a", // 注意：返回 "b" 也是正確的
			description: "測試兩個不同字符的情況",
		},
		{
			name:        "All same characters",
			input:       "aaaaa",
			expected:    "aaaaa",
			description: "測試全部相同字符的情況",
		},
		{
			name:        "Palindrome at beginning",
			input:       "racecarxyz",
			expected:    "racecar",
			description: "測試回文在字符串開始的情況",
		},
		{
			name:        "Palindrome at end",
			input:       "xyzracecar",
			expected:    "racecar",
			description: "測試回文在字符串結尾的情況",
		},
		{
			name:        "Palindrome in middle",
			input:       "xyzracecarzyx",
			expected:    "racecar",
			description: "測試回文在字符串中間的情況",
		},
		{
			name:        "Multiple palindromes",
			input:       "aacabdkacaa",
			expected:    "aca",
			description: "測試存在多個不同回文的情況",
		},
		{
			name:        "Long string",
			input:       "ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy",
			expected:    "fklkf",
			description: "測試長字符串性能",
		},
	}

	// 執行測試表格中的每個測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 調用被測函數
			result := longestPalindrome(tt.input)

			// 檢查結果是否符合預期
			// 注意：由於一些情況下可能有多個正確答案，需要檢查結果長度和是否為回文
			if !isPalindrome(result) || len(result) != len(tt.expected) {
				t.Errorf("longestPalindrome(%q) = %q, want %q or equivalent length palindrome\n%s",
					tt.input, result, tt.expected, tt.description)
			}
		})
	}
}

// 輔助函數：檢查字符串是否為回文
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
