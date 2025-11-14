package Permutations

func permute(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			numsCopy := make([]int, len(nums))
			copy(numsCopy, nums)
			result = append(result, numsCopy)
			return
		}

		for i := start; i < len(nums); i++ {
			// 交換 nums[start] 和 nums[i]
			nums[start], nums[i] = nums[i], nums[start]

			// 遞歸
			backtrack(start + 1)

			// 回溯：交換回來
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	return result
}
