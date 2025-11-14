package main

import "testing"

// 輔助函數：建立有環的鏈結串列
// values: 節點值
// pos: 環的起始位置（-1 表示無環）
func createCycleList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	// 建立所有節點
	nodes := make([]*ListNode, len(values))
	for i, val := range values {
		nodes[i] = &ListNode{Val: val}
	}

	// 連接節點
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// 如果 pos >= 0，建立環
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}

	return nodes[0]
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		pos    int
		want   bool
	}{
		{
			name:   "範例1：有環，環從位置1開始",
			values: []int{3, 2, 0, -4},
			pos:    1,
			want:   true,
		},
		{
			name:   "範例2：有環，環從位置0開始",
			values: []int{1, 2},
			pos:    0,
			want:   true,
		},
		{
			name:   "範例3：無環，單一節點",
			values: []int{1},
			pos:    -1,
			want:   false,
		},
		{
			name:   "空鏈結串列",
			values: []int{},
			pos:    -1,
			want:   false,
		},
		{
			name:   "兩個節點無環",
			values: []int{1, 2},
			pos:    -1,
			want:   false,
		},
		{
			name:   "兩個節點有環",
			values: []int{1, 2},
			pos:    1,
			want:   true,
		},
		{
			name:   "長鏈結串列無環",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:    -1,
			want:   false,
		},
		{
			name:   "長鏈結串列有環，環從中間開始",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:    5,
			want:   true,
		},
		{
			name:   "自環（節點指向自己）",
			values: []int{1},
			pos:    0,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.values, tt.pos)
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycleAlternative(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		pos    int
		want   bool
	}{
		{
			name:   "範例1：有環",
			values: []int{3, 2, 0, -4},
			pos:    1,
			want:   true,
		},
		{
			name:   "範例2：有環",
			values: []int{1, 2},
			pos:    0,
			want:   true,
		},
		{
			name:   "範例3：無環",
			values: []int{1},
			pos:    -1,
			want:   false,
		},
		{
			name:   "空鏈結串列",
			values: []int{},
			pos:    -1,
			want:   false,
		},
		{
			name:   "長鏈結串列有環",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:    3,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createCycleList(tt.values, tt.pos)
			got := hasCycleAlternative(head)
			if got != tt.want {
				t.Errorf("hasCycleAlternative() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基準測試：比較兩種實作方式
func BenchmarkHasCycle(b *testing.B) {
	benchmarks := []struct {
		name   string
		values []int
		pos    int
	}{
		{
			name:   "無環_短鏈結串列",
			values: []int{1, 2, 3, 4, 5},
			pos:    -1,
		},
		{
			name:   "有環_短鏈結串列",
			values: []int{1, 2, 3, 4, 5},
			pos:    2,
		},
		{
			name:   "無環_長鏈結串列",
			values: make([]int, 1000),
			pos:    -1,
		},
		{
			name:   "有環_長鏈結串列",
			values: make([]int, 1000),
			pos:    500,
		},
	}

	for _, bm := range benchmarks {
		// 初始化長鏈結串列的值
		if len(bm.values) == 1000 {
			for i := range bm.values {
				bm.values[i] = i
			}
		}

		b.Run("標準實作_"+bm.name, func(b *testing.B) {
			head := createCycleList(bm.values, bm.pos)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hasCycle(head)
			}
		})

		b.Run("替代實作_"+bm.name, func(b *testing.B) {
			head := createCycleList(bm.values, bm.pos)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hasCycleAlternative(head)
			}
		})
	}
}
