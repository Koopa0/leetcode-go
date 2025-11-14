# 136. Single Number

## 題目描述

給定一個非空整數陣列，除了某個元素只出現一次以外，其餘每個元素均出現兩次。找出那個只出現一次的元素。

**注意**：你的演算法應該具有線性時間複雜度，且不使用額外空間。

## 範例

### 範例 1:
```
輸入：nums = [2,2,1]
輸出：1
```

### 範例 2:
```
輸入：nums = [4,1,2,1,2]
輸出：4
```

### 範例 3:
```
輸入：nums = [1]
輸出：1
```

## 限制條件

- `1 <= nums.length <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`
- 除了某個元素只出現一次以外，其餘每個元素均出現兩次

---

## 解法思路

### 方法一：位元運算 XOR（最佳解法）⭐

使用 XOR（互斥或）位元運算的數學性質：

**XOR 的關鍵性質：**
1. `a ^ a = 0`：相同的數字 XOR 結果為 0
2. `a ^ 0 = a`：任何數字與 0 XOR 結果為自己
3. **交換律和結合律**：`a ^ b ^ a = (a ^ a) ^ b = 0 ^ b = b`

**演算法步驟：**
1. 初始化 result = 0
2. 遍歷陣列，將所有數字與 result 進行 XOR 運算
3. 成對的數字會相互抵消（變成 0）
4. 最後剩下的就是只出現一次的數字

**範例演示：**
```
輸入：[4, 1, 2, 1, 2]

計算過程：
result = 0
result = 0 ^ 4 = 4
result = 4 ^ 1 = 5
result = 5 ^ 2 = 7
result = 7 ^ 1 = 6    (因為 4 ^ 1 ^ 1 = 4 ^ 0 = 4)
result = 6 ^ 2 = 4    (因為 4 ^ 2 ^ 2 = 4 ^ 0 = 4)

最終結果：4
```

### 複雜度分析

- **時間複雜度**：O(n)
  - 只需遍歷陣列一次

- **空間複雜度**：O(1)
  - 只使用一個變數儲存結果

### 方法二：雜湊表（直觀解法）

使用雜湊表統計每個數字出現的次數，找出出現一次的數字。

**複雜度分析：**
- 時間複雜度：O(n)
- 空間複雜度：O(n)

### 方法三：數學解法（集合）

利用數學公式：`2 * (a + b + c) - (a + a + b + b + c) = c`

**步驟：**
1. 計算所有不重複數字的和 sumUnique
2. 計算陣列中所有數字的和 sumAll
3. 結果 = 2 * sumUnique - sumAll

**複雜度分析：**
- 時間複雜度：O(n)
- 空間複雜度：O(n)（需要儲存集合）

### 解法比較

| 方法 | 時間複雜度 | 空間複雜度 | 優點 | 缺點 |
|------|-----------|-----------|------|------|
| XOR 位元運算 | O(n) | O(1) | 最優解，空間效率最高 | 需要理解位元運算 |
| 雜湊表 | O(n) | O(n) | 直觀易懂 | 需要額外空間 |
| 數學公式 | O(n) | O(n) | 有趣的思路 | 需要額外空間 |

## 實作細節

### XOR 實作（推薦）

```go
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}
```

**為什麼 XOR 有效？**

XOR 運算的真值表：
```
a | b | a ^ b
--|---|------
0 | 0 |  0
0 | 1 |  1
1 | 0 |  1
1 | 1 |  0
```

當兩個相同的數字進行 XOR 時，對應位元相同則為 0，最終結果為 0。

**位元層級的演示（範例：[2, 2, 1]）：**
```
  2: 010
^ 2: 010
------
  0: 000

  0: 000
^ 1: 001
------
  1: 001
```

### 雜湊表實作

```go
func singleNumberWithMap(nums []int) int {
    count := make(map[int]int)

    for _, num := range nums {
        count[num]++
    }

    for num, cnt := range count {
        if cnt == 1 {
            return num
        }
    }

    return 0
}
```

### 數學公式實作

```go
func singleNumberWithSet(nums []int) int {
    set := make(map[int]bool)
    sumUnique := 0
    sumAll := 0

    for _, num := range nums {
        sumAll += num
        if !set[num] {
            set[num] = true
            sumUnique += num
        }
    }

    return 2*sumUnique - sumAll
}
```

## 效能測試結果

根據基準測試（小/中/大陣列）：

```
XOR 方法：
- 小陣列（3 元素）：~2-3 ns/op
- 中陣列（11 元素）：~8-10 ns/op
- 大陣列（10001 元素）：~6000-7000 ns/op

Map 方法：約為 XOR 的 10-20 倍時間
Set 方法：約為 XOR 的 15-25 倍時間
```

**結論**：XOR 方法在所有情況下都是最快且最節省空間的解法。

## 延伸問題

1. **Single Number II (LeetCode 137)**：每個元素出現三次，只有一個出現一次
   - 解法：使用位元運算統計每個位元出現的次數

2. **Single Number III (LeetCode 260)**：有兩個元素只出現一次
   - 解法：先 XOR 全部得到兩個數的 XOR，再利用差異位元分組

## 相關題目

- [137. Single Number II](https://leetcode.com/problems/single-number-ii/) - 每個元素出現三次
- [260. Single Number III](https://leetcode.com/problems/single-number-iii/) - 兩個元素只出現一次
- [268. Missing Number](https://leetcode.com/problems/missing-number/) - 同樣可用 XOR 解決

---

## Description

Given a non-empty array of integers `nums`, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

## Constraints

- `1 <= nums.length <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`
- Each element in the array appears twice except for one element which appears only once.

## Solution

Use XOR bit manipulation. The key properties:
- `a ^ a = 0`
- `a ^ 0 = a`
- XOR is commutative and associative

All pairs will cancel out (XOR to 0), leaving only the single number.

**Time Complexity:** O(n)
**Space Complexity:** O(1)
