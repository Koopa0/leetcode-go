package GroupAnagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "示例1",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "示例2",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "示例3",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "無字母異位詞",
			input:    []string{"a", "b", "c", "d"},
			expected: [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
		},
		{
			name:     "全部相同",
			input:    []string{"a", "a", "a"},
			expected: [][]string{{"a", "a", "a"}},
		},
		{
			name:     "空輸入",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "長字串",
			input:    []string{"aaaaaaaaaa", "aaaaaaaaaa"},
			expected: [][]string{{"aaaaaaaaaa", "aaaaaaaaaa"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)

			// 由於返回順序不重要，我們需要比較結果是否等價
			if !areEquivalentResults(result, tt.expected) {
				t.Errorf("groupAnagrams() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 輔助函數：檢查兩個二維字串數組是否在內容上等價
// （即它們包含相同的字串組，但順序可能不同）
func areEquivalentResults(result, expected [][]string) bool {
	if len(result) != len(expected) {
		return false
	}

	// 先對每個子數組進行排序，然後排序外層數組
	sortedResult := sortGroups(result)
	sortedExpected := sortGroups(expected)

	return reflect.DeepEqual(sortedResult, sortedExpected)
}

// 輔助函數：排序二維字串數組
func sortGroups(groups [][]string) [][]string {
	// 對每個子數組進行排序
	for i := range groups {
		sort.Strings(groups[i])
	}

	// 將每個子數組轉換為字串，用於排序
	groupStrs := make([]string, len(groups))
	for i, group := range groups {
		for _, s := range group {
			groupStrs[i] += s + ","
		}
	}

	// 根據字串表示對組進行排序
	sortedIndices := make([]int, len(groups))
	for i := range sortedIndices {
		sortedIndices[i] = i
	}
	sort.SliceStable(sortedIndices, func(i, j int) bool {
		return groupStrs[sortedIndices[i]] < groupStrs[sortedIndices[j]]
	})

	// 創建排序後的結果
	sortedGroups := make([][]string, len(groups))
	for i, idx := range sortedIndices {
		sortedGroups[i] = groups[idx]
	}

	return sortedGroups
}
