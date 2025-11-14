package GenerateParentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	// 測試用例表
	tests := []struct {
		name string   // 測試用例名稱
		n    int      // 輸入：括號對數
		want []string // 預期輸出：有效括號組合的列表
	}{
		{
			name: "零對括號",
			n:    0,
			want: []string{""},
		},
		{
			name: "一對括號",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "兩對括號",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "三對括號",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	// 執行所有測試用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			// 由於生成的括號順序可能不同，需要檢查內容是否匹配而不關心順序
			if !areEqualSets(got, tt.want) {
				t.Errorf("generateParenthesis(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// areEqualSets 檢查兩個字串集合是否包含相同的元素（不考慮順序）
func areEqualSets(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// 將 a 的元素添加到映射中
	aMap := make(map[string]int)
	for _, val := range a {
		aMap[val]++
	}

	// 將 b 的元素添加到映射中
	bMap := make(map[string]int)
	for _, val := range b {
		bMap[val]++
	}

	// 比較兩個映射是否相同
	for key, aVal := range aMap {
		if bMap[key] != aVal {
			return false
		}
	}

	return true
}
