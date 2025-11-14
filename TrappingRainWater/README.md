# Trapping Rain Water - LeetCode 解題指南

## 1. 原始問題

### Original Problem:
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

**Example 1:**
```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
```

**Example 2:**
```
Input: height = [4,2,0,3,2,5]
Output: 9
```

**Constraints:**
- n == height.length
- 1 <= n <= 2 * 10^4
- 0 <= height[i] <= 10^5

### 繁體中文翻譯:
給定 n 個非負整數表示高度圖，其中每個條形的寬度為 1，計算下雨後能夠捕獲多少雨水。

**示例 1:**
```
輸入: height = [0,1,0,2,1,0,1,3,2,1,2,1]
輸出: 6
解釋: 上面的高度圖（黑色部分）由數組 [0,1,0,2,1,0,1,3,2,1,2,1] 表示。在這種情況下，可以接 6 個單位的雨水（藍色部分）。
```

**示例 2:**
```
輸入: height = [4,2,0,3,2,5]
輸出: 9
```

**約束條件:**
- n == height.length
- 1 <= n <= 2 * 10^4
- 0 <= height[i] <= 10^5

## 2. 初步問題解析與思維模型建立

當我們第一次看到這個問題時，應該如何思考？

首先，讓我們仔細閱讀問題陳述，提取關鍵信息：
- 我們有一個整數數組，代表一個地形的高度圖
- 每個數字代表一個寬度為1的柱子的高度
- 我們需要計算這個地形能夠儲存多少單位的雨水

觀察例子1：`[0,1,0,2,1,0,1,3,2,1,2,1]`
這個數組代表了一個地形，我們可以想像成一個直方圖。雨水會落在這個地形上，並被低處捕獲。

從直觀上看，雨水會被兩邊較高的柱子之間的低處捕獲。例如，位置2的高度為0，但它兩邊的高度分別為1和2，所以這個位置可以捕獲高度為1的雨水。

初步思維模型：我們可以想像為每個位置上方可能存在的"水柱"，其高度取決於這個位置左右兩側的最高柱子中的較小者，減去當前位置的高度。

## 3. 問題理解與核心挑戰

破解這個問題的核心在於理解：對於任何一個位置 i，能夠存儲的雨水量取決於：
1. 該位置左邊最高的柱子高度（leftMax）
2. 該位置右邊最高的柱子高度（rightMax）
3. 這兩個高度中的較小值減去當前位置的高度

表達式可以寫為：`min(leftMax, rightMax) - height[i]`（如果結果為正）

邊界情況：
- 如果數組長度為0或1，無法捕獲雨水，返回0
- 最左邊和最右邊的位置無法捕獲雨水（因為至少需要兩側都有柱子才能形成"容器"）

核心挑戰在於如何高效地計算每個位置的leftMax和rightMax。

## 4. 視覺問題表示

讓我們用圖表來表示問題。以示例1為例：`[0,1,0,2,1,0,1,3,2,1,2,1]`

```
    |                     
    |                #    
    |        #       #    #
    |    #   #   #   # #  #
    | #  # # # # # # # # #
    -----------------------
      0 1 0 2 1 0 1 3 2 1 2 1
```

我們可以使用不同的符號來表示柱子和雨水：

```
    |                     
    |                #    
    |        #       #    #
    |    #   #   #   # #  #
    | #  # ~ # ~ # ~ # # #
    -----------------------
      0 1 0 2 1 0 1 3 2 1 2 1
```

其中 "#" 代表柱子，"~" 代表捕獲的雨水。

通過這個視覺表示，我們可以清晰地看到，雨水只會被捕獲在兩個較高柱子之間的低處。

## 5. 問題模式識別

這個問題具有以下特徵：
- 需要考慮元素兩側的信息
- 對於每個位置，其結果取決於某種"最大值"或"邊界"
- 需要從數組中找出某種模式或關係

這些特徵表明，該問題可能適合使用以下策略：
1. 雙指針技巧
2. 預計算和存儲（例如，動態規劃或前綴/後綴數組）
3. 單調棧

特別是，這個問題是"雙邊約束"問題的典型代表，即一個位置的結果取決於其左右兩側的某種極值。

## 6. 策略性解題框架

針對這類雙邊約束問題，我們可以採用以下框架：

1. **預計算方法**：
    - 計算每個位置左側的最大值（leftMax）
    - 計算每個位置右側的最大值（rightMax）
    - 利用這兩個數組計算每個位置能捕獲的雨水量

2. **雙指針方法**：
    - 從數組的兩端向中間遍歷
    - 維護左右兩側當前遇到的最大高度
    - 根據左右最大高度的較小值決定當前處理的位置能捕獲多少雨水

3. **單調棧方法**：
    - 使用棧來維護一個單調遞減的高度序列
    - 當遇到較高的柱子時，計算能形成的"水槽"

這些方法各有優缺點，但都能有效解決問題。我們會逐一探討。

## 7. 算法設計（編碼前）

讓我們設計三種主要的算法方案，先用偽代碼表示：

### 方案一：動態規劃（預計算）

```
function trap(height):
    if height.length <= 2:
        return 0
    
    n = height.length
    leftMax = 新建長度為n的數組，初始值為0
    rightMax = 新建長度為n的數組，初始值為0
    
    // 預計算leftMax
    leftMax[0] = height[0]
    for i from 1 to n-1:
        leftMax[i] = max(leftMax[i-1], height[i])
    
    // 預計算rightMax
    rightMax[n-1] = height[n-1]
    for i from n-2 down to 0:
        rightMax[i] = max(rightMax[i+1], height[i])
    
    totalWater = 0
    for i from 0 to n-1:
        water = min(leftMax[i], rightMax[i]) - height[i]
        if water > 0:
            totalWater += water
    
    return totalWater
```

### 方案二：雙指針

```
function trap(height):
    if height.length <= 2:
        return 0
    
    left = 0
    right = height.length - 1
    leftMax = height[left]
    rightMax = height[right]
    totalWater = 0
    
    while left < right:
        if leftMax < rightMax:
            left++
            if height[left] > leftMax:
                leftMax = height[left]
            else:
                totalWater += leftMax - height[left]
        else:
            right--
            if height[right] > rightMax:
                rightMax = height[right]
            else:
                totalWater += rightMax - height[right]
    
    return totalWater
```

### 方案三：單調棧

```
function trap(height):
    if height.length <= 2:
        return 0
    
    stack = 新建空棧
    totalWater = 0
    
    for i from 0 to height.length-1:
        while stack非空 且 height[i] > height[stack.top()]:
            top = stack.pop()
            
            if stack為空:
                break
            
            width = i - stack.top() - 1
            boundedHeight = min(height[i], height[stack.top()]) - height[top]
            totalWater += width * boundedHeight
        
        stack.push(i)
    
    return totalWater
```

這三種方案都能解決問題，但有不同的時間和空間複雜度權衡。

## 8. 視覺説明

讓我們通過圖解來說明雙指針方法（方案二）的工作原理，因為它最為直觀。

以示例1為例：`[0,1,0,2,1,0,1,3,2,1,2,1]`

初始狀態：
```
left=0, right=11
leftMax=0, rightMax=1
totalWater=0

     |                     
     |                #    
     |        #       #    #
     |    #   #   #   # #  #
     | #  # ~ # ~ # ~ # # #
     -----------------------
       0 1 0 2 1 0 1 3 2 1 2 1
       L                   R
```

因為 leftMax(0) < rightMax(1)，我們移動左指針並更新 leftMax。

經過多次迭代後，雙指針會逐漸向中間靠攏，每次移動時計算當前位置能捕獲的雨水量。最終，所有位置都被處理完畢，得到總的雨水量。

## 9. 解決方案發展歷程

### 初始直觀方法（暴力解）
最直觀的方法是對每個位置 i，分別向左右兩側尋找最高的柱子，然後計算該位置能捕獲的雨水量。這種方法時間複雜度為 O(n²)，對於大型輸入效率較低。

### 優化思路：預計算
我們發現，對每個位置重複計算左右最大值是不必要的。可以預先計算並存儲每個位置的左右最大值，這樣每個位置只需 O(1) 時間就能計算出雨水量。這種方法將時間複雜度降低到 O(n)，但需要 O(n) 的額外空間。

### 進一步優化：雙指針
觀察到我們並不總是需要知道確切的左右最大值，而只需要知道它們中的較小者。利用這一點，可以使用雙指針方法在 O(n) 時間內解決問題，且只需 O(1) 的額外空間。

### 另一種思路：單調棧
通過維護一個單調遞減的柱子高度棧，我們可以在遇到較高柱子時立即計算能形成的"水槽"。這種方法也能在 O(n) 時間內解決問題，但實現略微複雜。

## 10. 實際例子演示

讓我們以示例1：`[0,1,0,2,1,0,1,3,2,1,2,1]` 為例，使用雙指針方法進行詳細演示：

初始狀態：
- left = 0, right = 11
- leftMax = 0, rightMax = 1
- totalWater = 0

迭代過程：
1. leftMax(0) < rightMax(1)，移動左指針：left = 1
    - 更新 leftMax = max(0, 1) = 1
    - 計算水量：max(1 - 1, 0) = 0
    - totalWater = 0
2. leftMax(1) = rightMax(1)，我們選擇移動左指針：left = 2
    - leftMax 保持不變 = 1
    - 計算水量：max(1 - 0, 0) = 1
    - totalWater = 0 + 1 = 1
3. 繼續迭代...

最終，當 left 和 right 相遇時，我們得到總水量 totalWater = 6。

## 11. Golang 實現

```go
func trap(height []int) int {
    // 邊界情況處理
    if len(height) <= 2 {
        return 0
    }
    
    left, right := 0, len(height)-1
    leftMax, rightMax := height[left], height[right]
    totalWater := 0
    
    for left < right {
        // 決定移動哪個指針
        if leftMax < rightMax {
            left++
            // 更新leftMax或計算捕獲的水
            if height[left] > leftMax {
                leftMax = height[left]
            } else {
                totalWater += leftMax - height[left]
            }
        } else {
            right--
            // 更新rightMax或計算捕獲的水
            if height[right] > rightMax {
                rightMax = height[right]
            } else {
                totalWater += rightMax - height[right]
            }
        }
    }
    
    return totalWater
}
```

讓我們也實現動態規劃（預計算）方法作為參考：

```go
func trap(height []int) int {
    n := len(height)
    if n <= 2 {
        return 0
    }
    
    // 預計算左右最大值
    leftMax := make([]int, n)
    rightMax := make([]int, n)
    
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])
    }
    
    rightMax[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i])
    }
    
    // 計算總水量
    totalWater := 0
    for i := 0; i < n; i++ {
        water := min(leftMax[i], rightMax[i]) - height[i]
        totalWater += water
    }
    
    return totalWater
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 12. 實現執行演示

以雙指針方法為例，讓我們演示代碼如何處理示例1：`[0,1,0,2,1,0,1,3,2,1,2,1]`

```
初始化：
- height = [0,1,0,2,1,0,1,3,2,1,2,1]
- left = 0, right = 11
- leftMax = 0, rightMax = 1
- totalWater = 0

第1次迭代：
- leftMax(0) < rightMax(1)，所以移動左指針
- left = 1
- 更新 leftMax = max(0, 1) = 1
- totalWater += max(0, leftMax - height[left]) = max(0, 1 - 1) = 0
- totalWater = 0

第2次迭代：
- leftMax(1) = rightMax(1)，我們選擇移動左指針
- left = 2
- leftMax 不變 = 1
- totalWater += max(0, leftMax - height[left]) = max(0, 1 - 0) = 1
- totalWater = 1

...（繼續迭代）

最終：
- left 和 right 相遇
- totalWater = 6
```

## 13. 複雜度分析

### 動態規劃（預計算）方法
- 時間複雜度：O(n)，我們需要遍歷數組三次（計算leftMax、rightMax和計算水量）
- 空間複雜度：O(n)，需要兩個額外的數組來存儲leftMax和rightMax

### 雙指針方法
- 時間複雜度：O(n)，最多遍歷數組一次
- 空間複雜度：O(1)，只需要幾個變數

### 單調棧方法
- 時間複雜度：O(n)，每個元素最多入棧、出棧各一次
- 空間複雜度：O(n)，最壞情況下棧可能存儲所有元素

在這三種方法中，雙指針方法在保持線性時間複雜度的同時，將空間複雜度降低到了常數級別，因此是最優的解決方案。

## 14. 優化與改進

已經討論了三種主要的解決方案，其中雙指針方法在時間和空間複雜度上都是最優的。

對於實際應用，可以考慮以下優化：
- 提前檢測特殊情況（如空數組、長度小於3的數組）
- 在適當情況下使用並行計算（例如，在預計算方法中並行計算leftMax和rightMax）

相關問題推薦：
- Container With Most Water（盛最多水的容器）
- Largest Rectangle in Histogram（柱狀圖中最大的矩形）
- Product of Array Except Self（除自身以外數組的乘積）

## 15. 通用解題智慧

從這個問題中，我們可以學到以下解題技巧：

1. **雙向思考**：有時從兩個方向考慮問題可以帶來突破性的解決方案，就像雙指針方法一樣。

2. **預計算價值**：有時提前計算和存儲某些值可以顯著提高效率，即使需要額外的空間。

3. **問題轉換**：將複雜問題分解為更簡單的子問題，如將"捕獲雨水"轉換為"找左右最大值"。

4. **視覺化幫助**：通過畫圖或模擬可以更好地理解問題和解決方案。

5. **空間-時間權衡**：認識到算法設計中常常存在空間和時間效率的權衡。
