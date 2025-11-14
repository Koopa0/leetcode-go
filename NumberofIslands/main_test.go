package main

import "testing"

// 輔助函數：複製網格，避免測試之間互相影響
func copyGrid(grid [][]byte) [][]byte {
	if grid == nil {
		return nil
	}
	result := make([][]byte, len(grid))
	for i := range grid {
		result[i] = make([]byte, len(grid[i]))
		copy(result[i], grid[i])
	}
	return result
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "範例1：一個大島嶼",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "範例2：三個島嶼",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "空網格",
			grid: [][]byte{},
			want: 0,
		},
		{
			name: "全是水",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "全是陸地",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "單一格子 - 陸地",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "單一格子 - 水",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: "對角線不相連",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
		{
			name: "環形島嶼",
			grid: [][]byte{
				{'1', '1', '1', '1'},
				{'1', '0', '0', '1'},
				{'1', '0', '0', '1'},
				{'1', '1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "多個小島嶼",
			grid: [][]byte{
				{'1', '0', '1', '0', '1'},
				{'0', '0', '0', '0', '0'},
				{'1', '0', '1', '0', '1'},
			},
			want: 6,
		},
		{
			name: "L形島嶼",
			grid: [][]byte{
				{'1', '0', '0'},
				{'1', '0', '0'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "條狀島嶼",
			grid: [][]byte{
				{'1'},
				{'1'},
				{'1'},
				{'1'},
				{'1'},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 複製網格給 DFS 使用
			gridCopy := copyGrid(tt.grid)
			got := numIslands(gridCopy)
			if got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumIslandsBFS(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "範例1：一個大島嶼",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "範例2：三個島嶼",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "全是陸地",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "對角線不相連",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridCopy := copyGrid(tt.grid)
			got := numIslandsBFS(gridCopy)
			if got != tt.want {
				t.Errorf("numIslandsBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumIslandsDFSIterative(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "範例1：一個大島嶼",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "範例2：三個島嶼",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "環形島嶼",
			grid: [][]byte{
				{'1', '1', '1', '1'},
				{'1', '0', '0', '1'},
				{'1', '0', '0', '1'},
				{'1', '1', '1', '1'},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridCopy := copyGrid(tt.grid)
			got := numIslandsDFSIterative(gridCopy)
			if got != tt.want {
				t.Errorf("numIslandsDFSIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基準測試：比較三種解法的效能
func BenchmarkNumIslands(b *testing.B) {
	benchmarks := []struct {
		name string
		grid [][]byte
	}{
		{
			name: "小網格_3x3",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
		},
		{
			name: "中網格_10x10_多個島嶼",
			grid: [][]byte{
				{'1', '1', '0', '0', '0', '1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0', '1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0', '0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
				{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
				{'1', '1', '0', '0', '0', '1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0', '1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0', '0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
				{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
			},
		},
		{
			name: "大網格_20x20_全是陸地",
			grid: func() [][]byte {
				grid := make([][]byte, 20)
				for i := range grid {
					grid[i] = make([]byte, 20)
					for j := range grid[i] {
						grid[i][j] = '1'
					}
				}
				return grid
			}(),
		},
	}

	for _, bm := range benchmarks {
		b.Run("DFS遞迴_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridCopy := copyGrid(bm.grid)
				numIslands(gridCopy)
			}
		})

		b.Run("BFS_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridCopy := copyGrid(bm.grid)
				numIslandsBFS(gridCopy)
			}
		})

		b.Run("DFS迭代_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridCopy := copyGrid(bm.grid)
				numIslandsDFSIterative(gridCopy)
			}
		})
	}
}
