## 1. Original Problem

### English Version
**Find the Index of the First Occurrence in a String**

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

**Example 1:**
```
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
```

**Example 2:**
```
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
```

**Constraints:**
- 1 <= haystack.length, needle.length <= 10^4
- haystack and needle consist of only lowercase English characters.

### 繁體中文翻譯
**找出字串中第一個匹配項的下標**

給你兩個字串 `haystack` 和 `needle`，請你在 `haystack` 字串中找出 `needle` 字串的第一個匹配項的下標（下標從 0 開始）。如果 `needle` 不是 `haystack` 的一部分，則返回 `-1`。

**示例 1：**
```
輸入：haystack = "sadbutsad", needle = "sad"
輸出：0
解釋："sad" 在下標 0 和 6 處匹配。
第一個匹配項的下標是 0，所以返回 0。
```

**示例 2：**
```
輸入：haystack = "leetcode", needle = "leeto"
輸出：-1
解釋："leeto" 沒有在 "leetcode" 中出現，所以返回 -1。
```

**約束條件：**
- 1 <= haystack.length, needle.length <= 10^4
- haystack 和 needle 僅由小寫英文字組成。

## 2. 問題理解

這個問題要求我們在一個主字串（haystack）中找到一個子字串（needle）第一次出現的位置。若找到了，返回該位置的索引值；若沒找到，則返回 -1。

核心要求：
- 我們需要檢查 needle 是否是 haystack 的子串
- 如果是，返回 needle 在 haystack 中第一次出現的起始索引
- 如果不是，返回 -1

主要挑戰：
- 如何高效地在主字串中查找子字串
- 處理各種邊界情況，比如當 needle 為空時
- 考慮時間複雜度和空間複雜度的優化

需要注意的邊界情況：
- 當 needle 的長度大於 haystack 的長度時，肯定找不到匹配
- 當 needle 為空字串時（題目約束條件說明這種情況不會發生）

## 3. 視覺解釋

讓我們用一個例子來說明這個問題：

haystack = "sadbutsad", needle = "sad"

我們需要檢查 "sad" 在 "sadbutsad" 中第一次出現的位置。

```
索引：   0 1 2 3 4 5 6 7 8
haystack: s a d b u t s a d
           ↑ ↑ ↑
needle:    s a d
```

我們看到 needle "sad" 在 haystack 的索引 0 處首次出現。

如果我們考慮另一個例子：

haystack = "leetcode", needle = "leeto"

```
索引：   0 1 2 3 4 5 6 7
haystack: l e e t c o d e
           ↑ ↑ ↑ ↑ ×
needle:    l e e t o
```

在這個例子中，我們看到 "leeto" 與 "leetcode" 不匹配，因為第 5 個字不同，所以返回 -1。

## 4. 思考過程

對於這個字串匹配問題，我有幾種解決方法：

### 方法 1: 暴力解法
最直接的方法是檢查 haystack 的每個可能起始位置，看看從該位置開始的子串是否與 needle 匹配。

步驟：
1. 遍歷 haystack 的每個字（直到 haystack.length - needle.length）
2. 對於每個位置 i，檢查從 i 開始的 needle.length 個字是否與 needle 相匹配
3. 如果找到匹配，返回索引 i
4. 如果遍歷完成後沒有找到匹配，返回 -1

時間複雜度：O(n*m)，其中 n 是 haystack 的長度，m 是 needle 的長度。

### 方法 2: KMP 算法
Knuth-Morris-Pratt (KMP) 算法是一種更高效的字串匹配算法，它利用已經匹配的信息來避免不必要的比較。

KMP 算法的關鍵是構建一個「部分匹配表」（也稱為「前綴函數」或「失配函數」），用來記錄每個位置的最長相同前綴後綴長度，從而在不匹配時知道應該回退到哪個位置繼續匹配。

時間複雜度：O(n+m)，其中 n 是 haystack 的長度，m 是 needle 的長度。

### 方法 3: 內置函數
在實際編程中，許多語言提供了內置的字串查找函數。在 Go 中，我們可以使用 `strings.Index` 函數。然而，為了學習目的，我們應該理解並實現自己的解決方案。

對於這個問題，我認為 KMP 算法是最優解，但暴力解法更容易理解和實現，且在大多數情況下也能高效工作。

## 5. 最優解決方案開發

考慮到問題的規模和約束，我將先實現簡單的暴力解法，然後再介紹 KMP 算法作為優化。

### 暴力解法詳解：

1. 如果 needle 的長度大於 haystack 的長度，直接返回 -1（不可能找到匹配）
2. 遍歷 haystack 的每個可能起始位置 i（從 0 到 haystack.length - needle.length）
3. 對於每個位置 i，檢查從該位置開始的 needle.length 個字是否與 needle 匹配
4. 如果找到完全匹配，返回索引 i
5. 如果遍歷完成後沒有找到匹配，返回 -1

讓我們使用示例 haystack = "sadbutsad", needle = "sad" 來追踪這個算法的執行：

1. 首先檢查位置 i = 0:
   ```
   haystack[0:3] = "sad"
   needle = "sad"
   ```
   匹配成功，返回 0。

如果沒有找到匹配，我們會繼續檢查下一個位置，直到找到匹配或遍歷完所有可能的起始位置。

### KMP 算法詳解：

KMP 算法由以下步驟組成：

1. 構建 needle 的部分匹配表（PMT）或稱為失配函數（next 數組）
2. 使用該表來指導在 haystack 中搜索 needle 的過程

部分匹配表的每個元素 PMT[i] 表示 needle[0...i] 的最長相同前綴後綴的長度。

以 needle = "sad" 為例：
- PMT[0] = 0（字 "s" 沒有相同的前綴後綴）
- PMT[1] = 0（字串 "sa" 沒有相同的前綴後綴）
- PMT[2] = 0（字串 "sad" 沒有相同的前綴後綴）

所以 PMT = [0, 0, 0]。

對於更複雜的例子，比如 needle = "abacab"：
- PMT[0] = 0（字 "a" 沒有相同的前綴後綴）
- PMT[1] = 0（字串 "ab" 沒有相同的前綴後綴）
- PMT[2] = 1（字串 "aba" 的前綴 "a" 與後綴 "a" 相同）
- PMT[3] = 0（字串 "abac" 沒有相同的前綴後綴）
- PMT[4] = 1（字串 "abaca" 的前綴 "a" 與後綴 "a" 相同）
- PMT[5] = 2（字串 "abacab" 的前綴 "ab" 與後綴 "ab" 相同）

所以 PMT = [0, 0, 1, 0, 1, 2]。

使用這個 PMT 表，KMP 算法可以在發生不匹配時，利用已知的信息來跳過一些不必要的比較，從而提高效率。

## 6. Golang 實現

讓我們首先實現暴力解法，然後再實現 KMP 算法。

### 暴力解法：

```go
func strStr(haystack string, needle string) int {
    // 獲取兩個字串的長度
    n, m := len(haystack), len(needle)
    
    // 如果 needle 為空，返回 0（題目約束已確保 needle 不為空）
    if m == 0 {
        return 0
    }
    
    // 如果 needle 長度大於 haystack 長度，不可能找到匹配
    if m > n {
        return -1
    }
    
    // 遍歷 haystack 中所有可能的起始位置
    for i := 0; i <= n-m; i++ {
        // 假設從當前位置開始匹配
        matched := true
        
        // 檢查從位置 i 開始的子串是否與 needle 匹配
        for j := 0; j < m; j++ {
            if haystack[i+j] != needle[j] {
                matched = false
                break
            }
        }
        
        // 如果完全匹配，返回起始索引
        if matched {
            return i
        }
    }
    
    // 如果沒有找到匹配，返回 -1
    return -1
}
```

### KMP 算法：

```go
func strStr(haystack string, needle string) int {
    // 獲取兩個字串的長度
    n, m := len(haystack), len(needle)
    
    // 如果 needle 為空，返回 0（題目約束已確保 needle 不為空）
    if m == 0 {
        return 0
    }
    
    // 如果 needle 長度大於 haystack 長度，不可能找到匹配
    if m > n {
        return -1
    }
    
    // 構建部分匹配表（PMT）
    // next[i] 表示 needle[0...i] 的最長相同前綴後綴長度
    next := make([]int, m)
    
    // 初始化
    next[0] = 0
    
    // 填充 next 數組
    for i, j := 1, 0; i < m; {
        if needle[i] == needle[j] {
            j++
            next[i] = j
            i++
        } else if j > 0 {
            j = next[j-1]
        } else {
            next[i] = 0
            i++
        }
    }
    
    // 使用 KMP 算法在 haystack 中搜索 needle
    for i, j := 0, 0; i < n; {
        if haystack[i] == needle[j] {
            i++
            j++
            if j == m {
                // 找到完全匹配，返回起始索引
                return i - j
            }
        } else if j > 0 {
            // 部分匹配失敗，使用 next 數組回退
            j = next[j-1]
        } else {
            // 第一個字就不匹配，直接往前移動
            i++
        }
    }
    
    // 如果沒有找到匹配，返回 -1
    return -1
}
```

## 7. 複雜度分析

### 暴力解法：

**時間複雜度**：
- 最壞情況下：O(n*m)，其中 n 是 haystack 的長度，m 是 needle 的長度。
  例如，當 haystack = "aaaaaaab"，needle = "aaaab" 時，我們需要多次比較才能發現不匹配。
- 最好情況下：O(n)，當 needle 很短或很快就能找到匹配時。

**空間複雜度**：
- O(1)，只使用了常數額外空間。

### KMP 算法：

**時間複雜度**：
- 構建 next 數組：O(m)，其中 m 是 needle 的長度。
- 在 haystack 中搜索：O(n)，其中 n 是 haystack 的長度。
- 總體時間複雜度：O(n+m)。

**空間複雜度**：
- O(m)，用於存儲 next 數組。

## 8. 優化與改進

### 暴力解法的優化：

我們可以通過早期退出來優化暴力解法：
- 當發現不匹配時立即停止當前位置的比較
- 使用 Go 的字串處理函數進行子串比較

### KMP 算法的優化：

KMP 算法已經是一個高效的算法，但在某些情況下，我們可以進一步優化 next 數組的構建：
- 當 needle[i] == needle[next[i-1]] 時，可以設置 next[i] = next[next[i-1]]，這樣可以避免一些不必要的比較。

### 其他算法：

除了 KMP 算法，還有其他一些字串匹配算法，如 Boyer-Moore 算法和 Rabin-Karp 算法，在某些情況下可能更有效。

### 學習建議：

- 深入理解 KMP 算法的原理和實現
- 學習其他字串匹配算法，如 Boyer-Moore 和 Rabin-Karp
- 練習更多的字串處理問題，尤其是與子串搜索相關的問題

## 9. 測試策略

這些測試用例涵蓋了：
- 基本功能測試（來自問題示例）
- 邊界情況（空字串、完全匹配等）
- 特殊情況（多次出現、無匹配等）
- 性能測試（長字串）
- 特殊模式（重複字）

通過這組測試，我們可以確保我們的解決方案在各種情況下都能正確工作。

## 總結

深入分析了「Find the Index of the First Occurrence in a String」問題，提供了兩種解決方案：暴力解法和 KMP 算法。

暴力解法簡單直接，容易理解和實現，適用於大多數情況，尤其是當字串不太長或者模式較短時。

KMP 算法更加高效，特別是對於長字串和複雜模式，它能夠避免不必要的比較，時間複雜度為 O(n+m)。