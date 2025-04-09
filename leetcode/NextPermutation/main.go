package NextPermutation

func nextPermutation(nums []int) {
	// 特殊情況處理：長度小於2的數組沒有其他排列
	if len(nums) < 2 {
		return
	}

	n := len(nums)

	// 步驟1：從右向左尋找第一個遞減的元素
	// 找到索引i，使得nums[i] < nums[i+1]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 步驟2：如果找到了遞減元素（i不為-1）
	if i >= 0 {
		// 從右向左尋找第一個大於nums[i]的元素
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}

		// 交換nums[i]和nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 步驟3：反轉從i+1到末尾的子數組
	// 不論是否找到遞減元素，都需要反轉
	// 如果i為-1（沒有找到遞減元素），則反轉整個數組
	reverse(nums, i+1, n-1)
}

// // reverse 反轉數組中指定範圍的元素
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
