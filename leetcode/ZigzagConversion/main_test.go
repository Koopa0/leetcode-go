package ZigzagConversion

import "testing"

func TestConvert(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		// 基本功能測試
		{
			name:     "Example 1",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Example 2",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},

		// 邊界情況測試
		{
			name:     "Single Row",
			s:        "PAYPALISHIRING",
			numRows:  1,
			expected: "PAYPALISHIRING",
		},
		{
			name:     "Single Character",
			s:        "A",
			numRows:  3,
			expected: "A",
		},
		{
			name:     "Empty String",
			s:        "",
			numRows:  3,
			expected: "",
		},

		// 特殊情況測試
		{
			name:     "Rows Equal to String Length",
			s:        "ABC",
			numRows:  3,
			expected: "ABC",
		},
		{
			name:     "Rows Greater than String Length",
			s:        "ABC",
			numRows:  5,
			expected: "ABC",
		},

		// 性能測試（較大輸入）
		{
			name:     "Large Input",
			s:        generateLargeInput(1000),
			numRows:  100,
			expected: "", // 這裡不預先計算預期結果，而是在測試中單獨處理
		},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convert(tt.s, tt.numRows)

			// 對於性能測試，不比較結果
			if tt.name == "Large Input" {
				return
			}

			if result != tt.expected {
				t.Errorf("convert(%q, %d) = %q, want %q",
					tt.s, tt.numRows, result, tt.expected)
			}
		})
	}
}

// 生成大型測試輸入
func generateLargeInput(n int) string {
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = chars[i%len(chars)]
	}
	return string(result)
}

// 基準測試
func BenchmarkConvert(b *testing.B) {
	s := generateLargeInput(1000)
	numRows := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert(s, numRows)
	}
}
