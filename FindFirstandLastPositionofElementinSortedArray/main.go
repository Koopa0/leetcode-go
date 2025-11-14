package FindFirstandLastPositionofElementinSortedArray

func searchRange(nums []int, target int) []int {
	// 輔助函數：查找邊界
	findBoundary := func(findFirst bool) int {
		left, right := 0, len(nums)-1
		index := -1

		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target {
				index = mid
				if findFirst {
					right = mid - 1 // 查找第一個位置
				} else {
					left = mid + 1 // 查找最後一個位置
				}
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return index
	}

	// 處理空數組情況
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 查找第一個和最後一個位置
	first := findBoundary(true)
	last := findBoundary(false)

	return []int{first, last}
}
