# Group Anagrams - LeetCode Solution

## 1. Original Problem

### English:
Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

An **anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example 1:**
```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

**Example 2:**
```
Input: strs = [""]
Output: [[""]]
```

**Example 3:**
```
Input: strs = ["a"]
Output: [["a"]]
```

**Constraints:**
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters.

### 繁體中文:
給定一個字串數組 `strs`，將所有的字母異位詞（anagram）組合在一起。你可以按任意順序返回結果。

**字母異位詞**是由重新排列另一個詞或短語的字母形成的詞或短語，通常使用所有原始字母恰好一次。

**示例 1:**
```
輸入: strs = ["eat","tea","tan","ate","nat","bat"]
輸出: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

**示例 2:**
```
輸入: strs = [""]
輸出: [[""]]
```

**示例 3:**
```
輸入: strs = ["a"]
輸出: [["a"]]
```

**約束條件:**
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] 僅包含小寫英文字母。

## 2. 初步問題解析與心智模型構建

### 問題初步解析
當我們第一次看到這個問題時，需要仔細閱讀並提取關鍵信息：
- 我們有一個字串數組
- 需要找出所有互為字母異位詞的字串
- 將它們分組在一起

### 解釋範例
以範例1為例："eat", "tea", "tan", "ate", "nat", "bat"
- "eat", "tea", "ate" 互為字母異位詞，因為它們都由字母 'a', 'e', 't' 組成
- "tan" 和 "nat" 互為字母異位詞，由字母 'a', 'n', 't' 組成
- "bat" 是獨立的，由字母 'a', 'b', 't' 組成

### 輸入和輸出分析
- 輸入：一個字串數組
- 輸出：一個二維字串數組，每個子數組包含互為字母異位詞的字串

### 心智模型
我們可以想像一個分類過程：
- 字母異位詞的本質是：雖然字母順序不同，但字母組成完全相同
- 我們需要一種方式來識別「相同字母組成」的字串
- 然後建立映射關係：字母組成 → 對應的所有字串

## 3. 問題理解與核心挑戰

### 問題本質
字母異位詞問題的核心是：如何有效地判斷兩個字串是否互為字母異位詞？

### 核心挑戰
1. **如何表示字母組成**：我們需要一種方法來唯一表示一組字母（不考慮順序）
2. **如何高效分組**：處理大量字串時，需要高效的分組方法
3. **如何處理邊界情況**：例如空字串或單字字串

### 關鍵洞察
- 字母異位詞具有相同的字母組成，意味著：
    - 它們排序後完全相同
    - 它們的字母頻率計數完全相同
- 這提供了兩種可能的表示方法

## 4. 視覺化問題表示

讓我們視覺化展示如何識別和分組字母異位詞：

假設輸入：["eat", "tea", "tan", "ate", "nat", "bat"]

方法1：排序後分組
```
原始字串 -> 排序後 -> 分組
"eat"     -> "aet"   -> Group 1
"tea"     -> "aet"   -> Group 1
"tan"     -> "ant"   -> Group 2
"ate"     -> "aet"   -> Group 1
"nat"     -> "ant"   -> Group 2
"bat"     -> "abt"   -> Group 3
```

方法2：字母頻率計數
```
原始字串 -> 字母頻率表示     -> 分組
"eat"     -> "a1e1t1"       -> Group 1
"tea"     -> "a1e1t1"       -> Group 1
"tan"     -> "a1n1t1"       -> Group 2
"ate"     -> "a1e1t1"       -> Group 1
"nat"     -> "a1n1t1"       -> Group 2
"bat"     -> "a1b1t1"       -> Group 3
```

最終結果：
- Group 1: ["eat", "tea", "ate"]
- Group 2: ["tan", "nat"]
- Group 3: ["bat"]

## 5. 問題模式識別

### 模式特徵
這是一個**分類/分組問題**，具有以下特徵：
- 需要根據某種等價關係將元素分組
- 等價關係在這裡是「具有相同字母組成」
- 需要一個映射結構來維護分組關係

### 適用的常見模式
- **哈希表映射**：使用哈希表將相似元素映射到同一組
- **排序+分組**：通過排序創建標準形式，然後分組

### 從約束條件獲得的提示
- 字串長度最多為100，排序單個字串不會成為性能瓶頸
- 字串數量可達10^4，需要高效的分組方法

## 6. 策略性問題解決框架

對於字母異位詞分組問題，可以採用以下策略框架：

1. **選擇合適的字母異位詞表示方法**：
    - 方法A：將字串排序，相同排序結果的是字母異位詞
    - 方法B：計算每個字出現的頻率，創建頻率表示

2. **使用哈希表進行分組**：
    - 鍵：字母異位詞的唯一表示（排序結果或頻率表示）
    - 值：具有相同表示的原始字串列表

3. **構建結果**：
    - 將哈希表中的所有值（字串列表）收集起來

### 思考啟發式
- 當需要分組相似項時，思考「什麼使它們相似」
- 尋找能夠唯一標識這種相似性的表示方法
- 使用哈希表將相同表示映射到同一組

## 7. 算法設計（編碼前）

### 排序法偽代碼
```
function groupAnagrams(strs):
    創建一個映射 map: 排序後的字串 -> 原始字串列表
    
    for 每個字串 str in strs:
        將 str 按字母順序排序得到 sortedStr
        將 str 添加到 map[sortedStr] 列表中
    
    創建結果數組 result
    for 每個 列表 in map.values():
        將該列表添加到 result 中
    
    return result
```

### 字頻率法偽代碼
```
function groupAnagrams(strs):
    創建一個映射 map: 字頻率表示 -> 原始字串列表
    
    for 每個字串 str in strs:
        計算 str 中每個字的頻率，生成頻率表示 freq
        將 str 添加到 map[freq] 列表中
    
    創建結果數組 result
    for 每個 列表 in map.values():
        將該列表添加到 result 中
    
    return result
```

我們將採用排序法，因為它實現簡單且對於這個問題足夠高效。

## 8. 視覺化解釋

讓我們以示例 `["eat", "tea", "tan", "ate", "nat", "bat"]` 為例，視覺化排序法的執行過程：

```
步驟1: 創建空映射
map = {}

步驟2: 遍歷每個字串
- "eat" 排序後 -> "aet"
  map = {"aet": ["eat"]}
  
- "tea" 排序後 -> "aet"
  map = {"aet": ["eat", "tea"]}
  
- "tan" 排序後 -> "ant"
  map = {"aet": ["eat", "tea"], "ant": ["tan"]}
  
- "ate" 排序後 -> "aet"
  map = {"aet": ["eat", "tea", "ate"], "ant": ["tan"]}
  
- "nat" 排序後 -> "ant"
  map = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"]}
  
- "bat" 排序後 -> "abt"
  map = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"], "abt": ["bat"]}

步驟3: 收集所有值列表
result = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

這個映射關係可以用下圖表示：

```
                   ┌─────────┐
                   │   Map   │
                   └─────────┘
                         │
          ┌──────────────┼───────────────┐
          ▼              ▼               ▼
    ┌──────────┐   ┌──────────┐     ┌────────┐
    │Key: "aet"│   │Key: "ant"│     │Key: "abt"│
    └──────────┘   └──────────┘     └────────┘
          │              │               │
          ▼              ▼               ▼
┌──────────────────┐ ┌────────────┐ ┌────────┐
│["eat","tea","ate"]│ │["tan","nat"]│ │["bat"]│
└──────────────────┘ └────────────┘ └────────┘
```

## 9. 解決方案發展歷程

### 初始直覺：暴力法
最直接的方法是比較每對字串，檢查它們是否互為字母異位詞：
- 對於每對字串，排序並比較
- 時間複雜度：O(n²k log k)，其中n是字串數量，k是平均字串長度
- 這種方法效率很低，不適合大規模輸入

### 關鍵洞察：唯一表示
關鍵優化洞察：我們不需要兩兩比較，而是可以為每個字串創建一個唯一的「指紋」，使得字母異位詞具有相同的指紋。

### 改進方案1：排序法
- 將每個字串排序作為指紋
- 使用哈希表根據這個指紋進行分組
- 時間複雜度：O(n k log k)，其中n是字串數量，k是平均字串長度

### 改進方案2：字頻率法
- 計算每個字串中26個小寫字母的出現頻率
- 創建一個包含這些頻率的字串作為指紋
- 使用哈希表根據這個指紋進行分組
- 時間複雜度：O(n k)，其中n是字串數量，k是平均字串長度

### 最終選擇
雖然字頻率法在理論上更快（特別是對於長字串），但排序法實現更簡單，且在實踐中對於這個問題的約束條件足夠高效。因此，我們選擇排序法作為最終解決方案。

## 10. 實例詳細解析

讓我以一個更複雜的例子來演示算法執行過程：
輸入：`["bdddddddddd", "bbbbbbbbbbc", "ddddddddddz", "abbbc", "abbc", "abc"]`

1. 創建空映射 `map = {}`

2. 處理 "bdddddddddd"
    - 排序後 -> "bdddddddddd"
    - 更新映射：`map = {"bdddddddddd": ["bdddddddddd"]}`

3. 處理 "bbbbbbbbbbc"
    - 排序後 -> "bbbbbbbbbbcc"
    - 更新映射：`map = {"bdddddddddd": ["bdddddddddd"], "bbbbbbbbbbcc": ["bbbbbbbbbbc"]}`

4. 處理 "ddddddddddz"
    - 排序後 -> "ddddddddddz"
    - 更新映射：`map = {"bdddddddddd": ["bdddddddddd"], "bbbbbbbbbbcc": ["bbbbbbbbbbc"], "ddddddddddz": ["ddddddddddz"]}`

5. 處理 "abbbc"
    - 排序後 -> "abbbcc"
    - 更新映射：`map = {"bdddddddddd": ["bdddddddddd"], "bbbbbbbbbbcc": ["bbbbbbbbbbc"], "ddddddddddz": ["ddddddddddz"], "abbbcc": ["abbbc"]}`

6. 處理 "abbc"
    - 排序後 -> "abbcc"
    - 更新映射：`map = {"bdddddddddd": ["bdddddddddd"], "bbbbbbbbbbcc": ["bbbbbbbbbbc"], "ddddddddddz": ["ddddddddddz"], "abbbcc": ["abbbc"], "abbcc": ["abbc"]}`

7. 處理 "abc"
    - 排序後 -> "abcc"
    - 更新映射：`map = {"bdddddddddd": ["bdddddddddd"], "bbbbbbbbbbcc": ["bbbbbbbbbbc"], "ddddddddddz": ["ddddddddddz"], "abbbcc": ["abbbc"], "abbcc": ["abbc"], "abcc": ["abc"]}`

8. 收集所有值列表作為結果：
   `result = [["bdddddddddd"], ["bbbbbbbbbbc"], ["ddddddddddz"], ["abbbc"], ["abbc"], ["abc"]]`

在這個例子中，沒有兩個字串互為字母異位詞，所以每個字串都獨自形成一個組。

## 11. Golang 實現

```go
package main

import (
	"sort"
	"strings"
)

// 將字串排序（用於創建字母異位詞的唯一標識）
func sortString(s string) string {
	// 將字串轉換為字切片
	chars := strings.Split(s, "")
	// 對字進行排序
	sort.Strings(chars)
	// 將排序後的字重新連接為字串
	return strings.Join(chars, "")
}

// 主函數：對字母異位詞進行分組
func groupAnagrams(strs []string) [][]string {
	// 創建映射：排序後的字串 -> 原始字串列表
	anagramMap := make(map[string][]string)
	
	// 遍歷每個輸入字串
	for _, str := range strs {
		// 獲取排序後的字串作為鍵
		sortedStr := sortString(str)
		
		// 將原始字串添加到對應的列表中
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	
	// 創建結果數組
	result := make([][]string, 0, len(anagramMap))
	
	// 收集所有的字串組
	for _, group := range anagramMap {
		result = append(result, group)
	}
	
	return result
}
```

## 12. 實現執行步驟解析

讓我們使用先前的例子 `["eat", "tea", "tan", "ate", "nat", "bat"]` 追踪代碼的執行過程：

1. 創建空映射 `anagramMap = {}`

2. 處理第一個字串 "eat"
    - 排序後 -> sortString("eat") = "aet"
    - 更新映射：`anagramMap = {"aet": ["eat"]}`

3. 處理 "tea"
    - 排序後 -> sortString("tea") = "aet"
    - 更新映射：`anagramMap = {"aet": ["eat", "tea"]}`

4. 處理 "tan"
    - 排序後 -> sortString("tan") = "ant"
    - 更新映射：`anagramMap = {"aet": ["eat", "tea"], "ant": ["tan"]}`

5. 處理 "ate"
    - 排序後 -> sortString("ate") = "aet"
    - 更新映射：`anagramMap = {"aet": ["eat", "tea", "ate"], "ant": ["tan"]}`

6. 處理 "nat"
    - 排序後 -> sortString("nat") = "ant"
    - 更新映射：`anagramMap = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"]}`

7. 處理 "bat"
    - 排序後 -> sortString("bat") = "abt"
    - 更新映射：`anagramMap = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"], "abt": ["bat"]}`

8. 創建結果數組 `result = []`

9. 遍歷 anagramMap 的值
    - 添加 ["eat", "tea", "ate"] 到 result
    - 添加 ["tan", "nat"] 到 result
    - 添加 ["bat"] 到 result

10. 最終結果：`result = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]`

注意：Go 的 map 遍歷順序是隨機的，所以實際輸出的組的順序可能與上述不同，但這不影響正確性，因為題目允許以任意順序返回結果。

## 13. 複雜度分析

### 時間複雜度

- **排序字串**：對於每個長度為 k 的字串，排序需要 O(k log k) 時間
- **遍歷所有字串**：有 n 個字串，所以這步需要 O(n) 時間
- **哈希表操作**：在平均情況下，插入和查找操作為 O(1)
- **收集結果**：遍歷映射中的所有值列表，總體為 O(n)

總時間複雜度：O(n \* k log k)，其中 n 是字串數量，k 是最長字串的長度。

### 空間複雜度

- **哈希表**：存儲所有字串的分組，最壞情況下需要 O(n \* k) 空間
- **結果數組**：存儲相同的內容，需要 O(n \* k) 空間
- **排序過程**：需要 O(k) 臨時空間

總空間複雜度：O(n \* k)

### 複雜度分析解釋

時間複雜度主要由字串排序決定。對於每個字串，我們需要 O(k log k) 的時間來排序，其中 k 是字串長度。由於我們需要處理 n 個字串，總時間複雜度為 O(n \* k log k)。

空間複雜度主要來自於存儲哈希表和結果數組。在最壞情況下（每個字串都不是其他字串的字母異位詞），我們需要存儲 n 個鍵值對，每個鍵值對包含長度為 k 的排序後字串和單個原始字串，總共需要 O(n \* k) 的空間。

## 14. 優化與改進

### 可能的優化方向

1. **計數字頻率**
    - 使用大小為 26 的數組記錄每個字母的出現次數
    - 將數組轉換為字串作為哈希表鍵
    - 時間複雜度改進為 O(n \* k)，對於長字串效果更好
   ```go
   func countFrequency(s string) string {
       count := make([]int, 26)
       for _, ch := range s {
           count[ch-'a']++
       }
       // 將計數轉換為字串
       var sb strings.Builder
       for i, c := range count {
           if c > 0 {
               sb.WriteByte(byte(i + 'a'))
               sb.WriteString(strconv.Itoa(c))
           }
       }
       return sb.String()
   }
   ```

2. **自定義哈希函數**
    - 為字頻率計算自定義哈希值，避免創建中間字串
    - 這可以進一步減少空間使用，但實現複雜度增加

### 不同解決方案的比較

| 方法 | 時間複雜度 | 空間複雜度 | 實現複雜度 | 適用場景 |
|------|------------|------------|------------|----------|
| 排序法 | O(n \* k log k) | O(n \* k) | 簡單 | 短字串，易實現 |
| 計數法 | O(n \* k) | O(n \* k) | 中等 | 長字串，對性能敏感 |
| 素數乘積法* | O(n \* k) | O(n) | 複雜 | 理論上更優但有數值溢出風險 |

*素數乘積法：將每個字母映射到唯一素數，字串的「指紋」是字母對應素數的乘積。這種方法理論上很好，但容易導致數值溢出。

### 推薦

對於大多數情況，排序方法提供了良好的平衡：它易於實現、效率合理，且不受字串內容的影響。如果發現性能瓶頸且主要處理長字串，可以考慮切換到計數方法。

## 15. 通用問題解決智慧

### 核心教訓

1. **尋找等價關係**：此問題的關鍵在於識別出字母異位詞之間的等價關係。在許多分類問題中，找出正確的等價關係是解決的第一步。

2. **唯一表示的重要性**：為相似對象創建唯一標識符是一種強大的技術，可以將複雜的分組問題轉化為簡單的哈希表查找。

3. **哈希表的強大應用**：哈希表是解決分組問題的理想數據結構，它使我們能夠以接近常數時間複雜度進行查找和插入操作。

### 相關問題

這個問題的解決方案可以應用於許多其他需要找出相似性或等價關係的問題：
- 字串匹配問題
- 模式識別問題
- 集合分類問題

### 算法思維培養

1. **問題轉化**：將「判斷兩個字串是否互為字母異位詞」轉化為「它們是否有相同的唯一表示」
2. **預處理數據**：通過排序或計數預處理數據，簡化後續操作
3. **分而治之**：將大問題（分組所有字串）分解為小問題（處理每個字串並放入正確的組）

## 16. 測試策略

### 使用表驅動測試

這個測試套件涵蓋了多種情況：
- 基本案例（示例1-3）
- 邊界情況（空輸入、單字輸入）
- 極端情況（無字母異位詞、全部相同、長字串）