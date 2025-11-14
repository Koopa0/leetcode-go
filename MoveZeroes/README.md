# 283. Move Zeroes

## 題目描述

給定一個整數陣列 `nums`，將所有 `0` 移動到陣列末尾，同時保持非零元素的相對順序。

**注意**：必須在原陣列上進行操作（原地修改），不能複製額外的陣列。

## 範例

### 範例 1:
```
輸入：nums = [0,1,0,3,12]
輸出：[1,3,12,0,0]
```

### 範例 2:
```
輸入：nums = [0]
輸出：[0]
```

## 限制條件

- `1 <= nums.length <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

## 進階

你能最小化操作次數嗎？

---

## 解法思路

這是一道經典的雙指標題目，考察原地修改陣列的能力。

### 方法一：雙指標（標準解法）⭐

**核心思想：**
使用兩個指標：
- **left**: 指向下一個非零元素應該放置的位置
- **right**: 遍歷整個陣列，尋找非零元素

**演算法步驟：**
1. 初始化 left = 0
2. 用 right 遍歷整個陣列：
   - 如果 `nums[right] != 0`
   - 交換 `nums[left]` 和 `nums[right]`
   - left 向右移動一位
3. 遍歷結束後，left 之前都是非零元素，之後都是零

**為什麼有效？**
- 每找到一個非零元素，就將它移到 left 位置
- left 始終指向「下一個非零元素應該放置的位置」
- 交換操作確保了非零元素的相對順序不變

**視覺化過程（範例：[0,1,0,3,12]）：**

```
初始狀態：
[0, 1, 0, 3, 12]
 ↑
left=0, right=0, nums[0]=0（是零，不交換）

步驟 1：right=1
[0, 1, 0, 3, 12]
 ↑  ↑
left=0, right=1, nums[1]=1（非零）
交換 nums[0] 和 nums[1]
[1, 0, 0, 3, 12]
    ↑
left=1

步驟 2：right=2
[1, 0, 0, 3, 12]
    ↑  ↑
left=1, right=2, nums[2]=0（是零，不交換）

步驟 3：right=3
[1, 0, 0, 3, 12]
    ↑     ↑
left=1, right=3, nums[3]=3（非零）
交換 nums[1] 和 nums[3]
[1, 3, 0, 0, 12]
       ↑
left=2

步驟 4：right=4
[1, 3, 0, 0, 12]
       ↑      ↑
left=2, right=4, nums[4]=12（非零）
交換 nums[2] 和 nums[4]
[1, 3, 12, 0, 0]
          ↑
left=3

完成！
```

### 複雜度分析

- **時間複雜度**：O(n)
  - 只需遍歷陣列一次

- **空間複雜度**：O(1)
  - 只使用兩個指標變數，原地修改

### 方法二：雙指標優化版

**改進思路：**
分兩步處理，避免不必要的交換：
1. 第一遍：將所有非零元素移到前面
2. 第二遍：將剩餘位置填充為 0

**優點：**
- 當 left == right 時，不會進行自我交換
- 在某些情況下可能更快（沒有零或零很少時）

**缺點：**
- 需要遍歷兩次
- 程式碼稍長

### 方法三：雪球法（Snowball）

**有趣的視角：**
把所有的零想像成一個雪球，不斷往後滾。

**演算法：**
- 維護一個 `snowballSize` 變數，記錄雪球大小（遇到的零的數量）
- 遇到零時，雪球變大
- 遇到非零元素時，如果有雪球，讓非零元素「穿過」雪球

**優點：**
- 概念有趣易懂
- 效能與標準雙指標相當

### 方法四：暴力法（僅供對比）

**思路：**
1. 創建臨時陣列存儲所有非零元素
2. 複製回原陣列
3. 填充剩餘位置為 0

**缺點：**
- 需要 O(n) 額外空間
- 不符合題目要求（原地修改）

### 解法比較

| 方法 | 時間複雜度 | 空間複雜度 | 操作次數 | 優點 | 缺點 |
|------|-----------|-----------|---------|------|------|
| 雙指標（標準） | O(n) | O(1) | 最優 | 簡潔，效率高 | 有自我交換 |
| 雙指標（優化） | O(n) | O(1) | 較多 | 無自我交換 | 兩次遍歷 |
| 雪球法 | O(n) | O(1) | 最優 | 概念有趣 | 稍難理解 |
| 暴力法 | O(n) | O(n) | 較多 | 直觀 | 額外空間 |

## 實作細節

### 雙指標標準實作（推薦）

```go
func moveZeroes(nums []int) {
    if len(nums) <= 1 {
        return
    }

    left := 0  // 下一個非零元素要放置的位置

    for right := 0; right < len(nums); right++ {
        if nums[right] != 0 {
            // 交換
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
    }
}
```

**關鍵點：**
- Go 的多重賦值可以一行完成交換
- `left` 始終指向下一個非零元素應該放置的位置
- 交換確保了相對順序不變

### 雙指標優化版

```go
func moveZeroesOptimized(nums []int) {
    if len(nums) <= 1 {
        return
    }

    left := 0

    // 第一遍：移動非零元素
    for right := 0; right < len(nums); right++ {
        if nums[right] != 0 {
            nums[left] = nums[right]
            left++
        }
    }

    // 第二遍：填充零
    for i := left; i < len(nums); i++ {
        nums[i] = 0
    }
}
```

### 雪球法實作

```go
func moveZeroesSnowball(nums []int) {
    if len(nums) <= 1 {
        return
    }

    snowballSize := 0  // 雪球大小（零的數量）

    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            snowballSize++  // 雪球變大
        } else if snowballSize > 0 {
            // 非零元素穿過雪球
            nums[i-snowballSize] = nums[i]
            nums[i] = 0
        }
    }
}
```

## 常見錯誤

### 錯誤 1：破壞了相對順序

```go
// ❌ 錯誤：使用雙指標從兩端向中間移動會破壞順序
func moveZeroes(nums []int) {
    left, right := 0, len(nums)-1
    for left < right {
        if nums[left] == 0 && nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
        }
        if nums[left] != 0 {
            left++
        }
        if nums[right] == 0 {
            right--
        }
    }
}
// 這會破壞非零元素的相對順序！
```

### 錯誤 2：使用額外空間

```go
// ❌ 錯誤：創建新陣列不符合原地修改要求
func moveZeroes(nums []int) {
    result := []int{}
    zeroCount := 0

    for _, num := range nums {
        if num != 0 {
            result = append(result, num)
        } else {
            zeroCount++
        }
    }

    for i := 0; i < zeroCount; i++ {
        result = append(result, 0)
    }

    copy(nums, result)  // 需要額外 O(n) 空間
}
```

### 錯誤 3：不必要的複雜邏輯

```go
// ❌ 過度複雜
func moveZeroes(nums []int) {
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            for j := i + 1; j < len(nums); j++ {
                if nums[j] != 0 {
                    nums[i], nums[j] = nums[j], nums[i]
                    break
                }
            }
        }
    }
}
// 時間複雜度退化為 O(n²)
```

## 效能測試結果

根據基準測試（不同規模和零的比例）：

```
大陣列（10000 元素，20% 零）：
- 雪球法：31167 ns/op ⭐ 最快
- 優化版：33166 ns/op
- 標準雙指標：33734 ns/op
- 暴力法：71512 ns/op（最慢，且需要雙倍記憶體）

全是零（1000 元素）：
- 標準雙指標：3581 ns/op ⭐
- 雪球法：3853 ns/op
- 優化版：4211 ns/op
- 暴力法：5390 ns/op

沒有零（1000 元素）：
- 雪球法：4548 ns/op ⭐
- 標準雙指標：4572 ns/op
- 優化版：5007 ns/op
- 暴力法：9450 ns/op
```

**結論：**
- **雪球法**在大多數情況下最快
- **標準雙指標**在全是零的情況下最優
- **優化版**效能穩定，但不一定最快
- **暴力法**在所有情況下都最慢

實際應用中，**標準雙指標**程式碼最簡潔，效能也很好，是首選解法。

## 相關題目

- [27. Remove Element](https://leetcode.com/problems/remove-element/) - 移除特定元素
- [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) - 移除重複元素
- [80. Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/) - 保留部分重複
- [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) - 合併有序陣列

## 延伸思考

1. **如果要移動到開頭呢？**
   - 從右到左遍歷，原理相同

2. **如果有多種元素要移動呢？**
   - 可以使用多個指標或計數器
   - 參考「荷蘭國旗問題」（3-way partitioning）

3. **如果要穩定排序呢？**
   - 當前解法已經是穩定的（保持相對順序）

---

## Description

Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

**Note** that you must do this in-place without making a copy of the array.

## Constraints

- `1 <= nums.length <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

## Follow up

Could you minimize the total number of operations done?

## Solution

Use two pointers: one to track the position for the next non-zero element, and one to traverse the array.

**Algorithm:**
1. Initialize `left = 0` (position for next non-zero element)
2. Iterate through the array with `right`:
   - If `nums[right] != 0`, swap `nums[left]` and `nums[right]`, then increment `left`
3. All non-zero elements will be moved to the front, maintaining their relative order

**Time Complexity:** O(n) - single pass through the array
**Space Complexity:** O(1) - only two pointer variables
