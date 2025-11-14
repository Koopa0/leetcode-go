# 141. Linked List Cycle

## 題目描述

給定一個鏈結串列的頭節點 `head`，判斷鏈結串列中是否有環。

如果鏈結串列中存在某個節點，可以通過連續跟蹤 `next` 指標再次到達該節點，則鏈結串列中存在環。內部使用整數 `pos` 表示鏈結串列尾部連接到鏈結串列中的位置（索引從 0 開始）。注意，`pos` 不作為參數傳遞。

如果鏈結串列中存在環，則回傳 `true`。否則，回傳 `false`。

## 範例

### 範例 1:
```
輸入：head = [3,2,0,-4], pos = 1
輸出：true
解釋：鏈結串列中有一個環，尾部連接到第二個節點。
```

### 範例 2:
```
輸入：head = [1,2], pos = 0
輸出：true
解釋：鏈結串列中有一個環，尾部連接到第一個節點。
```

### 範例 3:
```
輸入：head = [1], pos = -1
輸出：false
解釋：鏈結串列中沒有環。
```

## 限制條件

- 鏈結串列中節點的數量範圍是 `[0, 10^4]`
- `-10^5 <= Node.val <= 10^5`
- `pos` 為 `-1` 或鏈結串列中的一個有效索引

## 進階

你能用 O(1) 空間複雜度解決此問題嗎？

---

## 解法思路

### 方法：Floyd's Cycle Detection Algorithm（快慢指標法）

使用兩個指標以不同速度遍歷鏈結串列：
- **slow 指標**：每次移動一步
- **fast 指標**：每次移動兩步

**核心概念：**
- 如果鏈結串列無環，fast 指標會先到達末尾（null）
- 如果鏈結串列有環，fast 指標最終會在環內追上 slow 指標

**為什麼 fast 一定會追上 slow？**
- 在環內，fast 每次比 slow 多走一步，相對距離每次縮短 1
- 無論環的大小如何，fast 最終都會追上 slow

### 複雜度分析

- **時間複雜度**：O(n)
  - 無環：fast 走完整個串列，最多 n/2 步
  - 有環：最壞情況下，slow 需要走 n 步進入環，fast 在環內追上 slow 最多需要環長度的步數

- **空間複雜度**：O(1)
  - 只使用兩個指標變數

### 其他解法比較

| 方法 | 時間複雜度 | 空間複雜度 | 說明 |
|------|-----------|-----------|------|
| 雜湊表 | O(n) | O(n) | 儲存訪問過的節點 |
| 快慢指標 | O(n) | O(1) | Floyd 演算法，最優解 |

## 實作細節

### 標準實作
```go
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head.Next

    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }

    return true
}
```

### 替代實作
```go
func hasCycleAlternative(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            return true
        }
    }

    return false
}
```

兩種實作的差異：
- **標準實作**：slow 和 fast 從不同位置開始，使用 `while (slow != fast)` 迴圈
- **替代實作**：兩個指標都從 head 開始，先移動再檢查是否相遇

效能差異微小，都是 O(1) 空間的最優解。

## 相關題目

- [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/) - 找出環的起始節點
- [202. Happy Number](https://leetcode.com/problems/happy-number/) - 同樣使用快慢指標檢測循環
- [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/) - Floyd 演算法的應用

---

## Description

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. Note that `pos` is not passed as a parameter.

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

## Constraints

- The number of the nodes in the list is in the range `[0, 10^4]`.
- `-10^5 <= Node.val <= 10^5`
- `pos` is `-1` or a valid index in the linked-list.

## Solution

Use Floyd's Cycle Detection Algorithm (two pointers moving at different speeds). If there's a cycle, the fast pointer will eventually meet the slow pointer.

**Time Complexity:** O(n)
**Space Complexity:** O(1)
