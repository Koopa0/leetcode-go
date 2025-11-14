package RestoreIPAddresses

import (
	"reflect"
	"sort"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		desc     string // 測試目的描述
	}{
		{
			name:     "範例 1",
			input:    "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
			desc:     "基本功能測試",
		},
		{
			name:     "範例 2",
			input:    "0000",
			expected: []string{"0.0.0.0"},
			desc:     "全零測試",
		},
		{
			name:     "範例 3",
			input:    "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
			desc:     "多解測試",
		},
		{
			name:     "邊界測試 1",
			input:    "123",
			expected: []string{},
			desc:     "字串太短",
		},
		{
			name:     "邊界測試 2",
			input:    "12345678901234567890",
			expected: []string{},
			desc:     "字串太長",
		},
		{
			name:     "前導零",
			input:    "010010",
			expected: []string{"0.10.0.10", "0.100.1.0"},
			desc:     "前導零處理",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := restoreIpAddressesIterative(tt.input)
			// 排序結果以確保比較一致性
			sort.Strings(got)
			sort.Strings(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 用於基準測試的輸入範例
	input := "25525511135"

	b.Run("回溯法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			restoreIpAddresses(input)
		}
	})

	b.Run("迭代法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			restoreIpAddressesIterative(input)
		}
	})
}
