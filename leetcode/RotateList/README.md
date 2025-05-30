# LeetCode 61: Rotate List (旋轉鏈結串列)

## 1. 問題定義

### 原始問題 (英文)
```
Given the head of a linked list, rotate the list to the right by k places.

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]

Example 2:
Input: head = [0,1,2], k = 4
Output: [2,0,1]

Constraints:
- The number of nodes in the list is in the range [0, 500].
- -100 <= Node.val <= 100
- 0 <= k <= 2 * 10^9
```

### 問題翻譯 (繁體中文)
```
給定一個鏈結串列的頭節點，將該串列向右旋轉 k 個位置。

範例 1：
輸入：head = [1,2,3,4,5], k = 2
輸出：[4,5,1,2,3]

範例 2：
輸入：head = [0,1,2], k = 4
輸出：[2,0,1]

限制條件：
- 串列中的節點數量在 [0, 500] 範圍內。
- -100 <= 節點值 <= 100
- 0 <= k <= 2 * 10^9
```

### 範例與限制條件
- **範例 1:**
  ```
  輸入: head = [1,2,3,4,5], k = 2
  輸出: [4,5,1,2,3]
  解釋: 向右旋轉 2 次，原串列變為 [4,5,1,2,3]
  ```
- **範例 2:**
  ```
  輸入: head = [0,1,2], k = 4
  輸出: [2,0,1]
  解釋: 向右旋轉 4 次，原串列變為 [2,0,1]
  ```

- **限制條件:**
    - 串列中的節點數量在 [0, 500] 範圍內
    - -100 <= 節點值 <= 100
    - 0 <= k <= 2 * 10^9

## 2. 問題理解

### 初始反應與心智模型
- 第一反應：這是個鏈結串列操作問題，需要將串列末端的節點移至前端
- 範例解釋：旋轉 k 次意味著將最後 k 個節點移到串列的開頭
- 輸入是鏈結串列的頭節點，輸出是旋轉後的新頭節點
- 視覺上，可以想像鏈結串列是一個環，我們只是改變了「頭」的位置

### 問題分解
- 核心子問題：
    1. 找到新的頭節點位置（即原始串列的第 (length-k%length) 個節點）
    2. 將原串列的尾節點連接到原始頭節點
    3. 切斷新的尾節點連接，並將其 next 設為 null
- 需要處理的邊緣情況：
    - 空串列或單節點串列（直接返回原串列）
    - k 為 0 或 k 為串列長度的倍數（不需要旋轉）
    - k 可能非常大，需要使用 k % length 來簡化

### 視覺呈現
```
原始串列: 1 -> 2 -> 3 -> 4 -> 5 -> NULL, k = 2

步驟 1: 計算長度並找到新的斷開點
長度 = 5
k % 長度 = 2
新的斷開點 = 5 - 2 = 3 (即節點值為 3 的節點)

步驟 2: 找到尾節點並形成環
1 -> 2 -> 3 -> 4 -> 5
                    |
                    v
                    1

步驟 3: 從斷開點切斷並重新定義頭節點
原始：1 -> 2 -> 3 -> 4 -> 5 -> NULL
                 |
斷開：1 -> 2 -> 3    4 -> 5
                    |
新頭節點：          4 -> 5 -> 1 -> 2 -> 3 -> NULL
```

## 3. 模式識別與知識映射

### 演算法模式分類
- 這個問題屬於以下常見演算法模式：
    - [x] 鏈結串列操作
    - [ ] 雙指針/滑動窗口
    - [ ] 二分搜尋/二元答案
    - [ ] 深度優先搜尋 (DFS)
    - [ ] 廣度優先搜尋 (BFS)
    - [ ] 回溯法
    - [ ] 動態規劃 (DP)
    - [ ] 貪婪演算法
    - [ ] 分治演算法
    - [ ] 圖演算法
    - [ ] 樹狀問題
    - [ ] 堆疊/佇列
    - [ ] 優先佇列/堆積
    - [ ] 雜湊表/集合
    - [ ] 排序演算法
    - [ ] 位元操作
    - [ ] 其他：________

- 識別依據：
    - 問題直接要求操作鏈結串列結構
    - 需要重新安排鏈結串列的連接方式
    - 涉及找到特定節點（新的頭節點和尾節點）並重新連接

### 知識連結
- 基本電腦科學概念：指針操作、鏈結資料結構
- 相關資料結構：單向鏈結串列
- 需要應用的理論知識：鏈結串列的遍歷與修改、模運算

### 相似問題比較
- LeetCode #19：刪除鏈結串列的倒數第 N 個節點（同樣需要找到特定位置的節點）
- LeetCode #206：反轉鏈結串列（改變鏈結方向）
- LeetCode #92：反轉鏈結串列 II（部分反轉，涉及斷開並重新連接）
- 相似之處：都需要遍歷鏈結串列並重新安排連接關係
- 差異：本題特別涉及「旋轉」操作，即將尾部節點移到前端

## 4. 演算法直覺發展

### 直覺建立
- 真實世界類比：想像一條繩子，我們在特定位置剪開，然後將尾端接到頭部
- 初始直覺：需要先計算串列長度，找到新的頭節點位置，然後重新連接
- 驗證方法：手動追蹤幾個簡單的例子，確保新的連接關係正確

### 多角度思考
- 不同視角：
    - 可以將問題視為：找到新的頭節點和尾節點，然後重新連接
    - 或者視為：將鏈結串列變成環，然後在正確的位置斷開
    - 迭代實現較為直觀，但遞迴方法在此問題中並不合適
- 最有前景的視角：將串列視為環然後斷開，這樣思考最直觀且易於實現

## 5. 解決方案發展過程

### 方法 1：暴力解法

#### 思考過程
- 最直接的方法是逐步旋轉 k 次，每次將最後一個節點移到前端
- 雖然簡單，但當 k 很大時效率低下
- 對於限制條件中 k 最大可達 2 * 10^9，此方法可能會超時

#### 演算法設計
```
1. 對於每次旋轉 (共 k 次)：
   a. 找到當前鏈結串列的最後一個節點和倒數第二個節點
   b. 將最後一個節點的 next 指向頭節點
   c. 倒數第二個節點的 next 設為 null
   d. 更新頭節點為原來的最後一個節點
2. 返回最終的頭節點
```

#### 實作細節
- 使用 Go 的指針操作處理鏈結串列
- 需要特別注意邊緣情況（空串列、單節點串列）
- 針對 k 值可能很大的情況進行了優化（k % length）
- 每次旋轉需要找到最後和倒數第二個節點，增加了時間複雜度

#### 複雜度分析
- **時間複雜度**：O(n * k % n) — 首先遍歷一次串列計算長度 O(n)，然後執行 k % n 次旋轉，每次需要 O(n) 時間找到最後節點，總體為 O(n * k % n)
- **空間複雜度**：O(1) — 只使用了常數額外空間

#### 解法評估
- 優點：思路簡單直接，容易理解和實現
- 缺點：對於大的 k 值效率低下，特別是當串列較長時
- 在面試中接受度：可以作為初始解法，但應該立即提出優化方案

### 方法 2：優化解法

#### 關鍵洞見
- 我們不需要逐步旋轉，可以一次性找到新的頭節點位置
- 新的頭節點是原串列中的第 (length - k % length) 個節點
- 只需要一次遍歷找到新的切割點，然後重新連接

#### 優化策略
```
1. 計算鏈結串列長度
2. 計算實際需要旋轉的次數：k % length
3. 如果實際旋轉次數為 0，直接返回原串列
4. 找到新的尾節點（原串列中的第 length - k % length - 1 個節點）
5. 找到當前的尾節點
6. 重新連接串列：
   a. 將當前尾節點指向原頭節點
   b. 新的頭節點是新尾節點的下一個節點
   c. 新尾節點的 next 設為 null
7. 返回新的頭節點
```

#### 實作改進
- 只需遍歷一次串列找到長度和尾節點
- 直接計算新的切割位置，避免多次旋轉
- 使用 Go 的指針操作，一次性完成重新連接

#### 複雜度分析
- **時間複雜度**：O(n) — 需要兩次遍歷，一次計算長度和找尾節點，一次找新的尾節點，總體為 O(n)
- **空間複雜度**：O(1) — 只使用了常數額外空間

#### 解法評估
- 優點：大幅提高效率，特別是對於大的 k 值
- 缺點：需要兩次遍歷串列
- 為什麼優化有效：消除了重複操作，直接找到最終位置
- 在面試中接受度：良好，展示了對問題的深入理解

### 方法 3：最佳解法

#### 突破性思考
- 我們可以結合第一次遍歷和找新尾節點的過程
- 關鍵洞見：將串列首尾相連形成環，然後在正確位置切斷
- 這種方法只需要一次完整遍歷

#### 最佳演算法
```
1. 處理邊緣情況（空串列、單節點串列、k = 0）
2. 計算鏈結串列長度並找到尾節點
3. 計算實際需要旋轉的次數：k % length
4. 如果實際旋轉次數為 0，直接返回原串列
5. 將尾節點與頭節點相連，形成環狀鏈結串列
6. 找到新的尾節點（往前走 length - k % length 步）
7. 新的頭節點是新尾節點的下一個節點
8. 切斷新尾節點的連接
9. 返回新的頭節點
```

#### 實作卓越點
- 將串列視為環然後在正確位置切斷，思路清晰
- 充分利用 Go 的指針操作，代碼簡潔
- 邊緣情況處理全面，包括 k 為 0 或 length 的倍數

#### 複雜度分析
- **時間複雜度**：O(n) — 只需要一次完整遍歷計算長度，一次部分遍歷找新尾節點，總體為 O(n)
- **空間複雜度**：O(1) — 只使用了常數額外空間

#### 從暴力到最佳的思考演進
- 暴力解法：逐步旋轉，直觀但效率低
- 優化解法：一次找到新頭節點，減少旋轉次數
- 最佳解法：將問題轉化為「找到切斷點」，視鏈結串列為環
- 這種思考模式（將線性結構視為環形）可應用於其他鏈結串列問題

## 6. 範例演練與 Go 實現

### 完整範例追蹤
追蹤範例輸入：head = [1,2,3,4,5], k = 2 使用最佳解法：

1. 初始狀態：
    - 輸入：head = [1,2,3,4,5], k = 2
    - 串列：1 -> 2 -> 3 -> 4 -> 5 -> NULL

2. 計算長度和尾節點：
    - 遍歷串列：length = 5
    - tail = 指向值為 5 的節點

3. 計算實際旋轉次數：
    - k % length = 2 % 5 = 2

4. 形成環：
    - tail.Next = head
    - 現在串列：1 -> 2 -> 3 -> 4 -> 5 -> 1 -> ...

5. 找到新的尾節點：
    - 計算位置：length - k - 1 = 5 - 2 - 1 = 2
    - 從 head 向前走 2 步：current 指向值為 3 的節點

6. 切斷環並定義新頭節點：
    - newHead = current.Next (值為 4 的節點)
    - current.Next = nil
    - 最終串列：4 -> 5 -> 1 -> 2 -> 3 -> NULL

7. 最終狀態：
    - 輸出：[4,5,1,2,3]

### 所有方法的效能比較
```
| 方法           | 時間複雜度     | 空間複雜度 | 範例執行時間 |
|---------------|--------------|-----------|------------|
| 暴力解法       | O(n * k % n) | O(1)      | 較慢        |
| 優化解法       | O(n)         | O(1)      | 中等        |
| 最佳解法       | O(n)         | O(1)      | 最快        |
```

## 7. Go 最佳實踐與測試

### Go 慣用解法
- 解法遵循 Go 的簡潔風格，避免不必要的複雜邏輯
- 使用 Go 的指針操作處理鏈結串列，代碼清晰
- 變數命名符合 Go 的慣例，易於理解

### 錯誤處理與邊緣情況
- 處理了各種邊緣情況：
    - 空串列返回 nil
    - 單節點串列直接返回原串列
    - k 為 0 或串列長度的倍數時直接返回原串列
- 避免了潛在的空指針錯誤

## 8. 面試模擬

### 時間管理規劃
- 問題理解：~2-3 分鐘
- 初始解法提案：~3-5 分鐘
- 優化討論：~5-7 分鐘
- 程式碼編寫：~10 分鐘
- 測試與除錯：~5 分鐘

### 面試官互動模擬
- 如何向面試官解釋我的 Go 實現：
  "這個問題涉及鏈結串列的旋轉。我的解法考慮將鏈結串列視為環，然後在適當的位置切斷。首先，我會處理幾個邊緣情況，然後計算串列長度和實際需要旋轉的次數。將串列首尾相連形成環後，找到新的切斷點並重新定義頭節點。這種方法只需要 O(n) 時間複雜度和 O(1) 空間複雜度。"

- 可能的提示和引導問題：
    - 問："當 k 很大時，你的解法會有效率問題嗎？"
    - 答："不會，因為我使用了 k % length 來減少實際旋轉次數，無論 k 多大，實際計算都不會超過串列長度。"

    - 問："能否不使用額外空間達到 O(n) 的時間複雜度？"
    - 答："是的，我的解法只使用了常數額外空間 O(1)，主要是通過指針操作實現的。"

### 潛在的延伸問題
- 如果輸入規模增加，如何優化您的 Go 代碼？
  "對於非常大的串列，基本邏輯仍然適用。但可以考慮並行處理計算長度的部分，不過在這個特定問題中，瓶頸主要是鏈結串列的線性遍歷，難以通過並行化實現顯著提升。"

- 如果記憶體受限，如何調整您的解法？
  "我的解法已經是記憶體優化的，只使用 O(1) 額外空間。如果對記憶體有極端限制，可以考慮原地修改而不是創建新節點，但在目前的實現中已經做到了這一點。"

- 如何擴展您的解法以處理雙向鏈結串列？
  "對於雙向鏈結串列，核心邏輯相似，但需要同時更新 prev 和 next 指針。找到新的頭節點後，還需要將其 prev 設為 nil，並確保所有節點的 prev 和 next 指針正確連接。基本思路不變，但實現細節需要調整。"

## 9. 知識整合與學習

### 問題解決洞見
- 這個問題教會我利用鏈結串列的特性（轉化為環）來簡化問題
- 深入理解了 Go 中指針操作在鏈結串列中的應用
- 需要加強的領域：更多複雜鏈結串列操作，如合併、分割等

### 心智模型建構
- 通用解題框架：
    1. 先考慮極端/邊緣情況
    2. 將複雜操作轉化為基本操作（如將旋轉問題轉化為切割問題）
    3. 尋找問題的數學特性（如使用模運算處理循環）
- 這種框架可以應用於其他需要改變結構連接關係的問題

### 錯誤模式識別
- 常見錯誤：忘記處理特殊情況如空串列或單節點串列
- 盲點：沒有認識到可以利用環形結構簡化問題
- 避免方法：畫圖模擬解法，檢查各種邊緣情況

### 知識圖譜擴展
- 相關 Go 學習資源：
    - 《Effective Go》中關於指針和結構的部分
    - Go 標準庫中的容器實現
- 進階技術：自定義迭代器、函數式編程在 Go 中的應用
- 這個實現與其他鏈結串列問題形成知識網絡
