package SimplifyPath

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		desc     string
	}{
		{
			name:     "基本測試1",
			input:    "/home/",
			expected: "/home",
			desc:     "測試末尾斜線的移除",
		},
		{
			name:     "基本測試2",
			input:    "/../",
			expected: "/",
			desc:     "測試向上超出根目錄的情況",
		},
		{
			name:     "連續斜線測試",
			input:    "/home//foo/",
			expected: "/home/foo",
			desc:     "測試連續斜線的處理",
		},
		{
			name:     "複雜路徑測試",
			input:    "/a/./b/../../c/",
			expected: "/c",
			desc:     "測試複雜路徑的簡化",
		},
		{
			name:     "僅根目錄測試",
			input:    "/",
			expected: "/",
			desc:     "測試僅有根目錄的情況",
		},
		{
			name:     "多個點測試",
			input:    "/a/b/c/.../",
			expected: "/a/b/c/...",
			desc:     "測試包含多個點的目錄名稱",
		},
		{
			name:     "空路徑測試",
			input:    "",
			expected: "/",
			desc:     "測試空路徑的處理",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPathOptimal(tt.input)
			if got != tt.expected {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 用於基準測試的範例輸入
	input := "/a/./b/../../c/d/.././e/../../f/./g/h/./../../i/./"

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			simplifyPath(input)
		}
	})

	b.Run("優化解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			simplifyPathOptimized(input)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			simplifyPathOptimal(input)
		}
	})
}
