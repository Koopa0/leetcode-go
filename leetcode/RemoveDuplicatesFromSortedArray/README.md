## 1. Original Problem

### English Version:
Given an integer array `nums` sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in `nums`.

Consider the number of unique elements of `nums` to be `k`, to get accepted, you need to do the following things:
- Change the array `nums` such that the first `k` elements of `nums` contain the unique elements in the order they were present in `nums` initially. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

Custom Judge:

The judge will test your solution with the following code:
```
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
If all assertions pass, then your solution will be accepted.

### 繁體中文版本：
給你一個 非遞減排序 的整數數組 `nums` ，請你 原地 刪除重複出現的元素，使每個元素 只出現一次 ，返回刪除後數組的新長度。元素的 相對順序 應該保持 一致 。然後返回 `nums` 中唯一元素的數量。

考慮 `nums` 的唯一元素的數量為 `k` ，你需要做以下事情得到通過：
- 更改數組 `nums` ，使 `nums` 的前 `k` 個元素包含唯一元素，並按照它們最初在 `nums` 中出現的順序排列。`nums` 的其餘元素與 `nums` 的大小不重要。
- 返回 `k` 。

判題標準：

系統會用下面的程式碼來測試你的解：
```
int[] nums = [...]; // 輸入數組
int[] expectedNums = [...]; // 長度正確的期望答案

int k = removeDuplicates(nums); // 調用你的實現

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
如果所有斷言都通過，那麼你的解決方案將被 通過。

**Example 1:**
```
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

**Example 2:**
```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

**Constraints:**
- `1 <= nums.length <= 3 * 10^4`
- `-100 <= nums[i] <= 100`
- `nums` is sorted in non-decreasing order.

## 2. 問題理解

這道題目的核心要求是：
- 給定一個已排序（非遞減）的整數數組
- 需要原地刪除重複元素
- 保持元素的相對順序不變
- 返回去重後的數組長度
- 前 k 個元素必須是去重後的元素，且順序與原數組一致

關鍵條件：
- 必須是原地操作（in-place），不能複製到新數組
- 數組已經是排序的，這是一個重要的提示
- 關心的是前 k 個元素，其餘元素不重要

邊界情況：
- 空數組？題目約束條件說明長度至少為 1
- 全部元素都相同？例如 [1,1,1,1]
- 沒有重複元素？例如 [1,2,3,4]

主要難點：
- 如何在不使用額外空間的情況下原地移除重複元素
- 如何高效地處理已排序數組的特性

## 3. 視覺解釋

我們來看一個例子來理解算法流程：

對於數組 `[0,0,1,1,1,2,2,3,3,4]`

我們使用兩個指針的方法：
1. 一個慢指針 `slow` 指向當前已處理的唯一元素的最後位置
2. 一個快指針 `fast` 用於遍歷整個數組

```
初始狀態：
[0,0,1,1,1,2,2,3,3,4]
 s
 f

快指針向前移動，發現 nums[1] 與 nums[0] 相同，繼續移動：
[0,0,1,1,1,2,2,3,3,4]
 s
   f

發現 nums[2] != nums[slow]，因此 slow 前進一位，然後將 nums[fast] 複製到 nums[slow]：
[0,1,1,1,1,2,2,3,3,4]
   s
     f

繼續移動 fast，發現 nums[3] == nums[slow]，繼續移動：
[0,1,1,1,1,2,2,3,3,4]
   s
       f

同樣 nums[4] == nums[slow]，繼續移動：
[0,1,1,1,1,2,2,3,3,4]
   s
         f

發現 nums[5] != nums[slow]，此時 slow 前進一位，將 nums[fast] 複製到 nums[slow]：
[0,1,2,1,1,2,2,3,3,4]
     s
         f

...依此類推，最終結果為：
[0,1,2,3,4,2,2,3,3,4]
         s
                   f
```

最終 slow+1 的值為 5，表示有 5 個唯一元素。

## 4. 思考過程

針對這個問題，我們可以考慮幾種不同的解決方案：

**方法 1：使用額外空間**
- 創建一個新數組來存儲唯一元素
- 遍歷原數組，若元素不在新數組中，則添加到新數組
- 將新數組的元素複製回原數組

缺點：使用了 O(n) 的額外空間，不符合原地修改的要求。

**方法 2：雙指針法**
- 使用一個慢指針 `slow` 指向當前已處理的不重複元素的最後位置
- 使用一個快指針 `fast` 遍歷整個數組
- 當 `nums[fast] != nums[slow]` 時，將 `slow` 指針前進一位，然後將 `nums[fast]` 的值賦給 `nums[slow]`

優點：
- 原地操作，不需要額外空間
- 適合已排序數組的情況
- 時間複雜度為 O(n)，只需一次遍歷

**方法 3：計數法**
- 遍歷數組，記錄每個元素出現的次數
- 再次遍歷數組，僅保留每個元素的第一個出現

缺點：需要兩次遍歷，且仍需要額外空間來記錄計數。

對於這個問題，方法 2 的雙指針法是最優解。它利用了數組已排序的特性，只需一次遍歷就可以完成，且是原地操作。

## 5. 最優解決方案開發

讓我們逐步完善雙指針解法：

1. 特殊情況處理：
    - 如果數組為空或只有一個元素，可以直接返回數組長度

2. 初始化兩個指針：
    - 慢指針 `slow` 初始為 0
    - 快指針 `fast` 初始為 1

3. 遍歷數組：
    - 當 `fast` 指針指向的元素與 `slow` 指針指向的元素不同時：
        - 將 `slow` 指針向前移動一位
        - 將 `fast` 指針指向的元素值賦給 `slow` 指針指向的位置
    - 無論如何，`fast` 指針都向前移動一位

4. 返回 `slow + 1`，表示唯一元素的數量

以 `[0,0,1,1,1,2,2,3,3,4]` 為例，具體步驟為：

- 初始狀態：`slow = 0`, `fast = 1`
- `nums[0] == nums[1]`，即都是 0，因此只移動 `fast = 2`
- `nums[0] != nums[2]`，即 0 != 1，所以 `slow = 1`，將 `nums[2]` 複製到 `nums[1]`，然後 `fast = 3`
- 持續這個過程...

最終數組變為 `[0,1,2,3,4,2,2,3,3,4]`，返回 `slow + 1 = 5`。

## 7. 複雜度分析

**時間複雜度**：
- 最好情況：O(n)，當數組中沒有重複元素時，每個元素只需處理一次
- 平均情況：O(n)，需要遍歷整個數組一次
- 最壞情況：O(n)，當數組中所有元素都相同時，仍然需要遍歷整個數組一次

**空間複雜度**：
- O(1)，只使用了兩個指針變量，不需要額外空間
- 這符合題目要求的原地操作

## 8. 優化與改進

這個解決方案已經是時間和空間複雜度上的最優解，主要是因為：
1. 已經實現了 O(n) 的時間複雜度，這是無法再優化的，因為至少需要遍歷一次數組
2. 已經實現了 O(1) 的空間複雜度，使用了原地操作

其他可能的考慮點：
- 對於特別大的數組，可以考慮並行處理，但這超出了題目的範圍
- 在實際應用中，如果不需要保持元素的相對順序，可以考慮使用更高效的方法，如排序或哈希表

相關題目：
- "Remove Element"
- "Move Zeroes"
- "Remove Duplicates from Sorted Array II"（允許每個元素最多出現兩次）

## 9. 測試策略

這個測試覆蓋了以下情況：
1. 題目提供的示例
2. 沒有重複元素的數組
3. 所有元素都相同的數組
4. 只有一個元素的數組
5. 包含邊界值的數組 (-100 和 100)

每個測試都會檢查：
- 返回值 k 是否正確
- 數組前 k 個元素是否符合預期