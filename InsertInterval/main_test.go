package InsertInterval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
		desc        string
	}{
		{
			name:        "範例 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
			desc:        "新區間部分重疊",
		},
		{
			name:        "範例 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
			desc:        "新區間重疊多個現有區間",
		},
		{
			name:        "空輸入",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
			desc:        "輸入為空陣列",
		},
		{
			name:        "無重疊前",
			intervals:   [][]int{{3, 5}, {6, 7}},
			newInterval: []int{1, 2},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 7}},
			desc:        "新區間在所有現有區間之前",
		},
		{
			name:        "無重疊後",
			intervals:   [][]int{{3, 5}, {6, 7}},
			newInterval: []int{8, 10},
			expected:    [][]int{{3, 5}, {6, 7}, {8, 10}},
			desc:        "新區間在所有現有區間之後",
		},
		{
			name:        "無重疊中",
			intervals:   [][]int{{1, 2}, {6, 7}},
			newInterval: []int{3, 5},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 7}},
			desc:        "新區間在現有區間之間且不重疊",
		},
		{
			name:        "完全覆蓋",
			intervals:   [][]int{{3, 5}, {6, 7}},
			newInterval: []int{2, 8},
			expected:    [][]int{{2, 8}},
			desc:        "新區間完全覆蓋所有現有區間",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("insert() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	// 範例輸入
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	b.Run("暴力解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			insertBruteForce(intervals, newInterval)
		}
	})

	b.Run("優化解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			insertOptimized(intervals, newInterval)
		}
	})

	b.Run("最佳解法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			insert(intervals, newInterval)
		}
	})
}
