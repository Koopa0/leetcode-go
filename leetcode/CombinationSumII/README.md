# LeetCode 問題解析：Combination Sum II

## 1. Original Problem

### English Version:
**40. Combination Sum II**

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used **once** in the combination.

**Note**: The solution set must not contain duplicate combinations.

**Example 1:**
```
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: 
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

**Example 2:**
```
Input: candidates = [2,5,2,1,2], target = 5
Output: 
[
[1,2,2],
[5]
]
```

**Constraints:**
- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

### 繁體中文版：
**40. 組合總和 II**

給定一個候選數字的集合（`candidates`）和一個目標數字（`target`），找出 `candidates` 中所有可以使數字和為 `target` 的組合。

`candidates` 中的每個數字在每個組合中只能使用**一次**。

**注意**：解集不能包含重複的組合。

**示例 1：**
```
輸入: candidates = [10,1,2,7,6,1,5], target = 8
輸出: 
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

**示例 2：**
```
輸入: candidates = [2,5,2,1,2], target = 5
輸出: 
[
[1,2,2],
[5]
]
```

**約束條件：**
- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## 2. 初始問題解析和心智建模

當我第一次看到這個問題時，我會這樣分析：

1. **問題類型識別**：這是一個組合問題，目標是找出所有總和為特定值的數字組合。

2. **關鍵要素提取**：
    - 輸入是一個整數數組 `candidates` 和一個目標值 `target`
    - 需要找出所有總和為 `target` 的組合
    - 每個數字在每個組合中只能使用一次
    - 不能有重複的組合

3. **示例理解**：以示例 1 為例，`candidates = [10,1,2,7,6,1,5]`，`target = 8`
    - 注意到 `candidates` 中有重複元素（兩個 1）
    - 輸出 `[1,1,6], [1,2,5], [1,7], [2,6]` 顯示組合中可以使用重複元素（如兩個 1），但每個元素在索引層面只能使用一次
    - 所有組合的和都等於 8

4. **心智模型**：我們可以想像一個決策樹，每個節點表示選擇或不選擇當前數字，但我們需要一個機制來避免產生重複組合。

5. **初步思路**：可能需要回溯法（Backtracking）來探索所有可能的組合，同時需要某種方式來處理重複元素，以避免重複組合。

## 3. 問題理解和核心挑戰

### 問題分析

1. **輸入特徵**：
    - 候選數字陣列可能包含重複元素
    - 所有數字都是正整數
    - 數組長度和各值都有明確範圍

2. **輸出要求**：
    - 所有和為 `target` 的唯一組合
    - 每個組合中每個數字只能用一次

3. **邊界情況**：
    - 陣列中有重複元素
    - 目標值小於陣列中最小元素
    - 陣列長度為最小值（1）

### 核心挑戰

1. **避免重複組合**：這是最核心的挑戰。當 `candidates` 包含重複元素時，如何確保不產生重複的組合？

2. **高效搜索空間**：如何高效地搜索所有可能的組合，避免不必要的搜索分支？

3. **組合選擇的有序性**：如何確保組合是按照某種順序生成的，以方便去重？

## 4. 視覺問題表示

為了更好地理解問題，我們可以將搜索過程視覺化為一棵決策樹：

以 `candidates = [1,1,2,5,6,7,10]`（排序後），`target = 8` 為例：

```
                          []
                        /    \
           選擇1      /      \     不選擇1
                    /        \
                [1]          []
               /   \        /   \
       選擇1  /     \ 不選1 /     \ 不選2
            /       \     /       \
         [1,1]      [1]  []        []
        /    \     /   \  ...      ...
    選2/      \不選2
      /        \
  [1,1,2]     [1,1]
 超過target    ...
```

通過這種決策樹的表示，我們可以看到：
1. 每個節點表示當前已選擇的組合
2. 每個分支代表選擇或不選擇當前數字
3. 當組合總和等於 `target` 時，我們找到一個有效組合
4. 當組合總和超過 `target` 時，我們可以剪枝（不繼續向下探索）

## 5. 問題模式識別

這個問題具有以下特徵：

1. **組合問題**：尋找滿足特定條件的所有可能組合，通常使用回溯法解決。

2. **有約束條件的搜索**：組合總和必須等於目標值，且不允許重複組合。

3. **集合中有重複元素**：需要特殊處理以避免生成重複組合。

4. **剪枝機會**：可以通過排序和提前判斷來減少搜索空間。

這屬於典型的「回溯＋剪枝」問題模式，類似於 LeetCode 中的其他組合問題（如 Combination Sum, Subsets II 等）。

## 6. 策略性問題解決框架

對於這類組合搜索問題，我們可以採用以下通用框架：

1. **預處理**：
    - 首先對候選數組進行排序，使相同元素相鄰，便於去重
    - 可以先檢查是否有快速的無解判斷方法

2. **回溯框架**：
    - 維護一個當前組合路徑
    - 在每一步，決定選擇或不選擇當前元素
    - 使用遞迴實現深度優先搜索

3. **剪枝策略**：
    - 和超過目標值時提前返回
    - 對於相同元素，只在特定條件下繼續搜索以避免重複
    - 當剩餘元素總和不足以達到目標時提前返回

4. **結果收集**：
    - 當當前組合總和等於目標值時，將組合添加到結果列表中

## 7. 演算法設計（編碼前）

讓我們設計一個回溯算法來解決這個問題：

**偽代碼**：
```
function combinationSum2(candidates, target):
    結果 = []
    當前組合 = []
    
    // 排序候選數組，使相同元素相鄰
    排序(candidates)
    
    function backtrack(起始索引, 剩餘目標值):
        if 剩餘目標值 == 0:
            添加當前組合副本到結果
            return
        
        if 剩餘目標值 < 0:
            return  // 剪枝：總和超過目標
        
        for i from 起始索引 to candidates.length - 1:
            // 跳過重複項以避免重複組合
            if i > 起始索引 且 candidates[i] == candidates[i-1]:
                continue
            
            // 選擇當前數字
            添加 candidates[i] 到當前組合
            
            // 遞迴，注意索引增加1，因為每個數字只能用一次
            backtrack(i + 1, 剩餘目標值 - candidates[i])
            
            // 回溯，移除最後添加的數字
            移除當前組合的最後一個元素
    
    backtrack(0, target)
    return 結果
```

這個算法的關鍵點是：
1. 排序數組以便於去重
2. 在回溯過程中跳過重複元素（當 `i > 起始索引 且 candidates[i] == candidates[i-1]` 時）
3. 使用起始索引來確保每個元素只使用一次
4. 當剩餘目標值為0時，我們找到一個有效組合

## 8. 視覺解釋

讓我們通過一個具體例子來視覺化算法的執行：
以 `candidates = [1,1,2,5,6,7,10]`（已排序），`target = 8` 為例：

1. 初始狀態：當前組合 = [], 剩餘目標 = 8
2. 選擇或不選擇第一個 1：
    - 選擇：當前組合 = [1], 剩餘目標 = 7
        - 選擇或不選擇第二個 1：
            - 選擇：當前組合 = [1,1], 剩餘目標 = 6
                - 選擇或不選擇 2：
                    - 選擇：當前組合 = [1,1,2], 剩餘目標 = 4
                        - 選擇或不選擇 5：...（繼續遞迴）
                    - 不選擇 2：...（繼續遞迴）
            - 不選擇第二個 1：當前組合 = [1], 剩餘目標 = 7
                - 注意：我們跳過第二個 1，因為這會產生重複組合
                - 直接考慮 2：...（繼續遞迴）
    - 不選擇第一個 1：當前組合 = [], 剩餘目標 = 8
        - 繼續考慮從第二個 1 開始的情況...

通過這種選擇與不選擇的遞迴結構，我們可以系統地探索所有可能的組合，同時通過跳過重複元素來避免重複組合。

## 9. 解決方案發展歷程

### 初始方法：暴力回溯

最直接的想法是使用純回溯法，嘗試所有可能的組合，然後使用集合去重。但這種方法效率低下，尤其是當有大量重複元素時。

### 改進方法：排序 + 剪枝

1. **排序**：首先排序數組，使相同元素相鄰。
2. **在搜索過程中去重**：在回溯過程中，如果當前元素與前一個元素相同，且前一個元素沒有被選擇，則跳過當前元素。
3. **其他剪枝策略**：當剩餘總和變為負數時提前返回。

### 關鍵優化：同層級元素去重

最關鍵的優化是在同一遞迴層級中跳過重複元素。當 `i > 起始索引 且 candidates[i] == candidates[i-1]` 時，我們跳過當前元素，這確保了對於相同值的元素，我們只在第一次出現時考慮選擇它，從而避免了重複組合。

## 10. 實例演練

讓我們用 `candidates = [1,1,2,5,6,7,10]`（已排序），`target = 8` 手動追蹤算法的執行過程：

1. **初始調用**：`backtrack(0, 8)`
    - 當前組合 = []

2. **迭代索引 i=0**（值為1）：
    - 選擇：當前組合 = [1]
    - 遞迴調用：`backtrack(1, 7)`

3. **在 `backtrack(1, 7)` 中**：
    - 迭代索引 i=1（值為1）：
        - 選擇：當前組合 = [1,1]
        - 遞迴調用：`backtrack(2, 6)`

4. **在 `backtrack(2, 6)` 中**：
    - 迭代索引 i=2（值為2）：
        - 選擇：當前組合 = [1,1,2]
        - 遞迴調用：`backtrack(3, 4)`

5. **在 `backtrack(3, 4)` 中**：
    - 迭代索引 i=3（值為5）：（超過目標，跳過）
    - 迭代索引 i=4（值為6）：（超過目標，跳過）
    - 迭代索引 i=5（值為7）：（超過目標，跳過）
    - 迭代索引 i=6（值為10）：（超過目標，跳過）
    - 回溯：當前組合 = [1,1]

6. **回到 `backtrack(2, 6)`**：
    - 迭代索引 i=3（值為5）：
        - 選擇：當前組合 = [1,1,5]
        - 遞迴調用：`backtrack(4, 1)`

7. **在 `backtrack(4, 1)` 中**：
    - 迭代索引 i=4（值為6）：（超過目標，跳過）
    - ...（其他索引同樣超過目標）
    - 回溯：當前組合 = [1,1]

... 依此類推，繼續探索所有可能的組合路徑。

這個過程最終會找到所有和為 8 的組合：`[1,1,6]`, `[1,2,5]`, `[1,7]`, `[2,6]`。

## 12. 實現執行演練

以輸入 `candidates = [1,1,2,5,6,7,10]`，`target = 8` 為例，讓我們詳細追蹤 Golang 代碼的執行過程：

1. **初始化**：
    - 創建空結果集 `result = []`
    - 排序候選數組 `candidates = [1,1,2,5,6,7,10]`
    - 創建空路徑 `path = []`
    - 調用 `backtrack(0, 8, [])`

2. **backtrack(0, 8, [])**：
    - 檢查 `remain == 0`：否
    - 檢查 `remain < 0 || start >= len`：否
    - 循環 `i = 0` 到 `6`：
        - `i = 0`：檢查是否跳過：否（這是該層第一個元素）
        - 檢查 `candidates[0] > remain`：否（1 < 8）
        - 將 `candidates[0]` 添加到路徑：`path = [1]`
        - 調用 `backtrack(1, 7, [1])`

3. **backtrack(1, 7, [1])**：
    - 檢查 `remain == 0`：否
    - 檢查 `remain < 0 || start >= len`：否
    - 循環 `i = 1` 到 `6`：
        - `i = 1`：檢查是否跳過：否（索引 1 是該層第一個元素）
        - 檢查 `candidates[1] > remain`：否（1 < 7）
        - 將 `candidates[1]` 添加到路徑：`path = [1,1]`
        - 調用 `backtrack(2, 6, [1,1])`

4. **backtrack(2, 6, [1,1])**：
    - 循環 `i = 2` 到 `6`：
        - `i = 2`：選擇 `candidates[2]` (2)
        - 路徑：`path = [1,1,2]`
        - 調用 `backtrack(3, 4, [1,1,2])`

5. **backtrack(3, 4, [1,1,2])**：
    - 循環 `i = 3` 到 `6`：
        - `i = 3`：選擇 `candidates[3]` (5)
        - 檢查 `candidates[3] > remain`：是（5 > 4），跳出循環
    - 回溯：`path = [1,1]`

6. **回到 backtrack(2, 6, [1,1])**：
    - 繼續循環，選擇 `candidates[4]` (6)
    - 路徑：`path = [1,1,6]`
    - 調用 `backtrack(5, 0, [1,1,6])`

7. **backtrack(5, 0, [1,1,6])**：
    - 檢查 `remain == 0`：是
    - 將副本 `[1,1,6]` 添加到結果集
    - 回溯：`path = [1,1]`

繼續這個過程，最終我們會找到所有有效組合 `[1,1,6]`、`[1,2,5]`、`[1,7]`、`[2,6]`。

## 13. 複雜度分析

### 時間複雜度

回溯算法的時間複雜度分析通常比較複雜，因為它取決於搜索樹的形狀和剪枝效率。

1. **最壞情況**：O(2^n)，其中 n 是 `candidates` 的長度。這是因為每個元素有兩種選擇（選或不選）。

2. **實際情況**：由於我們有多種剪枝策略（排序後的提前退出、跳過重複元素等），實際運行時間會大大低於理論上界。

3. **排序**：排序數組需要 O(n log n) 的時間。

綜合起來，時間複雜度為 O(2^n + n log n)，簡化為 O(2^n)，因為當 n 較大時，指數項主導。

### 空間複雜度

1. **遞迴堆棧**：在最壞情況下，遞迴深度可達 n，因此遞迴堆棧空間為 O(n)。

2. **臨時路徑**：`path` 數組最多包含 n 個元素，空間為 O(n)。

3. **結果集**：結果集大小取決於有效組合的數量，最壞情況下可以是 O(2^n * n)，但實際上會因為問題約束（如 `target` 值）而小得多。

綜合起來，空間複雜度為 O(n + output_size)，其中 output_size 是結果集的大小。

## 14. 優化與改進

雖然我們的解決方案已經包含了多種優化策略，但還可以考慮以下改進：

1. **初始檢查**：在開始回溯前，檢查特殊情況，如所有候選數字之和是否小於目標值，如果是，則可以直接返回空結果。

2. **更進一步的剪枝**：計算剩餘數字的總和，如果總和小於剩餘目標值，可以提前結束搜索。

3. **迭代而非遞迴**：在極端情況下，可以考慮使用迭代方法代替遞迴以避免堆棧溢出，但這會使代碼更加複雜。

4. **位掩碼優化**：對於較小的輸入，可以考慮使用位掩碼來表示組合，這可能提高效率並減少內存使用。

## 15. 通用問題解決智慧

從這個問題中，我們可以提取以下通用解題思路和原則：

1. **排序預處理**：在處理具有重複元素且需要避免重複組合的問題時，排序通常是第一步，使相同元素相鄰便於去重。

2. **回溯模板**：大多數組合/排列問題都可以使用回溯法解決，關鍵是設計好狀態轉移和終止條件。

3. **剪枝策略**：識別並應用問題特定的剪枝策略可以顯著提高效率：
    - 排序後的提前退出
    - 同層元素去重
    - 目標值或條件不滿足時的提前返回

4. **選擇表示**：在解決組合問題時，常見的選擇表示方法是：選擇當前元素、跳過當前元素，這構成了決策樹的兩個分支。

5. **問題轉化**：將問題轉化為已知模式（如回溯、動態規劃等）是解題的關鍵一步。

## 16. 測試策略

1. **LeetCode 官方示例**：驗證與問題描述中的示例一致。
2. **無解情況**：測試當沒有符合條件的組合時的行為。
3. **單一元素等於目標值**：測試只有一個元素等於目標值的情況。
4. **多個重複元素**：測試有多個相同元素的情況。
5. **最小輸入大小**：測試數組只有一個元素的邊界情況。
6. **所有元素都大於目標值**：測試快速剪枝優化的情況。
7. **最大目標值**：測試接近約束上限的情況。

測試還包括一個輔助函數 `sortResult`，用於對二維切片進行排序，以便進行結果比較，因為組合的順序可能不同但內容相同。
