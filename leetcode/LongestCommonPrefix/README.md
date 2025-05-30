## 1. Original Problem:

**English:**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

**Example 1:**
```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

**Example 2:**
```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

**Constraints:**
- 1 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] consists of only lowercase English letters.

**繁體中文:**
編寫一個函數來查找字串數組中的最長公共前綴。

如果不存在公共前綴，返回空字串 ""。

**示例 1:**
```
輸入: strs = ["flower","flow","flight"]
輸出: "fl"
```

**示例 2:**
```
輸入: strs = ["dog","racecar","car"]
輸出: ""
解釋: 輸入不存在公共前綴。
```

**約束:**
- 1 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] 僅由小寫英文字母組成。

## 2. 問題理解:

這個問題要求我們找出一個字串數組中所有字串共有的最長前綴。

**核心要求:**
- 我們需要找出所有字串共同擁有的前綴部分
- 如果存在公共前綴，我們必須返回最長的那一個
- 如果不存在公共前綴，我們必須返回空字串

**輸入/輸出特徵:**
- 輸入是一個字串數組
- 輸出是一個字串（公共前綴或空字串）

**邊緣情況:**
- 如果數組中只有一個字串，那麼最長公共前綴就是該字串本身
- 如果數組中存在空字串，則最長公共前綴必定為空字串
- 如果數組中字串的首字母就不相同，則最長公共前綴為空字串

**挑戰:**
- 需要有效地比較所有字串的相應位置的字
- 需要處理長度不一的字串情況

## 3. 視覺解釋:

讓我們用例子來說明算法流程:

對於輸入 `["flower", "flow", "flight"]`：

```
字串:   f l o w e r
          | | | |
          f l o w
          | |
          f l i g h t
結果:     f l (最長公共前綴: "fl")
```

對於輸入 `["dog", "racecar", "car"]`：

```
字串:   d o g
          |
          r a c e c a r
          |
          c a r
結果:     (無公共前綴: "")
```

## 4. 思考過程:

為解決這個問題，我們可以考慮以下幾種方法:

**方法 1: 水平掃描**
- 取第一個字串作為基準前綴
- 將其與其他字串逐一比較，每次比較后縮短前綴
- 直到找到所有字串的公共前綴或前綴變為空

**方法 2: 垂直掃描**
- 按列（字位置）而不是按行（字串）比較
- 逐一檢查每個字串的相同位置的字是否相同
- 遇到不同字或某字串結束時停止

**方法 3: 分治法**
- 將問題分解為查找兩個字串的最長公共前綴
- 然後合併結果

方法 1 和方法 2 都是直接的解決方案，時間複雜度相似。方法 1 對於短字串數組通常更加高效，而方法 2 對於有長公共前綴的字串數組會更有效率。考慮到問題特性和 Go 的字串處理，我會選擇方法 1 作為主要解決方案。

## 5. 最優解決方案:

我們採用水平掃描方法，步驟如下:

1. 如果字串數組為空或包含空字串，返回空字串
2. 選擇第一個字串作為初始前綴
3. 遍歷剩余字串，檢查當前前綴是否是每個字串的前綴
4. 如果不是，則從當前前綴中移除最後一個字，重複此步驟
5. 直到找到公共前綴或前綴變為空

以 `["flower", "flow", "flight"]` 為例:
- 初始前綴 = "flower"
- 檢查 "flow"：前綴不匹配，縮短為 "flowe"
- 循環直至前綴變為 "flow"
- 檢查 "flight"：前綴不匹配，縮短為 "flo"
- 循環直至前綴變為 "fl"
- "fl" 是所有三個字串的公共前綴，返回 "fl"

## 7. 複雜度分析:

**時間複雜度:**
- 最壞情況：O(S)，其中 S 是所有字串中的字總數
- 在最壞情況下，我們需要比較每個字串的每個字
- 例如，如果 n 個字串的平均長度為 m，則時間複雜度為 O(n*m)

**空間複雜度:**
- O(1)，我們只使用了常數額外空間（prefix 變量）
- 我們沒有使用隨輸入大小增長的額外數據結構

**推導過程:**
- 第一個循環遍歷字串數組，最多執行 n-1 次（其中 n 是字串數量）
- 嵌套循環最多執行 m 次（其中 m 是第一個字串的長度）
- 因此總操作數上限為 O(n*m)，等同於所有字總數 S

## 8. 優化與改進:

**當前解決方案的優點:**
- 程式碼簡潔易懂
- 在大多數情況下效率較高，特別是當公共前綴短或不存在時
- 無需額外空間

**可能的改進:**
1. **提前終止:**
    - 我們可以先找出最短字串的長度，以減少不必要的比較
    - 這對於包含極短字串的數組會有幫助

2. **二分查找優化:**
    - 如果字串很長，我們可以使用二分查找來加速前綴搜索
    - 而不是每次只縮短一個字

3. **字典樹 (Trie) 方法:**
    - 對於大量字串或需要重複查詢的情況，可以考慮使用字典樹
    - 但這會增加空間複雜度和實現複雜度

**相關問題:**
- 最長公共子串（不同於前綴，可以在任意位置）
- 字串匹配和前綴樹實現

## 9. 測試策略:

我們的測試方案使用了 Go 的表格驅動測試方法，包含了以下測試案例：

1. **基本測試：** 測試正常公共前綴情況
2. **無公共前綴：** 測試無公共前綴情況
3. **單一字串：** 測試數組只有一個字串的情況
4. **空數組：** 測試空數組的情況
5. **包含空字串：** 測試數組包含空字串的情況
6. **相同字串：** 測試所有字串都相同的情況
7. **長字串：** 測試長字串的情況
8. **不同長度：** 測試不同長度字串的情況
9. **Unicode 字：** 測試包含 Unicode 字的情況

這些測試案例涵蓋了正常使用情況、邊緣情況和特殊情況，確保了我們的解決方案的正確性和穩健性。

總結來說，這個「最長公共前綴」問題是一個經典的字串處理問題，我們透過水平掃描的方法高效地解決了它。這種方法簡單直觀，且在大多數情況下性能良好。理解這種問題和解決方案對於掌握字串處理和算法技巧非常有幫助。