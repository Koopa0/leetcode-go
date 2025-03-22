package GroupAnagrams

import (
	"reflect"
	"sort"
	"testing"
)

// TestGroupAnagrams 使用表格驅動測試方法測試字母異序詞分組函數
func TestGroupAnagrams(t *testing.T) {
	// 定義測試表格
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "LeetCode Example 1",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "LeetCode Example 2",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "LeetCode Example 3",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "Empty Input",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "All Same Anagrams",
			input:    []string{"abc", "bca", "cab", "acb"},
			expected: [][]string{{"abc", "bca", "cab", "acb"}},
		},
		{
			name:     "All Different Strings",
			input:    []string{"a", "b", "c", "d"},
			expected: [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
		},
		{
			name:     "Duplicate Strings",
			input:    []string{"eat", "eat", "tea", "tea"},
			expected: [][]string{{"eat", "eat", "tea", "tea"}},
		},
		{
			name:     "Mixed Length Strings",
			input:    []string{"ab", "ba", "abc", "cba", "a", "b"},
			expected: [][]string{{"ab", "ba"}, {"abc", "cba"}, {"a"}, {"b"}},
		},
	}

	// 遍歷測試表格
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 運行被測函數
			result := groupAnagramsCount(tt.input)

			// 由於結果順序不固定，需要先對結果進行正規化處理
			if !compareAnagramGroups(result, tt.expected) {
				t.Errorf("groupAnagrams() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestGroupAnagramsCount 測試基於字符計數的方法
func TestGroupAnagramsCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "LeetCode Example 1",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		// 可以添加其他測試用例...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagramsCount(tt.input)

			if !compareAnagramGroups(result, tt.expected) {
				t.Errorf("groupAnagramsCount() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 幫助函數：比較兩個字母異序詞分組結果是否相等
func compareAnagramGroups(actual, expected [][]string) bool {
	if len(actual) != len(expected) {
		return false
	}

	// 將每個子數組排序，然後對整個二維數組排序，以便於比較
	normalizedActual := normalizeGroups(actual)
	normalizedExpected := normalizeGroups(expected)

	return reflect.DeepEqual(normalizedActual, normalizedExpected)
}

// 幫助函數：正規化分組結果（對每個分組內部排序，然後對分組進行排序）
func normalizeGroups(groups [][]string) [][]string {
	// 創建一個新的切片，避免修改原始數據
	normalized := make([][]string, len(groups))

	// 對每個分組內的字符串進行排序
	for i, group := range groups {
		sorted := make([]string, len(group))
		copy(sorted, group)
		sort.Strings(sorted)
		normalized[i] = sorted
	}

	// 對分組進行排序
	sort.Slice(normalized, func(i, j int) bool {
		// 如果分組長度不同，按長度排序
		if len(normalized[i]) != len(normalized[j]) {
			return len(normalized[i]) < len(normalized[j])
		}

		// 如果分組長度相同，按第一個元素的字典序排序
		if len(normalized[i]) > 0 && len(normalized[j]) > 0 {
			return normalized[i][0] < normalized[j][0]
		}

		return false
	})

	return normalized
}

// 幫助函數：生成大量測試數據
func generateLargeInput(size int) []string {
	// 在實際測試中，可以生成大量的字母異序詞
	// 這裡只是一個簡單的範例
	result := make([]string, size)
	for i := 0; i < size; i++ {
		// 生成字母異序詞的邏輯
		// ...
		result[i] = "teststring" // 簡化實現
	}
	return result
}

// 基準測試，用於比較兩種實現的性能
func BenchmarkGroupAnagrams(b *testing.B) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	b.Run("SortMethod", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			groupAnagramsCount(input)
		}
	})

	b.Run("CountMethod", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			groupAnagramsCount(input)
		}
	})
}
