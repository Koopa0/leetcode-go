## 1. Original Problem

**Remove Element**

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` in-place. The order of the elements may be changed. Then return the number of elements in `nums` which are not equal to `val`.

Consider the number of elements in `nums` which are not equal to `val` be `k`, to get accepted, you need to do the following things:
- Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

**Custom Judge:**

The judge will test your solution with the following code:
```
int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                        // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
If all assertions pass, then your solution will be accepted.

**Example 1:**
```
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

**Example 2:**
```
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

**Constraints:**
- 0 <= nums.length <= 100
- 0 <= nums[i] <= 50
- 0 <= val <= 100

**移除元素**

給定一個整數數組 `nums` 和一個整數 `val`，要求您移除數組中所有等於 `val` 的元素，並在原地完成。元素的順序可以改變。然後返回 `nums` 中不等於 `val` 的元素數量。

假設 `nums` 中不等於 `val` 的元素數量為 `k`，為了獲得通過，您需要執行以下操作：
- 修改數組 `nums`，使 `nums` 的前 `k` 個元素包含不等於 `val` 的元素。`nums` 的剩餘元素及數組大小不重要。
- 返回 `k`。

**自定義評判：**

評判將使用以下代碼測試您的解決方案：
```
int[] nums = [...]; // 輸入數組
int val = ...; // 需要移除的值
int[] expectedNums = [...]; // 預期答案，具有正確的長度。
                         // 它已排序，且不包含值等於 val 的元素。

int k = removeElement(nums, val); // 調用您的實現

assert k == expectedNums.length;
sort(nums, 0, k); // 對 nums 的前 k 個元素進行排序
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
如果所有斷言都通過，則您的解決方案將被接受。

**示例 1：**
```
輸入：nums = [3,2,2,3], val = 3
輸出：2, nums = [2,2,_,_]
解釋：您的函數應該返回 k = 2，nums 的前兩個元素為 2。
返回的 k 之後的元素不重要（因此它們是下劃線）。
```

**示例 2：**
```
輸入：nums = [0,1,2,2,3,0,4,2], val = 2
輸出：5, nums = [0,1,4,0,3,_,_,_]
解釋：您的函數應該返回 k = 5，nums 的前五個元素包含 0, 0, 1, 3 和 4。
請注意，這五個元素可以以任何順序返回。
返回的 k 之後的元素不重要（因此它們是下劃線）。
```

**約束條件：**
- 0 <= nums.length <= 100
- 0 <= nums[i] <= 50
- 0 <= val <= 100

## 2. 問題理解

這個問題的核心要求是：
- 移除數組中所有等於特定值 `val` 的元素
- 必須在原地（in-place）完成，不能使用額外的數組空間
- 返回處理後數組中不等於 `val` 的元素數量
- 數組中前 k 個位置必須包含所有不等於 `val` 的元素
- 剩餘位置的元素內容不重要

值得注意的是：
- 元素的順序可以改變，這給了我們更多的靈活性
- 題目的測試方式是：確保返回的 k 值正確，並檢查數組前 k 個元素（排序後）是否與預期相符
- 數組大小範圍為 0 至 100，元素值範圍為 0 至 50
- 需要移除的值範圍為 0 至 100

邊界情況包括：
- 空數組（nums.length = 0）
- 所有元素都等於 val（返回 k = 0）
- 沒有元素等於 val（返回 k = nums.length）

## 3. 視覺解釋

讓我們通過視覺方式來說明算法的執行流程。以示例 2 為例：

初始數組：[0, 1, 2, 2, 3, 0, 4, 2]，要移除的值 val = 2

我們可以使用兩個指針的方法來解決這個問題：
- 一個指針 i 用於遍歷數組
- 一個指針 k 用於指向下一個應該被填充的位置

```
初始狀態：
[0, 1, 2, 2, 3, 0, 4, 2]
 k  i
 
檢查 nums[i=0] = 0 != val(2)，將其保留在位置 k
[0, 1, 2, 2, 3, 0, 4, 2]
    k  i

檢查 nums[i=1] = 1 != val(2)，將其保留在位置 k
[0, 1, 2, 2, 3, 0, 4, 2]
       k  i

檢查 nums[i=2] = 2 == val(2)，跳過
[0, 1, 2, 2, 3, 0, 4, 2]
       k     i

檢查 nums[i=3] = 2 == val(2)，跳過
[0, 1, 2, 2, 3, 0, 4, 2]
       k        i

檢查 nums[i=4] = 3 != val(2)，將其放到位置 k
[0, 1, 3, 2, 3, 0, 4, 2]
          k        i

檢查 nums[i=5] = 0 != val(2)，將其放到位置 k
[0, 1, 3, 0, 3, 0, 4, 2]
             k        i

檢查 nums[i=6] = 4 != val(2)，將其放到位置 k
[0, 1, 3, 0, 4, 0, 4, 2]
                k        i

檢查 nums[i=7] = 2 == val(2)，跳過
[0, 1, 3, 0, 4, 0, 4, 2]
                k           i

最終結果：
[0, 1, 3, 0, 4, ?, ?, ?] （前 k=5 個元素是不等於 val 的元素）
```

## 4. 思考過程

對於這個問題，我們可以考慮幾種不同的解決方案：

**方法 1：雙指針（快慢指針）**
- 使用兩個指針：一個慢指針指向當前要填充的位置，一個快指針用於遍歷數組
- 當快指針遇到不等於 val 的元素時，將其複製到慢指針位置，然後慢指針向前移動
- 這種方法保持了除 val 之外所有元素的相對順序

**方法 2：雙向指針**
- 使用左右兩個指針，左指針從頭開始，右指針從尾部開始
- 當左指針遇到值為 val 的元素時，與右指針指向的元素交換，然後右指針向左移動
- 當左右指針相遇時，算法結束
- 這種方法不保持原始元素的相對順序，但通常需要較少的操作

**方法 3：計數並重寫**
- 先遍歷一次數組，統計不等於 val 的元素數量
- 再遍歷一次，將不等於 val 的元素填充到數組前面
- 這種方法需要兩次遍歷，效率較低

考慮到問題的要求（元素順序可以改變）和效率，我們可以選擇方法 1 或方法 2。方法 1 更為直觀，而且可以保持元素的相對順序（雖然題目沒有要求）。方法 2 在某些情況下可能需要較少的操作，特別是當 val 出現頻率較低時。

## 5. 最佳解決方案開發

讓我們選擇雙指針方法（方法 1）作為我們的解決方案，因為它簡單直觀且易於理解。

具體步驟如下：
1. 初始化一個慢指針 k = 0，用於指示下一個應該被填充的位置
2. 使用快指針 i 遍歷整個數組
3. 對於每個元素 nums[i]：
    - 如果 nums[i] != val，則將其複製到 nums[k]，然後 k++
    - 如果 nums[i] == val，則跳過
4. 遍歷結束後，數組的前 k 個元素就是所有不等於 val 的元素
5. 返回 k 作為結果

讓我們用示例 2 再次演示這個算法：
```
nums = [0,1,2,2,3,0,4,2], val = 2, k = 0

i=0: nums[0]=0 != val, 所以 nums[k]=nums[0], k++
[0,1,2,2,3,0,4,2], k=1

i=1: nums[1]=1 != val, 所以 nums[k]=nums[1], k++
[0,1,2,2,3,0,4,2], k=2

i=2: nums[2]=2 == val, 跳過
[0,1,2,2,3,0,4,2], k=2

i=3: nums[3]=2 == val, 跳過
[0,1,2,2,3,0,4,2], k=2

i=4: nums[4]=3 != val, 所以 nums[k]=nums[4], k++
[0,1,3,2,3,0,4,2], k=3

i=5: nums[5]=0 != val, 所以 nums[k]=nums[5], k++
[0,1,3,0,3,0,4,2], k=4

i=6: nums[6]=4 != val, 所以 nums[k]=nums[6], k++
[0,1,3,0,4,0,4,2], k=5

i=7: nums[7]=2 == val, 跳過
[0,1,3,0,4,0,4,2], k=5

最終結果：k=5, nums 的前 5 個元素為 [0,1,3,0,4]
```

這個算法的時間複雜度為 O(n)，其中 n 是數組長度，因為我們只需要遍歷一次數組。空間複雜度為 O(1)，因為我們是原地修改數組，沒有使用額外的空間。

## 6. Golang 實現

```go
func removeElement(nums []int, val int) int {
    // 初始化慢指針，用於指示下一個應該被填充的位置
    k := 0
    
    // 使用快指針遍歷數組
    for i := 0; i < len(nums); i++ {
        // 如果當前元素不等於 val，則保留該元素
        if nums[i] != val {
            // 將該元素複製到 k 位置，然後 k 向前移動
            nums[k] = nums[i]
            k++
        }
        // 如果當前元素等於 val，則跳過（不做任何操作）
    }
    
    // 返回不等於 val 的元素數量
    return k
}
```

## 7. 複雜度分析

**時間複雜度：**
- 最佳情況：O(n)，即使數組中沒有等於 val 的元素，我們仍然需要遍歷整個數組
- 平均情況：O(n)，我們需要遍歷整個數組一次
- 最壞情況：O(n)，即使數組中所有元素都等於 val，我們仍然需要遍歷整個數組

這裡的 n 是指數組的長度。無論數組中有多少元素等於 val，我們都需要遍歷整個數組一次，所以時間複雜度始終是 O(n)。

**空間複雜度：**
- O(1)，我們只使用了常數級別的額外空間（幾個指針變量）
- 所有操作都是原地完成的，沒有使用額外的數組空間

## 8. 優化與改進

我們使用的雙指針解法已經是最優的解法之一，因為：
- 時間複雜度已經是最優的 O(n)
- 空間複雜度已經是最優的 O(1)
- 算法簡單直觀，容易實現和理解

如果題目要求保持元素的相對順序，那麼我們的方法 1 正好滿足要求。如果不需要保持順序，我們也可以考慮使用方法 2（雙向指針），特別是在 val 出現頻率較低的情況下，它可能需要較少的元素移動操作。

雙向指針的實現：

```go
func removeElementAlt(nums []int, val int) int {
    left, right := 0, len(nums) - 1
    
    for left <= right {
        if nums[left] == val {
            // 將左指針指向的元素與右指針指向的元素交換
            nums[left] = nums[right]
            right--
        } else {
            // 左指針指向的元素不等於 val，向右移動
            left++
        }
    }
    
    return left
}
```

這種方法的優勢在於：當 val 在數組中出現頻率較低時，需要的交換操作更少。但缺點是不保持元素的原始順序。

## 9. 測試策略

這個測試包含了以下幾類測試用例：
1. 標準用例：LeetCode 提供的示例
2. 邊界用例：空數組、全部元素等於 val、沒有元素等於 val
3. 特殊用例：單一元素數組

測試策略確保了解決方案在各種情況下的正確性和穩健性。測試步驟遵循了 LeetCode 自定義評判的邏輯：
1. 調用 `removeElement` 函數
2. 檢查返回值 k 是否正確
3. 對數組前 k 個元素排序
4. 驗證排序後的前 k 個元素是否與預期相符