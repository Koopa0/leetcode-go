package threeSumClosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "Example 2",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			name:   "Minimal Array",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   6,
		},
		{
			name:   "All Negative",
			nums:   []int{-5, -10, -8, -15},
			target: -30,
			want:   -30,
		},
		{
			name:   "Large Difference",
			nums:   []int{1, 2, 3, 4, 5},
			target: 100,
			want:   12,
		},
		{
			name:   "Exact Match",
			nums:   []int{5, 10, 15, 20, 25},
			target: 40,
			want:   40,
		},
		{
			name:   "Random Elements",
			nums:   []int{-3, 0, 1, -5, 4, 10},
			target: 7,
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.nums, tt.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
