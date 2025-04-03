## 1. Original Problem

### English:
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

```
Symbol       Value
I            1
V            5
X            10
L            50
C            100
D            500
M            1000
```

For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX.

There are six instances where subtraction is used:
- I can be placed before V (5) and X (10) to make 4 and 9.
- X can be placed before L (50) and C (100) to make 40 and 90.
- C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

**Example 1:**
```
Input: s = "III"
Output: 3
```

**Example 2:**
```
Input: s = "IV"
Output: 4
```

**Example 3:**
```
Input: s = "IX"
Output: 9
```

**Example 4:**
```
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
```

**Example 5:**
```
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

**Constraints:**
- 1 <= s.length <= 15
- s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
- It is guaranteed that s is a valid roman numeral in the range [1, 3999].

### 繁體中文：
羅馬數字由七個不同的符號表示：I、V、X、L、C、D 和 M。

```
符號       值
I          1
V          5
X          10
L          50
C          100
D          500
M          1000
```

例如，2 在羅馬數字中寫作 II，即兩個一相加。12 寫作 XII，即 X + II。27 寫作 XXVII，即 XX + V + II。

羅馬數字通常從左到右按從大到小的順序書寫。然而，數字 4 不是 IIII。相反，數字 4 寫作 IV。因為 1 在 5 之前，所以我們減去它得到 4。同樣的原理適用於數字 9，它寫作 IX。

有六種情況使用減法：
- I 可以放在 V(5) 和 X(10) 之前，表示 4 和 9。
- X 可以放在 L(50) 和 C(100) 之前，表示 40 和 90。
- C 可以放在 D(500) 和 M(1000) 之前，表示 400 和 900。

給定一個羅馬數字，將其轉換為整數。

## 2. 問題理解

### 核心要求與限制：
- 需要將一個羅馬數字字串轉換為相應的整數值
- 羅馬數字由 I、V、X、L、C、D、M 這七個符號組成
- 羅馬數字通常按從大到小的順序從左到右排列
- 特殊情況是六種減法規則：IV(4)、IX(9)、XL(40)、XC(90)、CD(400)、CM(900)
- 輸入保證是有效的羅馬數字，範圍在 1 到 3999 之間

### 輸入/輸出特性與邊界情況：
- 輸入：一個由羅馬數字符號組成的字串，長度在 1 到 15 之間
- 輸出：對應的整數值
- 邊界情況包括：
    - 最小值如 "I"（1）
    - 最大值如 "MMMCMXCIX"（3999）
    - 包含所有減法規則的情況，如 "MCMXCIV"（1994）

### 潛在難點與關鍵挑戰：
1. 處理羅馬數字中的減法規則
2. 區分何時應該加上符號值，何時應該使用減法規則
3. 理解羅馬數字從左到右的讀取邏輯

## 3. 視覺化解釋

### 羅馬數字對照表：

### 範例解析流程：
以 "MCMXCIV" = 1994 為例，解析流程如下：

## 4. 思考過程

### 解題策略：

#### 方案 1：左到右掃描（基本方法）
對字串從左到右掃描，檢查每個字符：
- 如果當前字符表示的值大於等於下一個字符的值，則加上當前字符的值
- 如果當前字符的值小於下一個字符的值，則減去當前字符的值（處理減法規則）

優點：
- 直觀，容易理解和實現
- 模擬人類讀羅馬數字的方式

缺點：
- 需要處理邊界條件（如最後一個字符）

#### 方案 2：右到左掃描
從右到左掃描羅馬數字：
- 記錄上一個處理的字符的值
- 如果當前字符的值小於上一個處理的字符的值，則減去當前字符的值
- 否則加上當前字符的值

優點：
- 不需要檢查下一個字符，減少邊界處理
- 實作簡潔

缺點：
- 不是羅馬數字的自然讀法

#### 方案 3：預處理減法情況
預先處理所有可能的二字符減法組合（IV, IX, XL, XC, CD, CM）：
- 先檢查當前位置是否以這些二字符減法模式開始
- 如果是，加上相應的值，並跳過這兩個字符
- 否則，只加上當前字符的值

優點：
- 明確處理所有減法情況
- 程式碼結構清晰

缺點：
- 需要維護額外的減法規則映射

### 策略選擇：
考慮到羅馬數字轉換的特性，方案 1 的左到右掃描是最直觀的方法，但方案 2 的右到左掃描可以簡化邊界條件處理。方案 3 則更明確地處理了減法規則。

對於這個問題，建議使用方案 1 或方案 2，因為它們都能清晰地反映羅馬數字的轉換邏輯，且實作簡單。方案 3 雖然明確，但程式碼較複雜。

我們將選擇方案 2（右到左掃描）作為最終解法，因為它能有效處理減法規則且不需要處理字串邊界問題。

### 辨識類似問題的模式：
此類問題的特徵是：
1. 需要解析特定規則的符號
2. 符號之間有位置相關的轉換規則
3. 需要累積計算最終結果

類似問題包括：
- 數字系統轉換
- 字串解析問題
- 符號解釋問題

## 5. 最佳解決方案開發

### 解決方案演進：

#### 從基本解法開始：
最直觀的方法是從左到右掃描，依次累加每個羅馬數字符號的值。

```
例如 "III" = 1 + 1 + 1 = 3
```

但這不處理減法規則。

#### 改進：處理減法規則
改進方法是比較當前符號與下一個符號：
- 如果當前符號值 < 下一個符號值，執行減法
- 否則執行加法

```
例如 "IV" = -1 + 5 = 4
```

#### 最佳解法：從右到左掃描
從右到左掃描並與最後一個處理的值比較：
1. 初始化結果 result = 0，最後一個值 lastValue = 0
2. 從右到左遍歷每個字符：
    - 獲取當前字符的值 curValue
    - 如果 curValue ≥ lastValue，加上 curValue
    - 否則，減去 curValue
    - 更新 lastValue = curValue
3. 返回 result

這種方法的關鍵洞察是：在羅馬數字中，從右到左看，如果一個符號小於它右側的符號，那麼這個符號應該被減去，否則應該被加上。

##### 具體例子： "MCMXCIV" = 1994
1. 初始化 result = 0, lastValue = 0
2. 從右到左遍歷：
    - 'V': curValue = 5, 5 ≥ 0, result = 0 + 5 = 5, lastValue = 5
    - 'I': curValue = 1, 1 < 5, result = 5 - 1 = 4, lastValue = 1
    - 'C': curValue = 100, 100 ≥ 1, result = 4 + 100 = 104, lastValue = 100
    - 'X': curValue = 10, 10 < 100, result = 104 - 10 = 94, lastValue = 10
    - 'M': curValue = 1000, 1000 ≥ 10, result = 94 + 1000 = 1094, lastValue = 1000
    - 'C': curValue = 100, 100 < 1000, result = 1094 - 100 = 994, lastValue = 100
    - 'M': curValue = 1000, 1000 ≥ 100, result = 994 + 1000 = 1994, lastValue = 1000
3. 返回 result = 1994

## 7. 複雜度分析

### 時間複雜度：
- 我們只需遍歷一次羅馬數字字串，所以時間複雜度為 O(n)，其中 n 是字串的長度。
- map 的查詢操作在 Go 中的平均時間複雜度為 O(1)。
- 由於羅馬數字字串的長度受限（題目說明長度不超過 15，且值不超過 3999），所以實際上時間複雜度可視為常數 O(1)。

### 空間複雜度：
- 我們使用了一個固定大小的 map 來存儲羅馬數字符號到整數的映射，這需要 O(1) 的額外空間。
- 除此之外，我們只使用了幾個變數，所以總空間複雜度為 O(1)。

### 無法進一步優化的理由：
- 時間複雜度已經是 O(n)，這是必須的，因為我們至少需要讀取每個字符一次。
- 空間複雜度為 O(1)，已經是最優的。
- 雖然可以不使用 map 而改用 switch 語句，但這不會改變複雜度分析，只是實現細節的不同。

### 解決方案比較：

| 方案              | 優點                                   | 缺點                                 |
|-------------------|----------------------------------------|--------------------------------------|
| 右到左掃描        | 簡潔、邊界處理簡單、容易理解           | 不是羅馬數字的自然讀法               |
| 左到右掃描        | 直觀、符合羅馬數字的自然讀法           | 需要處理邊界條件                     |
| 預處理減法規則    | 明確處理所有減法情況、結構清晰         | 程式碼較複雜、需要額外的映射表         |

考慮到效率和程式碼簡潔性，右到左掃描的方法是最佳選擇。

## 9. 測試策略

這個測試策略包含：
1. **一般情況測試**：包括基本加法、簡單減法、混合情況
2. **特殊規則測試**：測試所有減法規則（IV、IX、XL、XC、CD、CM）
3. **邊界情況測試**：包括最小值 "I" 和最大值 "MMMCMXCIX"
4. **複雜情況測試**：測試多種規則組合的情況，如 "MCMXCIV"

這種表格驅動的測試方法符合 Go 語言的最佳實踐，易於擴展和維護，並且每個測試都有明確的描述，便於理解測試的目的。