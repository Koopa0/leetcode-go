package SubstringwithConcatenationofAllWords

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		words    []string
		expected []int
		desc     string
	}{
		{
			name:     "基本測試用例 1",
			s:        "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0, 9},
			desc:     "測試基本功能，單詞順序無關",
		},
		{
			name:     "基本測試用例 2",
			s:        "wordgoodgoodgoodbestword",
			words:    []string{"word", "good", "best", "word"},
			expected: []int{},
			desc:     "測試無匹配的情況",
		},
		{
			name:     "基本測試用例 3",
			s:        "barfoofoobarthefoobarman",
			words:    []string{"bar", "foo", "the"},
			expected: []int{6, 9, 12},
			desc:     "測試多個匹配",
		},
		{
			name:     "空字符串",
			s:        "",
			words:    []string{"foo", "bar"},
			expected: []int{},
			desc:     "測試空字符串邊界情況",
		},
		{
			name:     "空單詞數組",
			s:        "barfoothefoobarman",
			words:    []string{},
			expected: []int{},
			desc:     "測試空單詞數組邊界情況",
		},
		{
			name:     "重複單詞",
			s:        "foobarfoobar",
			words:    []string{"foo", "bar", "foo"},
			expected: []int{0},
			desc:     "測試包含重複單詞的情況",
		},
		{
			name:     "字符串長度剛好",
			s:        "foobar",
			words:    []string{"foo", "bar"},
			expected: []int{0},
			desc:     "測試字符串長度剛好等於所有單詞長度的情況",
		},
		{
			name:     "字符串長度不夠",
			s:        "foo",
			words:    []string{"foo", "bar"},
			expected: []int{},
			desc:     "測試字符串長度不夠的情況",
		},
		{
			name:     "大量重複單詞",
			s:        "aaaaaaaa",
			words:    []string{"aa", "aa", "aa"},
			expected: []int{0, 1, 2},
			desc:     "測試大量重複單詞的極端情況",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findSubstring(test.s, test.words)

			// 對結果排序以便比較
			sort.Ints(result)
			sort.Ints(test.expected)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("測試 '%s' 失敗: 期望 %v, 得到 %v\n描述: %s",
					test.name, test.expected, result, test.desc)
			}
		})
	}
}
