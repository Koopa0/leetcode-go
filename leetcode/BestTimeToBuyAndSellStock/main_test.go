package BestTimeToBuyAndSellStock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	// 定義測試用例表
	// 每個測試用例包含輸入、預期輸出以及測試描述
	testCases := []struct {
		prices      []int
		expected    int
		description string
	}{
		// 基本測試用例
		{[]int{7, 1, 5, 3, 6, 4}, 5, "普通測試用例 - 最大利潤是5"},
		{[]int{7, 6, 4, 3, 1}, 0, "下跌趨勢 - 無法獲利，應返回0"},

		// 邊緣情況
		{[]int{}, 0, "空數組 - 無法交易，應返回0"},
		{[]int{1}, 0, "只有一天 - 無法同時買入和賣出，應返回0"},

		// 特殊情況
		{[]int{1, 2}, 1, "只有兩天 - 利潤為1"},
		{[]int{2, 2, 2}, 0, "價格相同 - 無法獲利，應返回0"},
		{[]int{3, 2, 6, 5, 0, 3}, 4, "多次波動 - 最大利潤為4 (2買入6賣出)"},

		// 極端情況
		{[]int{10000}, 0, "一天，最大價格 - 無法交易，應返回0"},
		{[]int{0, 10000}, 10000, "最小價格到最大價格 - 利潤為10000"},

		// 性能測試用例
		{make([]int, 10000), 0, "大規模輸入，全部相同 - 應返回0"},
	}

	// 針對大規模測試用例生成遞增數組
	largeIncreasing := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		largeIncreasing[i] = i
	}
	testCases = append(testCases, struct {
		prices      []int
		expected    int
		description string
	}{largeIncreasing, 9999, "大規模輸入，持續上升 - 利潤為9999"})

	// 運行所有測試用例
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := maxProfit(tc.prices)
			if actual != tc.expected {
				t.Errorf("對於價格 %v，預期利潤為 %d，但得到 %d",
					tc.prices, tc.expected, actual)
			}
		})
	}
}
