# LeetCode 79: Word Search (單字搜尋)

## 1. 問題定義

### Original Problem (English)
```
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
```

### 問題翻譯 (繁體中文)
```
給定一個 m x n 的字元網格 board 和一個字串 word，如果 word 存在於網格中則返回 true。

單字可以由相鄰字元格的字母構成，其中相鄰字元格是指水平或垂直相鄰的格子。同一個字元格不能被使用多次。
```

### 範例與限制條件
- **範例 1:**
  ```
  Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
  Output: true
  Explanation: "ABCCED" 可以在網格中找到，路徑為 A → B → C → C → E → D
  ```
- **範例 2:**
  ```
  Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
  Output: true
  Explanation: "SEE" 可以在網格中找到，路徑為 S → E → E
  ```
- **範例 3:**
  ```
  Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
  Output: false
  Explanation: "ABCB" 無法在網格中找到，因為同一個 'B' 字元格不能使用兩次
  ```

- **限制條件:**
    - `m == board.length`
    - `n == board[i].length`
    - `1 <= m, n <= 6`
    - `1 <= word.length <= 15`
    - `board` 和 `word` 僅由大小寫英文字母組成

## 2. 問題理解

### 初步反應與心智模型
- 這是一個在二維網格中尋找路徑的問題
- 我們需要檢查是否存在一條路徑，使得沿著這條路徑的字元正好匹配給定的單字
- 路徑必須是連續的，且只能水平或垂直移動（不能對角移動）
- 同一個格子不能重複使用
- 直覺上，這看起來像是一個需要搜索的問題，可能需要回溯法

### 問題分解
- 核心子問題：從網格中的每個位置開始，檢查是否能夠找到指定的單字
- 必要操作：
    - 從網格中每個可能的起始位置進行搜索
    - 嘗試四個方向（上、下、左、右）進行移動
    - 標記已訪問的格子以避免重複使用
    - 當找到完整的單字時返回成功
- 邊界情況：
    - 空網格或空單字
    - 網格中不包含單字的第一個字元
    - 單字長度超過網格中可能的最長路徑

### 視覺表示
```
對於範例 1：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"

A B C E
S F C S
A D E E

路徑視覺化：
A → B → C
        ↓
    D ← E
```
- 這個視覺化幫助我理解我們需要尋找一條連續的路徑
- 從視覺模型中，我們可以看到需要在四個方向上進行搜索

## 3. 模式識別與知識映射

### 演算法模式分類
- 本問題屬於以下常見演算法模式：
    - [ ] 陣列/字串操作
    - [ ] 雙指針/滑動視窗
    - [ ] 二分搜索/二元答案
    - [x] 深度優先搜索 (DFS)
    - [ ] 廣度優先搜索 (BFS)
    - [x] 回溯法
    - [ ] 動態規劃 (DP)
    - [ ] 貪婪演算法
    - [ ] 分而治之
    - [ ] 圖演算法
    - [ ] 樹問題
    - [ ] 堆疊/佇列
    - [ ] 優先佇列/堆
    - [ ] 雜湊表/集合
    - [ ] 排序演算法
    - [ ] 位元操作
    - [ ] 其他: ________

- 識別依據：
    - 我們需要嘗試所有可能的路徑，這是 DFS 的典型應用
    - 我們需要標記已訪問的格子，如果當前路徑不成功，需要撤銷標記，這是回溯法的特點

### 知識連接
- 牽涉到的基本資訊科學概念：圖形搜索、遞迴、回溯法
- 這個問題與其他使用 DFS 的矩陣/網格問題相關
- 需要應用的理論知識：深度優先搜索演算法、回溯法原理

### 類似問題比較
- LeetCode 212 Word Search II：尋找網格中存在的多個單字，通常使用 Trie 結構進行優化
- LeetCode 200 Number of Islands：在網格中尋找連通區域
- LeetCode 130 Surrounded Regions：處理被包圍的區域，使用 DFS 或 BFS
- 這些問題都涉及在二維網格中進行搜索，但本問題專注於單一路徑的查找

## 4. 演算法直覺發展

### 直覺建立
- 對於每個格子，如果它匹配單字的第一個字元，我們就從這個格子開始嘗試
- 從一個格子開始，我們嘗試四個方向，看是否能匹配下一個字元
- 如果匹配，我們繼續嘗試下一個字元；如果不匹配，我們回溯
- 這個過程類似於走迷宮，我們嘗試所有可能的路徑，直到找到一條成功的路徑

### 多種視角
- 從不同角度看待這個問題：
    - 自頂向下與自底向上：在這個問題中，自頂向下的 DFS 更直觀
    - 遞迴與迭代：遞迴實作更簡潔，但迭代可能更有效率
    - 狀態視角：可以將每個位置視為在「已訪問」或「未訪問」兩種狀態間轉換
- 對於此類問題，深度優先搜索結合回溯法是最具前景的方法

## 5. 解決方案開發過程

### 方法 1：暴力解法（帶回溯的 DFS）

#### 思考過程
- 最直接的方法是嘗試從網格中的每個位置開始進行深度優先搜索
- 對於每個起始位置，如果它匹配單字的第一個字元，我們就嘗試四個方向
- 我們使用遞迴來實現 DFS，並使用一個矩陣來標記已訪問的格子
- 這種方法簡單直觀，但可能不是最優的

#### 演算法設計
```
1. 遍歷網格中的每個位置 (i, j)
2. 如果位置 (i, j) 的字元等於單字的第一個字元，則從此位置開始進行 DFS
3. 在 DFS 中：
   - 如果已經匹配了整個單字，返回 true
   - 標記當前位置為已訪問
   - 嘗試四個方向（上、下、左、右）
   - 如果任一方向能找到剩餘部分的單字，返回 true
   - 將當前位置標記為未訪問（回溯）
   - 如果所有方向都失敗，返回 false
4. 如果所有起始位置都無法找到單字，返回 false
```

#### 實作細節
- 使用 `[][]bool` 類型的 `visited` 矩陣來跟踪已訪問的格子
- 使用閉包函數實現 DFS，方便訪問外部變數
- 使用方向陣列來簡化四個方向的移動
- 在 DFS 過程中實現回溯，確保同一個格子不被重複使用

#### 複雜度分析
- **時間複雜度**: O(M×N×4^L) — 其中 M 和 N 是網格的尺寸，L 是單字的長度。對於每個起始位置，我們最多有 4^L 個可能的路徑（每一步有 4 個方向）。
- **空間複雜度**: O(M×N + L) — 使用 M×N 的訪問矩陣，遞迴棧的深度為 L。

#### 解法評估
- 優點：直觀，容易理解和實現
- 缺點：在最壞情況下，時間複雜度相當高
- 在面試環境中，這是一個可接受的解法，因為它清晰且正確

### 方法 2：優化解法（直接在原網格上標記）

#### 關鍵洞察
- 在暴力解法中，我們使用了額外的訪問矩陣來標記已訪問的格子
- 我們可以直接修改原網格來標記已訪問的格子，然後在回溯時恢復原值
- 這樣可以節省空間，但需要注意不要永久性地修改輸入

#### 優化策略
```
1. 不使用額外的訪問矩陣
2. 在訪問一個格子時，將其值暫時改為特殊字元（如 '#'）
3. 在回溯時，將其值恢復為原始值
4. 其他邏輯與暴力解法相同
```

#### 實作改進
- 不使用額外的訪問矩陣，直接在原網格上標記已訪問的格子
- 使用特殊字元（'#'）來暫時標記已訪問的格子
- 在回溯時恢復原始值
- 其他邏輯保持不變

#### 複雜度分析
- **時間複雜度**: O(M×N×4^L) — 與暴力解法相同
- **空間複雜度**: O(L) — 只需要考慮遞迴棧的深度，不需要額外的訪問矩陣

#### 解法評估
- 優點：節省了空間，實現仍然直觀
- 缺點：修改了輸入資料，雖然最終會恢復
- 這是一個在面試中很好的優化，展示了對空間複雜度的理解

### 方法 3：最佳化解法（提前剪枝與方向優化）

#### 突破性思考
- 我們可以提前檢查網格中的字母是否足夠構成單字，以快速排除不可能的情況
- 在搜索過程中，我們可以優先嘗試更有可能成功的方向
- 這些優化不會改變最壞情況下的時間複雜度，但可以在實踐中提高效率

#### 最佳演算法
```
1. 前處理：檢查網格中的字母是否足夠構成單字
2. 對於每個可能的起始位置：
   - 如果它匹配單字的第一個字元，則開始 DFS
   - 在 DFS 中優先嘗試與下一個字元匹配的方向
3. 其他邏輯與優化解法相同
```

#### 實作卓越性
- 添加前處理步驟，檢查網格中的字母是否足夠構成單字
- 使用 Go 的 map 作為字母頻率計數器
- 保持代碼的可讀性和結構清晰
- 遵循 Go 的編程風格和最佳實踐

#### 複雜度分析
- **時間複雜度**: O(M×N×4^L) — 最壞情況下與之前的解法相同，但平均情況下會更快
- **空間複雜度**: O(L + C) — L 是單字長度，C 是不同字的數量（頻率計數器）

#### 從暴力解法到最佳解法的思維演變
- 暴力解法：直接應用 DFS 和回溯的基本原理
- 優化解法：通過在原網格上標記減少空間使用
- 最佳解法：增加前處理步驟來提前排除不可能的情況
- 這種思維演變展示了如何從基本解法逐步優化到最佳解法

## 6. 使用 Go 實作的範例演練

### 完整範例追蹤
追蹤範例輸入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"

1. 初始狀態：
    - 前處理：檢查字母頻率，網格中有足夠的字母構成 "ABCCED"
    - 開始從位置 (0,0) 嘗試 DFS

2. DFS 過程：
    - 位置 (0,0)：'A' 匹配 word[0]，標記為已訪問
    - 位置 (0,1)：'B' 匹配 word[1]，標記為已訪問
    - 位置 (0,2)：'C' 匹配 word[2]，標記為已訪問
    - 位置 (1,2)：'C' 匹配 word[3]，標記為已訪問
    - 位置 (1,3)：不匹配 word[4]，回溯
    - 位置 (0,3)：'E' 匹配 word[4]，標記為已訪問
    - 位置 (1,3)：不匹配 word[5]，回溯
    - 位置 (0,2)：已訪問，回溯
    - 繼續尋找其他可能的路徑...
    - 最終找到路徑：(0,0) → (0,1) → (0,2) → (1,2) → (2,2) → (2,1)

3. 最終狀態：
    - 輸出：true（找到匹配的路徑）

### 所有方法的效能比較
```
| 方法           | 時間複雜度       | 空間複雜度    | 範例執行時間 |
|----------------|------------------|---------------|--------------|
| 暴力解法       | O(M×N×4^L)      | O(M×N + L)    | X ms         |
| 優化解法       | O(M×N×4^L)      | O(L)          | Y ms         |
| 最佳化解法     | O(M×N×4^L)      | O(L + C)      | Z ms         |
```

## 7. Go 最佳實踐與測試

### Go 慣用解法
- 使用閉包函數實現 DFS，便於訪問外部變數
- 使用方向陣列簡化四個方向的嘗試
- 使用 map 進行字母頻率計數
- 代碼結構清晰，變數命名符合 Go 的慣例

### 錯誤處理與邊界情況
- 處理空網格的情況
- 檢查索引是否越界
- 在回溯過程中恢復修改的值
- 處理字母頻率不匹配的情況

## 8. 面試模擬

### 時間管理規劃
- 問題理解：~2-3 分鐘
- 初步解法提案：~5 分鐘
- 優化討論：~5-7 分鐘
- 代碼編寫：~10-15 分鐘
- 測試與除錯：~5 分鐘

### 面試官互動模擬
- 如何向面試官解釋我的 Go 實作：
    - "我打算使用深度優先搜索結合回溯法來解決這個問題。首先，我會從網格中的每個位置開始嘗試，看是否能找到匹配的路徑。然後，我會使用 DFS 遞迴地尋找後續字..."
- 潛在的提示與引導問題：
    - 面試官可能會問："你如何確保不重複使用同一個格子？"
    - 回答："我會使用一個臨時標記來記錄已訪問的格子，在回溯時恢復原始值。"
- 如何清晰地溝通我的實作決策：
    - 解釋為什麼選擇 DFS 而不是 BFS
    - 說明使用閉包函數的優勢
    - 描述前處理步驟如何提高效率

### 潛在的後續問題
- 如果網格非常大，你會如何修改你的 Go 代碼？
    - 可以考慮使用並行處理，為不同的起始位置創建 goroutine
    - 使用記憶化來避免重複計算相同的子問題
- 如果記憶體受到限制，你會如何優化你的解決方案？
    - 繼續使用在原網格上標記的方法
    - 可以考慮迭代而非遞迴實現，以減少棧空間使用
- 如何擴展你的解決方案來處理相關但更複雜的問題？
    - 例如，如果允許對角線移動，只需擴展方向陣列
    - 如果需要找到所有可能的路徑，可以使用一個結果集來收集它們

## 9. 知識整合與學習

### 問題解決洞察
- 這個問題教會我如何在 Go 中優雅地實作回溯法
- 深入理解了 Go 中閉包函數和遞迴的使用
- 需要進一步加強的知識點：Go 的並行處理在圖搜索中的應用

### 心智模型構建
- 從這個問題中提煉出的一般框架：
    1. 明確問題需要搜索哪些可能的狀態
    2. 設計狀態轉移方式（如四個方向的移動）
    3. 實作 DFS/回溯來探索狀態空間
    4. 考慮優化來減少不必要的搜索
- 這個框架可以應用於其他需要探索狀態空間的問題

### 錯誤模式識別
- 實作過程中的常見錯誤：
    - 忘記在回溯時恢復已修改的值
    - 邊界條件檢查不完整
    - DFS 終止條件設置不正確
- 這些錯誤揭示了在實作回溯法時需要特別注意的地方

### 知識圖譜擴展
- 可以進一步探索的 Go 模式：
    - 使用 `context` 包來控制搜索超時
    - 使用 goroutine 並行化獨立的搜索任務
    - 使用 `sync.Pool` 重用暫存資料結構