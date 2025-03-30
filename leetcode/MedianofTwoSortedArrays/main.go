package MedianofTwoSortedArrays

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	// 確保 nums1 的長度不大於 nums2，這樣可以簡化邊界條件檢查
//	// 如果 nums1 比 nums2 長，則交換兩個數組
//	if len(nums1) > len(nums2) {
//		nums1, nums2 = nums2, nums1
//	}
//
//	// 取得兩個數組的長度
//	m, n := len(nums1), len(nums2)
//
//	// 計算合併後數組的總長度
//	totalLength := m + n
//
//	// 計算要找的中間位置（左半部分的元素個數）
//	// 對於奇數長度，左半部分比右半部分多一個元素
//	half := (totalLength + 1) / 2
//
//	// 在 nums1 中二分查找合適的分割位置
//	// left 和 right 是 nums1 中可能的分割位置範圍
//	left, right := 0, m
//
//	// 二分查找循環
//	for left <= right {
//		// 在 nums1 中的分割位置
//		i := (left + right) / 2
//
//		// 在 nums2 中的分割位置（確保兩部分總共有 half 個元素）
//		j := half - i
//
//		// 檢查分割位置是否合適
//		if i < m && nums2[j-1] > nums1[i] {
//
//			// 條件1: 如果 i < m 且 nums2[j-1] > nums1[i]
//			// 意味著 nums1 的分割點太靠左，需要右移
//			left = i + 1
//		} else if i > 0 && nums1[i-1] > nums2[j] {
//
//			// 條件2: 如果 i > 0 且 nums1[i-1] > nums2[j]
//			// 意味著 nums1 的分割點太靠右，需要左移
//			right = i - 1
//		} else {
//
//			// 找到合適的分割位置
//			// 計算左半部分的最大值
//			maxLeft := 0
//			if i == 0 { // 如果 nums1 的左半部分為空
//				maxLeft = nums2[j-1]
//			} else if j == 0 { // 如果 nums2 的左半部分為空
//				maxLeft = nums1[i-1]
//			} else { // 兩個數組的左半部分都不為空
//				maxLeft = max(nums1[i-1], nums2[j-1])
//			}
//
//			// 如果總長度為奇數，中位數就是左半部分的最大值
//			if totalLength%2 == 1 {
//				return float64(maxLeft)
//			}
//
//			// 計算右半部分的最小值
//			minRight := 0
//			if i == m { // 如果 nums1 的右半部分為空
//				minRight = nums2[j]
//			} else if j == n { // 如果 nums2 的右半部分為空
//				minRight = nums1[i]
//			} else { // 兩個數組的右半部分都不為空
//				minRight = min(nums1[i], nums2[j])
//			}
//
//			// 如果總長度為偶數，中位數是左半部分最大值和右半部分最小值的平均值
//			return float64(maxLeft+minRight) / 2.0
//		}
//
//	}
//
//	// 這裡不應該到達，但為了讓編譯器開心，返回 0
//	return 0.0
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 先確保
	// 1. len(nums2) > len(nums1)
	// 2. 找出 nums1,num2 的長度
	// 找出 half , 為假設 nums1, nums2 合併後的中位數的 index 位置
	// 找出 nums1 可能分割的位置, 最少為0 最大為 nums1 的長度

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)

	half := (m + n + 1) / 2

	left, right := 0, len(nums1)

	// 開始二分查找
	for left <= right {

		// 先假設 要在 nums1 和 nums2 要分割的位置 然後再開始慢慢做比對
		i := (left + right) / 2

		j := half - i

		if i < m && nums1[i] < nums2[j-1] {
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {

			maxLeft := 0

			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			// 計算右半部分的最小值
			minRight := 0
			if i == m { // 如果 nums1 的右半部分為空
				minRight = nums2[j]
			} else if j == n { // 如果 nums2 的右半部分為空
				minRight = nums1[i]
			} else { // 兩個數組的右半部分都不為空
				minRight = min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0

		}

	}

	return 0.0
}
