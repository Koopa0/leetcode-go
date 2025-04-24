package ValidNumber

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		desc     string // 測試目的描述
	}{
		// 基本功能測試
		{"Valid Integer", "123", true, "簡單整數測試"},
		{"Valid Decimal", "3.14", true, "簡單小數測試"},
		{"Valid Scientific", "2e10", true, "簡單科學記數法測試"},
		{"Valid Complex", "-123.456e789", true, "複雜有效數字測試"},

		// 邊緣情況測試
		{"Empty String", "", false, "空字串應該無效"},
		{"Only Sign", "+", false, "只有符號字應該無效"},
		{"Only Dot", ".", false, "只有小數點應該無效"},
		{"Only E", "e", false, "只有e應該無效"},

		// 特殊情況測試
		{"Decimal No Trailing", "4.", true, "小數點後沒有數字也是有效的"},
		{"Decimal No Leading", ".5", true, "小數點前沒有數字也是有效的"},
		{"E With Sign", "3e+7", true, "e後帶符號的科學記數法"},

		// 無效輸入測試
		{"Invalid Character", "abc", false, "含有字母的無效輸入"},
		{"Invalid Format", "1a", false, "數字後跟隨無效字"},
		{"Invalid E Format", "1e", false, "e後沒有數字的無效格式"},
		{"Invalid Double Sign", "--6", false, "連續符號的無效格式"},
		{"Invalid E Decimal", "99e2.5", false, "e後接小數的無效格式"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNumber(tt.input)
			if got != tt.expected {
				t.Errorf("isNumber(%q) = %v, 預期為 %v, 說明: %s",
					tt.input, got, tt.expected, tt.desc)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	validInput := "+3.14e-10"
	invalidInput := "99e2.5"

	b.Run("有限狀態機-有效輸入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isNumberFSM(validInput)
		}
	})

	b.Run("有限狀態機-無效輸入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isNumberFSM(invalidInput)
		}
	})

	b.Run("線性掃描-有效輸入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isNumberLinear(validInput)
		}
	})

	b.Run("線性掃描-無效輸入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isNumberLinear(invalidInput)
		}
	})

	b.Run("正則表達式-有效輸入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isNumber(validInput)
		}
	})

	b.Run("正則表達式-無效輸入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isNumber(invalidInput)
		}
	})
}
