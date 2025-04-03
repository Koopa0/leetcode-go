package TrappingRainWater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "Example 2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "Empty array",
			height: []int{},
			want:   0,
		},
		{
			name:   "Single element",
			height: []int{5},
			want:   0,
		},
		{
			name:   "Two elements",
			height: []int{5, 2},
			want:   0,
		},
		{
			name:   "No water can be trapped",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "All elements are same",
			height: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "Large test case",
			height: []int{0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0},
			want:   90,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
