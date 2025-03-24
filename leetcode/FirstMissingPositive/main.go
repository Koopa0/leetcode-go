package FirstMissingPositive

func firstMissingPositive(nums []int) int {
	// 第一步：將每個數字放到"正確"的位置上
	for i := range nums {
		// 當數字在有效範圍內且不在正確位置時進行交換
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 第二步：找出第一個不在正確位置上的數
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 如果所有數都在正確位置，返回n+1
	return len(nums) + 1
}
