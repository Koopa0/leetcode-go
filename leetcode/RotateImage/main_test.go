package RotateImage

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
		desc     string
	}{
		{
			name:     "3x3 矩陣",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
			desc:     "標準3x3矩陣，測試基本旋轉功能",
		},
		{
			name:     "4x4 矩陣",
			input:    [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
			desc:     "較大矩陣，測試算法在更多元素情況下的正確性",
		},
		{
			name:     "1x1 矩陣",
			input:    [][]int{{1}},
			expected: [][]int{{1}},
			desc:     "邊界情況：1x1矩陣旋轉後保持不變",
		},
		{
			name:     "2x2 矩陣",
			input:    [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{3, 1}, {4, 2}},
			desc:     "偶數尺寸的小矩陣，測試奇偶邊界情況",
		},
		{
			name:     "2x2 矩陣 - 負數",
			input:    [][]int{{-1, -2}, {-3, -4}},
			expected: [][]int{{-3, -1}, {-4, -2}},
			desc:     "包含負數的矩陣，測試算法對不同值域的處理",
		},
		{
			name:     "3x3 矩陣 - 零",
			input:    [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			expected: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			desc:     "全零矩陣，測試特殊情況下的處理",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 創建輸入矩陣的副本，因為rotate會原地修改矩陣
			input := make([][]int, len(tc.input))
			for i := range tc.input {
				input[i] = make([]int, len(tc.input[i]))
				copy(input[i], tc.input[i])
			}

			// 執行旋轉操作
			rotate(input)

			// 驗證結果
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("Test '%s' failed.\nDescription: %s\nInput: %v\nExpected: %v\nGot: %v",
					tc.name, tc.desc, tc.input, tc.expected, input)
			}
		})
	}
}
