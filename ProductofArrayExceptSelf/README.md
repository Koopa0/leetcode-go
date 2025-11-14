# LeetCode 238: 除自身以外數組的乘積（Product of Array Except Self）

## 問題定義

### 原始問題（英文）
```
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
```

### 問題翻譯（繁體中文）
```
給定一個整數數組 nums，返回一個數組 answer，其中 answer[i] 等於 nums 中除 nums[i] 之外所有元素的乘積。

題目保證 nums 的任何前綴或後綴的乘積都可以用 32 位整數表示。

你必須編寫一個在 O(n) 時間內運行且不使用除法運算的算法。
```

### 範例

**範例 1：**
```
輸入：nums = [1,2,3,4]
輸出：[24,12,8,6]
解釋：
answer[0] = 2*3*4 = 24
answer[1] = 1*3*4 = 12
answer[2] = 1*2*4 = 8
answer[3] = 1*2*3 = 6
```

**範例 2：**
```
輸入：nums = [-1,1,0,-3,3]
輸出：[0,0,9,0,0]
```

**限制條件：**
- 2 <= nums.length <= 10^5
- -30 <= nums[i] <= 30
- nums 的任何前綴或後綴的乘積都保證可以用 32 位整數表示

**進階要求：**
- 你能在 O(1) 額外空間複雜度內完成這個題目嗎？（輸出數組不被視為額外空間）

## 解法說明

### 方法：左右乘積列表

**核心思路：**
對於數組中的每個位置 i，`answer[i]` = 左側所有元素的乘積 × 右側所有元素的乘積

**算法步驟：**
1. 第一次遍歷（從左到右）：計算每個位置左側所有元素的乘積
2. 第二次遍歷（從右到左）：計算每個位置右側所有元素的乘積，並與左側乘積相乘

**複雜度分析：**
- 時間複雜度：O(n)，遍歷數組兩次
- 空間複雜度：O(1)，除了輸出數組外，只使用常數額外空間

**關鍵洞察：**
通過巧妙地使用輸出數組來存儲左側乘積，然後用一個變數追蹤右側乘積，我們可以在不使用除法的情況下，用 O(1) 額外空間完成計算。
