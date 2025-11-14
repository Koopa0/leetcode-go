# LeetCode 125: 驗證回文串（Valid Palindrome）

## 問題定義

### 原始問題（英文）
```
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
```

### 問題翻譯（繁體中文）
```
如果一個短語在將所有大寫字母轉換為小寫字母並移除所有非字母數字字符後，正著讀和反著讀都一樣，則它是回文串。字母數字字符包括字母和數字。

給定一個字串 s，如果它是回文串則返回 true，否則返回 false。
```

### 範例

**範例 1：**
```
輸入：s = "A man, a plan, a canal: Panama"
輸出：true
解釋："amanaplanacanalpanama" 是回文串。
```

**範例 2：**
```
輸入：s = "race a car"
輸出：false
解釋："raceacar" 不是回文串。
```

**範例 3：**
```
輸入：s = " "
輸出：true
解釋：移除非字母數字字符後得到空字串 ""，這被認為是回文串。
```

**限制條件：**
- 1 <= s.length <= 2 * 10^5
- s 僅由可打印的 ASCII 字符組成

## 解法說明

### 方法：雙指針

**核心思路：**
使用雙指針從字串兩端向中間移動，跳過非字母數字字符，比較對應位置的字符（忽略大小寫）。

**算法步驟：**
1. 初始化左指針 `left = 0`，右指針 `right = len(s) - 1`
2. 當 `left < right` 時：
   - 跳過左側的非字母數字字符
   - 跳過右側的非字母數字字符
   - 比較兩個字符（轉為小寫），如果不相等則返回 false
   - 移動兩個指針
3. 如果所有字符都匹配，返回 true

**複雜度分析：**
- 時間複雜度：O(n)，其中 n 是字串長度，最多遍歷一次字串
- 空間複雜度：O(1)，只使用常數額外空間

**關鍵洞察：**
雙指針技巧可以有效地驗證回文串，同時跳過非字母數字字符。通過自行實現字符判斷和大小寫轉換函數，避免了使用標準庫，提高了性能。
