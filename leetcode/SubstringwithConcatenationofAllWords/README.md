# Substring with Concatenation of All Words

## 1. Original Problem

**Original Problem (English):**
You are given a string `s` and an array of strings `words` of the same length. Return all starting indices of substring(s) in `s` that is a concatenation of each word in `words` exactly once, in any order, and without any intervening characters.

You can return the answer in any order.

**Example 1:**
```
Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
```

**Example 2:**
```
Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
```

**Example 3:**
```
Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]
```

**Constraints:**
- 1 <= s.length <= 10^4
- 1 <= words.length <= 5000
- 1 <= words[i].length <= 30
- `s` and `words[i]` consist of lowercase English letters.

**繁體中文翻譯：**

給定一個字符串 `s` 和一個字符串數組 `words`，所有字符串的長度都相同。找出 `s` 中所有是由 `words` 中每個單詞恰好一次連接形成的子串的起始位置（子串中單詞的順序任意，且不能有任何間隔）。

你可以按任意順序返回答案。

**示例 1：**
```
輸入：s = "barfoothefoobarman", words = ["foo","bar"]
輸出：[0,9]
解釋：從索引 0 和 9 開始的子串分別是 "barfoo" 和 "foobar"。
輸出的順序不重要，返回 [9,0] 也是可以的。
```

**示例 2：**
```
輸入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
輸出：[]
```

**示例 3：**
```
輸入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
輸出：[6,9,12]
```

**約束條件：**
- 1 <= s.length <= 10^4
- 1 <= words.length <= 5000
- 1 <= words[i].length <= 30
- `s` 和 `words[i]` 由小寫英文字母組成。

## 2. 問題理解

這道題目要求我們在一個字符串 `s` 中找出所有滿足特定條件的子串的起始索引。這些特定條件是：

1. 子串必須是由數組 `words` 中的所有單詞恰好連接一次形成的
2. 子串中單詞的順序可以是任意的
3. 子串中不能有任何額外字符（即單詞之間沒有間隔）

關鍵特性：
- 所有 `words` 中的單詞長度相同，這是一個重要的特性，可以簡化我們的解決方案
- 每個單詞必須恰好使用一次
- 單詞可以以任何順序連接

邊界情況：
- 如果 `words` 數組為空，應該返回空列表
- 如果 `s` 字符串長度小於 `words` 中所有單詞的總長度，則無法形成有效的子串
- 如果 `words` 數組中有重複的單詞，我們需要確保每個單詞都使用恰好一次

挑戰點：
- 需要處理單詞順序的任意性
- 可能需要檢查字符串 `s` 的許多子串，有效率地完成這一點是關鍵
- 重複單詞的處理可能會增加複雜性

## 3. 視覺解釋

讓我們透過視覺化來理解這個問題：

以示例 1 為例：
```
s = "barfoothefoobarman"
words = ["foo", "bar"]
```

單詞長度均為 3，所以我們尋找長度為 6 的子串（3 * 2 = 6）。

以下是我們如何在 s 中滑動窗口尋找匹配：

```
索引: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17
字符: b a r f o o t h e f  o  o  b  a  r  m  a  n
      |_______|
      bar+foo (匹配)
        |_______|
        arf+oot (不匹配)
          |_______|
          rfo+oth (不匹配)
            ...
                |_______|
                foo+bar (匹配)
```

我們找到了兩個匹配：
1. 起始索引 0：單詞 "bar" 和 "foo" 連接形成 "barfoo"
2. 起始索引 9：單詞 "foo" 和 "bar" 連接形成 "foobar"

所以結果是 [0, 9]。

## 4. 思考過程

讓我們考慮幾種解決這個問題的策略：

### 策略 1：暴力枚舉
最直接的方法是檢查 `s` 中每個可能的起始位置，看它是否可以形成一個有效的連接子串。

步驟：
1. 計算目標子串的總長度：單詞長度 * 單詞數量
2. 對於 `s` 中每個可能的起始位置（直到 `s.length - 總長度 + 1`）：
    - 提取目標長度的子串
    - 判斷這個子串是否可以由 `words` 中所有單詞恰好一次組成

### 策略 2：滑動窗口 + 哈希表
由於我們需要檢查單詞的出現情況，哈希表是一個很好的數據結構選擇：

步驟：
1. 創建一個哈希表 `wordCount`，記錄 `words` 中每個單詞的出現次數
2. 計算單個單詞的長度 `wordLen` 和所有單詞的總長度 `totalLen`
3. 遍歷 `s` 的每個可能起始位置（最多到 `s.length - totalLen + 1`）：
    - 創建一個當前窗口的哈希表 `currCount`
    - 將當前子串按單詞長度分割成多個單詞
    - 檢查這些分割出的單詞是否在 `wordCount` 中，並且出現次數匹配
    - 如果全部匹配，則將當前起始位置添加到結果中

### 策略 3：優化的滑動窗口
我們可以進一步優化策略 2，避免重複計算：

步驟：
1. 由於所有單詞長度相同，可以按照單詞長度分組進行檢查
2. 對於每個偏移量（0 到 單詞長度-1），使用滑動窗口技術：
    - 維護一個窗口，檢查連續的單詞
    - 當窗口內單詞不符合要求時，調整窗口位置
    - 當窗口大小等於 `words` 長度時，找到一個有效索引

### 選擇最佳策略
策略 3 是最優的，因為它避免了重複計算，並利用了單詞長度相同這一特性。雖然策略 2 也是可行的，但策略 3 在實際中可能會更高效，尤其是當 `s` 和 `words` 較大時。

## 5. 最優解決方案開發

讓我們逐步開發策略 3 的解決方案：

1. 首先，處理邊界情況：
    - 如果 `s` 為空或 `words` 為空，返回空數組
    - 計算總長度，如果 `s` 長度小於總長度，返回空數組

2. 初始化參數：
    - 單詞長度 `wordLen`
    - 單詞總數 `wordCount`
    - 總長度 `totalLen = wordLen * wordCount`
    - 創建哈希表 `wordFreq`，存儲 `words` 中每個單詞的頻率

3. 對於每個可能的起始偏移量（從 0 到 `wordLen-1`）：
    - 初始化左指針 `left = i` 和右指針 `right = i`
    - 初始化當前窗口的單詞頻率哈希表 `currFreq`

    - 當 `right + wordLen <= s.length` 時：
        - 提取當前單詞 `currWord = s[right:right+wordLen]`
        - 如果 `currWord` 不在 `wordFreq` 中，重置窗口
        - 否則，增加 `currFreq[currWord]`

        - 當 `currFreq[currWord] > wordFreq[currWord]` 時：
            - 移動左指針，直到 `currFreq[currWord]` 不再超過 `wordFreq[currWord]`

        - 如果當前窗口大小等於 `wordCount`：
            - 將 `left` 添加到結果中
            - 移除最左側的單詞，移動左指針

        - 移動右指針 `right += wordLen`

4. 返回結果數組

讓我們通過示例來驗證這個方法：

以 `s = "barfoothefoobarman", words = ["foo","bar"]` 為例：

初始化：
- `wordLen = 3`
- `wordCount = 2`
- `totalLen = 6`
- `wordFreq = {"foo": 1, "bar": 1}`

對於偏移量 0：
1. `left = 0, right = 0`
2. 提取 "bar"，在 `wordFreq` 中，`currFreq = {"bar": 1}`
3. `right += 3 = 3`
4. 提取 "foo"，在 `wordFreq` 中，`currFreq = {"bar": 1, "foo": 1}`
5. 窗口大小 = 2，等於 `wordCount`，將 `left = 0` 添加到結果
6. 移除 "bar"，`left += 3 = 3`，`currFreq = {"foo": 1}`
   ...

重複這個過程，最終我們會得到結果 [0, 9]。

## 6. Golang 實現

```go
func findSubstring(s string, words []string) []int {
    result := []int{}
    
    // 處理邊界情況
    if len(s) == 0 || len(words) == 0 {
        return result
    }
    
    // 初始化參數
    wordLen := len(words[0])
    wordCount := len(words)
    totalLen := wordLen * wordCount
    
    // 如果 s 的長度小於總長度，則不可能找到匹配
    if len(s) < totalLen {
        return result
    }
    
    // 創建單詞頻率哈希表
    wordFreq := make(map[string]int)
    for _, word := range words {
        wordFreq[word]++
    }
    
    // 對於每個可能的起始偏移量
    for i := 0; i < wordLen; i++ {
        // 初始化滑動窗口
        left, right := i, i
        currFreq := make(map[string]int)
        count := 0  // 窗口中有效單詞的數量
        
        // 滑動窗口
        for right + wordLen <= len(s) {
            // 提取當前單詞
            currWord := s[right : right + wordLen]
            right += wordLen
            
            // 如果當前單詞不在 wordFreq 中，重置窗口
            if _, exists := wordFreq[currWord]; !exists {
                currFreq = make(map[string]int)
                count = 0
                left = right
                continue
            }
            
            // 增加當前單詞頻率
            currFreq[currWord]++
            count++
            
            // 如果當前單詞出現次數超過要求，縮小窗口
            for currFreq[currWord] > wordFreq[currWord] {
                leftWord := s[left : left + wordLen]
                currFreq[leftWord]--
                count--
                left += wordLen
            }
            
            // 如果窗口大小等於 wordCount，找到一個匹配
            if count == wordCount {
                result = append(result, left)
                
                // 移除最左側的單詞，繼續滑動
                leftWord := s[left : left + wordLen]
                currFreq[leftWord]--
                count--
                left += wordLen
            }
        }
    }
    
    return result
}
```

### 程式碼說明：

1. 首先，我們處理邊界情況並初始化必要的參數。
2. 創建 `wordFreq` 哈希表來記錄每個單詞的出現頻率。
3. 對於每個可能的起始偏移量（0 到 wordLen-1），我們使用滑動窗口技術：
    - 使用左右指針定義窗口
    - 當右指針移動時，提取新單詞並更新窗口狀態
    - 如果遇到不在 wordFreq 中的單詞，重置窗口
    - 如果某個單詞頻率超過要求，縮小窗口
    - 當窗口中有效單詞數等於 wordCount 時，找到一個匹配
4. 返回所有匹配的起始索引

## 7. 複雜度分析

### 時間複雜度

- 創建 `wordFreq` 哈希表：O(m)，其中 m 是 words 數組的長度
- 外層循環執行 wordLen 次（單詞長度）
- 內層循環在最壞情況下遍歷整個字符串：O(n)，其中 n 是字符串 s 的長度
- 在內層循環中，每次操作（提取單詞、更新哈希表）都是 O(1) 或 O(wordLen)

總體時間複雜度：O(n * wordLen)，這裡 n 是字符串長度。

最壞情況下（當所有單詞都相同時），時間複雜度可能接近 O(n * m)，其中 m 是 words 數組的長度。

### 空間複雜度

- `wordFreq` 哈希表：O(k)，其中 k 是不同單詞的數量，最壞情況下 k = m
- `currFreq` 哈希表：O(k)
- 結果數組：O(n) 在最壞情況下（每個位置都是一個匹配）

總體空間複雜度：O(k + n)，通常可以簡化為 O(m + n)。

## 8. 優化與改進

儘管我們的解決方案已經相當高效，但還可以考慮一些優化：

1. **提前終止**：如果剩餘的字符串長度不足以容納所有單詞，可以提前終止循環。

2. **哈希優化**：對於較短的單詞，可以考慮使用整數哈希代替字符串哈希，減少哈希計算的開銷。

3. **並行處理**：由於每個偏移量的處理是獨立的，可以考慮並行處理不同的偏移量，尤其是在多核環境下。

4. **預處理優化**：對於較長的字符串，可以預先計算每個位置的單詞，避免重複提取。


## 9. 測試策略

這個測試涵蓋了各種情況：
- 基本功能測試
- 邊界情況（空字符串、空單詞數組）
- 重複單詞的處理
- 長度剛好和不足的情況
- 極端情況（大量重複單詞）

通過這些測試用例，我們可以確保我們的解決方案能夠正確處理各種輸入，並且在不同條件下都能得到期望的結果。

每個測試用例都包含以下信息：
- 測試名稱
- 輸入字符串 s
- 輸入單詞數組 words
- 期望的輸出結果
- 測試描述（說明測試的目的）