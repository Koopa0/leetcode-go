package BestTimeToBuyAndSellStock

// maxProfit 計算可能的最大利潤
// 時間複雜度: O(n)，其中n是價格數組的長度
// 空間複雜度: O(1)
func maxProfit(prices []int) int {
	// 處理邊緣情況：如果價格數組為空或只有一個元素，無法進行交易，返回0
	if len(prices) <= 1 {
		return 0
	}

	// 初始化最小價格為第一天的價格，最大利潤為0
	minPrice := prices[0]
	maxProfit := 0

	// 遍歷價格數組，從第二天開始
	for i := 1; i < len(prices); i++ {
		// 如果當前價格低於記錄的最低價格，更新最低價格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// 計算當前利潤（當前價格 - 最低價格）
			currentProfit := prices[i] - minPrice
			// 如果當前利潤大於記錄的最大利潤，更新最大利潤
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	// 返回最大利潤
	return maxProfit
}
