# LeetCode Solutions in Go

[繁體中文](#繁體中文) | [English](#english)

---

## 繁體中文

### 📖 專案簡介

這是一個使用 Go 語言實現的 LeetCode 題目解答集合。每個題目都包含：
- 詳細的問題定義（中英文對照）
- 清晰的解題思路和算法說明
- 完整的 Go 語言實現（含詳細註解）
- 全面的單元測試和性能基準測試

### 📊 統計信息

- **題目總數**：103 題
- **編程語言**：Go 1.24.1
- **測試覆蓋**：100%（每題都包含測試）
- **文檔完整度**：100%（每題都有 README）

### 🗂️ 專案結構

```
leetcode-go/
├── README.md           # 專案說明文件
├── go.mod             # Go 模組配置
├── .gitignore         # Git 忽略配置
│
├── TwoSum/            # 題目目錄範例
│   ├── README.md      # 題目說明和解法分析
│   ├── main.go        # 題目實現（含多種解法）
│   └── main_test.go   # 單元測試和基準測試
│
├── 3Sum/
│   ├── README.md
│   ├── main.go
│   └── main_test.go
│
└── ... (其他 101 題)
```

### 🚀 使用說明

#### 1. 克隆專案

```bash
git clone https://github.com/yourusername/leetcode-go.git
cd leetcode-go
```

#### 2. 運行單個題目的測試

```bash
# 運行某個題目的測試
go test ./TwoSum -v

# 運行基準測試
go test ./TwoSum -bench=. -benchmem
```

#### 3. 運行所有測試

```bash
# 測試所有題目
go test ./... -v
```

#### 4. 查看題目解法

每個題目目錄下都有 `README.md`，包含：
- 題目原文（中英對照）
- 解題思路和算法分析
- 複雜度分析
- 多種解法比較

### 📚 題目列表

以下是所有已完成的題目（按字母順序排列）：

| # | 題目 | 難度 | 主題標籤 |
|---|------|------|---------|
| 1 | [Two Sum](./TwoSum) | Easy | Array, Hash Table |
| 2 | [Add Two Numbers](./AddTwoNumbers) | Medium | Linked List, Math |
| 7 | [Reverse Integer](./ReverseInteger) | Medium | Math |
| 9 | [Palindrome Number](./PalindromeNumber) | Easy | Math |
| 13 | [Roman to Integer](./RomanToInteger) | Easy | Hash Table, Math, String |
| 15 | [3Sum](./3Sum) | Medium | Array, Two Pointers, Sorting |
| 16 | [3Sum Closest](./3SumClosest) | Medium | Array, Two Pointers, Sorting |
| 18 | [4Sum](./4Sum) | Medium | Array, Two Pointers, Sorting |
| 19 | [Remove Nth Node From End of List](./RemoveNthNodeFromEndOfList) | Medium | Linked List, Two Pointers |
| 20 | [Valid Parentheses](./ValidParentheses) | Easy | String, Stack |
| 21 | [Merge Two Sorted Lists](./MergeTwoSortedLists) | Easy | Linked List, Recursion |
| 23 | [Merge k Sorted Lists](./MergekSortedLists) | Hard | Linked List, Divide and Conquer, Heap |
| 24 | [Swap Nodes in Pairs](./SwapNodesinPairs) | Medium | Linked List, Recursion |
| 25 | [Reverse Nodes in k-Group](./ReverseNodesink-Group) | Hard | Linked List, Recursion |
| 26 | [Remove Duplicates from Sorted Array](./RemoveDuplicatesFromSortedArray) | Easy | Array, Two Pointers |
| 27 | [Remove Element](./RemoveElement) | Easy | Array, Two Pointers |
| 29 | [Divide Two Integers](./DivideTwoIntegers) | Medium | Math, Bit Manipulation |
| 31 | [Next Permutation](./NextPermutation) | Medium | Array, Two Pointers |
| 36 | [Valid Sudoku](./ValidSudoku) | Medium | Array, Hash Table, Matrix |
| 37 | [Sudoku Solver](./SudokuSolver) | Hard | Array, Backtracking, Matrix |
| 38 | [Count and Say](./CountandSay) | Medium | String |
| 39 | [Combination Sum](./CombinationSum) | Medium | Array, Backtracking |
| 40 | [Combination Sum II](./CombinationSumII) | Medium | Array, Backtracking |
| 41 | [First Missing Positive](./FirstMissingPositive) | Hard | Array, Hash Table |
| 44 | [Wildcard Matching](./WildcardMatching) | Hard | String, Dynamic Programming, Greedy, Recursion |
| 46 | [Permutations](./Permutations) | Medium | Array, Backtracking |
| 47 | [Permutations II](./PermutationsII) | Medium | Array, Backtracking |
| 48 | [Rotate Image](./RotateImage) | Medium | Array, Math, Matrix |
| 50 | [Pow(x, n)](./PowXN) | Medium | Math, Recursion |
| 51 | [N-Queens](./N-Queens) | Hard | Array, Backtracking |
| 52 | [N-Queens II](./N-Queens-II) | Hard | Backtracking |
| 53 | [Maximum Subarray](./MaximumSubarray) | Medium | Array, Divide and Conquer, Dynamic Programming |
| 54 | [Spiral Matrix](./SpiralMatrix) | Medium | Array, Matrix, Simulation |
| 59 | [Spiral Matrix II](./SpiralMatrixII) | Medium | Array, Matrix, Simulation |
| 60 | [Permutation Sequence](./PermutationSequence) | Hard | Math, Recursion |
| 61 | [Rotate List](./RotateList) | Medium | Linked List, Two Pointers |
| 62 | [Unique Paths](./UniquePaths) | Medium | Math, Dynamic Programming, Combinatorics |
| 63 | [Unique Paths II](./UniquePathsII) | Medium | Array, Dynamic Programming, Matrix |
| 64 | [Minimum Path Sum](./MinimumPathSum) | Medium | Array, Dynamic Programming, Matrix |
| 66 | [Plus One](./PlusOne) | Easy | Array, Math |
| 67 | [Add Binary](./AddBinary) | Easy | Math, String, Bit Manipulation, Simulation |
| 69 | [Sqrt(x)](./SqrtX) | Easy | Math, Binary Search |
| 70 | [Climbing Stairs](./ClimbingStairs) | Easy | Math, Dynamic Programming, Memoization |
| 72 | [Edit Distance](./EditDistance) | Medium | String, Dynamic Programming |
| 73 | [Set Matrix Zeroes](./SetMatrixZeroes) | Medium | Array, Hash Table, Matrix |
| 75 | [Sort Colors](./SortColors) | Medium | Array, Two Pointers, Sorting |
| 76 | [Minimum Window Substring](./MinimumWindowSubstring) | Hard | Hash Table, String, Sliding Window |
| 78 | [Subsets](./Subsets) | Medium | Array, Backtracking, Bit Manipulation |
| 79 | [Word Search](./WordSearch) | Medium | Array, Backtracking, Matrix |
| 80 | [Remove Duplicates from Sorted Array II](./RemoveDuplicatesFromSortedArrayII) | Medium | Array, Two Pointers |
| 81 | [Search in Rotated Sorted Array II](./SearchInRotatedSortedArrayII) | Medium | Array, Binary Search |
| 82 | [Remove Duplicates from Sorted List II](./RemoveDuplicatesFromSortedListII) | Medium | Linked List, Two Pointers |
| 83 | [Remove Duplicates from Sorted List](./RemoveDuplicatesFromSortedList) | Easy | Linked List |
| 84 | [Largest Rectangle in Histogram](./LargestRectangleInHistogram) | Hard | Array, Stack, Monotonic Stack |
| 85 | [Maximal Rectangle](./MaximalRectangle) | Hard | Array, Dynamic Programming, Stack, Matrix, Monotonic Stack |
| 86 | [Partition List](./PartitionList) | Medium | Linked List, Two Pointers |
| 87 | [Scramble String](./ScrambleString) | Hard | String, Dynamic Programming |
| 88 | [Merge Sorted Array](./MergeSortedArray) | Easy | Array, Two Pointers, Sorting |
| 90 | [Subsets II](./SubsetsII) | Medium | Array, Backtracking, Bit Manipulation |
| 92 | [Reverse Linked List II](./ReverseLinkedListII) | Medium | Linked List |
| 93 | [Restore IP Addresses](./RestoreIPAddresses) | Medium | String, Backtracking |
| 94 | [Binary Tree Inorder Traversal](./BinaryTreeInorderTraversal) | Easy | Stack, Tree, Depth-First Search, Binary Tree |
| 95 | [Unique Binary Search Trees II](./UniqueBinarySearchTreesII) | Medium | Dynamic Programming, Backtracking, Tree, Binary Search Tree, Binary Tree |
| 96 | [Unique Binary Search Trees](./UniqueBinarySearchTrees) | Medium | Math, Dynamic Programming, Tree, Binary Search Tree, Binary Tree |
| 97 | [Interleaving String](./InterleavingString) | Medium | String, Dynamic Programming |
| 98 | [Validate Binary Search Tree](./ValidateBinarySearchTree) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| 99 | [Recover Binary Search Tree](./RecoverBinarySearchTree) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| 100 | [Same Tree](./SameTree) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| 101 | [Symmetric Tree](./SymmetricTree) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| 121 | [Best Time to Buy and Sell Stock](./BestTimeToBuyAndSellStock) | Easy | Array, Dynamic Programming |
| 125 | [Valid Palindrome](./ValidPalindrome) | Easy | Two Pointers, String |
| 238 | [Product of Array Except Self](./ProductofArrayExceptSelf) | Medium | Array, Prefix Sum |

*更多題目請查看各自的目錄...*

### 🧪 測試說明

本專案所有題目都包含完整的測試：

- **單元測試**：驗證各種解法的正確性
- **基準測試**：比較不同解法的性能
- **邊界測試**：覆蓋各種邊界情況

測試範例：
```go
func TestTwoSum(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        target   int
        expected []int
    }{
        {"基本測試", []int{2, 7, 11, 15}, 9, []int{0, 1}},
        {"負數測試", []int{-1, -2, -3, -4}, -6, []int{1, 3}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := twoSum(tt.nums, tt.target)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("got %v, want %v", got, tt.expected)
            }
        })
    }
}
```

### 💡 代碼特色

- **多種解法**：大部分題目提供暴力解、優化解等多種實現
- **詳細註解**：每行關鍵代碼都有中文註解
- **性能優化**：使用高效的數據結構和算法
- **最佳實踐**：遵循 Go 語言編碼規範

### 📖 學習建議

1. **按主題學習**：相同類型的題目一起練習（如動態規劃、二叉樹等）
2. **理解原理**：不僅要會寫代碼，更要理解算法原理
3. **多做測試**：通過測試用例驗證理解
4. **性能分析**：使用基準測試比較不同解法的性能

### 🤝 貢獻

歡迎提交 Issue 或 Pull Request！

### 📄 授權

MIT License

---

## English

### 📖 Project Description

A comprehensive collection of LeetCode problem solutions implemented in Go. Each problem includes:
- Detailed problem definitions (bilingual: English & Traditional Chinese)
- Clear solution explanations and algorithm analysis
- Complete Go implementations with detailed comments
- Comprehensive unit tests and benchmark tests

### 📊 Statistics

- **Total Problems**: 103
- **Programming Language**: Go 1.24.1
- **Test Coverage**: 100% (all problems include tests)
- **Documentation**: 100% (all problems include README)

### 🗂️ Project Structure

```
leetcode-go/
├── README.md           # Project documentation
├── go.mod             # Go module configuration
├── .gitignore         # Git ignore configuration
│
├── TwoSum/            # Problem directory example
│   ├── README.md      # Problem description and solution analysis
│   ├── main.go        # Implementation (multiple approaches)
│   └── main_test.go   # Unit tests and benchmarks
│
├── 3Sum/
│   ├── README.md
│   ├── main.go
│   └── main_test.go
│
└── ... (101 more problems)
```

### 🚀 Getting Started

#### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/leetcode-go.git
cd leetcode-go
```

#### 2. Run Tests for a Single Problem

```bash
# Run tests
go test ./TwoSum -v

# Run benchmarks
go test ./TwoSum -bench=. -benchmem
```

#### 3. Run All Tests

```bash
# Test all problems
go test ./... -v
```

### 💡 Code Features

- **Multiple Solutions**: Most problems include brute force, optimized approaches
- **Detailed Comments**: Key code lines include explanatory comments
- **Performance Optimized**: Efficient data structures and algorithms
- **Best Practices**: Follows Go coding conventions

### 🤝 Contributing

Issues and Pull Requests are welcome!

### 📄 License

MIT License
