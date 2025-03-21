package MergeSortedArray

import "testing"

func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
		desc     string
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
			desc:     "範例 1",
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
			desc:     "範例 2：nums2 為空",
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
			desc:     "範例 3：nums1 為空",
		},
		{
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
			desc:     "nums2 的元素較小",
		},
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
			desc:     "nums2 的元素較大",
		},
		{
			nums1:    []int{1, 2, 3, 4, 5, 6},
			m:        6,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3, 4, 5, 6},
			desc:     "nums2 為空，nums1 已排序",
		},
		{
			nums1:    []int{0, 0, 0, 0, 0, 0},
			m:        0,
			nums2:    []int{1, 2, 3, 4, 5, 6},
			n:        6,
			expected: []int{1, 2, 3, 4, 5, 6},
			desc:     "nums1 為空，nums2 已排序",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			for i := 0; i < len(tc.expected); i++ {
				if tc.nums1[i] != tc.expected[i] {
					t.Errorf("expected nums1[%d] = %d, got %d", i, tc.expected[i], tc.nums1[i])
				}
			}
		})
	}
}
