# 1. Original Problem:

## Search in Rotated Sorted Array

There is an integer array `nums` sorted in ascending order (with distinct values).

Prior to being passed to your function, `nums` is possibly rotated at an unknown pivot index `k` (1 <= k < nums.length) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (0-indexed). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index 3 and become `[4,5,6,7,0,1,2]`.

Given the array `nums` after the possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not in `nums`.

You must write an algorithm with O(log n) runtime complexity.

**Example 1:**
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

**Example 2:**
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

**Example 3:**
```
Input: nums = [1], target = 0
Output: -1
```

**Constraints:**
- 1 <= nums.length <= 5000
- -10^4 <= nums[i] <= 10^4
- All values of `nums` are unique.
- `nums` is an ascending array that is possibly rotated.
- -10^4 <= target <= 10^4

## 搜尋旋轉排序數組

有一個按升序排列的整數數組 `nums`（具有不同的值）。

在傳遞給你的函數之前，`nums` 可能在未知的某個樞紐索引 `k` (1 <= k < nums.length) 處進行了旋轉，使得數組變為 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（索引從 0 開始）。例如，`[0,1,2,4,5,6,7]` 可能在樞紐索引 3 處旋轉並變成 `[4,5,6,7,0,1,2]`。

給定旋轉後的數組 `nums` 和一個整數 `target`，如果 `target` 存在於 `nums` 中，則返回 `target` 的索引，否則返回 `-1`。

你必須設計一個時間複雜度為 O(log n) 的算法。

# 2. 問題理解

這個問題的核心要點是：
- 我們有一個本來是升序排列的數組，但它在某個位置被旋轉了
- 數組中的所有值都是唯一的（沒有重複元素）
- 我們需要在這個旋轉後的數組中找到目標值的索引
- 解法必須具有 O(log n) 的時間複雜度，這暗示我們需要使用二分查找的變形

關鍵挑戰在於：普通的二分查找依賴於完全排序的數組，但在這個問題中，數組被旋轉了，不再是完全排序的。然而，我們可以觀察到，即使數組被旋轉，在任何時候至少有一半的數組仍是排序的。這個特性可以幫助我們修改二分查找算法。

# 3. 視覺化解釋

讓我們使用例子 `nums = [4,5,6,7,0,1,2]` 來可視化旋轉排序數組：

原始排序數組：`[0,1,2,4,5,6,7]`  
在索引 3 處旋轉後：`[4,5,6,7,0,1,2]`

```
索引： 0 1 2 3 4 5 6
數值： 4 5 6 7 0 1 2
      ↑       ↑
     左半部分  右半部分
```

我們可以看到旋轉後的數組實際上由兩個排序子數組組成：
- 左半部分：`[4,5,6,7]`（排序的）
- 右半部分：`[0,1,2]`（排序的）
- 旋轉點位於索引 3 和 4 之間

這個觀察對我們的算法至關重要，因為在二分查找的每一步，我們需要確定哪一半是排序的，並檢查目標值是否在該排序區間內。

# 4. 思考過程

當面對這個問題時，我們可以思考幾種解決方案：

**方法1：線性搜索**
- 直接遍歷整個數組查找目標值
- 時間複雜度：O(n)
- 這不符合問題要求的 O(log n) 時間複雜度

**方法2：先找到旋轉點，再使用二分查找**
- 使用二分查找找到旋轉點
- 確定目標值在哪個子數組中
- 對該子數組使用標準二分查找
- 時間複雜度：O(log n) + O(log n) = O(log n)

**方法3：修改的二分查找（最優解）**
- 直接在旋轉數組上應用二分查找的變體
- 利用關鍵觀察：在任何時候，數組的一半必定是排序的
- 確定哪一半是排序的，並檢查目標值是否在該區間內
- 時間複雜度：O(log n)

我選擇方法3作為最優解，因為它更簡潔直接，不需要先找到旋轉點。

# 5. 最優解決方案開發

讓我們開發修改後的二分查找算法：

1. 初始化左指針 `left = 0`，右指針 `right = len(nums) - 1`
2. 當 `left <= right` 時執行循環：
   a. 計算中間索引 `mid = (left + right) / 2`
   b. 如果 `nums[mid] == target`，返回 `mid`
   c. 確定哪一半是有序的：
    - 如果左半部分有序（`nums[left] <= nums[mid]`）：
        * 如果目標值在左半部分的範圍內（`nums[left] <= target < nums[mid]`），搜索左半部分
        * 否則，搜索右半部分
    - 如果右半部分有序：
        * 如果目標值在右半部分的範圍內（`nums[mid] < target <= nums[right]`），搜索右半部分
        * 否則，搜索左半部分
3. 如果未找到目標值，返回 -1

這個算法的關鍵在於識別哪一半是有序的，然後決定在哪一半繼續搜索。

# 6. 實例演練

讓我們用 `nums = [4,5,6,7,0,1,2]` 和 `target = 0` 來演示算法執行過程：

**詳細步驟分析：**

**第一輪迭代：**
- 初始化：left = 0, right = 6
- 中間點 mid = (0 + 6) / 2 = 3
- nums[mid] = 7，不等於目標值 0
- 檢查左半部分是否有序：nums[left] <= nums[mid]? 4 <= 7? 是
- 目標值是否在左半部分範圍內：nums[left] <= target < nums[mid]? 4 <= 0 < 7? 否
- 因此，搜索右半部分：left = mid + 1 = 4, right = 6

**第二輪迭代：**
- 中間點 mid = (4 + 6) / 2 = 5
- nums[mid] = 1，不等於目標值 0
- 檢查左半部分是否有序：nums[left] <= nums[mid]? 0 <= 1? 是
- 目標值是否在左半部分範圍內：nums[left] <= target < nums[mid]? 0 <= 0 < 1? 是
- 因此，搜索左半部分：left = 4, right = mid - 1 = 4

**第三輪迭代：**
- 中間點 mid = (4 + 4) / 2 = 4
- nums[mid] = 0，等於目標值 0
- 找到目標！返回 mid = 4

這個例子展示了我們的算法如何在旋轉排序數組中有效地找到目標值。關鍵在於每次迭代時，我們都能確定哪半部分是有序的，並據此決定搜索方向。

# 8. 實現執行詳解

我們繼續使用 `nums = [4, 5, 6, 7, 0, 1, 2]` 和 `target = 0` 逐行分析程式碼執行：

1. **初始化**：
   ```go
   left, right := 0, len(nums)-1  // left = 0, right = 6
   ```

2. **第一輪迭代**：
   ```go
   // 進入 for 循環，因為 left (0) <= right (6)
   mid := left + (right - left) / 2  // mid = 0 + (6 - 0) / 2 = 3
   // nums[mid] = nums[3] = 7 != target (0)
   
   // 判斷左半部分是否有序
   // nums[left] <= nums[mid]? nums[0] <= nums[3]? 4 <= 7? 是
   // 左半部分有序
   
   // 目標值在左半部分的有序區間內？
   // nums[left] <= target && target < nums[mid]? 4 <= 0 && 0 < 7? 否
   // 目標不在左半部分，搜索右半部分
   left = mid + 1  // left = 3 + 1 = 4
   // 現在 left = 4, right = 6
   ```

3. **第二輪迭代**：
   ```go
   // 進入 for 循環，因為 left (4) <= right (6)
   mid := left + (right - left) / 2  // mid = 4 + (6 - 4) / 2 = 5
   // nums[mid] = nums[5] = 1 != target (0)
   
   // 判斷左半部分是否有序
   // nums[left] <= nums[mid]? nums[4] <= nums[5]? 0 <= 1? 是
   // 左半部分有序
   
   // 目標值在左半部分的有序區間內？
   // nums[left] <= target && target < nums[mid]? 0 <= 0 && 0 < 1? 是
   // 目標在左半部分，搜索左半部分
   right = mid - 1  // right = 5 - 1 = 4
   // 現在 left = 4, right = 4
   ```

4. **第三輪迭代**：
   ```go
   // 進入 for 循環，因為 left (4) <= right (4)
   mid := left + (right - left) / 2  // mid = 4 + (4 - 4) / 2 = 4
   // nums[mid] = nums[4] = 0 == target (0)
   // 找到目標，返回 mid = 4
   ```

關鍵點：
- 我們使用 `mid := left + (right - left) / 2` 而不是 `mid := (left + right) / 2` 來避免整數溢出
- 在每一步中，我們都需要判斷哪一半是有序的，然後決定搜索方向
- 算法的核心在於有效縮小搜索範圍，每次迭代都能排除剩餘數組的一半

# 9. 複雜度分析

**時間複雜度：O(log n)**
- 二分查找算法每次迭代將搜索區間縮小一半
- 即使在旋轉排序數組上，我們仍然通過判斷哪一半有序來排除一半的搜索空間
- 最壞情況下，需要約 log₂n 次迭代才能找到目標或確定目標不存在

**空間複雜度：O(1)**
- 我們只使用了常數個變量（left, right, mid）
- 無論輸入數組大小如何，使用的額外空間都是固定的

# 10. 優化與改進

當前解決方案已經達到了最優的時間複雜度 O(log n)，這對於搜索問題來說是最佳的。不過，我們可以考慮一些情況進行小的優化：

1. **提前檢查邊界情況**：
    - 如果數組為空，直接返回 -1
    - 如果數組只有一個元素，直接檢查是否等於目標值

2. **處理數組未旋轉的情況**：
    - 如果 nums[left] < nums[right]，則數組已經排序，可以直接使用標準二分查找

3. **處理重複元素**（雖然題目說明不會有重複，但對於拓展是有用的）：
    - 當 nums[left] == nums[mid] 時，我們可能無法確定哪一半是有序的
    - 一個解決方案是將 left 遞增，即 left++，然後繼續二分查找

這些優化不會改變算法的時間複雜度，但可能在某些特定輸入下提高效率。

# 11. 測試策略

這個測試策略涵蓋了多種情況：
- 基本示例（包括目標存在和不存在的情況）
- 邊界情況（單元素數組）
- 特殊情況（未旋轉數組）
- 目標位於不同位置（旋轉點左側、右側）
- 極端旋轉情況（只旋轉一次）

通過這些測試，我們可以確保我們的算法在各種情況下都能正確工作。