package FirstMissingPositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
		desc   string
	}{
		{
			name:   "基本測試 1",
			input:  []int{1, 2, 0},
			expect: 3,
			desc:   "正常情況，數組包含一些連續的正整數",
		},
		{
			name:   "基本測試 2",
			input:  []int{3, 4, -1, 1},
			expect: 2,
			desc:   "數組包含一些非連續的正整數和負數",
		},
		{
			name:   "基本測試 3",
			input:  []int{7, 8, 9, 11, 12},
			expect: 1,
			desc:   "數組中沒有小的正整數",
		},
		{
			name:   "邊緣情況 1",
			input:  []int{1},
			expect: 2,
			desc:   "數組只有一個元素",
		},
		{
			name:   "邊緣情況 2",
			input:  []int{-1, -2, -3},
			expect: 1,
			desc:   "數組只包含負數",
		},
		{
			name:   "邊緣情況 3",
			input:  []int{0, 0, 0},
			expect: 1,
			desc:   "數組只包含零",
		},
		{
			name:   "特殊情況 1",
			input:  []int{1, 1, 1},
			expect: 2,
			desc:   "數組包含重複的正整數",
		},
		{
			name:   "特殊情況 2",
			input:  []int{1, 2, 3, 4, 5},
			expect: 6,
			desc:   "數組包含所有從 1 到 n 的正整數",
		},
		{
			name:   "特殊情況 3",
			input:  []int{5, 4, 3, 2, 1},
			expect: 6,
			desc:   "數組包含從 1 到 n 的正整數但順序不同",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 複製輸入以避免修改原始測試數據
			input := make([]int, len(test.input))
			copy(input, test.input)

			got := firstMissingPositive(input)
			if got != test.expect {
				t.Errorf("輸入: %v, 期望: %d, 得到: %d, 描述: %s",
					test.input, test.expect, got, test.desc)
			}
		})
	}
}
