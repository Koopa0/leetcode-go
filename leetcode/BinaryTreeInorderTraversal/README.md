# LeetCode 94: Binary Tree Inorder Traversal 二元樹的中序遍歷

## 1. 問題定義

### Original Problem (English)
```
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Constraints:
- The number of nodes in the tree is in the range [0, 100].
- -100 <= Node.val <= 100
```

### 問題翻譯（繁體中文）
```
給定一個二元樹的根節點，回傳其節點值的中序遍歷結果。

範例 1:
輸入: root = [1,null,2,3]
輸出: [1,3,2]

範例 2:
輸入: root = []
輸出: []

範例 3:
輸入: root = [1]
輸出: [1]

限制條件:
- 樹中的節點數量範圍在 [0, 100] 之間。
- -100 <= 節點值 <= 100
```

### 範例與限制條件
- **範例 1:**
  ```
  輸入: root = [1,null,2,3]
  輸出: [1,3,2]
  說明: 該二元樹結構如下:
       1
        \
         2
        /
       3
  中序遍歷順序為: 左->根->右，所以結果為 [1,3,2]
  ```
- **範例 2:**
  ```
  輸入: root = []
  輸出: []
  說明: 空樹，遍歷結果為空陣列。
  ```
- **範例 3:**
  ```
  輸入: root = [1]
  輸出: [1]
  說明: 只有一個根節點，遍歷結果為包含該節點值的陣列。
  ```

- **限制條件:**
    - 樹中的節點數量範圍在 [0, 100] 之間。
    - -100 <= 節點值 <= 100

## 2. 問題理解

### 初始反應與心智模型
- 這是一個基本的二元樹遍歷問題，屬於資料結構的基礎題型。
- 中序遍歷（Inorder Traversal）是指按照「左子樹->根節點->右子樹」的順序訪問所有節點。
- 對於一棵二元樹，中序遍歷可以得到一個有序的序列（如果是二元搜尋樹）。
- 我可以從遞迴和迭代兩種方式來思考此問題。

### 問題分解
- 核心子問題：如何在訪問根節點前，確保左子樹的所有節點都已被訪問。
- 基本操作：遞迴遍歷左子樹、訪問當前節點、遞迴遍歷右子樹。
- 必須處理的邊界情況：空樹和只有一個節點的樹。
- 可以分解為：先完整遍歷左子樹，再訪問當前節點，最後完整遍歷右子樹。

### 視覺表示
```
例如考慮以下二元樹:
     1
    / \
   2   3
  / \
 4   5

中序遍歷的過程如下:
1. 從根節點 1 開始
2. 遞迴訪問左子樹（以 2 為根）
   a. 遞迴訪問左子樹（以 4 為根）
      i. 4 沒有左子樹，訪問 4
      ii. 4 沒有右子樹，回溯到 2
   b. 訪問 2
   c. 遞迴訪問右子樹（以 5 為根）
      i. 5 沒有左子樹，訪問 5
      ii. 5 沒有右子樹，回溯到 1
3. 訪問 1
4. 遞迴訪問右子樹（以 3 為根）
   a. 3 沒有左子樹，訪問 3
   b. 3 沒有右子樹，結束

最終結果: [4, 2, 5, 1, 3]
```

- 這個視覺化幫助我理解遍歷的順序以及遞迴調用的過程。
- 從視覺模型中，我們可以看出中序遍歷總是先處理完左子樹，再處理當前節點，最後處理右子樹。

## 3. 模式識別與知識映射

### 演算法模式分類
- 這個問題屬於以下常見的演算法模式：
    - [ ] 陣列/字串操作
    - [ ] 雙指針/滑動窗口
    - [ ] 二分搜尋/二分答案
    - [x] 深度優先搜尋 (DFS)
    - [ ] 廣度優先搜尋 (BFS)
    - [ ] 回溯
    - [ ] 動態規劃 (DP)
    - [ ] 貪婪演算法
    - [ ] 分治法
    - [ ] 圖論演算法
    - [x] 樹問題
    - [x] 堆疊/佇列
    - [ ] 優先佇列/堆
    - [ ] 雜湊表/集合
    - [ ] 排序演算法
    - [ ] 位元操作
    - [ ] 其他: ________

- 識別基礎:
    - 問題明確要求對二元樹進行中序遍歷，這是典型的樹問題。
    - 中序遍歷本質上是深度優先搜尋的一種特殊形式，使用遞迴或堆疊來實現。
    - 我之前遇到過前序、中序和後序遍歷的問題，它們都有相似的解題模式。

### 知識連結
- 相關的基本資訊科學概念：樹資料結構、遞迴、堆疊、深度優先搜尋。
- 這個問題與其他樹遍歷問題（前序、後序）密切相關，只是訪問節點的順序不同。
- 需要應用的理論知識：樹的基本性質、遞迴思想、迭代方法中堆疊的使用。

### 類似問題比較
- LeetCode 144: Binary Tree Preorder Traversal（二元樹前序遍歷）
- LeetCode 145: Binary Tree Postorder Traversal（二元樹後序遍歷）
- LeetCode 102: Binary Tree Level Order Traversal（二元樹層序遍歷）
- 相似之處：都是關於樹的遍歷方式，核心思想類似。
- 不同之處：前序是「根->左->右」，中序是「左->根->右」，後序是「左->右->根」，層序是按層從上到下遍歷。

## 4. 演算法直覺發展

### 直覺建立
- 中序遍歷像是在二元樹中進行「左手法則」探索：總是盡可能往左走，直到無路可走，然後訪問當前節點，再嘗試往右走。
- 這就像閱讀一本展開的書：從左頁開始，然後是中間的書脊，最後是右頁。
- 我可以通過模擬手動執行來驗證這種直覺理解是否正確。

### 多角度思考
- 從不同角度思考此問題：
    - 遞迴方法：自然反映了「左->根->右」的訪問順序，代碼簡潔但可能有堆棧溢出的風險。
    - 迭代方法：使用明確的堆疊來模擬遞迴過程，避免了遞迴可能帶來的問題，但實現較複雜。
    - 莫里斯遍歷（Morris Traversal）：一種不使用額外空間的遍歷方法，通過臨時修改樹的結構來實現。
- 對於這種基礎問題，遞迴方法最為直觀，但在面試中展示迭代解法也很重要。

## 5. 解決方案開發歷程

### 方法 1: 遞迴解法

#### 思考過程
- 最直觀的方法是使用遞迴，因為中序遍歷的定義本身就是遞迴的：先遍歷左子樹，再訪問當前節點，最後遍歷右子樹。
- 遞迴簡潔且直觀地表達了問題的本質。
- 這種方法需要處理空樹的情況，這是遞迴的基本情況。

#### 演算法設計
```
1. 建立一個輔助函數 inorderTraversal，該函數接收一個節點和一個結果陣列。
2. 如果當前節點為空，直接返回。
3. 遞迴遍歷左子樹。
4. 將當前節點的值加入結果陣列。
5. 遞迴遍歷右子樹。
6. 主函數初始化一個空陣列，調用輔助函數，然後返回結果陣列。
```

#### 實現細節
- 使用了 Go 的閉包特性定義內部輔助函數，這樣可以直接訪問外部的結果陣列。
- 沒有使用額外的函數參數來傳遞結果陣列，使代碼更簡潔。
- 遞迴基本情況的處理是確保算法正確性的關鍵。

#### 複雜度分析
- **時間複雜度**: O(n) — 其中 n 是樹中的節點數量。每個節點恰好被訪問一次。
- **空間複雜度**: O(h) — 其中 h 是樹的高度。遞迴調用堆疊的最大深度等於樹的高度。在最壞情況下（樹退化為鏈表），空間複雜度為 O(n)；在平衡二元樹的情況下，空間複雜度為 O(log n)。

#### 解法評估
- **優勢**: 實現簡單直觀，代碼簡潔。
- **劣勢**: 對於非常深的樹可能導致堆疊溢出。
- 在面試環境中，這是一個可接受的解法，尤其是要快速展示對問題的理解。

### 方法 2: 迭代解法（使用堆疊）

#### 關鍵洞察
- 遞迴解法中，系統調用堆疊追踪了需要訪問的節點。我們可以顯式使用一個堆疊來模擬這個過程。
- 在中序遍歷中，我們需要先遍歷到最左邊的節點，然後回溯。
- 使用堆疊可以追踪已訪問的節點的路徑，使我們能夠在需要時回溯。

#### 優化策略
```
1. 初始化一個空堆疊和一個指向根節點的指針。
2. 當指針非空或堆疊非空時：
   a. 如果指針非空，將當前節點壓入堆疊，然後將指針移到左子節點（盡可能向左）。
   b. 如果指針為空，彈出堆疊頂部的節點，訪問它（加入結果陣列），然後將指針指向其右子節點。
3. 返回結果陣列。
```

#### 實現改進
- 使用 Go 的切片作為堆疊，避免了使用額外的堆疊數據結構。
- 通過兩個嵌套的循環清晰地表達了算法邏輯：外部循環控制整體過程，內部循環處理向左遍歷的部分。
- 迭代解法避免了遞迴調用可能導致的堆疊溢出問題。

#### 複雜度分析
- **時間複雜度**: O(n) — 每個節點仍然只被訪問一次。
- **空間複雜度**: O(h) — 堆疊的最大大小等於樹的高度。在最壞情況下為 O(n)，在平衡二元樹的情況下為 O(log n)。

#### 解法評估
- **優勢**: 避免了遞迴可能導致的堆疊溢出問題，對於非常深的樹更安全。
- **劣勢**: 代碼較遞迴解法稍微複雜一些，理解起來需要更多的思考。
- 在面試中展示這種解法可以表明你了解遞迴與迭代之間的關係，以及如何在實際中處理潛在的遞迴問題。

### 方法 3: 莫里斯遍歷（Morris Traversal）

#### 突破性思維
- 前兩種方法都需要 O(h) 的額外空間。莫里斯遍歷允許我們在 O(1) 空間複雜度下完成遍歷。
- 關鍵洞察：我們可以通過臨時修改樹的結構（建立臨時鏈接）來避免使用堆疊或遞迴。
- 具體而言，對於當前節點，我們找到其前驅節點（中序遍歷中當前節點的前一個節點），然後將前驅節點的右指針指向當前節點，這樣在遍歷完左子樹後，我們可以通過這個臨時鏈接回到當前節點。

#### 最佳演算法
```
1. 初始化當前節點為根節點。
2. 當當前節點非空時：
   a. 如果當前節點沒有左子節點：
      i. 訪問當前節點。
      ii. 移動到右子節點。
   b. 如果當前節點有左子節點：
      i. 找到當前節點的前驅（左子樹的最右節點）。
      ii. 如果前驅的右指針為空，將其指向當前節點，然後移動到左子節點。
      iii. 如果前驅的右指針已經指向當前節點，將其重置為空（恢復樹的原始結構），訪問當前節點，然後移動到右子節點。
3. 返回結果陣列。
```

#### 實現卓越性
- 這個 Go 實現達到了 O(1) 的額外空間複雜度，是一個真正的優化。
- 雖然代碼比前兩種解法更複雜，但它展示了對樹結構的深入理解。
- 在保持效能的同時，代碼通過清晰的註釋和結構化的邏輯保持了可讀性。

#### 複雜度分析
- **時間複雜度**: O(n) — 乍看之下，尋找前驅節點似乎需要 O(n) 時間，但實際上每條邊至多被訪問 3 次（建立臨時鏈接時、通過臨時鏈接回溯時、恢復原始結構時），所以總時間複雜度仍然是 O(n)。
- **空間複雜度**: O(1) — 只使用了有限的幾個變數，不依賴於樹的大小。

#### 從暴力解法到最佳解法的思維演變
- 我的思維從最直觀的遞迴解法開始，理解了問題的本質。
- 然後轉向迭代解法，顯式使用堆疊來避免遞迴可能帶來的問題。
- 最終，通過對樹結構的深入理解，找到了可以在不使用額外空間的情況下完成遍歷的莫里斯方法。
- 這種思維模式可以應用於其他樹問題：先找到最直觀的解法，然後考慮如何降低空間複雜度。

## 6. 範例演示與 Go 實現

### 完整範例追蹤
以範例輸入 `[1,null,2,3]` 使用我們的最佳解法（莫里斯遍歷）進行追蹤：

1. 初始狀態：
    - 輸入：`[1,null,2,3]`，樹結構為：
      ```
        1
         \
          2
         /
        3
      ```
    - `current` = 1
    - `result` = []

2. 第一步：
    - `current.Left` 為 `nil`，所以我們訪問當前節點，將 1 加入結果陣列。
    - `current` 更新為 `current.Right`，即 2。
    - `result` = [1]

3. 第二步：
    - `current` = 2
    - `current.Left` 不為 `nil`（是 3），所以我們找到 2 的前驅。
    - 前驅是 3。
    - 前驅的右指針為 `nil`，所以我們將其指向 2，然後將 `current` 更新為 `current.Left`，即 3。
    - 樹被臨時修改為：
      ```
        1
         \
          2
         /
        3 -> 2 (臨時鏈接)
      ```

4. 第三步：
    - `current` = 3
    - `current.Left` 為 `nil`，所以我們訪問當前節點，將 3 加入結果陣列。
    - `current` 更新為 `current.Right`，由於臨時鏈接，這是 2。
    - `result` = [1, 3]

5. 第四步：
    - `current` = 2
    - `current.Left` 不為 `nil`（是 3），所以我們找到 2 的前驅。
    - 前驅是 3。
    - 前驅的右指針已經指向 2，所以我們將其重置為 `nil`，恢復原始樹結構，然後訪問當前節點，將 2 加入結果陣列。
    - `current` 更新為 `current.Right`，即 `nil`。
    - `result` = [1, 3, 2]

6. 第五步：
    - `current` = `nil`，循環結束。

7. 最終狀態：
    - 輸出：`[1, 3, 2]`

### 所有方法的效能比較
```
| 方法            | 時間複雜度  | 空間複雜度  | 範例執行時間   |
|-----------------|------------|------------|--------------|
| 遞迴解法         | O(n)       | O(h)       | 0 ms         |
| 迭代解法（堆疊）  | O(n)       | O(h)       | 0 ms         |
| 莫里斯遍歷        | O(n)       | O(1)       | 0 ms         |
```

## 7. Go 最佳實踐與測試

### Go 慣用解決方案
- 我的解決方案遵循了 Go 的最佳實踐：代碼清晰、變數命名合理、註釋充分。
- 使用了 Go 的切片作為堆疊，這是 Go 中常見的做法。
- 確保了代碼簡潔而可維護，不使用不必要的變數或複雜結構。

### 錯誤處理與邊界情況
- 處理了空樹的情況：無論是遞迴、迭代還是莫里斯遍歷，都能正確處理空樹。
- 對於只有一個節點的樹，三種方法也都能正確處理。
- 解決方案對意外輸入（如極端不平衡的樹）也很健壯。


## 8. 面試模擬

### 時間管理規劃
- 問題理解: ~1-2 分鐘（基本樹遍歷問題，概念簡單）
- 初始解法提案: ~2-3 分鐘（直接提出遞迴解法）
- 優化討論: ~5 分鐘（討論從遞迴到迭代，再到莫里斯遍歷的過程）
- 代碼編寫: ~5-8 分鐘（實現遞迴解法和迭代解法）
- 測試與除錯: ~2-3 分鐘（檢查邊界情況）

### 面試官互動模擬
- 如何向面試官解釋我的 Go 實現：
  "首先，我會使用遞迴方法解決這個問題，因為中序遍歷本身就是遞迴定義的。在 Go 中，我可以使用閉包來定義一個內部的遞迴函數，這樣可以直接訪問外部的結果陣列。遞迴遍歷的過程是：先遞迴遍歷左子樹，再訪問當前節點，最後遞迴遍歷右子樹。"

  "但是，遞迴解法可能會在樹非常深時導致堆疊溢出。因此，我還可以提供一個迭代解法，使用明確的堆疊來模擬遞迴過程。在 Go 中，我可以使用切片作為堆疊。"

  "如果我們想進一步優化空間複雜度，可以使用莫里斯遍歷，它的空間複雜度是 O(1)。這種方法通過臨時修改樹的結構來避免使用額外的堆疊或遞迴。"

- 潛在的提示和引導問題：
    - "你能解釋一下中序遍歷的定義嗎？"
    - "遞迴解法的空間複雜度是多少？有沒有辦法優化？"
    - "你能不使用遞迴或堆疊來解決這個問題嗎？"
    - "莫里斯遍歷是如何工作的？它如何避免使用額外空間？"

### 潛在的後續問題
- 如果輸入規模大幅增加，你會如何修改你的 Go 代碼？
    - 對於非常大的樹，我會使用莫里斯遍歷，因為它的空間複雜度是 O(1)，不會因為樹太大而耗盡記憶體。

- 如果記憶體受限，你會應用哪些 Go 特定的優化？
    - 使用莫里斯遍歷，它不需要額外的堆疊或佇列。
    - 如果需要返回結果陣列，可以事先估計結果大小並預分配空間，避免多次動態調整陣列大小。

- 你的解決方案如何擴展以處理相關但更複雜的問題？
    - 這種遍歷方法可以擴展到其他樹遍歷問題，如前序或後序遍歷。
    - 對於更複雜的問題，如二元樹的序列化與反序列化，或者樹的路徑和問題，基本的遍歷方法都是基礎。

## 9. 知識整合與學習

### 問題解決洞察
- 這個問題教會我了 Go 中閉包的使用以及如何使用切片作為堆疊。
- 我更深入地理解了 Go 中遞迴與迭代的關係，以及如何在需要時從一種轉換到另一種。
- 我需要進一步加強的 Go 知識領域包括：更複雜的樹操作和使用 Go 進行圖算法。

### 心智模型構建
- 從這個經驗中，我抽象出的一般解題框架是：先找到最直觀的解法，然後再考慮優化。
- 這個框架適用於各種樹問題：先考慮遞迴解法，然後考慮如何轉換為迭代解法，最後考慮是否有特殊的空間優化技巧。
- 將這種經驗內化為直覺，有助於未來更快地解決類似問題。

### 錯誤模式識別
- 在實現過程中，常見的錯誤包括忘記處理空樹的情況或錯誤地實現堆疊操作。
- 這些錯誤揭示了我在處理邊界情況和使用數據結構方面的盲點。
- 為了防止類似錯誤，我應該在實現前更清晰地思考邊界情況，並在實現後使用具體例子進行驗證