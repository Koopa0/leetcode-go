package MultiplyStrings

import "testing"

func TestMultiply(t *testing.T) {
	// 定義測試表
	tests := []struct {
		name     string
		num1     string
		num2     string
		expected string
		desc     string
	}{
		{
			name:     "Example 1",
			num1:     "2",
			num2:     "3",
			expected: "6",
			desc:     "簡單的單位數乘法",
		},
		{
			name:     "Example 2",
			num1:     "123",
			num2:     "456",
			expected: "56088",
			desc:     "多位數乘法",
		},
		{
			name:     "Zero Multiplication",
			num1:     "0",
			num2:     "123",
			expected: "0",
			desc:     "任何數乘以零等於零",
		},
		{
			name:     "Single Digit Multiplication",
			num1:     "9",
			num2:     "9",
			expected: "81",
			desc:     "單位數乘法，有進位",
		},
		{
			name:     "Large Numbers",
			num1:     "9999",
			num2:     "9999",
			expected: "99980001",
			desc:     "大數相乘",
		},
		{
			name:     "Different Length",
			num1:     "1234",
			num2:     "56789",
			expected: "70077626",
			desc:     "不同長度的數字相乘",
		},
		{
			name:     "Leading Zeros in Result",
			num1:     "100",
			num2:     "10",
			expected: "1000",
			desc:     "結果中間包含零",
		},
		{
			name:     "Very Large Numbers",
			num1:     "123456789",
			num2:     "987654321",
			expected: "121932631112635269",
			desc:     "非常大的數字相乘，測試效能",
		},
	}

	// 執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.num1, tt.num2); got != tt.expected {
				t.Errorf("multiply() = %v, want %v", got, tt.expected)
			}
		})
	}
}
