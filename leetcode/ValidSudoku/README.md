# Valid Sudoku

## 1. Original Problem:

### English:
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.
- The Sudoku board is represented as a 9x9 2D array of characters where '.' represents empty cells and '1'-'9' represent filled cells.

### 繁體中文:
判斷一個 9 x 9 的數獨是否有效。只需要根據以下規則，驗證已經填入的數字是否有效即可：

1. 數字 1-9 在每一行只能出現一次。
2. 數字 1-9 在每一列只能出現一次。
3. 數字 1-9 在每一個以粗實線分隔的 3 x 3 宮格內只能出現一次。

注意:
- 一個有效的數獨（部分已填充）並不一定是可解的。
- 只需要根據上述規則，驗證已經填入的數字是否有效即可。
- 數獨由一個 9x9 的二維字數組表示，其中 '.' 表示空格，而 '1'-'9' 表示填入的數字。

## 2. 初始問題解析與心智模型建立

當我們初次看到這個數獨問題時，首先需要理解什麼是「有效的數獨」。根據問題描述，我們不需要解決數獨或判斷它是否可解，只需確認目前已填入的數字是否符合數獨的三個基本規則。

讓我們仔細分析問題提供的資訊：
- 輸入是一個 9x9 的二維字數組
- '.' 代表空格，'1'-'9' 代表已填入的數字
- 我們只需要檢查已填入的數字是否違反規則

在心智模型中，我可以將數獨板想像成由三種不同的集合所構成：
1. 9 條水平行
2. 9 條垂直列
3. 9 個 3x3 的子方格

我們需要確保在每一個集合中，數字 1-9 都不會重複出現。空格（'.'）不需要檢查，我們可以直接忽略它們。

## 3. 問題理解與核心挑戰

核心挑戰在於：
1. 如何有效地檢查每一行、每一列和每一個 3x3 子方格中是否有重複數字
2. 如何確定一個單元格屬於哪個 3x3 子方格

關鍵考慮因素：
- 我們需要跟蹤每行、每列和每個子方格中已經出現過的數字
- 對於每個非空格單元格，我們需要檢查其數字是否已在其所屬的行、列或子方格中出現過
- 我們需要一個高效的方法來計算單元格所屬的子方格索引

這個問題本質上是一個「約束檢查」問題，我們需要確保每個數字在特定的分組中只出現一次。

## 4. 視覺問題表示

以下是一個數獨板的視覺表示：

```
+-------+-------+-------+
| 5 3 . | . 7 . | . . . |
| 6 . . | 1 9 5 | . . . |
| . 9 8 | . . . | . 6 . |
+-------+-------+-------+
| 8 . . | . 6 . | . . 3 |
| 4 . . | 8 . 3 | . . 1 |
| 7 . . | . 2 . | . . 6 |
+-------+-------+-------+
| . 6 . | . . . | 2 8 . |
| . . . | 4 1 9 | . . 5 |
| . . . | . 8 . | . 7 9 |
+-------+-------+-------+
```

每個單元格可以是空的（表示為'.'）或包含數字 1-9。

在檢查時，我們需要：
1. 檢查每一行（水平線）
2. 檢查每一列（垂直線）
3. 檢查每一個 3x3 子方格（由粗線分隔）

對於子方格索引的計算，我們可以使用這個公式：`boxIndex = (row/3)*3 + (col/3)`，其中 row 和 col 是單元格的行列索引（從 0 開始）。這樣可以將 9x9 的網格映射到 9 個子方格，如下所示：

```
+-------+-------+-------+
| Box 0 | Box 1 | Box 2 |
+-------+-------+-------+
| Box 3 | Box 4 | Box 5 |
+-------+-------+-------+
| Box 6 | Box 7 | Box 8 |
+-------+-------+-------+
```

## 5. 問題模式識別

這個問題屬於「約束驗證」類型，我們需要檢查特定集合中是否有重複元素。類似的問題模式包括：
- 棋盤遊戲規則驗證
- 集合重複性檢查
- 具有區域約束的問題

對於這類問題，哈希集合（HashSet）或哈希映射（HashMap）是高效的數據結構，因為它們提供 O(1) 的查詢時間複雜度，可以快速檢查元素是否已存在。

## 6. 策略性解題框架

對於此類約束檢查問題，一個有效的策略是：

1. 明確定義違反規則的情況（即在同一行、同一列或同一子方格中出現重複數字）
2. 選擇適當的數據結構來跟蹤已見過的數字
3. 設計一個能夠同時檢查所有約束的單通解決方案
4. 高效地確定每個單元格所屬的子方格

解題思路：
- 使用哈希集合來跟蹤每行、每列和每個子方格中已經出現的數字
- 遍歷數獨板的每個單元格，只處理非空格單元格
- 對於每個非空格單元格，計算其所屬子方格的索引，然後檢查其數字是否已在相應的行、列或子方格中出現過
- 如果發現重複，則數獨無效；如果遍歷完整個數獨板沒有發現重複，則數獨有效

## 7. 算法設計

以下是我們的算法設計：

```
函數 isValidSudoku(board):
    # 初始化三組哈希集合來跟蹤已見數字
    rows = [空集合 x 9]     # 每行一個集合
    cols = [空集合 x 9]     # 每列一個集合
    boxes = [空集合 x 9]    # 每個3x3子方格一個集合
    
    # 遍歷數獨板的每個單元格
    對於 row 從 0 到 8:
        對於 col 從 0 到 8:
            # 跳過空格
            如果 board[row][col] == '.':
                繼續
                
            digit = board[row][col]
            
            # 計算3x3子方格的索引（0-8）
            boxIndex = (row/3)*3 + (col/3)
            
            # 檢查數字是否已經在行、列或子方格中出現過
            如果 digit 在 rows[row] 或 digit 在 cols[col] 或 digit 在 boxes[boxIndex]:
                return False  # 無效的數獨
                
            # 將數字添加到對應的集合中
            rows[row].add(digit)
            cols[col].add(digit)
            boxes[boxIndex].add(digit)
    
    # 如果遍歷完整個數獨板沒有發現重複，則數獨有效
    return True
```

這個算法的優點：
- 只需一次遍歷數獨板即可檢查所有三個約束
- 使用哈希集合進行 O(1) 時間複雜度的查詢和插入操作
- 使用整數除法有效地計算子方格索引

## 8. 視覺解釋

讓我們通過一個具體例子來視覺化算法的工作過程：

考慮數獨板的前幾個單元格：

```
+-------+-------+-------+
| 5 3 . | . 7 . | . . . |
| 6 . . | 1 9 5 | . . . |
| . 9 8 | . . . | . 6 . |
+-------+-------+-------+
...
```

初始化集合：
- rows = [set(), set(), ..., set()] (9個空集合)
- cols = [set(), set(), ..., set()] (9個空集合)
- boxes = [set(), set(), ..., set()] (9個空集合)

處理單元格 (0,0)，值為 '5'：
- boxIndex = (0/3)*3 + (0/3) = 0
- '5' 不在 rows[0]、cols[0] 或 boxes[0] 中
- 將 '5' 添加到 rows[0]、cols[0] 和 boxes[0]

處理單元格 (0,1)，值為 '3'：
- boxIndex = (0/3)*3 + (1/3) = 0
- '3' 不在 rows[0]、cols[1] 或 boxes[0] 中
- 將 '3' 添加到 rows[0]、cols[1] 和 boxes[0]

處理單元格 (0,2)，值為 '.'：
- 這是空格，跳過

如果遇到重複數字，例如假設在第一行中有另一個 '5'：

```
+-------+-------+-------+
| 5 3 5 | . 7 . | . . . |   ← 重複的 '5'
...
```

當處理第二個 '5' 時：
- 檢查 '5' 是否在 rows[0] 中 → 是的，它在！
- 立即返回 False，表示數獨無效

## 9. 解決方案發展歷程

我們可以考慮多種方法來解決這個問題：

1. 暴力解法（三次遍歷）：
    - 第一次遍歷檢查每一行
    - 第二次遍歷檢查每一列
    - 第三次遍歷檢查每一個子方格

   這種方法直觀但效率較低，需要三次遍歷。

2. 優化解法（單次遍歷）：
    - 使用一次遍歷同時檢查所有約束
    - 為每行、每列和每個子方格維護單獨的集合

   這種方法效率更高，只需一次遍歷。

3. 進一步優化（編碼法）：
    - 不使用多個集合，而是將位置和數字編碼為唯一的字串
    - 例如，將「第 1 行中的數字 5」編碼為「r1-5」

   這種方法可以使用單一集合，但可能降低可讀性。

我們選擇第二種方法，因為它在效率和可讀性之間取得了良好的平衡。

## 10. 實例步驟解析

讓我們詳細跟蹤算法如何處理前面的數獨板例子：

初始狀態：
- rows = [set(), set(), ..., set()] (9個空集合)
- cols = [set(), set(), ..., set()] (9個空集合)
- boxes = [set(), set(), ..., set()] (9個空集合)

第 1 步：處理單元格 (0,0)，值為 '5'
- boxIndex = (0/3)*3 + (0/3) = 0
- 檢查：'5' 不在 rows[0]、cols[0] 或 boxes[0] 中
- 操作：將 '5' 添加到 rows[0]、cols[0] 和 boxes[0]
- 結果：rows[0] = {'5'}, cols[0] = {'5'}, boxes[0] = {'5'}

第 2 步：處理單元格 (0,1)，值為 '3'
- boxIndex = (0/3)*3 + (1/3) = 0
- 檢查：'3' 不在 rows[0]、cols[1] 或 boxes[0] 中
- 操作：將 '3' 添加到 rows[0]、cols[1] 和 boxes[0]
- 結果：rows[0] = {'5','3'}, cols[1] = {'3'}, boxes[0] = {'5','3'}

第 3 步：處理單元格 (0,2)，值為 '.'
- 這是空格，跳過

第 4 步：處理單元格 (0,3)，值為 '.'
- 這是空格，跳過

第 5 步：處理單元格 (0,4)，值為 '7'
- boxIndex = (0/3)*3 + (4/3) = 1
- 檢查：'7' 不在 rows[0]、cols[4] 或 boxes[1] 中
- 操作：將 '7' 添加到 rows[0]、cols[4] 和 boxes[1]
- 結果：rows[0] = {'5','3','7'}, cols[4] = {'7'}, boxes[1] = {'7'}

繼續這樣的處理直到遍歷完整個數獨板。由於這個數獨板是有效的，我們不會發現任何違反規則的情況，最終返回 True。

## 12. 實現執行解析

讓我們使用前面的範例數獨板追蹤 Golang 實現的執行過程：

初始化：
- rows[0...8]：9 個空映射
- cols[0...8]：9 個空映射
- boxes[0...8]：9 個空映射

對於單元格 (0,0)，值為 '5'：
- boxIndex = (0/3)*3 + (0/3) = 0
- 檢查 rows[0]['5']、cols[0]['5'] 和 boxes[0]['5'] 是否為 true → 都是 false
- 設置 rows[0]['5'] = true, cols[0]['5'] = true, boxes[0]['5'] = true

對於單元格 (0,1)，值為 '3'：
- boxIndex = (0/3)*3 + (1/3) = 0
- 檢查 rows[0]['3']、cols[1]['3'] 和 boxes[0]['3'] 是否為 true → 都是 false
- 設置 rows[0]['3'] = true, cols[1]['3'] = true, boxes[0]['3'] = true

流程繼續，直到處理完所有單元格。由於數獨板是有效的，函數最終返回 true。

## 13. 複雜度分析

時間複雜度：
- 我們遍歷 9x9 數獨板的每個單元格，共 81 個單元格
- 對於每個單元格，我們執行常數時間的操作（哈希映射的查詢和插入）
- 因此，時間複雜度是 O(1)，因為數獨板的大小是固定的
- 如果推廣到 n×n 的數獨板，時間複雜度將是 O(n²)

空間複雜度：
- 我們使用 9 個行映射、9 個列映射和 9 個子方格映射，每個最多包含 9 個元素
- 總空間複雜度是 O(3*9*9) = O(243)，簡化為 O(1)，因為數獨板的大小是固定的
- 對於 n×n 的數獨板，空間複雜度將是 O(n²)

## 14. 優化與改進

我們的解決方案已經相當高效，但仍有一些可能的優化方向：

1. 使用位操作：
    - 由於我們只需要跟蹤數字 1-9，可以使用一個 9 位的整數來表示每行、每列和每個子方格中出現過的數字
    - 將第 i 位設為 1 表示數字 i+1 已出現
    - 這種方法更節省內存，但可能降低可讀性

2. 使用單一集合與編碼：
    - 不使用多個集合，而是將行、列和方格信息編碼為唯一的字串
    - 例如，使用像 "r1-5"、"c2-5"、"b0-5" 這樣的字串
    - 這樣可以只使用一個集合來跟蹤所有約束

3. 優化子方格索引計算：
    - 雖然 `(row/3)*3 + (col/3)` 已經是一個高效的公式，但在一些語言中，可以使用位操作來進一步優化
    - 例如，`(row/3)*3 + (col/3)` 可以重寫為 `(row/3<<1) + (row/3) + (col/3)`

在實際應用中，我們的原始解決方案在清晰度和效率之間取得了良好的平衡，通常是最佳選擇。

## 15. 一般解題智慧

這個問題教會我們幾個重要的算法設計原則：

1. 數據結構的選擇至關重要：
    - 哈希集合/映射是檢查重複項的理想選擇，提供 O(1) 的查詢時間
    - 適當的數據結構可以顯著簡化解決方案並提高效率

2. 多維索引的映射：
    - 將二維索引（行、列）映射到一維索引（子方格）的能力是一個重要技巧
    - 類似的映射在許多網格問題中都很有用

3. 單通算法的效率：
    - 通過聰明的設計，我們可以在一次遍歷中檢查多個約束
    - 這種方法避免了重複工作，提高了效率

4. 約束驗證模式：
    - 這個問題展示了「約束驗證」的通用模式
    - 類似的技術可應用於其他需要檢查集合屬性的問題

這些原則不僅適用於解決數獨問題，還可應用於許多其他算法問題：
- N 皇后問題
- 生命遊戲
- 單詞搜索
- 具有區域約束的其他網格問題

## 16. 測試策略

這個測試套件包含多種情況：
- 有效的部分填充數獨板
- 行中有重複的無效數獨
- 列中有重複的無效數獨
- 子方格中有重複的無效數獨
- 空數獨板（應該是有效的）
- 完全填充的有效數獨板
