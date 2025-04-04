# Find First and Last Position of Element in Sorted Array

## 1. Original Problem

**English:**
Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

**Example 2:**
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

**Example 3:**
```
Input: nums = [], target = 0
Output: [-1,-1]
```

**Constraints:**
- 0 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
- nums is a non-decreasing array
- -10^9 <= target <= 10^9

**繁體中文：**
給你一個按非遞減順序排列的整數數組 `nums`，請你找出給定目標值 `target` 的開始位置和結束位置。

如果數組中不存在目標值 `target`，返回 `[-1, -1]`。

你必須設計並實現時間複雜度為 `O(log n)` 的演算法解決此問題。

**示例 1：**
```
輸入：nums = [5,7,7,8,8,10], target = 8
輸出：[3,4]
```

**示例 2：**
```
輸入：nums = [5,7,7,8,8,10], target = 6
輸出：[-1,-1]
```

**示例 3：**
```
輸入：nums = [], target = 0
輸出：[-1,-1]
```

**約束條件：**
- 0 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
- nums 按非遞減順序排列
- -10^9 <= target <= 10^9

## 2. 問題理解與初步分析

這個問題要求我們在一個已排序的數組中找出目標值的第一個和最後一個位置。

**核心需求：**
- 找出目標值在排序數組中的開始和結束位置
- 如果目標值不存在，返回 [-1, -1]
- 算法的時間複雜度必須是 O(log n)

**輸入/輸出特徵：**
- 輸入是一個已排序（非遞減）的整數數組和一個目標值
- 輸出是一個包含兩個整數的數組，表示目標值的開始和結束位置

**邊界情況：**
- 空數組 → 返回 [-1, -1]
- 目標值不在數組中 → 返回 [-1, -1]
- 目標值只出現一次 → 開始和結束位置相同
- 目標值出現多次 → 需找出第一次和最後一次出現的位置

**潛在難點：**
1. 標準二分查找只能找到目標值的任一位置，而不是開始或結束位置
2. 需要設計方法來專門找到第一個和最後一個目標值的位置
3. 在保持 O(log n) 時間複雜度的同時處理所有邊界情況

## 3. 問題模式識別

這是一個**二分查找的變體問題**。關鍵指標如下：

1. 數組是已排序的（非遞減順序）
2. 要求 O(log n) 的時間複雜度
3. 需要在排序數組中查找特定元素的位置

標準二分查找通常只能找到目標值的任意一個位置，而這個問題需要找到目標值的開始和結束位置。這需要對二分查找進行修改或執行兩次不同的二分查找：
- 一次用於查找第一個出現的位置（開始位置）
- 一次用於查找最後一個出現的位置（結束位置）

從問題約束條件中可以獲得的提示：
- 數組大小可能很大（最多 10^5），暗示線性搜索是不夠高效的
- 數組可能為空，需要處理這種特殊情況
- 目標值可能不在數組中，需要返回 [-1, -1]

## 4. 策略性解題框架

對於此類二分查找變體問題，我們可以採用以下策略：

1. **將一個問題分解為兩個子問題**:
    - 查找目標值第一次出現的位置
    - 查找目標值最後一次出現的位置

2. **修改二分查找的決策邏輯**:
    - 查找第一個位置時：即使找到目標值，也繼續在左半部分搜索
    - 查找最後一個位置時：即使找到目標值，也繼續在右半部分搜索

3. **處理邊界情況**:
    - 特別注意空數組和目標值不存在的情況
    - 確保返回值正確反映搜索結果

4. **思考路徑**:
    - 先確認問題是否可以用二分查找解決
    - 確定需要找什麼（第一個和最後一個位置）
    - 修改標準二分查找以滿足特定需求
    - 驗證解決方案在各種情況下的正確性

## 5. 視覺化說明

讓我們以示例 `nums = [5,7,7,8,8,10], target = 8` 來視覺化解釋：

**找尋目標值 8 的第一個位置：**

```
初始狀態：
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
      ↑               ↑
      左              右
      中 = (0+5)/2 = 2
```

```
第一次迭代：
中間值 nums[2] = 7 < 目標值 8，移動左指針
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
            ↑         ↑
            左        右
            中 = (3+5)/2 = 4
```

```
第二次迭代：
中間值 nums[4] = 8 == 目標值 8，但我們要找第一個位置，所以移動右指針
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
            ↑  ↑
            左 右
            中 = (3+3)/2 = 3
```

```
第三次迭代：
中間值 nums[3] = 8 == 目標值 8，但仍需檢查左側，所以移動右指針
由於左 = 右 = 3，迭代結束
結果：第一個位置是索引 3
```

**找尋目標值 8 的最後一個位置：**

```
初始狀態：
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
      ↑               ↑
      左              右
      中 = (0+5)/2 = 2
```

```
第一次迭代：
中間值 nums[2] = 7 < 目標值 8，移動左指針
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
            ↑        ↑
            左       右
            中 = (3+5)/2 = 4
```

```
第二次迭代：
中間值 nums[4] = 8 == 目標值 8，但我們要找最後一個位置，所以移動左指針
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
                  ↑  ↑
                  左 右
                  中 = (5+5)/2 = 5
```

```
第三次迭代：
中間值 nums[5] = 10 > 目標值 8，移動右指針
索引： 0  1  2  3  4  5
數值： 5  7  7  8  8  10
                  ↑
                左,右
                  終止
結果：最後一個位置是索引 4
```

最終結果：[3, 4]

## 6. 解題思路發展過程

### 初始暴力解法

最直觀的方法是線性掃描整個數組：
1. 從左到右掃描，找到第一個目標值的位置
2. 從右到左掃描，找到最後一個目標值的位置

```go
func searchRange(nums []int, target int) []int {
    first, last := -1, -1
    
    // 找第一個位置
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            first = i
            break
        }
    }
    
    // 找最後一個位置
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] == target {
            last = i
            break
        }
    }
    
    return []int{first, last}
}
```

這種方法的時間複雜度是 O(n)，不符合題目要求的 O(log n)。

### 優化思路

為了達到 O(log n) 的時間複雜度，我們需要使用二分查找。但是，標準的二分查找只能找到目標值的任意一個位置，我們需要修改它來找到第一個和最後一個位置。

1. **找第一個位置的二分查找**：
    - 當 `nums[mid] == target` 時，記錄當前位置，但繼續在左側搜索
    - 當 `nums[mid] < target` 時，在右側搜索
    - 當 `nums[mid] > target` 時，在左側搜索

2. **找最後一個位置的二分查找**：
    - 當 `nums[mid] == target` 時，記錄當前位置，但繼續在右側搜索
    - 當 `nums[mid] < target` 時，在右側搜索
    - 當 `nums[mid] > target` 時，在左側搜索

關鍵洞察：修改標準二分查找的方式是，即使找到目標值，也不立即返回，而是繼續在特定方向搜索以找到第一個或最後一個位置。

## 7. 實例步驟演示

讓我們使用示例 `nums = [5,7,7,8,8,10], target = 8` 來演示完整的解決方案：

### 查找第一個位置的過程：

**初始狀態**：
- left = 0, right = 5
- 中間位置 mid = (0 + 5) / 2 = 2
- nums[mid] = 7 < target(8)，所以我們在右側搜索
- 更新 left = mid + 1 = 3

**第一次迭代**：
- left = 3, right = 5
- 中間位置 mid = (3 + 5) / 2 = 4
- nums[mid] = 8 == target(8)，我們找到了目標值
- 但為了找第一個位置，我們記錄 firstPos = 4 並在左側繼續搜索
- 更新 right = mid - 1 = 3

**第二次迭代**：
- left = 3, right = 3
- 中間位置 mid = (3 + 3) / 2 = 3
- nums[mid] = 8 == target(8)，我們找到了目標值
- 記錄 firstPos = 3 並在左側繼續搜索
- 更新 right = mid - 1 = 2

**第三次迭代**：
- left = 3, right = 2
- 由於 left > right，搜索結束
- 第一個位置為 firstPos = 3

### 查找最後一個位置的過程：

**初始狀態**：
- left = 0, right = 5
- 中間位置 mid = (0 + 5) / 2 = 2
- nums[mid] = 7 < target(8)，所以我們在右側搜索
- 更新 left = mid + 1 = 3

**第一次迭代**：
- left = 3, right = 5
- 中間位置 mid = (3 + 5) / 2 = 4
- nums[mid] = 8 == target(8)，我們找到了目標值
- 但為了找最後一個位置，我們記錄 lastPos = 4 並在右側繼續搜索
- 更新 left = mid + 1 = 5

**第二次迭代**：
- left = 5, right = 5
- 中間位置 mid = (5 + 5) / 2 = 5
- nums[mid] = 10 > target(8)，所以我們在左側搜索
- 更新 right = mid - 1 = 4

**第三次迭代**：
- left = 5, right = 4
- 由於 left > right，搜索結束
- 最後一個位置為 lastPos = 4

**最終結果**：[3, 4]

## 8. Golang 實現

```go
func searchRange(nums []int, target int) []int {
    // 初始化結果為 [-1, -1]
    result := []int{-1, -1}
    
    // 處理空數組情況
    if len(nums) == 0 {
        return result
    }
    
    // 查找第一個位置
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            result[0] = mid  // 找到一個目標值，暫存為可能的第一個位置
            right = mid - 1  // 繼續在左側搜索，尋找可能更早出現的目標值
        } else if nums[mid] < target {
            left = mid + 1   // 目標值在右側
        } else {
            right = mid - 1  // 目標值在左側
        }
    }
    
    // 如果沒有找到目標值，直接返回 [-1, -1]
    if result[0] == -1 {
        return result
    }
    
    // 查找最後一個位置
    left, right = 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            result[1] = mid  // 找到一個目標值，暫存為可能的最後一個位置
            left = mid + 1   // 繼續在右側搜索，尋找可能更晚出現的目標值
        } else if nums[mid] < target {
            left = mid + 1   // 目標值在右側
        } else {
            right = mid - 1  // 目標值在左側
        }
    }
    
    return result
}
```

## 9. 程式碼執行過程解析

讓我們使用示例 `nums = [5,7,7,8,8,10], target = 8` 來追蹤程式碼的執行過程：

**查找第一個位置**：
1. 初始化：`left = 0, right = 5, result = [-1, -1]`
2. 進入第一個循環（查找第一個位置）
    - 第一次迭代：`mid = 2, nums[mid] = 7 < target(8)`, 更新 `left = 3`
    - 第二次迭代：`mid = 4, nums[mid] = 8 == target(8)`, 設置 `result[0] = 4`, 更新 `right = 3`
    - 第三次迭代：`mid = 3, nums[mid] = 8 == target(8)`, 設置 `result[0] = 3`, 更新 `right = 2`
    - 第四次迭代：`left = 3, right = 2`, 循環結束
3. 此時 `result = [3, -1]`

**查找最後一個位置**：
1. 重置：`left = 0, right = 5`
2. 進入第二個循環（查找最後一個位置）
    - 第一次迭代：`mid = 2, nums[mid] = 7 < target(8)`, 更新 `left = 3`
    - 第二次迭代：`mid = 4, nums[mid] = 8 == target(8)`, 設置 `result[1] = 4`, 更新 `left = 5`
    - 第三次迭代：`mid = 5, nums[mid] = 10 > target(8)`, 更新 `right = 4`
    - 第四次迭代：`left = 5, right = 4`, 循環結束
3. 此時 `result = [3, 4]`

最終返回：`[3, 4]`

## 10. 複雜度分析

**時間複雜度**：
- 執行了兩次二分查找，每次查找的時間複雜度為 O(log n)
- 總時間複雜度為 O(log n) + O(log n) = O(log n)

**空間複雜度**：
- 只使用了常數額外空間來存儲變量
- 空間複雜度為 O(1)

**為什麼時間複雜度不能進一步優化**：
- 二分查找的時間複雜度已經是 O(log n)，這是在排序數組中查找元素的理論最佳時間複雜度
- 我們需要找到第一個和最後一個位置，這需要兩次獨立的搜索，但兩次 O(log n) 的操作仍然是 O(log n)

## 11. 優化與改進

雖然我們的解決方案已經達到了最佳的時間複雜度 O(log n)，但還有一些可能的優化和變體：

1. **合併兩次二分查找**：
    - 可以通過一次二分查找找到任意一個目標值的位置，然後向左和向右擴展來找到邊界
    - 但在最壞情況下（所有元素都是目標值），這種方法的時間複雜度會退化為 O(n)

2. **使用標準庫函數**：
    - 在實際項目中，可以考慮使用標準庫中的二分查找函數（如果可用），以減少實現錯誤的可能性

3. **更簡潔的實現**：
    - 可以抽取一個輔助函數來執行二分查找，使程式碼更加模塊化和易於維護

```go
func searchRange(nums []int, target int) []int {
    // 輔助函數：查找邊界
    findBoundary := func(findFirst bool) int {
        left, right := 0, len(nums)-1
        index := -1
        
        for left <= right {
            mid := left + (right-left)/2
            if nums[mid] == target {
                index = mid
                if findFirst {
                    right = mid - 1  // 查找第一個位置
                } else {
                    left = mid + 1   // 查找最後一個位置
                }
            } else if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        
        return index
    }
    
    // 處理空數組情況
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    
    // 查找第一個和最後一個位置
    first := findBoundary(true)
    last := findBoundary(false)
    
    return []int{first, last}
}
```

## 12. 一般解題智慧

從這個問題中，我們可以學到以下幾個重要的解題思路：

1. **問題分解**：將複雜問題分解為多個子問題（在這裡是將"找範圍"分解為"找第一個位置"和"找最後一個位置"）。

2. **算法變體**：學會修改基本算法（如二分查找）以適應不同的問題需求。標準算法通常需要根據具體問題進行調整。

3. **邊界條件處理**：特別注意處理邊界情況，如空數組、目標值不存在、目標值只出現一次等。

4. **決策邏輯**：關鍵在於何時移動左指針和右指針。在本問題中，即使找到目標值，也要根據是找第一個還是最後一個位置來決定指針的移動方向。

5. **模式識別**：學會識別排序數組中的搜索問題通常可以用二分查找解決，尤其是當要求 O(log n) 時間複雜度時。

6. **程式碼抽象**：善用輔助函數減少程式碼重複，提高可讀性和可維護性。

7. **最優解非常簡單**：這個問題的最優解決方案僅需約 20 行程式碼，這反映了一個事實：好的算法往往簡潔優雅。

## 13. 測試策略

這個測試策略覆蓋了各種情況：
- 標準情況（目標值存在多次）
- 目標值不存在
- 空數組
- 只有一個元素的數組
- 目標值在數組邊界（開頭或結尾）
- 所有元素都是目標值
- 大型數組（測試性能和穩定性）

每個測試用例都有清晰的描述，指明它測試的是什麼情況，並提供了預期的輸出結果。通過這種全面的測試，我們可以確保解決方案在各種情況下都能正確工作。