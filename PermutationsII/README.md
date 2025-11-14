# Permutations II (LeetCode 47) 全面分析與解決方案

## 1. 原始問題

**English:**
Given a collection of numbers, `nums`, that might contain duplicates, return all possible unique permutations in any order.

**Examples:**

Example 1:
```
Input: nums = [1,1,2]
Output: [[1,1,2],[1,2,1],[2,1,1]]
```

Example 2:
```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**Constraints:**
- 1 <= nums.length <= 8
- -10 <= nums[i] <= 10

**繁體中文:**
給定一個可能包含重複數字的集合 `nums`，以任意順序返回所有可能的唯一排列。

**示例:**

示例 1:
```
輸入: nums = [1,1,2]
輸出: [[1,1,2],[1,2,1],[2,1,1]]
```

示例 2:
```
輸入: nums = [1,2,3]
輸出: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**約束條件:**
- 1 <= nums.length <= 8
- -10 <= nums[i] <= 10

## 2. 初始問題解析與心智建模

看到這個問題時，我們首先需要理解幾個關鍵點：

- 這是一個排列問題，我們需要找出所有可能的排列方式
- 陣列中可能包含重複元素
- 我們只需要返回唯一的排列，不應該有重複的排列

從示例1可以看出，當輸入包含重複元素（如 [1,1,2]）時，我們需要處理這些重複，確保不會生成重複的排列。如果我們單純地進行排列而不處理重複問題，像 [1,1,2] 這樣的輸入會生成相同的排列多次。

輸入：陣列 `nums`，可能包含重複元素
輸出：包含所有唯一排列的二維陣列
轉換要求：生成所有可能的排列，但要過濾掉重複的排列

## 3. 問題理解與核心挑戰

這個問題的核心挑戰在於：
1. 如何有效地生成所有可能的排列
2. 如何處理重複元素，確保我們不會產生重複的排列

排列問題通常可以通過回溯法（Backtracking）來解決。基本思路是：選擇一個元素放在當前位置，然後遞迴地排列剩餘的元素。

處理重複元素的關鍵是：
- 在相同層級的搜索中，相同的數字不應該被重複使用在相同的位置
- 這不同於「排列 I」問題，在那個問題中，所有元素都是唯一的

邊緣情況分析：
- 空陣列：根據約束條件，nums.length 至少為 1，所以不需要處理空陣列
- 單元素陣列：返回該元素本身的排列
- 全部元素相同的陣列（如 [1,1,1]）：只有一種唯一排列

## 4. 視覺問題表示

讓我們以 `nums = [1,1,2]` 為例，視覺化展示排列的過程：

```
                    []
           /         |         \
         [1]        [1]*       [2]
       /    \       /   \      /  \
    [1,1]  [1,2]  [1,1]* [1,2]* [2,1]
     |       |      |      |      |
  [1,1,2] [1,2,1] [1,1,2]* [1,2,1]* [2,1,1]
```

注意：標記 * 的路徑代表在同一層重複選擇相同的數字，這些是我們需要避免的。

當我們在同一層選擇元素時，如果一個值已經被選擇過，就應該跳過相同的值。例如，在第一層選擇時，有兩個 1，但我們只應該選擇其中一個，否則會導致重複的排列。

## 5. 問題模式辨識

這是一個典型的「帶有重複元素的排列」問題，屬於回溯法（Backtracking）的範疇。回溯法是解決排列、組合等問題的常用策略。

識別特徵：
- 需要列舉所有可能的排列（組合）
- 有重複元素需要處理
- 結果不能有重複

類似的問題模式包括：
- 排列 I（無重複元素）
- 組合問題
- 子集問題

解決這類問題的關鍵在於設計一個有效的機制來避免重複。常見的策略包括：
- 排序+跳過重複
- 使用標記數組或哈希表記錄已使用的元素

## 6. 策略性問題解決框架

對於帶有重複元素的排列問題，我們可以採用以下框架：

1. **排序**：首先對輸入數組進行排序，這樣相同的元素會相鄰
2. **回溯**：使用回溯法生成所有排列
3. **剪枝**：在同一層的選擇中，跳過重複的元素

回溯法的一般框架：
```
func backtrack(路徑, 選擇列表):
    if 滿足結束條件:
        結果.add(路徑)
        return
    
    for 選擇 in 選擇列表:
        if 無效選擇:
            continue
        做選擇
        backtrack(路徑, 選擇列表)
        撤銷選擇
```

對於這個特定問題，我們需要添加一個條件來避免在同一層選擇重複的元素。

## 7. 演算法設計（編碼前）

我們的算法如下：

1. 將輸入數組 `nums` 排序，使相同的元素相鄰
2. 使用一個布爾數組 `used` 來標記哪些元素已經在當前路徑中使用
3. 使用回溯法生成所有排列：
    - 對於每個位置，嘗試放置一個未使用的元素
    - 如果當前元素與前一個元素相同，且前一個元素未被使用，則跳過當前元素（這是避免重複的關鍵）
4. 收集所有有效的排列

讓我們用偽代碼描述這個算法：

```
function permuteUnique(nums):
    sort(nums)  // 排序，使相同元素相鄰
    result = []
    used = [false, false, ..., false]  // 長度為 nums 的長度
    
    backtrack([], nums, used, result)
    
    return result

function backtrack(path, nums, used, result):
    if path.length == nums.length:
        result.append(path.copy())
        return
    
    for i = 0 to nums.length - 1:
        // 剪枝條件：當前元素與前一個元素相同，且前一個元素未被使用
        if used[i] or (i > 0 and nums[i] == nums[i-1] and !used[i-1]):
            continue
        
        used[i] = true
        path.append(nums[i])
        
        backtrack(path, nums, used, result)
        
        path.pop()
        used[i] = false
```

## 8. 視覺說明

使用示例 `nums = [1,1,2]` 來視覺化算法的執行過程：

排序後：`nums = [1,1,2]`（已經排序）

回溯過程：
```
初始狀態: path=[], used=[F,F,F]

1. 選擇 nums[0]=1:
   path=[1], used=[T,F,F]
   
   1.1 選擇 nums[1]=1:
       path=[1,1], used=[T,T,F]
       
       1.1.1 選擇 nums[2]=2:
             path=[1,1,2], used=[T,T,T]
             結果: [[1,1,2]]
             
   1.2 選擇 nums[2]=2:
       path=[1,2], used=[T,F,T]
       
       1.2.1 選擇 nums[1]=1:
             path=[1,2,1], used=[T,T,T]
             結果: [[1,1,2], [1,2,1]]

2. 選擇 nums[1]=1:
   因為 nums[1]==nums[0] 且 !used[0]，跳過（剪枝）

3. 選擇 nums[2]=2:
   path=[2], used=[F,F,T]
   
   3.1 選擇 nums[0]=1:
       path=[2,1], used=[T,F,T]
       
       3.1.1 選擇 nums[1]=1:
             path=[2,1,1], used=[T,T,T]
             結果: [[1,1,2], [1,2,1], [2,1,1]]
   
   3.2 選擇 nums[1]=1:
       因為 nums[1]==nums[0] 且 !used[0]，跳過（剪枝）
```

最終結果：`[[1,1,2], [1,2,1], [2,1,1]]`

## 9. 解決方案發展歷程

### 初始解法：暴力法

一個初始的方法是生成所有可能的排列，然後使用集合去重：

```golang
func permuteUnique(nums []int) [][]int {
    var result [][]int
    var path []int
    used := make([]bool, len(nums))
    
    backtrack(nums, used, path, &result)
    
    // 去重（效率低）
    uniqueResult := make(map[string][]int)
    for _, perm := range result {
        key := fmt.Sprintf("%v", perm)
        uniqueResult[key] = perm
    }
    
    result = [][]int{}
    for _, perm := range uniqueResult {
        result = append(result, perm)
    }
    
    return result
}

func backtrack(nums []int, used []bool, path []int, result *[][]int) {
    if len(path) == len(nums) {
        pathCopy := make([]int, len(path))
        copy(pathCopy, path)
        *result = append(*result, pathCopy)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }
        
        used[i] = true
        path = append(path, nums[i])
        
        backtrack(nums, used, path, result)
        
        path = path[:len(path)-1]
        used[i] = false
    }
}
```

這種方法的問題在於：
- 時間複雜度高，需要生成所有排列然後去重
- 空間效率低，需要存儲中間結果

### 優化解法：排序 + 剪枝

更高效的方法是在回溯過程中直接避免生成重複的排列：

```golang
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)  // 排序，使相同元素相鄰
    
    var result [][]int
    var path []int
    used := make([]bool, len(nums))
    
    backtrack(nums, used, path, &result)
    
    return result
}

func backtrack(nums []int, used []bool, path []int, result *[][]int) {
    if len(path) == len(nums) {
        pathCopy := make([]int, len(path))
        copy(pathCopy, path)
        *result = append(*result, pathCopy)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        // 剪枝條件：當前元素與前一個元素相同，且前一個元素未被使用
        if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
            continue
        }
        
        used[i] = true
        path = append(path, nums[i])
        
        backtrack(nums, used, path, result)
        
        path = path[:len(path)-1]
        used[i] = false
    }
}
```

這種方法的關鍵在於剪枝條件：`if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1])`。

當我們在同一層選擇元素時，如果當前元素與前一個元素相同，且前一個元素未被使用（意味著前一個元素已經被撤銷選擇），則跳過當前元素。這確保了在每一層，相同的數字只會被選擇一次。

## 10. 實用示例演練

讓我們以 `nums = [1,1,2]` 為例，詳細追蹤算法的執行過程：

初始狀態：
- 排序後：`nums = [1,1,2]`（已經排序）
- `path = []`
- `used = [false, false, false]`
- `result = []`

步驟 1：開始回溯
選擇 index 0 (值為1)：
- `path = [1]`
- `used = [true, false, false]`
- 遞迴

  步驟 1.1：第二層回溯
  選擇 index 1 (值為1)：
    - `path = [1,1]`
    - `used = [true, true, false]`
    - 遞迴

      步驟 1.1.1：第三層回溯
      選擇 index 2 (值為2)：
        - `path = [1,1,2]`
        - `used = [true, true, true]`
        - 路徑長度等於 nums 長度，添加到結果：`result = [[1,1,2]]`
        - 回溯，撤銷選擇

  步驟 1.2：嘗試 index 2 (值為2)
    - `path = [1,2]`
    - `used = [true, false, true]`
    - 遞迴

      步驟 1.2.1：嘗試 index 1 (值為1)
        - `path = [1,2,1]`
        - `used = [true, true, true]`
        - 添加到結果：`result = [[1,1,2], [1,2,1]]`
        - 回溯

步驟 2：嘗試從 index 1 開始 (值為1)
- 檢查：`i > 0 && nums[i] == nums[i-1] && !used[i-1]`
- 此時 `i = 1, nums[1] = 1, nums[0] = 1, used[0] = false`
- 條件成立，跳過此元素（剪枝）

步驟 3：嘗試從 index 2 開始 (值為2)
- `path = [2]`
- `used = [false, false, true]`
- 遞迴

  步驟 3.1：選擇 index 0 (值為1)
    - `path = [2,1]`
    - `used = [true, false, true]`
    - 遞迴

      步驟 3.1.1：選擇 index 1 (值為1)
        - `path = [2,1,1]`
        - `used = [true, true, true]`
        - 添加到結果：`result = [[1,1,2], [1,2,1], [2,1,1]]`
        - 回溯

  步驟 3.2：嘗試 index 1 (值為1)
    - 檢查：`i > 0 && nums[i] == nums[i-1] && !used[i-1]`
    - 此時 `i = 1, nums[1] = 1, nums[0] = 1, used[0] = true`
    - 條件不成立，可以選擇
    - 但我們已經在步驟 3.1.1 中處理過這種情況，這裡不再重複

最終結果：`[[1,1,2], [1,2,1], [2,1,1]]`

## 11. Golang 實現

```golang
func permuteUnique(nums []int) [][]int {
    // 首先對數組進行排序，使相同的元素相鄰
    sort.Ints(nums)
    
    result := [][]int{}
    path := []int{}
    used := make([]bool, len(nums))
    
    // 啟動回溯過程
    backtrack(nums, used, path, &result)
    
    return result
}

// 回溯函數
func backtrack(nums []int, used []bool, path []int, result *[][]int) {
    // 如果當前路徑長度等於nums長度，說明找到了一個完整排列
    if len(path) == len(nums) {
        // 創建path的副本並加入結果集
        pathCopy := make([]int, len(path))
        copy(pathCopy, path)
        *result = append(*result, pathCopy)
        return
    }
    
    // 遍歷所有可能的選擇
    for i := 0; i < len(nums); i++ {
        // 當前元素已經使用過，跳過
        if used[i] {
            continue
        }
        
        // 剪枝條件：當前元素與前一個元素相同，且前一個元素未被使用
        // 這確保在同一層不會重複選擇相同的元素
        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        
        // 做選擇
        used[i] = true
        path = append(path, nums[i])
        
        // 遞迴
        backtrack(nums, used, path, result)
        
        // 撤銷選擇（回溯）
        path = path[:len(path)-1]
        used[i] = false
    }
}
```

## 12. 實現執行演練

讓我們追蹤代碼對於輸入 `nums = [1,1,2]` 的執行過程：

1. 首先，對 `nums` 進行排序：`nums = [1,1,2]`（已經排序）
2. 初始化：`result = []`, `path = []`, `used = [false, false, false]`
3. 調用 `backtrack(nums, used, path, &result)`

**第一次調用 backtrack**:
- `path = []`, `used = [false, false, false]`
- 進入 for 循環，i = 0
- 檢查：`used[0]` 為 false，繼續
- 檢查：i = 0，不需要檢查重複
- 標記 `used[0] = true`，添加 `path = [1]`
- 遞迴調用 `backtrack(nums, used, path, &result)`

**第二次調用 backtrack**:
- `path = [1]`, `used = [true, false, false]`
- 進入 for 循環，i = 1
- 檢查：`used[1]` 為 false，繼續
- 檢查：`i > 0 && nums[1] == nums[0] && !used[0]`，結果為 false（因為 `used[0] = true`）
- 標記 `used[1] = true`，添加 `path = [1,1]`
- 遞迴調用 `backtrack(nums, used, path, &result)`

**第三次調用 backtrack**:
- `path = [1,1]`, `used = [true, true, false]`
- 進入 for 循環，i = 2
- 檢查：`used[2]` 為 false，繼續
- 標記 `used[2] = true`，添加 `path = [1,1,2]`
- 遞迴調用 `backtrack(nums, used, path, &result)`

**第四次調用 backtrack**:
- `path = [1,1,2]`, `used = [true, true, true]`
- `len(path) == len(nums)`，將 `path` 添加到 `result`：`result = [[1,1,2]]`
- 返回上一層

繼續回溯過程，直到遍歷完所有可能的選擇。

最終結果：`[[1,1,2], [1,2,1], [2,1,1]]`

## 13. 複雜度分析

### 時間複雜度

- 排序需要 O(n log n) 時間，其中 n 是數組長度
- 回溯過程最多生成 n! 個排列（每個位置有 n 種選擇）
- 但由於剪枝，實際生成的排列數量少於 n!
- 對於包含重複元素的情況，時間複雜度為 O(n! / (c1! * c2! * ... * ck!))，其中 ci 是第 i 個不同元素的出現次數
- 總時間複雜度：O(n log n + n!)

### 空間複雜度

- 排序可能需要 O(log n) 的額外空間（取決於排序算法）
- 遞迴調用棧的深度為 O(n)
- 用於存儲結果的空間為 O(n * n!)（n! 個排列，每個長度為 n）
- `used` 數組需要 O(n) 空間
- `path` 數組在最壞情況下需要 O(n) 空間
- 總空間複雜度：O(n * n!)

## 14. 優化和改進

我們的解決方案已經相當高效，但仍有一些可能的改進方向：

1. **使用計數器而非布爾數組**：對於包含大量重複元素的輸入，可以使用計數器來記錄每個元素的出現次數，而不是布爾數組標記已使用的元素。這種方法在某些情況下可能更高效。

```golang
func permuteUnique(nums []int) [][]int {
    // 使用 map 記錄每個元素的出現次數
    counter := make(map[int]int)
    for _, num := range nums {
        counter[num]++
    }
    
    result := [][]int{}
    path := []int{}
    
    // 提取不同的數字和它們的出現次數
    var uniqueNums []int
    var counts []int
    for num, count := range counter {
        uniqueNums = append(uniqueNums, num)
        counts = append(counts, count)
    }
    
    backtrackWithCounter(uniqueNums, counts, len(nums), 0, path, &result)
    
    return result
}

func backtrackWithCounter(nums []int, counts []int, n, pos int, path []int, result *[][]int) {
    if pos == n {
        pathCopy := make([]int, len(path))
        copy(pathCopy, path)
        *result = append(*result, pathCopy)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if counts[i] == 0 {
            continue
        }
        
        counts[i]--
        path = append(path, nums[i])
        
        backtrackWithCounter(nums, counts, n, pos+1, path, result)
        
        path = path[:len(path)-1]
        counts[i]++
    }
}
```

2. **迭代而非遞迴**：在某些情況下，使用迭代方法可能比遞迴更高效，但實現會更複雜。

## 15. 一般問題解決智慧

從「Permutations II」問題中，我們可以學到以下關鍵經驗：

1. **排序預處理**：對於需要處理重複元素的問題，排序是一個常用的預處理步驟，它可以使相同的元素相鄰，便於識別和處理重複。

2. **回溯法的剪枝**：在回溯法中，合理的剪枝策略可以顯著提高算法效率。對於排列問題，跳過在同一層重複選擇相同元素是一個關鍵的剪枝策略。

3. **狀態追蹤**：使用 `used` 數組或計數器來追蹤元素的使用狀態是解決排列、組合類問題的常用技巧。

4. **遞迴與回溯的關係**：回溯法本質上是一種深度優先搜索（DFS）的應用，通常通過遞迴實現。理解遞迴調用的執行過程和狀態變化是掌握回溯法的關鍵。

類似問題的解決思路：

- **排列問題**：使用回溯法，通過標記數組追蹤已使用的元素
- **組合問題**：使用回溯法，通常從固定起點開始，避免重複組合
- **子集問題**：可以看作是每個元素「選或不選」的決策樹問題

## 16. 測試策略

這個測試套件涵蓋了各種情況：
- 標準測試用例（示例1和示例2）
- 邊界情況（單元素數組）
- 特殊情況（全部元素相同）
- 包含負數的情況
- 大型輸入的性能測試