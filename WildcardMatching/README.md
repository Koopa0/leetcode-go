# 1. Original Problem:

## Wildcard Matching (LeetCode #44)

Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
- '?' Matches any single character.
- '*' Matches any sequence of characters (including an empty sequence).
  The matching should cover the entire input string (not partial).

**Examples:**
```
Example 1:
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:
Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence of characters.

Example 3:
Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
```

**Constraints:**
- 0 <= s.length, p.length <= 2000
- s contains only lowercase English letters.
- p contains only lowercase English letters, '?' or '*'.

## 萬用字元匹配 (LeetCode #44)

給定一個輸入字串 (s) 和一個模式 (p)，實現萬用字元模式匹配，支持 '?' 和 '*'，其中：
- '?' 匹配任何單個字。
- '*' 匹配任何字序列（包括空序列）。
  匹配應該覆蓋整個輸入字串（非部分匹配）。

**示例：**
```
示例 1：
輸入：s = "aa", p = "a"
輸出：false
解釋："a" 不匹配整個字串 "aa"。

示例 2：
輸入：s = "aa", p = "*"
輸出：true
解釋：'*' 匹配任何字序列。

示例 3：
輸入：s = "cb", p = "?a"
輸出：false
解釋：'?' 匹配 'c'，但第二個字母是 'a'，不匹配 'b'。
```

**約束條件：**
- 0 <= s.length, p.length <= 2000
- s 僅包含小寫英文字母。
- p 僅包含小寫英文字母、'?' 或 '*'。

# 2. 初始問題解析與心智模型建立

當我第一次看到這個問題時，我應該先理解什麼是萬用字元匹配。問題定義了兩個特殊字：
- '?' 匹配任何單個字
- '*' 匹配任何字序列（包括空序列）

關鍵在於確定模式是否能匹配整個輸入字串。這不是部分匹配 - 我們需要從頭到尾進行匹配。

通過分析示例，我可以更好地理解問題：
- 示例 1："aa" 不匹配 "a"，因為 "a" 只能匹配一個字，而 "aa" 有兩個字。
- 示例 2："aa" 匹配 "*"，因為 '*' 可以匹配任何字序列，包括 "aa"。
- 示例 3："cb" 不匹配 "?a"，因為雖然 '?' 匹配 'c'，但 'a' 不匹配 'b'。

輸入是兩個字串：源字串和模式字串。
輸出是一個布爾值：如果模式匹配字串則為 true，否則為 false。

這個轉換要求是檢查模式是否可以根據萬用字元規則擴展，以精確匹配輸入字串。

# 3. 問題理解與核心挑戰

此問題的核心挑戰是處理 '*' 字，它可以匹配任何字序列，包括空序列。這引入了很多歧義，因為單個 '*' 可以以多種不同方式匹配。

關鍵挑戰：
1. 確定每個 '*' 應該消耗多少字串（貪婪匹配 vs. 非貪婪匹配）
2. 處理連續的 '*' 字（它們實際上等同於一個 '*'）
3. 確保整個字串匹配，而不僅僅是部分匹配
4. 處理模式中 '?' 和 '*' 的組合
5. 處理邊緣情況，如空字串或僅由 '*' 組成的模式

計算挑戰是確定是否存在任何方法，可以在遵守萬用字元規則的情況下，將模式匹配到字串。

# 4. 視覺化問題表示

我們可以將此問題視覺化為一個網格，類似於動態規劃解決方案的視覺化方式：

- 行代表模式中的位置
- 列代表字串中的位置
- 每個單元格 (i,j) 回答問題："模式到位置 i 可以匹配字串到位置 j 嗎？"

例如，對於 s = "abc" 和 p = "a*c"：

```
    | '' | a | b | c |
--------------------- 
'' | T  | F | F | F |
a  | F  | T | F | F |
*  | F  | T | T | T |
c  | F  | F | F | T |
```

最終答案在右下角的單元格中，它告訴我們整個模式是否可以匹配整個字串。

# 5. 問題模式識別

這個問題具有動態規劃的特徵：
- 最優子結構：解決方案可以從更小子問題的解決方案構建
- 重疊子問題：同樣的子問題被多次解決

它也類似於正則表達式匹配，這是一個經典的動態規劃問題。

關鍵指標是：
- 需要做出影響未來匹配的決策（特別是使用 '*'）
- 匹配過程的遞歸性質
- 能夠將問題分解為更小的、類似的子問題

# 6. 策略性問題解決框架

對於帶有萬用字元或正則表達式的字串匹配問題，動態規劃通常是一種有效的方法。一般策略包括：

1. 定義狀態：dp[i][j] = 模式[0...i-1] 是否匹配字串[0...j-1]
2. 定義基本情況：
    - dp[0][0] = true（空模式匹配空字串）
    - dp[i][0] 取決於模式是否有 '*' 可以匹配空字串
    - dp[0][j] = false 對於 j > 0（空模式不能匹配非空字串）
3. 定義轉移函數：
    - 如果模式[i-1]是普通字，dp[i][j] = dp[i-1][j-1] && 模式[i-1] == 字串[j-1]
    - 如果模式[i-1]是 '?'，dp[i][j] = dp[i-1][j-1]
    - 如果模式[i-1]是 '*'，dp[i][j] = dp[i-1][j] || dp[i][j-1]
        - dp[i-1][j]：'*' 匹配空序列
        - dp[i][j-1]：'*' 匹配一個或多個字

另外，也可以使用帶有記憶化的遞歸方法，遵循相同的邏輯但實現方式不同。

# 7. 編碼前的演算法設計

讓我們勾勒一個動態規劃演算法：

```
function isMatch(s, p):
    // 創建一個 2D DP 表，尺寸為 (p.length+1) x (s.length+1)
    // dp[i][j] = 模式 p[0...i-1] 是否匹配字串 s[0...j-1]
    
    // 基本情況
    dp[0][0] = true  // 空模式匹配空字串
    
    // 空模式不能匹配非空字串
    for j from 1 to s.length:
        dp[0][j] = false
    
    // 處理可以匹配空字串的帶有 '*' 的模式
    for i from 1 to p.length:
        if p[i-1] == '*':
            dp[i][0] = dp[i-1][0]
        else:
            dp[i][0] = false
    
    // 填充 DP 表
    for i from 1 to p.length:
        for j from 1 to s.length:
            if p[i-1] == '*':
                // '*' 匹配空字串或 '*' 匹配一個或多個字
                dp[i][j] = dp[i-1][j] || dp[i][j-1]
            else if p[i-1] == '?' || p[i-1] == s[j-1]:
                // '?' 匹配任何字或字直接匹配
                dp[i][j] = dp[i-1][j-1]
            else:
                dp[i][j] = false
    
    return dp[p.length][s.length]
```

讓我們用一個例子驗證這個演算法：
s = "abc", p = "a*c"

我們的 DP 表開始於：
```
    | '' | a | b | c |
--------------------- 
'' | T  | F | F | F |
```

處理模式的第一行（'a'）後：
```
    | '' | a | b | c |
--------------------- 
'' | T  | F | F | F |
a  | F  | T | F | F |
```

處理模式的第二行（'*'）後：
```
    | '' | a | b | c |
--------------------- 
'' | T  | F | F | F |
a  | F  | T | F | F |
*  | F  | T | T | T |
```

處理模式的第三行（'c'）後：
```
    | '' | a | b | c |
--------------------- 
'' | T  | F | F | F |
a  | F  | T | F | F |
*  | F  | T | T | T |
c  | F  | F | F | T |
```

最終結果是 dp[3][3] = true，這意味著模式 "a*c" 匹配字串 "abc"，這是正確的。

# 8. 視覺化解釋

上圖展示了動態規劃表格的填充過程，我們可以通過它來理解演算法的執行過程：

1. 初始化第一行和第一列：
    - dp[0][0] = true（空模式匹配空字串）
    - dp[0][j] = false 對於 j > 0（空模式不匹配非空字串）
    - 對於模式中的 '*'，我們特殊處理其能匹配空字串的情況

2. 逐步填充表格：
    - 對於普通字和 '?'，我們檢查是否與當前字匹配
    - 對於 '*'，我們考慮它匹配空序列或匹配當前字並可能匹配更多字的情況

3. 最終結果在右下角單元格，即 dp[3][3] = true，表示 "a*c" 匹配 "abc"

這個可視化幫助我們直觀地理解了動態規劃算法如何解決萬用字元匹配問題。

# 9. 解決方案開發過程

## 暴力解法
最簡單的解決方案是嘗試所有可能的匹配方式，特別是對於 '*' 字。這將涉及一個遞歸函數，嘗試不同的方式來匹配模式。

```
function isMatch(s, p, i, j):
    if j == p.length:
        return i == s.length
    
    if p[j] != '*':
        return i < s.length && (p[j] == '?' || p[j] == s[i]) && isMatch(s, p, i+1, j+1)
    
    while i <= s.length:
        if isMatch(s, p, i, j+1):
            return true
        i++
    
    return false
```

這種方法在最壞情況下具有指數時間複雜度，因為我們可能需要嘗試許多不同的匹配方式。

## 遞歸 + 記憶化
改進的方法是使用帶有記憶化的遞歸，避免重新計算相同的子問題。

```
function isMatch(s, p):
    memo = {}
    
    function dp(i, j):
        if (i, j) in memo:
            return memo[(i, j)]
        
        if j == p.length:
            return i == s.length
        
        if p[j] != '*':
            match = i < s.length && (p[j] == '?' || p[j] == s[i])
            memo[(i, j)] = match && dp(i+1, j+1)
        else:
            k = i
            while k <= s.length:
                if dp(k, j+1):
                    memo[(i, j)] = true
                    return true
                k++
            memo[(i, j)] = false
        
        return memo[(i, j)]
    
    return dp(0, 0)
```

這種方法的時間複雜度是 O(m * n²)，其中 m 是模式的長度，n 是字串的長度。

## 動態規劃
最有效的方法是使用動態規劃，從小問題構建解決方案，如我們在前面的演算法設計中概述的那樣。這允許我們系統地考慮所有可能性，而不必重復計算。

時間複雜度：O(m * n)
空間複雜度：O(m * n)

# 10. 實際示例演練

讓我們使用示例 s = "abc", p = "a*c" 進行詳細演練：

1. 初始化 DP 表：
   dp[0][0] = true（空模式匹配空字串）
   dp[0][1] = dp[0][2] = dp[0][3] = false（空模式不匹配非空字串）

2. 處理模式的第一行（i=1，p[0]='a'）：
   dp[1][0] = false（'a' 不匹配空字串）
   dp[1][1]：因為 p[0]=='a' 且 s[0]=='a'，所以 dp[1][1] = dp[0][0] = true
   dp[1][2] = dp[1][3] = false（'a' 不匹配 'ab' 或 'abc'）

3. 處理模式的第二行（i=2，p[1]='*'）：
   dp[2][0] = dp[1][0] = false（在這種情況下，'a*' 不匹配空字串）
   dp[2][1] = dp[1][1] || dp[2][0] = true || false = true（'a*' 匹配 'a'）
   dp[2][2] = dp[1][2] || dp[2][1] = false || true = true（'a*' 匹配 'ab'）
   dp[2][3] = dp[1][3] || dp[2][2] = false || true = true（'a*' 匹配 'abc'）

4. 處理模式的第三行（i=3，p[2]='c'）：
   dp[3][0] = false（'a*c' 不匹配空字串）
   dp[3][1] = dp[3][2] = false（'a*c' 不匹配 'a' 或 'ab'）
   dp[3][3]：因為 p[2]=='c' 且 s[2]=='c'，所以 dp[3][3] = dp[2][2] = true

最終結果是 dp[3][3] = true，這意味著模式 "a*c" 匹配字串 "abc"。

# 11. Golang 實現

```go
func isMatch(s string, p string) bool {
    m, n := len(p), len(s)
    
    // 創建 DP 表
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    
    // 基本情況：空模式匹配空字串
    dp[0][0] = true
    
    // 處理可以匹配空字串的帶有 '*' 的模式
    for i := 1; i <= m; i++ {
        if p[i-1] == '*' {
            dp[i][0] = dp[i-1][0]
        }
    }
    
    // 填充 DP 表
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[i-1] == '*' {
                // '*' 匹配空序列或可以匹配當前字和可能更多
                dp[i][j] = dp[i-1][j] || dp[i][j-1]
            } else if p[i-1] == '?' || p[i-1] == s[j-1] {
                // '?' 匹配任何字或字直接匹配
                dp[i][j] = dp[i-1][j-1]
            }
            // 如果上述條件都不符合，dp[i][j] 保持為 false
        }
    }
    
    return dp[m][n]
}
```

# 12. 實現執行演練

讓我們使用示例 s = "abc", p = "a*c" 追蹤我們的 Go 實現：

1. 初始化 m = 3, n = 3（p 和 s 的長度）
2. 創建 dp 表，尺寸為 (4 x 4)
3. 設置 dp[0][0] = true（空模式匹配空字串）
4. 處理帶有 '*' 的模式：
   對於 i = 1，p[0] = 'a'，不是 '*'，所以 dp[1][0] 保持為 false
   對於 i = 2，p[1] = '*'，所以 dp[2][0] = dp[1][0] = false
   對於 i = 3，p[2] = 'c'，不是 '*'，所以 dp[3][0] 保持為 false
5. 填充 DP 表：
   對於 i = 1（p[0] = 'a'）：
   對於 j = 1（s[0] = 'a'）：
   因為 p[0] == s[0]，dp[1][1] = dp[0][0] = true
   對於 j = 2,3：
   因為 p[0] != '?' 且 p[0] != '*' 且 p[0] != s[1] 且 p[0] != s[2]，dp[1][2] 和 dp[1][3] 保持為 false

   對於 i = 2（p[1] = '*'）：
   對於 j = 1,2,3：
   因為 p[1] == '*'，我們應用規則：
   dp[2][1] = dp[1][1] || dp[2][0] = true || false = true
   dp[2][2] = dp[1][2] || dp[2][1] = false || true = true
   dp[2][3] = dp[1][3] || dp[2][2] = false || true = true

   對於 i = 3（p[2] = 'c'）：
   對於 j = 1,2：
   因為 p[2] != '?' 且 p[2] != '*' 且 p[2] != s[0] 且 p[2] != s[1]，dp[3][1] 和 dp[3][2] 保持為 false
   對於 j = 3：
   因為 p[2] == s[2]，dp[3][3] = dp[2][2] = true

6. 返回 dp[m][n] = dp[3][3] = true

# 13. 複雜度分析

## 時間複雜度
- 演算法遍歷模式和字串中的每個字一次，導致 O(m * n) 操作，其中 m 是模式的長度，n 是字串的長度。
- 構建 DP 表需要 O(m * n) 時間。

## 空間複雜度
- 演算法使用尺寸為 (m+1) x (n+1) 的 2D DP 表，導致 O(m * n) 空間。

在使用動態規劃方法時，時間和空間複雜度無法進一步改進。有一種貪婪演算法可以在 O(m + n) 時間和 O(1) 空間中解決此問題，但它更難以理解和正確實現。

# 14. 優化與改進

## 優化方向
1. **空間優化**：DP 方法可以優化為僅使用 O(n) 空間，使用 1D 數組而不是 2D 數組，因為我們只需要前一行來計算當前行。

```go
func isMatch(s string, p string) bool {
    m, n := len(p), len(s)
    
    // 當前行和前一行
    curr := make([]bool, n+1)
    prev := make([]bool, n+1)
    
    // 基本情況
    prev[0] = true
    
    for i := 1; i <= m; i++ {
        // 更新當前行的第一個元素
        if p[i-1] == '*' {
            curr[0] = prev[0]
        } else {
            curr[0] = false
        }
        
        for j := 1; j <= n; j++ {
            if p[i-1] == '*' {
                curr[j] = prev[j] || curr[j-1]
            } else if p[i-1] == '?' || p[i-1] == s[j-1] {
                curr[j] = prev[j-1]
            } else {
                curr[j] = false
            }
        }
        
        // 交換當前行和前一行
        prev, curr = curr, prev
    }
    
    return prev[n]
}
```

2. **貪婪演算法**：對於某些模式，貪婪方法可能比 DP 更快。具體來說，我們可以從左到右貪婪地匹配模式，必要時進行回溯。

3. **字跳過**：我們可以通過跳過模式中連續的 '*' 字來優化演算法，因為它們等同於一個 '*'。

```go
func isMatch(s string, p string) bool {
    // 預處理：去除連續的 '*'
    var newP strings.Builder
    for i := 0; i < len(p); i++ {
        if i == 0 || p[i] != '*' || p[i-1] != '*' {
            newP.WriteByte(p[i])
        }
    }
    
    p = newP.String()
    
    // 然後使用之前的 DP 解決方案
    // ...
}
```

# 15. 一般問題解決智慧

這個問題教會我們：
1. 動態規劃在字串匹配問題中的強大能力
2. 將複雜問題分解為更簡單子問題的重要性
3. 需要仔細處理萬用字元並理解它們的含義
4. 使用網格或表格可視化問題空間的好處

要為類似問題開發演算法直覺：
1. 首先手動處理示例，以了解模式匹配規則
2. 識別問題的遞歸結構
3. 考慮邊緣情況，如空字串和僅包含萬用字元的模式
4. 當問題具有最優子結構和重疊子問題時，使用動態規劃等系統方法

在實際編程面試中，這個問題教導我們如何將理論知識應用於實際情況，以及如何在有多種可能解決方案時選擇最合適的方法。

# 16. 測試策略

這個測試套件包含了各種情況：
1. 基本匹配和不匹配情況
2. 空字串和模式的邊緣情況
3. 僅包含萬用字元的模式
4. 包含多個萬用字元的複雜模式
5. 非常長的輸入以測試性能和正確性

通過這種表驅動測試方法，我們可以系統地驗證我們的解決方案在各種情況下的正確性。