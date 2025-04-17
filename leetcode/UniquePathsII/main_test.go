package UniquePathsII

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		expected     int
		desc         string
	}{
		{
			name:         "Example 1",
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expected:     2,
			desc:         "3x3 grid with one obstacle in the middle",
		},
		{
			name:         "Example 2",
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			expected:     1,
			desc:         "2x2 grid with one obstacle at top-right",
		},
		{
			name:         "Start is obstacle",
			obstacleGrid: [][]int{{1, 0}, {0, 0}},
			expected:     0,
			desc:         "Start position is an obstacle",
		},
		{
			name:         "End is obstacle",
			obstacleGrid: [][]int{{0, 0}, {0, 1}},
			expected:     0,
			desc:         "End position is an obstacle",
		},
		{
			name:         "Path blocked",
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 0, 0}},
			expected:     2,
			desc:         "All possible paths are blocked by obstacles",
		},
		{
			name:         "1x1 grid",
			obstacleGrid: [][]int{{0}},
			expected:     1,
			desc:         "Single cell with no obstacle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePathsWithObstacles(tt.obstacleGrid)
			if got != tt.expected {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkUniquePathsWithObstacles(b *testing.B) {
	// 準備測試資料
	obstacleGrid1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}} // 小型網格
	obstacleGrid2 := make([][]int, 10)                        // 中型網格
	for i := range obstacleGrid2 {
		obstacleGrid2[i] = make([]int, 10)
	}
	obstacleGrid2[5][5] = 1

	// 執行基準測試
	b.Run("BruteForce-Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bruteForceSolution(obstacleGrid1)
		}
	})

	b.Run("Optimized-Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			optimizedSolution(obstacleGrid1)
		}
	})

	b.Run("Optimal-Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uniquePathsWithObstacles(obstacleGrid1)
		}
	})

	b.Run("BruteForce-Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bruteForceSolution(obstacleGrid2)
		}
	})

	b.Run("Optimized-Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			optimizedSolution(obstacleGrid2)
		}
	})

	b.Run("Optimal-Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			uniquePathsWithObstacles(obstacleGrid2)
		}
	})
}
