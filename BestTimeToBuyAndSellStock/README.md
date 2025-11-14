# LeetCode 121: 買賣股票的最佳時機（Best Time to Buy and Sell Stock）

## 問題定義

### 原始問題（英文）
```
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
```

### 問題翻譯（繁體中文）
```
給定一個數組 prices，其中 prices[i] 表示股票在第 i 天的價格。

你想通過選擇某一天購買一支股票，並在未來的某一天出售該股票來最大化利潤。

返回你可以從這筆交易中獲得的最大利潤。如果你無法獲得任何利潤，則返回 0。
```

### 範例

**範例 1：**
```
輸入：prices = [7,1,5,3,6,4]
輸出：5
解釋：在第 2 天（股票價格 = 1）買入，在第 5 天（股票價格 = 6）賣出，利潤 = 6-1 = 5。
     注意利潤不能是 7-1 = 6，因為賣出價格需要大於買入價格；同時，你不能在買入前賣出股票。
```

**範例 2：**
```
輸入：prices = [7,6,4,3,1]
輸出：0
解釋：在這種情況下，沒有交易完成，即最大利潤為 0。
```

**限制條件：**
- 1 <= prices.length <= 10^5
- 0 <= prices[i] <= 10^4

## 解法說明

### 方法：一次遍歷

**核心思路：**
- 維護兩個變數：`minPrice`（最低價格）和 `maxProfit`（最大利潤）
- 遍歷價格數組，對於每一天：
  - 更新遇到的最低價格
  - 計算當天賣出的利潤（當天價格 - 最低價格）
  - 更新最大利潤

**複雜度分析：**
- 時間複雜度：O(n)，只需遍歷一次數組
- 空間複雜度：O(1)，只使用常數額外空間

**關鍵洞察：**
問題的核心是找到「最低點買入，後續最高點賣出」的組合。通過一次遍歷，我們可以同時追蹤最低買入價和最大利潤。
