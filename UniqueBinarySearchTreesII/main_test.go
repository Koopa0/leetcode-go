package UniqueBinarySearchTreesII

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int // 期望的樹數量
		desc     string
	}{
		{"測試用例1", 1, 1, "只有一個節點應該只有一種樹結構"},
		{"測試用例2", 2, 2, "兩個節點應該有兩種樹結構"},
		{"測試用例3", 3, 5, "三個節點應該有五種樹結構"},
		{"測試用例4", 4, 14, "四個節點應該有十四種樹結構"},
		{"測試用例5", 0, 0, "零個節點應該返回空列表"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTreesMemo(tt.input)
			if len(got) != tt.expected {
				t.Errorf("generateTreesMemo() 生成了 %v 棵樹, 期望 %v", len(got), tt.expected)
			}
		})
	}
}

// 測試生成的樹是否為有效的二元搜尋樹
func TestValidBST(t *testing.T) {
	for n := 1; n <= 4; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			trees := generateTreesMemo(n)
			for i, tree := range trees {
				if !isValidBST(tree, nil, nil) {
					t.Errorf("第 %d 棵樹不是有效的二元搜尋樹", i)
				}
			}
		})
	}
}

// 測試生成的樹是否包含所有從 1 到 n 的數字
func TestTreeContainsAllValues(t *testing.T) {
	for n := 1; n <= 4; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			trees := generateTreesMemo(n)
			for i, tree := range trees {
				values := make(map[int]bool)
				collectValues(tree, values)

				if len(values) != n {
					t.Errorf("第 %d 棵樹節點數量錯誤: 應為 %d, 實際為 %d", i, n, len(values))
				}

				for j := 1; j <= n; j++ {
					if !values[j] {
						t.Errorf("第 %d 棵樹缺少值 %d", i, j)
					}
				}
			}
		})
	}
}

// 測試生成的樹結構是否唯一
func TestUniqueTrees(t *testing.T) {
	for n := 1; n <= 4; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			trees := generateTreesMemo(n)
			seen := make(map[string]bool)

			for i, tree := range trees {
				serialized := serializeTree(tree)
				if seen[serialized] {
					t.Errorf("第 %d 棵樹結構重複", i)
				}
				seen[serialized] = true
			}

			// 確認樹的數量符合卡特蘭數
			expectedCount := catalanNumber(n)
			if len(trees) != expectedCount {
				t.Errorf("生成的樹數量錯誤: 應為 %d, 實際為 %d", expectedCount, len(trees))
			}
		})
	}
}

// 基準測試輔助函數
func BenchmarkCatalan(b *testing.B) {
	for i := 1; i <= 20; i++ {
		b.Run(fmt.Sprintf("n=%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				catalanNumber(i)
			}
		})
	}
}

// 輔助函數: 檢查樹是否為有效的二元搜尋樹
func isValidBST(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Val <= *min) || (max != nil && root.Val >= *max) {
		return false
	}

	return isValidBST(root.Left, min, &root.Val) && isValidBST(root.Right, &root.Val, max)
}

// 輔助函數: 收集樹中的所有值
func collectValues(root *TreeNode, values map[int]bool) {
	if root == nil {
		return
	}
	values[root.Val] = true
	collectValues(root.Left, values)
	collectValues(root.Right, values)
}

// 輔助函數: 將樹序列化為字串以便比較結構
func serializeTree(root *TreeNode) string {
	if root == nil {
		return "X"
	}

	return fmt.Sprintf("(%d,%s,%s)", root.Val, serializeTree(root.Left), serializeTree(root.Right))
}

// 輔助函數: 計算卡特蘭數
func catalanNumber(n int) int {
	if n <= 1 {
		return 1
	}

	// 計算 C(n) = (2n)! / ((n+1)! * n!)
	result := 1
	for i := 0; i < n; i++ {
		result = result * 2 * (2*i + 1) / (i + 2)
	}
	return result
}

// 比較不同大小輸入的執行時間關係
func BenchmarkScaling(b *testing.B) {
	sizes := []int{1, 2, 4, 6, 8}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				generateTreesMemo(size)
			}
		})
	}
}

// 測試記憶體使用情況
func BenchmarkMemoryUsage(b *testing.B) {
	b.Run("記憶體使用-遞迴分治法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateTreesMemo(8)
		}
	})

	b.Run("記憶體使用-記憶化遞迴法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateTreesMemo(8)
		}
	})

	b.Run("記憶體使用-動態規劃法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateTreesDP(8)
		}
	})
}
