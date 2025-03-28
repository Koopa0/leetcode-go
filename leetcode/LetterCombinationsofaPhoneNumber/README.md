## 1. Original Problem

### English:
**Letter Combinations of a Phone Number**

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

```
2 -> abc
3 -> def
4 -> ghi
5 -> jkl
6 -> mno
7 -> pqrs
8 -> tuv
9 -> wxyz
```

**Example 1:**
```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Example 2:**
```
Input: digits = ""
Output: []
```

**Example 3:**
```
Input: digits = "2"
Output: ["a","b","c"]
```

**Constraints:**
- 0 <= digits.length <= 4
- digits[i] is a digit in the range ['2', '9'].

### 繁體中文：
**電話號碼的字母組合**

給定一個包含 2-9 的字符串，返回所有它能表示的字母組合。答案可以按任意順序返回。

數字到字母的映射如下（與電話按鍵相同）。注意，1 不對應任何字母。

```
2 -> abc
3 -> def
4 -> ghi
5 -> jkl
6 -> mno
7 -> pqrs
8 -> tuv
9 -> wxyz
```

**示例 1：**
```
輸入：digits = "23"
輸出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**示例 2：**
```
輸入：digits = ""
輸出：[]
```

**示例 3：**
```
輸入：digits = "2"
輸出：["a","b","c"]
```

**約束條件：**
- 0 <= digits.length <= 4
- digits[i] 是範圍 ['2', '9'] 中的一個數字。

## 2. 問題理解

這個問題本質上是要求我們根據手機鍵盤上的數字到字母的映射，找出所有可能的字母組合。

關鍵點：
- 每個數字對應多個字母
- 我們需要返回所有可能的組合
- 輸入字符串長度不超過 4 個字符
- 只處理 2-9 的數字（不包括 0、1）

邊緣情況：
- 空字符串應返回空列表
- 單個數字應返回其對應的所有字母

這個問題實際上是一個排列組合問題，我們需要找出所有可能的組合方式。例如，對於 "23"，我們需要把 2 對應的每個字母 ('a', 'b', 'c') 與 3 對應的每個字母 ('d', 'e', 'f') 組合起來，生成所有可能的組合。

## 3. 視覺解釋

讓我們用一個決策樹來說明 "23" 的所有可能組合：

```
                 ""
               /  |  \
              a   b   c  (對應 2 的字母)
            / | \/ | \/ | \
           ad ae af bd be bf cd ce cf  (對應 3 的字母)
```

這個樹顯示了我們如何從空字符串開始，首先添加 2 對應的每個字母，然後再添加 3 對應的每個字母，最終得到所有可能的組合。

## 4. 思考過程

對於這個問題，我們可以考慮幾種解決方案：

1. **回溯法（DFS）**：我們可以使用回溯算法，從空字符串開始，逐步添加每個數字對應的字母，直到處理完所有數字。這是一種深度優先的方法。

2. **迭代法**：我們可以使用迭代方法，維護一個結果列表，然後逐步處理每個數字，更新結果列表。

3. **隊列（BFS）**：我們可以使用隊列實現廣度優先搜索，逐層處理每個數字對應的字母。

對於這種排列組合問題，回溯法是最直觀且常用的方法，因為它能夠系統地生成所有可能的組合。迭代法也是一個不錯的選擇，尤其是當我們需要避免遞歸調用的開銷時。考慮到問題規模較小（最多 4 個數字），任何方法都應該能夠高效解決這個問題。

## 5. 最佳解決方案開發

讓我們從回溯法開始，這是解決這類問題的常用且優雅的方法。

回溯法的基本思路：
1. 建立數字到字母的映射
2. 從空字符串開始，逐步添加字母
3. 使用遞歸函數來嘗試所有可能的組合

讓我們以 "23" 為例，演示整個過程：

1. 初始化結果列表 `result = []`，當前組合 `current = ""`
2. 對於第一個數字 '2'：
    - 嘗試 'a'：current = "a"，繼續處理下一個數字
    - 嘗試 'b'：current = "b"，繼續處理下一個數字
    - 嘗試 'c'：current = "c"，繼續處理下一個數字
3. 對於第二個數字 '3'：
    - 當 current = "a" 時，嘗試 'd'、'e'、'f'，得到 "ad"、"ae"、"af"
    - 當 current = "b" 時，嘗試 'd'、'e'、'f'，得到 "bd"、"be"、"bf"
    - 當 current = "c" 時，嘗試 'd'、'e'、'f'，得到 "cd"、"ce"、"cf"
4. 添加所有生成的組合到結果列表中

最終，我們得到 ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]。

## 7. 複雜度分析

**時間複雜度：**
- 假設輸入字符串長度為 n，每個數字平均對應 m 個字母（m 約為 3-4）
- 每個數字可以選擇其對應的 m 個字母中的一個
- 總共有 m^n 種可能的組合
- 生成每個組合需要 O(n) 的時間（字符串連接）
- 因此，總時間複雜度為 O(m^n * n)

對於本題約束條件（n <= 4），最壞情況下時間複雜度為 O(4^4 * 4) = O(1024)，這是可接受的。

**空間複雜度：**
- 除了輸出結果外，遞歸調用棧的深度為 O(n)
- 輸出結果需要 O(m^n * n) 的空間
- 因此，總空間複雜度為 O(m^n * n)

## 9. 測試策略

這個測試涵蓋了各種情況：
- 空字符串
- 單個數字
- 兩個數字
- 三個數字
- 包含有 4 個字母的數字（7 和 9）
- 最大長度的輸入（4 個數字）


## 回溯法詳細解釋

回溯法是一種通過嘗試所有可能性來解決問題的算法。在這個問題中，我們需要嘗試每個數字對應的所有字母的組合。

### 1. 定義回溯函數

```go
var backtrack func(index int, current string)
```

這行代碼聲明了一個名為 `backtrack` 的函數變量，它接受兩個參數：
- `index`：當前處理到輸入字符串的第幾個數字（索引）
- `current`：到目前為止已經形成的字母組合

### 2. 實現回溯函數

```go
backtrack = func(index int, current string) {
    // 如果當前組合長度等於輸入長度，表示已經處理完所有數字
    if len(current) == len(digits) {
        result = append(result, current)
        return
    }
    
    // 獲取當前數字對應的字母
    digit := digits[index]
    letters := phoneMap[digit]
    
    // 嘗試每個字母
    for i := 0; i < len(letters); i++ {
        // 添加當前字母到組合中
        backtrack(index+1, current+string(letters[i]))
    }
}
```

這個函數的工作方式如下：

1. **終止條件**：
   ```go
   if len(current) == len(digits) {
       result = append(result, current)
       return
   }
   ```
   當我們已經處理完所有數字時（即 `current` 的長度等於 `digits` 的長度），我們將當前組合添加到結果列表中，然後返回。

2. **處理當前數字**：
   ```go
   digit := digits[index]
   letters := phoneMap[digit]
   ```
   獲取當前正在處理的數字（`digits[index]`），以及該數字對應的所有可能字母（`phoneMap[digit]`）。

3. **嘗試所有可能性**：
   ```go
   for i := 0; i < len(letters); i++ {
       backtrack(index+1, current+string(letters[i]))
   }
   ```
   對於當前數字對應的每個字母，我們將其添加到當前組合中，然後遞歸地處理下一個數字（`index+1`）。

### 3. 啟動回溯

```go
backtrack(0, "")
```

我們從空字符串開始，處理第一個數字（索引為 0）。

### 詳細執行過程示例

讓我們以輸入 `"23"` 為例，詳細跟蹤回溯函數的執行過程：

1. 調用 `backtrack(0, "")`：
    - `index = 0`，`current = ""`
    - 獲取 `digits[0] = '2'` 對應的字母 `"abc"`
    - 遍歷 `"abc"` 中的每個字母：
        - 對於 'a'：調用 `backtrack(1, "a")`
        - 對於 'b'：調用 `backtrack(1, "b")`
        - 對於 'c'：調用 `backtrack(1, "c")`

2. 調用 `backtrack(1, "a")`：
    - `index = 1`，`current = "a"`
    - 獲取 `digits[1] = '3'` 對應的字母 `"def"`
    - 遍歷 `"def"` 中的每個字母：
        - 對於 'd'：調用 `backtrack(2, "ad")`
        - 對於 'e'：調用 `backtrack(2, "ae")`
        - 對於 'f'：調用 `backtrack(2, "af")`

3. 調用 `backtrack(2, "ad")`：
    - `index = 2`，`current = "ad"`
    - 因為 `len(current) = 2` 等於 `len(digits) = 2`，所以將 `"ad"` 添加到結果列表中，然後返回

4. 調用 `backtrack(2, "ae")`：
    - `index = 2`，`current = "ae"`
    - 因為 `len(current) = 2` 等於 `len(digits) = 2`，所以將 `"ae"` 添加到結果列表中，然後返回

5. 調用 `backtrack(2, "af")`：
    - `index = 2`，`current = "af"`
    - 因為 `len(current) = 2` 等於 `len(digits) = 2`，所以將 `"af"` 添加到結果列表中，然後返回

6. 返回到 `backtrack(0, "")`，繼續處理 'b'：
    - 調用 `backtrack(1, "b")`
    - 依次生成 `"bd"`, `"be"`, `"bf"`

7. 返回到 `backtrack(0, "")`，繼續處理 'c'：
    - 調用 `backtrack(1, "c")`
    - 依次生成 `"cd"`, `"ce"`, `"cf"`

最終，我們得到結果：`["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]`

### 回溯樹

可以將整個過程想像成一棵決策樹：

```
                       ""
                    /  |  \
                   a   b   c   (處理數字 2)
                 / | \/ | \/ | \
                ad ae af bd be bf cd ce cf  (處理數字 3)
```

在這棵樹中，每一層代表一個數字，每個節點代表到當前為止形成的組合。當我們到達葉子節點時，就得到了一個完整的組合。

### 回溯法與迭代法的比較

回溯法通過遞歸和"嘗試-回退"的方式來尋找所有可能的解，它的思想是：
1. 嘗試一種可能性
2. 如果這種可能性不行，就回退到上一步
3. 嘗試下一種可能性

而迭代法則是通過維護一個當前結果的列表，不斷地擴展這個列表來得到所有可能的組合。

回溯法通常更直觀地反映了問題的結構，特別是對於這種排列組合問題。而迭代法可能更高效，因為它避免了函數調用的開銷。