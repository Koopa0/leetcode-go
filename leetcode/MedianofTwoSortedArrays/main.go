package MedianofTwoSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 確保 nums1 的長度不大於 nums2，這樣可以簡化邊界條件檢查
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	// 計算合併後數組的總長度
	totalLength := m + n
	// 計算要找的中間位置
	half := (totalLength + 1) / 2

	// 在 nums1 中二分查找合適的分割位置
	left, right := 0, m

	for left <= right {
		// 在 nums1 中的分割位置
		i := (left + right) / 2
		// 在 nums2 中的分割位置（確保兩部分總共有 half 個元素）
		j := half - i

		// 檢查分割位置是否合適
		if i < m && nums2[j-1] > nums1[i] {
			// nums1 的分割位置太小，需要右移
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// nums1 的分割位置太大，需要左移
			right = i - 1
		} else {
			// 找到合適的分割位置

			// 計算左半部分的最大值
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			// 如果總長度為奇數，中位數就是左半部分的最大值
			if totalLength%2 == 1 {
				return float64(maxLeft)
			}

			// 計算右半部分的最小值
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			// 如果總長度為偶數，中位數是左半部分最大值和右半部分最小值的平均值
			return float64(maxLeft+minRight) / 2.0
		}
	}

	// 這裡不應該到達，但為了讓編譯器開心，返回 0
	return 0.0
}
