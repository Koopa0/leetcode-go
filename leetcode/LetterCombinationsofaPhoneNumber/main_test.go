package LetterCombinationsofaPhoneNumber

import (
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "空字符串",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "單個數字",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "兩個數字",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "三個數字",
			digits: "234",
			expected: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
		{
			name:   "包含 7 和 9 (有 4 個字母的數字)",
			digits: "79",
			expected: []string{
				"pw", "px", "py", "pz", "qw", "qx", "qy", "qz",
				"rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz",
			},
		},
		{
			name:     "最大長度 4",
			digits:   "2345",
			expected: nil, // 這裡我們不列出所有預期結果，只檢查長度
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := letterCombinations(tt.digits)

			// 對於較長的輸入，只檢查結果長度
			if tt.name == "最大長度 4" {
				expected := 3 * 3 * 3 * 3 // 3*3*3*3 = 81 種組合
				if len(result) != expected {
					t.Errorf("預期 %d 種組合，得到 %d 種", expected, len(result))
				}
				return
			}

			// 檢查結果是否包含所有預期的組合
			if len(result) != len(tt.expected) {
				t.Errorf("預期 %d 種組合，得到 %d 種", len(tt.expected), len(result))
				return
			}

			// 將結果排序以便比較
			sort.Strings(result)
			expected := make([]string, len(tt.expected))
			copy(expected, tt.expected)
			sort.Strings(expected)

			for i := 0; i < len(result); i++ {
				if result[i] != expected[i] {
					t.Errorf("第 %d 個組合預期為 %s，得到 %s", i, expected[i], result[i])
				}
			}
		})
	}
}
