package TextJustification

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		expected []string
		desc     string
	}{
		{
			name:     "基本測試",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
			desc: "測試基本功能",
		},
		{
			name:     "最後一行特殊處理",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
			desc: "測試最後一行左對齊",
		},
		{
			name:     "長句子",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
			desc: "測試更複雜的長句子",
		},
		{
			name:     "單詞行",
			words:    []string{"a"},
			maxWidth: 10,
			expected: []string{"a         "},
			desc:     "測試單一單詞且長度遠小於最大寬度",
		},
		{
			name:     "長單詞",
			words:    []string{"pneumonoultramicroscopicsilicovolcanoconiosis"},
			maxWidth: 45,
			expected: []string{"pneumonoultramicroscopicsilicovolcanoconiosis"},
			desc:     "測試單一長單詞剛好等於最大寬度",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fullJustify(tt.words, tt.maxWidth)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 範例輸入
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20

	b.Run("蠻力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fullJustifyForce(words, maxWidth)
		}
	})

	b.Run("優化解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fullJustifyForce(words, maxWidth)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fullJustify(words, maxWidth)
		}
	})
}
