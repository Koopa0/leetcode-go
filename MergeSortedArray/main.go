package MergeSortedArray

// 使用額外空間的解決方案
func mergeExtraSpace(nums1 []int, m int, nums2 []int, n int) {
	// 創建一個新的陣列來存儲合併結果
	result := make([]int, m+n)

	// 指向兩個陣列的指針
	i, j, k := 0, 0, 0

	// 依次比較兩個陣列的元素，將較小的元素加入結果陣列
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result[k] = nums1[i]
			i++
		} else {
			result[k] = nums2[j]
			j++
		}
		k++
	}

	// 如果 nums1 還有剩餘元素，加入到結果陣列
	for i < m {
		result[k] = nums1[i]
		i++
		k++
	}

	// 如果 nums2 還有剩餘元素，加入到結果陣列
	for j < n {
		result[k] = nums2[j]
		j++
		k++
	}

	// 將結果複製回 nums1
	for i := 0; i < m+n; i++ {
		nums1[i] = result[i]
	}
}

// 從後向前合併的最優解決方案
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 三個指針，分別指向 nums1 的末尾，nums2 的末尾，以及合併結果的末尾
	p1, p2, p := m-1, n-1, m+n-1

	// 從後向前合併
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// 如果 nums2 還有剩餘元素，放入 nums1 的前部
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}

	// 注意：如果 p1 >= 0，不需要額外操作，因為這些元素已經在正確的位置
}
